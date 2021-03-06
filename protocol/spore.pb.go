// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: spore.proto

package protocol

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request_Type int32

const (
	Request_SEND_TRANSACTION Request_Type = 0
	Request_CREATE_CONTRACT  Request_Type = 1
	Request_UPDATE_PEER      Request_Type = 2
)

// Enum value maps for Request_Type.
var (
	Request_Type_name = map[int32]string{
		0: "SEND_TRANSACTION",
		1: "CREATE_CONTRACT",
		2: "UPDATE_PEER",
	}
	Request_Type_value = map[string]int32{
		"SEND_TRANSACTION": 0,
		"CREATE_CONTRACT":  1,
		"UPDATE_PEER":      2,
	}
)

func (x Request_Type) Enum() *Request_Type {
	p := new(Request_Type)
	*p = x
	return p
}

func (x Request_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Request_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_spore_proto_enumTypes[0].Descriptor()
}

func (Request_Type) Type() protoreflect.EnumType {
	return &file_spore_proto_enumTypes[0]
}

func (x Request_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Request_Type.Descriptor instead.
func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return file_spore_proto_rawDescGZIP(), []int{0, 0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Request_Type `protobuf:"varint,1,opt,name=type,proto3,enum=main.Request_Type" json:"type,omitempty"`
	Transaction *Transaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	UpdatePeer  *UpdatePeer  `protobuf:"bytes,3,opt,name=updatePeer,proto3" json:"updatePeer,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_spore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_spore_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetType() Request_Type {
	if x != nil {
		return x.Type
	}
	return Request_SEND_TRANSACTION
}

func (x *Request) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *Request) GetUpdatePeer() *UpdatePeer {
	if x != nil {
		return x.UpdatePeer
	}
	return nil
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Created  int64  `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Id       []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	To       []byte `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	From     []byte `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
	Gas      int64  `protobuf:"varint,6,opt,name=gas,proto3" json:"gas,omitempty"`
	GasPrice int64  `protobuf:"varint,7,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Nonce    int32  `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Contract bool   `protobuf:"varint,9,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_spore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_spore_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Transaction) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Transaction) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Transaction) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *Transaction) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *Transaction) GetGas() int64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *Transaction) GetGasPrice() int64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *Transaction) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetContract() bool {
	if x != nil {
		return x.Contract
	}
	return false
}

// The response message containing the greetings
type TransactionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId []byte `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

func (x *TransactionResponse) Reset() {
	*x = TransactionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionResponse) ProtoMessage() {}

func (x *TransactionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionResponse.ProtoReflect.Descriptor instead.
func (*TransactionResponse) Descriptor() ([]byte, []int) {
	return file_spore_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionResponse) GetTransactionId() []byte {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

type UpdatePeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserHandle []byte `protobuf:"bytes,1,opt,name=userHandle,proto3" json:"userHandle,omitempty"`
}

func (x *UpdatePeer) Reset() {
	*x = UpdatePeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePeer) ProtoMessage() {}

func (x *UpdatePeer) ProtoReflect() protoreflect.Message {
	mi := &file_spore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePeer.ProtoReflect.Descriptor instead.
func (*UpdatePeer) Descriptor() ([]byte, []int) {
	return file_spore_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePeer) GetUserHandle() []byte {
	if x != nil {
		return x.UserHandle
	}
	return nil
}

var File_spore_proto protoreflect.FileDescriptor

var file_spore_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x70, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d,
	0x61, 0x69, 0x6e, 0x22, 0xdc, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x26, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x65, 0x72, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x22, 0x42,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x45, 0x45, 0x52,
	0x10, 0x02, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x2c, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x32,
	0x81, 0x01, 0x0a, 0x05, 0x53, 0x70, 0x6f, 0x72, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x53, 0x65, 0x6e,
	0x64, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x65, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x73, 0x70, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spore_proto_rawDescOnce sync.Once
	file_spore_proto_rawDescData = file_spore_proto_rawDesc
)

func file_spore_proto_rawDescGZIP() []byte {
	file_spore_proto_rawDescOnce.Do(func() {
		file_spore_proto_rawDescData = protoimpl.X.CompressGZIP(file_spore_proto_rawDescData)
	})
	return file_spore_proto_rawDescData
}

var file_spore_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_spore_proto_goTypes = []interface{}{
	(Request_Type)(0),           // 0: main.Request.Type
	(*Request)(nil),             // 1: main.Request
	(*Transaction)(nil),         // 2: main.Transaction
	(*TransactionResponse)(nil), // 3: main.TransactionResponse
	(*UpdatePeer)(nil),          // 4: main.UpdatePeer
}
var file_spore_proto_depIdxs = []int32{
	0, // 0: main.Request.type:type_name -> main.Request.Type
	2, // 1: main.Request.transaction:type_name -> main.Transaction
	4, // 2: main.Request.updatePeer:type_name -> main.UpdatePeer
	2, // 3: main.Spore.Send:input_type -> main.Transaction
	2, // 4: main.Spore.CreateContract:input_type -> main.Transaction
	3, // 5: main.Spore.Send:output_type -> main.TransactionResponse
	3, // 6: main.Spore.CreateContract:output_type -> main.TransactionResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spore_proto_init() }
func file_spore_proto_init() {
	if File_spore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePeer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spore_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spore_proto_goTypes,
		DependencyIndexes: file_spore_proto_depIdxs,
		EnumInfos:         file_spore_proto_enumTypes,
		MessageInfos:      file_spore_proto_msgTypes,
	}.Build()
	File_spore_proto = out.File
	file_spore_proto_rawDesc = nil
	file_spore_proto_goTypes = nil
	file_spore_proto_depIdxs = nil
}