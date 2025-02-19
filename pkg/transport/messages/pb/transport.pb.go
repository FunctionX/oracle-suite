// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: transport.proto

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

type Price struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Price:
	Wat string `protobuf:"bytes,1,opt,name=wat,proto3" json:"wat,omitempty"`  // asset name
	Val []byte `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`  // big.Int encoded as bytes
	Age int64  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"` // timestamp
	// Ethereum Signature:
	Vrs []byte `protobuf:"bytes,4,opt,name=vrs,proto3" json:"vrs,omitempty"` // v, r, s combined into one byte array
	// Additional data:
	Trace   []byte `protobuf:"bytes,8,opt,name=trace,proto3" json:"trace,omitempty"`
	Version string `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Price) Reset() {
	*x = Price{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Price) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Price) ProtoMessage() {}

func (x *Price) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Price.ProtoReflect.Descriptor instead.
func (*Price) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{0}
}

func (x *Price) GetWat() string {
	if x != nil {
		return x.Wat
	}
	return ""
}

func (x *Price) GetVal() []byte {
	if x != nil {
		return x.Val
	}
	return nil
}

func (x *Price) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Price) GetVrs() []byte {
	if x != nil {
		return x.Vrs
	}
	return nil
}

func (x *Price) GetTrace() []byte {
	if x != nil {
		return x.Trace
	}
	return nil
}

func (x *Price) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             string                      `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id               []byte                      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Index            []byte                      `protobuf:"bytes,3,opt,name=index,proto3" json:"index,omitempty"`
	EventTimestamp   int64                       `protobuf:"varint,4,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	MessageTimestamp int64                       `protobuf:"varint,5,opt,name=messageTimestamp,proto3" json:"messageTimestamp,omitempty"`
	Data             map[string][]byte           `protobuf:"bytes,6,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signatures       map[string]*Event_Signature `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Event) GetIndex() []byte {
	if x != nil {
		return x.Index
	}
	return nil
}

func (x *Event) GetEventTimestamp() int64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

func (x *Event) GetMessageTimestamp() int64 {
	if x != nil {
		return x.MessageTimestamp
	}
	return 0
}

func (x *Event) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetSignatures() map[string]*Event_Signature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

type DataPointMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model     string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Value     []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DataPointMessage) Reset() {
	*x = DataPointMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPointMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPointMessage) ProtoMessage() {}

func (x *DataPointMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPointMessage.ProtoReflect.Descriptor instead.
func (*DataPointMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2}
}

func (x *DataPointMessage) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *DataPointMessage) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *DataPointMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type MuSigInitializeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID          []byte            `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	StartedAtTimestamp int64             `protobuf:"varint,2,opt,name=startedAtTimestamp,proto3" json:"startedAtTimestamp,omitempty"`
	MsgType            string            `protobuf:"bytes,3,opt,name=msgType,proto3" json:"msgType,omitempty"`
	MsgBody            []byte            `protobuf:"bytes,4,opt,name=msgBody,proto3" json:"msgBody,omitempty"`
	MsgMeta            map[string][]byte `protobuf:"bytes,5,rep,name=msgMeta,proto3" json:"msgMeta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signers            [][]byte          `protobuf:"bytes,6,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (x *MuSigInitializeMessage) Reset() {
	*x = MuSigInitializeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigInitializeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigInitializeMessage) ProtoMessage() {}

func (x *MuSigInitializeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigInitializeMessage.ProtoReflect.Descriptor instead.
func (*MuSigInitializeMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{3}
}

func (x *MuSigInitializeMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigInitializeMessage) GetStartedAtTimestamp() int64 {
	if x != nil {
		return x.StartedAtTimestamp
	}
	return 0
}

func (x *MuSigInitializeMessage) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *MuSigInitializeMessage) GetMsgBody() []byte {
	if x != nil {
		return x.MsgBody
	}
	return nil
}

func (x *MuSigInitializeMessage) GetMsgMeta() map[string][]byte {
	if x != nil {
		return x.MsgMeta
	}
	return nil
}

func (x *MuSigInitializeMessage) GetSigners() [][]byte {
	if x != nil {
		return x.Signers
	}
	return nil
}

type MuSigTerminateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Reason    string `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *MuSigTerminateMessage) Reset() {
	*x = MuSigTerminateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigTerminateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigTerminateMessage) ProtoMessage() {}

func (x *MuSigTerminateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigTerminateMessage.ProtoReflect.Descriptor instead.
func (*MuSigTerminateMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{4}
}

func (x *MuSigTerminateMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigTerminateMessage) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type MuSigCommitmentMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID  []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Commitment []byte `protobuf:"bytes,2,opt,name=commitment,proto3" json:"commitment,omitempty"`
	PubKeyX    []byte `protobuf:"bytes,3,opt,name=pubKeyX,proto3" json:"pubKeyX,omitempty"`
	PubKeyY    []byte `protobuf:"bytes,4,opt,name=pubKeyY,proto3" json:"pubKeyY,omitempty"`
}

func (x *MuSigCommitmentMessage) Reset() {
	*x = MuSigCommitmentMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigCommitmentMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigCommitmentMessage) ProtoMessage() {}

func (x *MuSigCommitmentMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigCommitmentMessage.ProtoReflect.Descriptor instead.
func (*MuSigCommitmentMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{5}
}

func (x *MuSigCommitmentMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigCommitmentMessage) GetCommitment() []byte {
	if x != nil {
		return x.Commitment
	}
	return nil
}

func (x *MuSigCommitmentMessage) GetPubKeyX() []byte {
	if x != nil {
		return x.PubKeyX
	}
	return nil
}

func (x *MuSigCommitmentMessage) GetPubKeyY() []byte {
	if x != nil {
		return x.PubKeyY
	}
	return nil
}

type MuSigNonceMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Nonce     []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *MuSigNonceMessage) Reset() {
	*x = MuSigNonceMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigNonceMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigNonceMessage) ProtoMessage() {}

func (x *MuSigNonceMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigNonceMessage.ProtoReflect.Descriptor instead.
func (*MuSigNonceMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{6}
}

func (x *MuSigNonceMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigNonceMessage) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

type MuSigPartialSignatureMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID        []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	PartialSignature []byte `protobuf:"bytes,2,opt,name=partialSignature,proto3" json:"partialSignature,omitempty"`
}

func (x *MuSigPartialSignatureMessage) Reset() {
	*x = MuSigPartialSignatureMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigPartialSignatureMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigPartialSignatureMessage) ProtoMessage() {}

func (x *MuSigPartialSignatureMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigPartialSignatureMessage.ProtoReflect.Descriptor instead.
func (*MuSigPartialSignatureMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{7}
}

func (x *MuSigPartialSignatureMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigPartialSignatureMessage) GetPartialSignature() []byte {
	if x != nil {
		return x.PartialSignature
	}
	return nil
}

type MuSigSignatureMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *MuSigSignatureMessage) Reset() {
	*x = MuSigSignatureMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MuSigSignatureMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MuSigSignatureMessage) ProtoMessage() {}

func (x *MuSigSignatureMessage) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MuSigSignatureMessage.ProtoReflect.Descriptor instead.
func (*MuSigSignatureMessage) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{8}
}

func (x *MuSigSignatureMessage) GetSessionID() []byte {
	if x != nil {
		return x.SessionID
	}
	return nil
}

func (x *MuSigSignatureMessage) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MuSigSignatureMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MuSigSignatureMessage) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type Event_Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer    []byte `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Event_Signature) Reset() {
	*x = Event_Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Signature) ProtoMessage() {}

