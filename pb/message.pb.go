// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: message.proto

package pb

import (
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

type MessageType int32

const (
	MessageType_Invalid  MessageType = 0
	MessageType_Keygen1  MessageType = 1
	MessageType_Keygen2  MessageType = 2
	MessageType_Keygen3  MessageType = 3
	MessageType_Refresh1 MessageType = 4
	MessageType_Refresh2 MessageType = 5
	MessageType_Refresh3 MessageType = 6
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "Invalid",
		1: "Keygen1",
		2: "Keygen2",
		3: "Keygen3",
		4: "Refresh1",
		5: "Refresh2",
		6: "Refresh3",
	}
	MessageType_value = map[string]int32{
		"Invalid":  0,
		"Keygen1":  1,
		"Keygen2":  2,
		"Keygen3":  3,
		"Refresh1": 4,
		"Refresh2": 5,
		"Refresh3": 6,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=pb.MessageType" json:"type,omitempty"`
	From uint32      `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	To   uint32      `protobuf:"varint,3,opt,name=to,proto3" json:"to,omitempty"`
	// Types that are assignable to Content:
	//	*Message_Keygen1
	//	*Message_Keygen2
	//	*Message_Keygen3
	//	*Message_Refresh1
	//	*Message_Refresh2
	//	*Message_Refresh3
	Content isMessage_Content `protobuf_oneof:"content"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_Invalid
}

func (x *Message) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Message) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

func (m *Message) GetContent() isMessage_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Message) GetKeygen1() *KeygenMessage1 {
	if x, ok := x.GetContent().(*Message_Keygen1); ok {
		return x.Keygen1
	}
	return nil
}

func (x *Message) GetKeygen2() *KeygenMessage2 {
	if x, ok := x.GetContent().(*Message_Keygen2); ok {
		return x.Keygen2
	}
	return nil
}

func (x *Message) GetKeygen3() *KeygenMessage3 {
	if x, ok := x.GetContent().(*Message_Keygen3); ok {
		return x.Keygen3
	}
	return nil
}

func (x *Message) GetRefresh1() *RefreshMessage1 {
	if x, ok := x.GetContent().(*Message_Refresh1); ok {
		return x.Refresh1
	}
	return nil
}

func (x *Message) GetRefresh2() *RefreshMessage2 {
	if x, ok := x.GetContent().(*Message_Refresh2); ok {
		return x.Refresh2
	}
	return nil
}

func (x *Message) GetRefresh3() *RefreshMessage3 {
	if x, ok := x.GetContent().(*Message_Refresh3); ok {
		return x.Refresh3
	}
	return nil
}

type isMessage_Content interface {
	isMessage_Content()
}

type Message_Keygen1 struct {
	Keygen1 *KeygenMessage1 `protobuf:"bytes,4,opt,name=keygen1,proto3,oneof"`
}

type Message_Keygen2 struct {
	Keygen2 *KeygenMessage2 `protobuf:"bytes,5,opt,name=keygen2,proto3,oneof"`
}

type Message_Keygen3 struct {
	Keygen3 *KeygenMessage3 `protobuf:"bytes,6,opt,name=keygen3,proto3,oneof"`
}

type Message_Refresh1 struct {
	Refresh1 *RefreshMessage1 `protobuf:"bytes,7,opt,name=refresh1,proto3,oneof"`
}

type Message_Refresh2 struct {
	Refresh2 *RefreshMessage2 `protobuf:"bytes,8,opt,name=refresh2,proto3,oneof"`
}

type Message_Refresh3 struct {
	Refresh3 *RefreshMessage3 `protobuf:"bytes,9,opt,name=refresh3,proto3,oneof"`
}

func (*Message_Keygen1) isMessage_Content() {}

func (*Message_Keygen2) isMessage_Content() {}

func (*Message_Keygen3) isMessage_Content() {}

func (*Message_Refresh1) isMessage_Content() {}

func (*Message_Refresh2) isMessage_Content() {}

func (*Message_Refresh3) isMessage_Content() {}

type KeygenMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *KeygenMessage1) Reset() {
	*x = KeygenMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeygenMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeygenMessage1) ProtoMessage() {}

func (x *KeygenMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeygenMessage1.ProtoReflect.Descriptor instead.
func (*KeygenMessage1) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *KeygenMessage1) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type KeygenMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rid []byte `protobuf:"bytes,1,opt,name=rid,proto3" json:"rid,omitempty"`
	X   *Point `protobuf:"bytes,2,opt,name=X,proto3" json:"X,omitempty"`
	A   *Point `protobuf:"bytes,3,opt,name=A,proto3" json:"A,omitempty"`
	U   []byte `protobuf:"bytes,4,opt,name=u,proto3" json:"u,omitempty"`
}

func (x *KeygenMessage2) Reset() {
	*x = KeygenMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeygenMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeygenMessage2) ProtoMessage() {}

func (x *KeygenMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeygenMessage2.ProtoReflect.Descriptor instead.
func (*KeygenMessage2) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *KeygenMessage2) GetRid() []byte {
	if x != nil {
		return x.Rid
	}
	return nil
}

func (x *KeygenMessage2) GetX() *Point {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *KeygenMessage2) GetA() *Point {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *KeygenMessage2) GetU() []byte {
	if x != nil {
		return x.U
	}
	return nil
}

type KeygenMessage3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchX *Scalar `protobuf:"bytes,1,opt,name=schX,proto3" json:"schX,omitempty"`
}

func (x *KeygenMessage3) Reset() {
	*x = KeygenMessage3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeygenMessage3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeygenMessage3) ProtoMessage() {}

func (x *KeygenMessage3) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeygenMessage3.ProtoReflect.Descriptor instead.
func (*KeygenMessage3) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *KeygenMessage3) GetSchX() *Scalar {
	if x != nil {
		return x.SchX
	}
	return nil
}

type RefreshMessage1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *RefreshMessage1) Reset() {
	*x = RefreshMessage1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessage1) ProtoMessage() {}

func (x *RefreshMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessage1.ProtoReflect.Descriptor instead.
func (*RefreshMessage1) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshMessage1) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type RefreshMessage2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X   []*Point `protobuf:"bytes,1,rep,name=X,proto3" json:"X,omitempty"`
	A   []*Point `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	Y   *Point   `protobuf:"bytes,3,opt,name=Y,proto3" json:"Y,omitempty"`
	B   *Point   `protobuf:"bytes,4,opt,name=B,proto3" json:"B,omitempty"`
	N   *Int     `protobuf:"bytes,5,opt,name=N,proto3" json:"N,omitempty"`
	S   *Int     `protobuf:"bytes,6,opt,name=S,proto3" json:"S,omitempty"`
	T   *Int     `protobuf:"bytes,7,opt,name=T,proto3" json:"T,omitempty"`
	Rho []byte   `protobuf:"bytes,8,opt,name=rho,proto3" json:"rho,omitempty"`
	U   []byte   `protobuf:"bytes,9,opt,name=u,proto3" json:"u,omitempty"`
}

func (x *RefreshMessage2) Reset() {
	*x = RefreshMessage2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessage2) ProtoMessage() {}

func (x *RefreshMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessage2.ProtoReflect.Descriptor instead.
func (*RefreshMessage2) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshMessage2) GetX() []*Point {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *RefreshMessage2) GetA() []*Point {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *RefreshMessage2) GetY() *Point {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *RefreshMessage2) GetB() *Point {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *RefreshMessage2) GetN() *Int {
	if x != nil {
		return x.N
	}
	return nil
}

func (x *RefreshMessage2) GetS() *Int {
	if x != nil {
		return x.S
	}
	return nil
}

func (x *RefreshMessage2) GetT() *Int {
	if x != nil {
		return x.T
	}
	return nil
}

func (x *RefreshMessage2) GetRho() []byte {
	if x != nil {
		return x.Rho
	}
	return nil
}

func (x *RefreshMessage2) GetU() []byte {
	if x != nil {
		return x.U
	}
	return nil
}

type RefreshMessage3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mod  *ZKMod      `protobuf:"bytes,1,opt,name=mod,proto3" json:"mod,omitempty"`
	Prm  *ZKPrm      `protobuf:"bytes,2,opt,name=prm,proto3" json:"prm,omitempty"`
	C    *Ciphertext `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	SchX []*Scalar   `protobuf:"bytes,4,rep,name=schX,proto3" json:"schX,omitempty"`
	SchY *Scalar     `protobuf:"bytes,5,opt,name=schY,proto3" json:"schY,omitempty"`
}

