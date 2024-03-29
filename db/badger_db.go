package db

import (
	"context"
	"fmt"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

const (
	// Default BadgerDB discardRatio. It represents the discard ratio for the
	// BadgerDB GC.
	//
	// Ref: https://godoc.org/github.com/dgraph-io/badger#DB.RunValueLogGC
	badgerDiscardRatio = 0.5

	// Default BadgerDB GC interval
	badgerGCInterval = 10 * time.Minute
)

var (
	// BadgerAlertNamespace defines the alerts BadgerDB namespace.
	BadgerAlertNamespace = []byte("alerts")
)

type (
	// DB defines an embedded key/value store database interface.
	DB interface {
		Get(namespace, key []byte) (value []byte, err error)
		Set(namespace, key, value []byte) error
		Has(namespace, key []byte) (bool, error)
		Close() error
	}

	// BadgerDB is a wrapper around a BadgerDB backend database that implements
	// the DB interface.
	BadgerDB struct {
		db         *badger.DB
		ctx        context.Context
		cancelFunc context.CancelFunc
	}
)

// NewBadgerDB returns a new initialized BadgerDB database implementing the DB
// interface. If the database cannot be initialized, an error will be returned.
func NewBadgerDB(dataDir string) (DB, error) {
	badgerDB, err := badger.Open(badger.DefaultOptions(dataDir))
	if err != nil {
		return nil, err
	}

	bdb := &BadgerDB{
		db: badgerDB,
	}
	bdb.ctx, bdb.cancelFunc = context.WithCancel(context.Background())

	go bdb.runGC()
	return bdb, nil
}

// Get implements the DB interface. It attempts to get a value for a given key
// and namespace. If the key does not exist in the provided namespace, an error
// is returned, otherwise the retrieved value.
func (bdb *BadgerDB) Get(namespace, key []byte) (value []byte, err error) {
	err = bdb.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(badgerNamespaceKey(namespace, key))
		if err != nil {
			return err
		}

		// Copy the value as the value provided Badger is only valid while the
		// transaction is open.
		//value = make([]byte, len(tmpValue))
		//copy(value, tmpValue)

		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return value, nil
}

// Set implements the DB interface. It attempts to store a value for a given key
// and namespace. If the key/value pair cannot be saved, an error is returned.
func (bdb *BadgerDB) Set(namespace, key, value []byte) error {
	err := bdb.db.Update(func(txn *badger.Txn) error {
		return txn.Set(badgerNamespaceKey(namespace, key), value)
	})

	if err != nil {
		bdb.db.Opts().Logger.Errorf("failed to set key %s for namespace %s: %v", key, namespace, err)
		return err
	}

	return nil
}

// Has implements the DB interface. It returns a boolean reflecting if the
// datbase has a given key for a namespace or not. An error is only returned if
// an error to Get would be returned that is not of type badger.ErrKeyNotFound.
func (bdb *BadgerDB) Has(namespace, key []byte) (ok bool, err error) {
	_, err = bdb.Get(namespace, key)
	switch err {
	case badger.ErrKeyNotFound:
		ok, err = false, nil
	case nil:
		ok, err = true, nil
	}

	return
}

// Close implements the DB interface. It closes the connection to the underlying
// BadgerDB database as well as invoking the context's cancel function.
func (bdb *BadgerDB) Close() error {
	bdb.cancelFunc()
	return bdb.db.Close()
}

// runGC triggers the garbage collection for the BadgerDB backend database. It
// should be run in a goroutine.
func (bdb *BadgerDB) runGC() {
	ticker := time.NewTicker(badgerGCInterval)
	for {
		select {
		case <-ticker.C:
			err := bdb.db.RunValueLogGC(badgerDiscardRatio)
			if err != nil {
				// don't report error when GC didn't result in any cleanup
				if err == badger.ErrNoRewrite {
					bdb.db.Opts().Logger.Debugf("no BadgerDB GC occurred: %v", err)
				} else {
					bdb.db.Opts().Logger.Errorf("failed to GC BadgerDB: %v", err)
				}
			}

		case <-bdb.ctx.Done():
			return
		}
	}
}

// badgerNamespaceKey returns a composite key used for lookup and storage for a
// given namespace and key.
func badgerNamespaceKey(namespace, key []byte) []byte {
	prefix := []byte(fmt.Sprintf("%s/", namespace))
	return append(prefix, key...)
}
