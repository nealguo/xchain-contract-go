// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/contract_peer.proto

package peer

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
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

type ContractPeerReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContractPeerReply) Reset()         { *m = ContractPeerReply{} }
func (m *ContractPeerReply) String() string { return proto.CompactTextString(m) }
func (*ContractPeerReply) ProtoMessage()    {}
func (*ContractPeerReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dcae9e62b44e6e0, []int{0}
}

func (m *ContractPeerReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractPeerReply.Unmarshal(m, b)
}
func (m *ContractPeerReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractPeerReply.Marshal(b, m, deterministic)
}
func (m *ContractPeerReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractPeerReply.Merge(m, src)
}
func (m *ContractPeerReply) XXX_Size() int {
	return xxx_messageInfo_ContractPeerReply.Size(m)
}
func (m *ContractPeerReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractPeerReply.DiscardUnknown(m)
}

var xxx_messageInfo_ContractPeerReply proto.InternalMessageInfo

func (m *ContractPeerReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ContractPeerReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ContractPeerReply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

type ContractPeerRequest struct {
	ContractId           *ContractID `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	Method               string      `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Key                  string      `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Payload              string      `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Sign                 string      `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	NodeId               string      `protobuf:"bytes,6,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ChannelName          string      `protobuf:"bytes,7,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ContractPeerRequest) Reset()         { *m = ContractPeerRequest{} }
func (m *ContractPeerRequest) String() string { return proto.CompactTextString(m) }
func (*ContractPeerRequest) ProtoMessage()    {}
func (*ContractPeerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2dcae9e62b44e6e0, []int{1}
}

func (m *ContractPeerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractPeerRequest.Unmarshal(m, b)
}
func (m *ContractPeerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractPeerRequest.Marshal(b, m, deterministic)
}
func (m *ContractPeerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractPeerRequest.Merge(m, src)
}
func (m *ContractPeerRequest) XXX_Size() int {
	return xxx_messageInfo_ContractPeerRequest.Size(m)
}
func (m *ContractPeerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractPeerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContractPeerRequest proto.InternalMessageInfo

func (m *ContractPeerRequest) GetContractId() *ContractID {
	if m != nil {
		return m.ContractId
	}
	return nil
}

func (m *ContractPeerRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ContractPeerRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ContractPeerRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *ContractPeerRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *ContractPeerRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *ContractPeerRequest) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func init() {
	proto.RegisterType((*ContractPeerReply)(nil), "peer.ContractPeerReply")
	proto.RegisterType((*ContractPeerRequest)(nil), "peer.ContractPeerRequest")
}

func init() { proto.RegisterFile("peer/contract_peer.proto", fileDescriptor_2dcae9e62b44e6e0) }

var fileDescriptor_2dcae9e62b44e6e0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4e, 0x02, 0x31,
	0x10, 0xc6, 0x5d, 0xf9, 0xa7, 0x83, 0x07, 0xac, 0x51, 0x2a, 0x27, 0xe4, 0xc4, 0x69, 0x0d, 0xf8,
	0x02, 0x06, 0xbc, 0x70, 0x31, 0x64, 0x39, 0x78, 0xf0, 0x40, 0x86, 0x76, 0xb2, 0x10, 0x77, 0xdb,
	0xb5, 0x5b, 0x4d, 0xf6, 0x5d, 0x7d, 0x18, 0xd3, 0xb2, 0x6e, 0x68, 0xe2, 0xed, 0xfb, 0xe6, 0x4b,
	0x67, 0x7e, 0xd3, 0x01, 0x5e, 0x10, 0x99, 0x47, 0xa1, 0x95, 0x35, 0x28, 0xec, 0xd6, 0xb9, 0xb8,
	0x30, 0xda, 0x6a, 0xd6, 0x76, 0x7a, 0x74, 0x1b, 0xe4, 0xab, 0x97, 0x63, 0x38, 0x79, 0x87, 0xeb,
	0x65, 0x5d, 0x5b, 0x13, 0x99, 0x84, 0x8a, 0xac, 0x62, 0x0c, 0xda, 0x42, 0x4b, 0xe2, 0xd1, 0x38,
	0x9a, 0x76, 0x12, 0xaf, 0x19, 0x87, 0x5e, 0x4e, 0x65, 0x89, 0x29, 0xf1, 0xf3, 0x71, 0x34, 0xbd,
	0x4c, 0xfe, 0xac, 0x4b, 0x0a, 0xac, 0x32, 0x8d, 0x92, 0xb7, 0x8e, 0x49, 0x6d, 0x27, 0x3f, 0x11,
	0xdc, 0x84, 0xdd, 0x3f, 0xbf, 0xa8, 0xb4, 0x6c, 0x06, 0xfd, 0x06, 0xf4, 0x20, 0xfd, 0x98, 0xfe,
	0x7c, 0x10, 0x7b, 0xe6, 0x65, 0x43, 0x98, 0x40, 0x43, 0x2b, 0xd9, 0x1d, 0x74, 0x73, 0xb2, 0x7b,
	0x2d, 0xeb, 0xe9, 0xb5, 0x63, 0x03, 0x68, 0x7d, 0x50, 0x55, 0x0f, 0x76, 0xf2, 0x14, 0xa7, 0x1d,
	0xe0, 0xb8, 0xb5, 0xca, 0x43, 0xaa, 0x78, 0xc7, 0x97, 0xbd, 0x66, 0x43, 0xe8, 0x29, 0x2d, 0xc9,
	0x61, 0x74, 0x8f, 0x8d, 0x9d, 0x5d, 0x49, 0xf6, 0x00, 0x57, 0x62, 0x8f, 0x4a, 0x51, 0xb6, 0x55,
	0x98, 0x13, 0xef, 0xf9, 0xb4, 0x5f, 0xd7, 0x5e, 0x31, 0xa7, 0xf9, 0x5b, 0xb8, 0xdd, 0x86, 0xcc,
	0xf7, 0x41, 0x10, 0x7b, 0x86, 0x8b, 0x94, 0xec, 0xc6, 0xa2, 0x25, 0x76, 0x1f, 0x2e, 0x75, 0xf2,
	0x09, 0xa3, 0xe1, 0x7f, 0x51, 0x91, 0x55, 0x93, 0xb3, 0xc5, 0x0c, 0x46, 0x42, 0xc5, 0xb9, 0x31,
	0x58, 0xc5, 0x06, 0xab, 0x1d, 0x62, 0x19, 0xa7, 0xa6, 0x10, 0xfe, 0xc5, 0x22, 0x38, 0xd8, 0xda,
	0x5d, 0x71, 0x1d, 0xed, 0xba, 0xfe, 0x9c, 0x4f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xee,
	0xbc, 0x19, 0x07, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContractPeerServiceClient is the client API for ContractPeerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContractPeerServiceClient interface {
	GetState(ctx context.Context, in *ContractPeerRequest, opts ...grpc.CallOption) (*ContractPeerReply, error)
}

type contractPeerServiceClient struct {
	cc *grpc.ClientConn
}

func NewContractPeerServiceClient(cc *grpc.ClientConn) ContractPeerServiceClient {
	return &contractPeerServiceClient{cc}
}

func (c *contractPeerServiceClient) GetState(ctx context.Context, in *ContractPeerRequest, opts ...grpc.CallOption) (*ContractPeerReply, error) {
	out := new(ContractPeerReply)
	err := c.cc.Invoke(ctx, "/peer.ContractPeerService/getState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractPeerServiceServer is the server API for ContractPeerService service.
type ContractPeerServiceServer interface {
	GetState(context.Context, *ContractPeerRequest) (*ContractPeerReply, error)
}

// UnimplementedContractPeerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedContractPeerServiceServer struct {
}

func (*UnimplementedContractPeerServiceServer) GetState(ctx context.Context, req *ContractPeerRequest) (*ContractPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetState not implemented")
}

func RegisterContractPeerServiceServer(s *grpc.Server, srv ContractPeerServiceServer) {
	s.RegisterService(&_ContractPeerService_serviceDesc, srv)
}

func _ContractPeerService_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContractPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractPeerServiceServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/peer.ContractPeerService/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractPeerServiceServer).GetState(ctx, req.(*ContractPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContractPeerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "peer.ContractPeerService",
	HandlerType: (*ContractPeerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getState",
			Handler:    _ContractPeerService_GetState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer/contract_peer.proto",
}
