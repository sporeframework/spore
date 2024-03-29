package protocol

import (
	"context"
	"crypto/sha256"
	hex "encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	reflect "reflect"
	"strconv"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/crypto"
)

var ps *pubsub.PubSub
var pubsubTopic string

// server is used to implement helloworld.GreeterServer.
type server struct {
	UnimplementedSporeServer
}

func StartRPCServer(topic string, pubsub *pubsub.PubSub, p *int) {
	ps = pubsub
	pubsubTopic = topic
	port := ":" + strconv.Itoa(*p)
	fmt.Println("RPC interface listenting on tcp", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterSporeServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) CreateContract(ctx context.Context, in *Transaction) (*TransactionResponse, error) {
	log.Printf("Received: %v", hex.EncodeToString(in.GetData()))
	if checkSignature(in) == false {
		return nil, errors.New("Could not validate signature")
	}
	setMetadata(in)
	req := &Request{
		Type:        Request_CREATE_CONTRACT,
		Transaction: in,
	}
	msgBytes, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = ps.Publish(pubsubTopic, msgBytes)
	dataHash := sha256.Sum256(in.Data)
	return &TransactionResponse{TransactionId: dataHash[:]}, nil
}

func (s *server) GetTransaction(ctx context.Context, in *TransactionId) (*Transaction, error) {
	log.Printf("Received: %v", hex.EncodeToString(in.GetTransactionId()))

	log.Println("Querying database with txn id: ", hex.EncodeToString(in.GetTransactionId()))
	// we don't have to broadcast this call to the network, it is a local query
	txnBytes, err := Database.Get([]byte(DatabaseNamespace), in.GetTransactionId())
	if err != nil {
		return nil, err
	}

	txn := &Transaction{}
	err = proto.Unmarshal(txnBytes, txn)
	if err != nil {
		return nil, err
	}
	return txn, nil
}

// Send implements Spore.Send
func (s *server) Send(ctx context.Context, in *Transaction) (*TransactionResponse, error) {
	log.Printf("Received: %v", hex.EncodeToString(in.GetData()))
	if !checkSignature(in) {
		return nil, errors.New("Could not validate signature")
	}
	setMetadata(in)
	req := &Request{
		Type:        Request_SEND_TRANSACTION,
		Transaction: in,
	}
	msgBytes, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	err = ps.Publish(pubsubTopic, msgBytes)

	return &TransactionResponse{TransactionId: in.GetId()}, nil
}

func setMetadata(in *Transaction) error {

	// set the id to the transaction's hash
	txnBytes, _ := proto.Marshal(in)
	dataHash := sha256.Sum256(txnBytes)
	in.Id = dataHash[:]

	now := time.Now().Unix()
	in.Created = now
	return nil
}

func checkSignature(txn *Transaction) bool {
	sig := make([]byte, len(txn.Signature))
	copy(sig, txn.Signature)
	txn.Signature = nil
	txnBytes, _ := proto.Marshal(txn)
	pSum := sha256.Sum256(txnBytes)
	//fmt.Printf("pSum: %s\n", hex.EncodeToString(pSum[:]))

	// reset sig
	txn.Signature = sig
	publicKey, _ := crypto.SigToPub(pSum[:], txn.Signature)
	//publicKey, _ := crypto.Ecrecover(pSum[:], txn.Signature)
	address := crypto.PubkeyToAddress(*publicKey)
	//fmt.Printf("sig: %s\n", hex.EncodeToString(sig))
	//fmt.Printf("address: %s\n", hex.EncodeToString(address.Bytes()))
	//fmt.Printf("public tx.from: %s\n", hex.EncodeToString(txn.GetFrom()))
	return reflect.DeepEqual(address.Bytes(), txn.GetFrom())

}