func (x *RefreshMessage3) Reset() {
	*x = RefreshMessage3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshMessage3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshMessage3) ProtoMessage() {}

func (x *RefreshMessage3) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshMessage3.ProtoReflect.Descriptor instead.
func (*RefreshMessage3) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshMessage3) GetMod() *ZKMod {
	if x != nil {
		return x.Mod
	}
	return nil
}

func (x *RefreshMessage3) GetPrm() *ZKPrm {
	if x != nil {
		return x.Prm
	}
	return nil
}

func (x *RefreshMessage3) GetC() *Ciphertext {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *RefreshMessage3) GetSchX() []*Scalar {
	if x != nil {
		return x.SchX
	}
	return nil
}

func (x *RefreshMessage3) GetSchY() *Scalar {
	if x != nil {
		return x.SchY
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x08, 0x7a, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x03, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x2e, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x31, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x31, 0x12,
	0x2e, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x32, 0x12,
	0x2e, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x33, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x33, 0x48, 0x00, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x33, 0x12,
	0x31, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x31, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x32, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x32, 0x12, 0x31, 0x0a, 0x08, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x33, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x48, 0x00, 0x52, 0x08,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x33, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x62, 0x0a, 0x0e, 0x4b, 0x65, 0x79,
	0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x10, 0x0a, 0x03, 0x72,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x01, 0x58, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x01, 0x58, 0x12, 0x17, 0x0a, 0x01, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x01, 0x41, 0x12,
	0x0c, 0x0a, 0x01, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x75, 0x22, 0x30, 0x0a,
	0x0e, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x12,
	0x1e, 0x0a, 0x04, 0x73, 0x63, 0x68, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x04, 0x73, 0x63, 0x68, 0x58, 0x22,
	0x25, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xda, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x17, 0x0a, 0x01, 0x58, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x01, 0x58, 0x12, 0x17, 0x0a, 0x01, 0x41, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x01, 0x41, 0x12, 0x17, 0x0a, 0x01,
	0x59, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x01, 0x59, 0x12, 0x17, 0x0a, 0x01, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x01, 0x42, 0x12, 0x15,
	0x0a, 0x01, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x01, 0x4e, 0x12, 0x15, 0x0a, 0x01, 0x53, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x01, 0x53, 0x12, 0x15, 0x0a, 0x01,
	0x54, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x74,
	0x52, 0x01, 0x54, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x68, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x72, 0x68, 0x6f, 0x12, 0x0c, 0x0a, 0x01, 0x75, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x01, 0x75, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x12, 0x1b, 0x0a, 0x03, 0x6d, 0x6f, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x5a, 0x4b, 0x4d, 0x6f, 0x64, 0x52,
	0x03, 0x6d, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x03, 0x70, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x5a, 0x4b, 0x50, 0x72, 0x6d, 0x52, 0x03, 0x70, 0x72,
	0x6d, 0x12, 0x1c, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x52, 0x01, 0x63, 0x12,
	0x1e, 0x0a, 0x04, 0x73, 0x63, 0x68, 0x58, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x04, 0x73, 0x63, 0x68, 0x58, 0x12,
	0x1e, 0x0a, 0x04, 0x73, 0x63, 0x68, 0x59, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x04, 0x73, 0x63, 0x68, 0x59, 0x2a,
	0x6b, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4b,
	0x65, 0x79, 0x67, 0x65, 0x6e, 0x31, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x67,
	0x65, 0x6e, 0x32, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x67, 0x65, 0x6e, 0x33,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x31, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x32, 0x10, 0x05, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x33, 0x10, 0x06, 0x42, 0x25, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x61, 0x75, 0x72, 0x75,
	0x73, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x63, 0x6d, 0x70, 0x2d, 0x65, 0x63, 0x64, 0x73, 0x61,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_message_proto_goTypes = []interface{}{
	(MessageType)(0),        // 0: pb.MessageType
	(*Message)(nil),         // 1: pb.Message
	(*KeygenMessage1)(nil),  // 2: pb.KeygenMessage1
	(*KeygenMessage2)(nil),  // 3: pb.KeygenMessage2
	(*KeygenMessage3)(nil),  // 4: pb.KeygenMessage3
	(*RefreshMessage1)(nil), // 5: pb.RefreshMessage1
	(*RefreshMessage2)(nil), // 6: pb.RefreshMessage2
	(*RefreshMessage3)(nil), // 7: pb.RefreshMessage3
	(*Point)(nil),           // 8: pb.Point
	(*Scalar)(nil),          // 9: pb.Scalar
	(*Int)(nil),             // 10: pb.Int
	(*ZKMod)(nil),           // 11: pb.ZKMod
	(*ZKPrm)(nil),           // 12: pb.ZKPrm
	(*Ciphertext)(nil),      // 13: pb.Ciphertext
}
var file_message_proto_depIdxs = []int32{
	0,  // 0: pb.Message.type:type_name -> pb.MessageType
	2,  // 1: pb.Message.keygen1:type_name -> pb.KeygenMessage1
	3,  // 2: pb.Message.keygen2:type_name -> pb.KeygenMessage2
	4,  // 3: pb.Message.keygen3:type_name -> pb.KeygenMessage3
	5,  // 4: pb.Message.refresh1:type_name -> pb.RefreshMessage1
	6,  // 5: pb.Message.refresh2:type_name -> pb.RefreshMessage2
	7,  // 6: pb.Message.refresh3:type_name -> pb.RefreshMessage3
	8,  // 7: pb.KeygenMessage2.X:type_name -> pb.Point
	8,  // 8: pb.KeygenMessage2.A:type_name -> pb.Point
	9,  // 9: pb.KeygenMessage3.schX:type_name -> pb.Scalar
	8,  // 10: pb.RefreshMessage2.X:type_name -> pb.Point
	8,  // 11: pb.RefreshMessage2.A:type_name -> pb.Point
	8,  // 12: pb.RefreshMessage2.Y:type_name -> pb.Point
	8,  // 13: pb.RefreshMessage2.B:type_name -> pb.Point
	10, // 14: pb.RefreshMessage2.N:type_name -> pb.Int
	10, // 15: pb.RefreshMessage2.S:type_name -> pb.Int
	10, // 16: pb.RefreshMessage2.T:type_name -> pb.Int
	11, // 17: pb.RefreshMessage3.mod:type_name -> pb.ZKMod
	12, // 18: pb.RefreshMessage3.prm:type_name -> pb.ZKPrm
	13, // 19: pb.RefreshMessage3.c:type_name -> pb.Ciphertext
	9,  // 20: pb.RefreshMessage3.schX:type_name -> pb.Scalar
	9,  // 21: pb.RefreshMessage3.schY:type_name -> pb.Scalar
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	22, // [22:22] is the sub-list for extension type_name
	22, // [22:22] is the sub-list for extension extendee
	0,  // [0:22] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_types_proto_init()
	file_zk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeygenMessage1); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeygenMessage2); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeygenMessage3); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessage1); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessage2); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshMessage3); i {
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
	file_message_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_Keygen1)(nil),
		(*Message_Keygen2)(nil),
		(*Message_Keygen3)(nil),
		(*Message_Refresh1)(nil),
		(*Message_Refresh2)(nil),
		(*Message_Refresh3)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