func (x *Event_Signature) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Signature.ProtoReflect.Descriptor instead.
func (*Event_Signature) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Event_Signature) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *Event_Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type DataPointMessage_Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signer    []byte `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *DataPointMessage_Signature) Reset() {
	*x = DataPointMessage_Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataPointMessage_Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPointMessage_Signature) ProtoMessage() {}

func (x *DataPointMessage_Signature) ProtoReflect() protoreflect.Message {
	mi := &file_transport_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPointMessage_Signature.ProtoReflect.Descriptor instead.
func (*DataPointMessage_Signature) Descriptor() ([]byte, []int) {
	return file_transport_proto_rawDescGZIP(), []int{2, 0}
}

func (x *DataPointMessage_Signature) GetSigner() []byte {
	if x != nil {
		return x.Signer
	}
	return nil
}

func (x *DataPointMessage_Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_transport_proto protoreflect.FileDescriptor

var file_transport_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7f, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x61,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x76,
	0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x22, 0xc0, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a,
	0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x36, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x1a, 0x41, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xb0, 0x02, 0x0a, 0x16, 0x4d, 0x75, 0x53, 0x69,
	0x67, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x2e, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73,
	0x67, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x3e, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x4d, 0x75, 0x53, 0x69, 0x67, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x1a, 0x3a,
	0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x15, 0x4d, 0x75,
	0x53, 0x69, 0x67, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x16, 0x4d, 0x75,
	0x53, 0x69, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x58, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x58, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x59, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x59, 0x22, 0x47, 0x0a, 0x11, 0x4d, 0x75, 0x53, 0x69, 0x67, 0x4e,
	0x6f, 0x6e, 0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22,
	0x68, 0x0a, 0x1c, 0x4d, 0x75, 0x53, 0x69, 0x67, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x2a, 0x0a,
	0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x7b, 0x0a, 0x15, 0x4d, 0x75, 0x53,
	0x69, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x72, 0x6f, 0x6e, 0x69, 0x63, 0x6c, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2d, 0x73, 0x75,
	0x69, 0x74, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transport_proto_rawDescOnce sync.Once
	file_transport_proto_rawDescData = file_transport_proto_rawDesc
)

func file_transport_proto_rawDescGZIP() []byte {
	file_transport_proto_rawDescOnce.Do(func() {
		file_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_proto_rawDescData)
	})
	return file_transport_proto_rawDescData
}

var file_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_transport_proto_goTypes = []interface{}{
	(*Price)(nil),                        // 0: Price
	(*Event)(nil),                        // 1: Event
	(*DataPointMessage)(nil),             // 2: DataPointMessage
	(*MuSigInitializeMessage)(nil),       // 3: MuSigInitializeMessage
	(*MuSigTerminateMessage)(nil),        // 4: MuSigTerminateMessage
	(*MuSigCommitmentMessage)(nil),       // 5: MuSigCommitmentMessage
	(*MuSigNonceMessage)(nil),            // 6: MuSigNonceMessage
	(*MuSigPartialSignatureMessage)(nil), // 7: MuSigPartialSignatureMessage
	(*MuSigSignatureMessage)(nil),        // 8: MuSigSignatureMessage
	(*Event_Signature)(nil),              // 9: Event.Signature
	nil,                                  // 10: Event.DataEntry
	nil,                                  // 11: Event.SignaturesEntry
	(*DataPointMessage_Signature)(nil),   // 12: DataPointMessage.Signature
	nil,                                  // 13: MuSigInitializeMessage.MsgMetaEntry
}
var file_transport_proto_depIdxs = []int32{
	10, // 0: Event.data:type_name -> Event.DataEntry
	11, // 1: Event.signatures:type_name -> Event.SignaturesEntry
	13, // 2: MuSigInitializeMessage.msgMeta:type_name -> MuSigInitializeMessage.MsgMetaEntry
	9,  // 3: Event.SignaturesEntry.value:type_name -> Event.Signature
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_transport_proto_init() }
func file_transport_proto_init() {
	if File_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Price); i {
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
		file_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPointMessage); i {
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
		file_transport_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigInitializeMessage); i {
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
		file_transport_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigTerminateMessage); i {
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
		file_transport_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigCommitmentMessage); i {
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
		file_transport_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigNonceMessage); i {
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
		file_transport_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigPartialSignatureMessage); i {
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
		file_transport_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MuSigSignatureMessage); i {
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
		file_transport_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Signature); i {
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
		file_transport_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPointMessage_Signature); i {
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
			RawDescriptor: file_transport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_proto_goTypes,
		DependencyIndexes: file_transport_proto_depIdxs,
		MessageInfos:      file_transport_proto_msgTypes,
	}.Build()
	File_transport_proto = out.File
	file_transport_proto_rawDesc = nil
	file_transport_proto_goTypes = nil
	file_transport_proto_depIdxs = nil
}
