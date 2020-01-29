// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package anycable

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_ERROR   Status = 0
	Status_SUCCESS Status = 1
	Status_FAILURE Status = 2
)

var Status_name = map[int32]string{
	0: "ERROR",
	1: "SUCCESS",
	2: "FAILURE",
}

var Status_value = map[string]int32{
	"ERROR":   0,
	"SUCCESS": 1,
	"FAILURE": 2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

type Env struct {
	Path                 string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Env) Reset()         { *m = Env{} }
func (m *Env) String() string { return proto.CompactTextString(m) }
func (*Env) ProtoMessage()    {}
func (*Env) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Env) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Env.Unmarshal(m, b)
}
func (m *Env) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Env.Marshal(b, m, deterministic)
}
func (m *Env) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Env.Merge(m, src)
}
func (m *Env) XXX_Size() int {
	return xxx_messageInfo_Env.Size(m)
}
func (m *Env) XXX_DiscardUnknown() {
	xxx_messageInfo_Env.DiscardUnknown(m)
}

var xxx_messageInfo_Env proto.InternalMessageInfo

func (m *Env) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Env) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type ConnectionRequest struct {
	Path                 string            `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Headers              map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env                  *Env              `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (m *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(m, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ConnectionRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *ConnectionRequest) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type ConnectionResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	Identifiers          string   `protobuf:"bytes,2,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	Transmissions        []string `protobuf:"bytes,3,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,4,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionResponse) Reset()         { *m = ConnectionResponse{} }
func (m *ConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*ConnectionResponse) ProtoMessage()    {}
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *ConnectionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionResponse.Unmarshal(m, b)
}
func (m *ConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionResponse.Marshal(b, m, deterministic)
}
func (m *ConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionResponse.Merge(m, src)
}
func (m *ConnectionResponse) XXX_Size() int {
	return xxx_messageInfo_ConnectionResponse.Size(m)
}
func (m *ConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionResponse proto.InternalMessageInfo

func (m *ConnectionResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *ConnectionResponse) GetIdentifiers() string {
	if m != nil {
		return m.Identifiers
	}
	return ""
}

func (m *ConnectionResponse) GetTransmissions() []string {
	if m != nil {
		return m.Transmissions
	}
	return nil
}

func (m *ConnectionResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type CommandMessage struct {
	Command               string   `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Identifier            string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	ConnectionIdentifiers string   `protobuf:"bytes,3,opt,name=connection_identifiers,json=connectionIdentifiers,proto3" json:"connection_identifiers,omitempty"`
	Data                  string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Env                   *Env     `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *CommandMessage) Reset()         { *m = CommandMessage{} }
func (m *CommandMessage) String() string { return proto.CompactTextString(m) }
func (*CommandMessage) ProtoMessage()    {}
func (*CommandMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *CommandMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandMessage.Unmarshal(m, b)
}
func (m *CommandMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandMessage.Marshal(b, m, deterministic)
}
func (m *CommandMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandMessage.Merge(m, src)
}
func (m *CommandMessage) XXX_Size() int {
	return xxx_messageInfo_CommandMessage.Size(m)
}
func (m *CommandMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommandMessage proto.InternalMessageInfo

func (m *CommandMessage) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandMessage) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CommandMessage) GetConnectionIdentifiers() string {
	if m != nil {
		return m.ConnectionIdentifiers
	}
	return ""
}

func (m *CommandMessage) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *CommandMessage) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type CommandResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	Disconnect           bool     `protobuf:"varint,2,opt,name=disconnect,proto3" json:"disconnect,omitempty"`
	StopStreams          bool     `protobuf:"varint,3,opt,name=stop_streams,json=stopStreams,proto3" json:"stop_streams,omitempty"`
	Streams              []string `protobuf:"bytes,4,rep,name=streams,proto3" json:"streams,omitempty"`
	Transmissions        []string `protobuf:"bytes,5,rep,name=transmissions,proto3" json:"transmissions,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,6,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandResponse) Reset()         { *m = CommandResponse{} }
func (m *CommandResponse) String() string { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()    {}
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *CommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandResponse.Unmarshal(m, b)
}
func (m *CommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandResponse.Marshal(b, m, deterministic)
}
func (m *CommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandResponse.Merge(m, src)
}
func (m *CommandResponse) XXX_Size() int {
	return xxx_messageInfo_CommandResponse.Size(m)
}
func (m *CommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommandResponse proto.InternalMessageInfo

func (m *CommandResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *CommandResponse) GetDisconnect() bool {
	if m != nil {
		return m.Disconnect
	}
	return false
}

func (m *CommandResponse) GetStopStreams() bool {
	if m != nil {
		return m.StopStreams
	}
	return false
}

func (m *CommandResponse) GetStreams() []string {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *CommandResponse) GetTransmissions() []string {
	if m != nil {
		return m.Transmissions
	}
	return nil
}

func (m *CommandResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type DisconnectRequest struct {
	Identifiers          string            `protobuf:"bytes,1,opt,name=identifiers,proto3" json:"identifiers,omitempty"`
	Subscriptions        []string          `protobuf:"bytes,2,rep,name=subscriptions,proto3" json:"subscriptions,omitempty"`
	Path                 string            `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Headers              map[string]string `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Env                  *Env              `protobuf:"bytes,5,opt,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DisconnectRequest) Reset()         { *m = DisconnectRequest{} }
func (m *DisconnectRequest) String() string { return proto.CompactTextString(m) }
func (*DisconnectRequest) ProtoMessage()    {}
func (*DisconnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *DisconnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectRequest.Unmarshal(m, b)
}
func (m *DisconnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectRequest.Marshal(b, m, deterministic)
}
func (m *DisconnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectRequest.Merge(m, src)
}
func (m *DisconnectRequest) XXX_Size() int {
	return xxx_messageInfo_DisconnectRequest.Size(m)
}
func (m *DisconnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectRequest proto.InternalMessageInfo

func (m *DisconnectRequest) GetIdentifiers() string {
	if m != nil {
		return m.Identifiers
	}
	return ""
}

func (m *DisconnectRequest) GetSubscriptions() []string {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

func (m *DisconnectRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *DisconnectRequest) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DisconnectRequest) GetEnv() *Env {
	if m != nil {
		return m.Env
	}
	return nil
}

type DisconnectResponse struct {
	Status               Status   `protobuf:"varint,1,opt,name=status,proto3,enum=anycable.Status" json:"status,omitempty"`
	ErrorMsg             string   `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DisconnectResponse) Reset()         { *m = DisconnectResponse{} }
func (m *DisconnectResponse) String() string { return proto.CompactTextString(m) }
func (*DisconnectResponse) ProtoMessage()    {}
func (*DisconnectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *DisconnectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DisconnectResponse.Unmarshal(m, b)
}
func (m *DisconnectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DisconnectResponse.Marshal(b, m, deterministic)
}
func (m *DisconnectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisconnectResponse.Merge(m, src)
}
func (m *DisconnectResponse) XXX_Size() int {
	return xxx_messageInfo_DisconnectResponse.Size(m)
}
func (m *DisconnectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DisconnectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DisconnectResponse proto.InternalMessageInfo

func (m *DisconnectResponse) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ERROR
}

func (m *DisconnectResponse) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("anycable.Status", Status_name, Status_value)
	proto.RegisterType((*Env)(nil), "anycable.Env")
	proto.RegisterMapType((map[string]string)(nil), "anycable.Env.HeadersEntry")
	proto.RegisterType((*ConnectionRequest)(nil), "anycable.ConnectionRequest")
	proto.RegisterMapType((map[string]string)(nil), "anycable.ConnectionRequest.HeadersEntry")
	proto.RegisterType((*ConnectionResponse)(nil), "anycable.ConnectionResponse")
	proto.RegisterType((*CommandMessage)(nil), "anycable.CommandMessage")
	proto.RegisterType((*CommandResponse)(nil), "anycable.CommandResponse")
	proto.RegisterType((*DisconnectRequest)(nil), "anycable.DisconnectRequest")
	proto.RegisterMapType((map[string]string)(nil), "anycable.DisconnectRequest.HeadersEntry")
	proto.RegisterType((*DisconnectResponse)(nil), "anycable.DisconnectResponse")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xee, 0xda, 0xf9, 0x3b, 0x69, 0xfb, 0x4b, 0x47, 0x3f, 0x90, 0x49, 0x51, 0x09, 0x16, 0x07,
	0x0b, 0x89, 0x1c, 0x02, 0x48, 0xa8, 0x27, 0x20, 0xb8, 0xa2, 0x12, 0x15, 0x68, 0xa3, 0x9e, 0x38,
	0x54, 0x1b, 0x67, 0x49, 0x2d, 0x9a, 0xb5, 0xf1, 0x6e, 0x22, 0xe5, 0x19, 0xb8, 0xf0, 0x0e, 0xbc,
	0x03, 0xcf, 0xc0, 0x33, 0xc0, 0xcb, 0x20, 0xaf, 0xed, 0x64, 0x13, 0x87, 0x22, 0x40, 0xdc, 0x76,
	0xbf, 0x99, 0xcc, 0x7c, 0xdf, 0xcc, 0x7e, 0x31, 0x34, 0x93, 0x38, 0xe8, 0xc5, 0x49, 0xa4, 0x22,
	0x6c, 0x30, 0xb1, 0x08, 0xd8, 0xe8, 0x8a, 0xbb, 0x1f, 0x09, 0xd8, 0xbe, 0x98, 0x23, 0x42, 0x25,
	0x66, 0xea, 0xd2, 0x21, 0x5d, 0xe2, 0x35, 0xa9, 0x3e, 0xe3, 0x23, 0xa8, 0x5f, 0x72, 0x36, 0xe6,
	0x89, 0x74, 0xac, 0xae, 0xed, 0xb5, 0xfa, 0x9d, 0x5e, 0xf1, 0xbb, 0x9e, 0x2f, 0xe6, 0xbd, 0x97,
	0x59, 0xd0, 0x17, 0x2a, 0x59, 0xd0, 0x22, 0xb5, 0x73, 0x0c, 0xbb, 0x66, 0x00, 0xdb, 0x60, 0xbf,
	0xe7, 0x8b, 0xbc, 0x70, 0x7a, 0xc4, 0xff, 0xa1, 0x3a, 0x67, 0x57, 0x33, 0xee, 0x58, 0x1a, 0xcb,
	0x2e, 0xc7, 0xd6, 0x13, 0xe2, 0x7e, 0x25, 0x70, 0x30, 0x88, 0x84, 0xe0, 0x81, 0x0a, 0x23, 0x41,
	0xf9, 0x87, 0x19, 0x97, 0x6a, 0x2b, 0xb7, 0xe7, 0x9b, 0xdc, 0xbc, 0x15, 0xb7, 0x52, 0x85, 0xed,
	0x4c, 0xf1, 0x0e, 0xd8, 0x5c, 0xcc, 0x1d, 0xbb, 0x4b, 0xbc, 0x56, 0x7f, 0x6f, 0x4d, 0x1b, 0x4d,
	0x23, 0x7f, 0x25, 0xe5, 0x33, 0x01, 0x34, 0x89, 0xc8, 0x38, 0x12, 0x92, 0xa3, 0x07, 0x35, 0xa9,
	0x98, 0x9a, 0x49, 0x5d, 0x65, 0xbf, 0xdf, 0x5e, 0xb5, 0x1d, 0x6a, 0x9c, 0xe6, 0x71, 0xec, 0x42,
	0x2b, 0x1c, 0x73, 0xa1, 0xc2, 0x77, 0x61, 0xa6, 0x32, 0x6d, 0x60, 0x42, 0x78, 0x0f, 0xf6, 0x54,
	0xc2, 0x84, 0x9c, 0x86, 0x52, 0x86, 0x91, 0x90, 0x8e, 0xdd, 0xb5, 0xbd, 0x26, 0x5d, 0x07, 0xf1,
	0x10, 0x9a, 0x3c, 0x49, 0xa2, 0xe4, 0x62, 0x2a, 0x27, 0x4e, 0x45, 0x57, 0x69, 0x68, 0xe0, 0x4c,
	0x4e, 0xdc, 0x2f, 0x04, 0xf6, 0x07, 0xd1, 0x74, 0xca, 0xc4, 0xf8, 0x8c, 0x4b, 0xc9, 0x26, 0x1c,
	0x1d, 0xa8, 0x07, 0x19, 0x92, 0x0b, 0x2d, 0xae, 0x78, 0x04, 0xb0, 0x6a, 0x9f, 0x13, 0x32, 0x10,
	0x7c, 0x0c, 0x37, 0x83, 0xa5, 0xe2, 0x0b, 0x93, 0xbc, 0xad, 0x73, 0x6f, 0xac, 0xa2, 0xa7, 0x86,
	0x0c, 0x84, 0xca, 0x98, 0x29, 0x96, 0x73, 0xd3, 0xe7, 0x62, 0x35, 0xd5, 0x9f, 0xad, 0xc6, 0xfd,
	0x4e, 0xe0, 0xbf, 0x9c, 0xf8, 0x1f, 0xcc, 0xf6, 0x08, 0x60, 0x1c, 0xca, 0x9c, 0x8e, 0x56, 0xd2,
	0xa0, 0x06, 0x82, 0x77, 0x61, 0x57, 0xaa, 0x28, 0xbe, 0x90, 0x2a, 0xe1, 0x6c, 0x9a, 0xf1, 0x6f,
	0xd0, 0x56, 0x8a, 0x0d, 0x33, 0x28, 0x1d, 0x53, 0x11, 0xad, 0xe8, 0xb1, 0x17, 0xd7, 0xf2, 0x5a,
	0xaa, 0xbf, 0x5c, 0x4b, 0x6d, 0x63, 0x2d, 0x9f, 0x2c, 0x38, 0x78, 0xb1, 0xa4, 0x53, 0xf8, 0x60,
	0xe3, 0x45, 0x90, 0xad, 0x2f, 0x42, 0xce, 0x46, 0x32, 0x48, 0xc2, 0x58, 0xe9, 0xd6, 0x56, 0xd6,
	0x7a, 0x0d, 0x5c, 0xfa, 0xc9, 0xde, 0xee, 0xa7, 0xca, 0xa6, 0x9f, 0x4a, 0x4c, 0xae, 0xf7, 0x53,
	0xf5, 0x9f, 0xf8, 0xe9, 0x2d, 0xa0, 0xc9, 0xe3, 0xb7, 0x57, 0xbe, 0x36, 0x6f, 0x6b, 0x7d, 0xde,
	0xf7, 0x1f, 0x40, 0x2d, 0x4b, 0xc7, 0x26, 0x54, 0x7d, 0x4a, 0x5f, 0xd3, 0xf6, 0x0e, 0xb6, 0xa0,
	0x3e, 0x3c, 0x1f, 0x0c, 0xfc, 0xe1, 0xb0, 0x4d, 0xd2, 0xcb, 0xc9, 0xb3, 0xd3, 0x57, 0xe7, 0xd4,
	0x6f, 0x5b, 0xfd, 0x6f, 0x04, 0x6c, 0xfa, 0x66, 0x80, 0x27, 0x50, 0xcf, 0x2d, 0x8e, 0x87, 0xd7,
	0xfc, 0xfd, 0x74, 0x6e, 0x6f, 0x0f, 0x66, 0x1a, 0xdc, 0x1d, 0x7c, 0x9a, 0xd6, 0xc9, 0x3c, 0xe6,
	0x98, 0xa9, 0xa6, 0x2f, 0x3b, 0xb7, 0x4a, 0x11, 0xa3, 0xc2, 0x29, 0xc0, 0x6a, 0x3a, 0x26, 0x99,
	0xd2, 0xee, 0x4c, 0x32, 0xe5, 0x81, 0xba, 0x3b, 0xa3, 0x9a, 0xfe, 0x44, 0x3c, 0xfc, 0x11, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0xc4, 0xf9, 0x39, 0x2f, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RPCClient interface {
	Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error)
	Command(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandResponse, error)
	Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error)
}

type rPCClient struct {
	cc *grpc.ClientConn
}

func NewRPCClient(cc *grpc.ClientConn) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) Connect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionResponse, error) {
	out := new(ConnectionResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Command(ctx context.Context, in *CommandMessage, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Disconnect(ctx context.Context, in *DisconnectRequest, opts ...grpc.CallOption) (*DisconnectResponse, error) {
	out := new(DisconnectResponse)
	err := c.cc.Invoke(ctx, "/anycable.RPC/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
type RPCServer interface {
	Connect(context.Context, *ConnectionRequest) (*ConnectionResponse, error)
	Command(context.Context, *CommandMessage) (*CommandResponse, error)
	Disconnect(context.Context, *DisconnectRequest) (*DisconnectResponse, error)
}

// UnimplementedRPCServer can be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (*UnimplementedRPCServer) Connect(ctx context.Context, req *ConnectionRequest) (*ConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (*UnimplementedRPCServer) Command(ctx context.Context, req *CommandMessage) (*CommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Command not implemented")
}
func (*UnimplementedRPCServer) Disconnect(ctx context.Context, req *DisconnectRequest) (*DisconnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}

func RegisterRPCServer(s *grpc.Server, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
}

func _RPC_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Connect(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Command(ctx, req.(*CommandMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisconnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anycable.RPC/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).Disconnect(ctx, req.(*DisconnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "anycable.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _RPC_Connect_Handler,
		},
		{
			MethodName: "Command",
			Handler:    _RPC_Command_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _RPC_Disconnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}
