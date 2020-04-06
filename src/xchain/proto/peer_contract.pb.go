// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer_contract.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
//const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PeerContractRequest struct {
	Method               string        `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Payload              string        `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ChannelName          string        `protobuf:"bytes,3,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	Sign                 string        `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
	AppId                string        `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	PeerIp               string        `protobuf:"bytes,6,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	Caller               []*ContractID `protobuf:"bytes,7,rep,name=caller,proto3" json:"caller,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PeerContractRequest) Reset()         { *m = PeerContractRequest{} }
func (m *PeerContractRequest) String() string { return proto.CompactTextString(m) }
func (*PeerContractRequest) ProtoMessage()    {}
func (*PeerContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd926e208b226158, []int{0}
}

func (m *PeerContractRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerContractRequest.Unmarshal(m, b)
}
func (m *PeerContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerContractRequest.Marshal(b, m, deterministic)
}
func (m *PeerContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerContractRequest.Merge(m, src)
}
func (m *PeerContractRequest) XXX_Size() int {
	return xxx_messageInfo_PeerContractRequest.Size(m)
}
func (m *PeerContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PeerContractRequest proto.InternalMessageInfo

func (m *PeerContractRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *PeerContractRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PeerContractRequest) GetChannelName() string {
	if m != nil {
		return m.ChannelName
	}
	return ""
}

func (m *PeerContractRequest) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *PeerContractRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *PeerContractRequest) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func (m *PeerContractRequest) GetCaller() []*ContractID {
	if m != nil {
		return m.Caller
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerContractRequest)(nil), "api.PeerContractRequest")
}

func init() { proto.RegisterFile("peer_contract.proto", fileDescriptor_cd926e208b226158) }

var fileDescriptor_cd926e208b226158 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x29, 0x6d, 0x1d, 0xe1, 0x16, 0x51, 0x5d, 0x05, 0x58, 0x99, 0x4a, 0x17, 0x3a, 0x65,
	0x68, 0x27, 0xe6, 0xb2, 0x64, 0x41, 0x28, 0x3c, 0x40, 0x74, 0x38, 0x27, 0x6a, 0x91, 0xd8, 0x87,
	0x6b, 0x2a, 0xf5, 0x01, 0x78, 0x46, 0x5e, 0x07, 0xe1, 0x24, 0x0c, 0x88, 0x85, 0xcd, 0xff, 0xff,
	0xdd, 0x49, 0xfe, 0x4e, 0xce, 0x99, 0xc8, 0x97, 0xda, 0xd9, 0xe0, 0x51, 0x87, 0x8c, 0xbd, 0x0b,
	0x0e, 0x86, 0xc8, 0x26, 0x9d, 0x79, 0xd6, 0xa5, 0x76, 0x4d, 0xe3, 0x6c, 0x5b, 0xa7, 0xb3, 0x7e,
	0x2c, 0xbf, 0x6f, 0x9b, 0xe5, 0xe7, 0x40, 0xce, 0x1f, 0x89, 0xfc, 0xb6, 0x03, 0x05, 0xbd, 0xbd,
	0xd3, 0x3e, 0xc0, 0x95, 0x14, 0x0d, 0x85, 0x9d, 0xab, 0xd4, 0x60, 0x31, 0x58, 0x9d, 0x15, 0x5d,
	0x02, 0x25, 0x13, 0xc6, 0x63, 0xed, 0xb0, 0x52, 0xa7, 0x11, 0xf4, 0x11, 0x6e, 0xe4, 0x54, 0xef,
	0xd0, 0x5a, 0xaa, 0x4b, 0x8b, 0x0d, 0xa9, 0x61, 0xc4, 0x93, 0xae, 0x7b, 0xc0, 0x86, 0x00, 0xe4,
	0x68, 0x6f, 0x5e, 0xac, 0x1a, 0x45, 0x14, 0xdf, 0x70, 0x29, 0x05, 0x32, 0x97, 0xa6, 0x52, 0xe3,
	0xd8, 0x8e, 0x91, 0x39, 0xaf, 0xe0, 0x5a, 0x26, 0xd1, 0xcb, 0xb0, 0x12, 0xed, 0x07, 0xbe, 0x63,
	0xce, 0x70, 0x2b, 0x85, 0xc6, 0xba, 0x26, 0xaf, 0x92, 0xc5, 0x70, 0x35, 0x59, 0x5f, 0x64, 0xc8,
	0x26, 0xdb, 0xfe, 0x78, 0x15, 0x1d, 0x5e, 0x7f, 0xfc, 0x32, 0x7b, 0x22, 0x7f, 0x30, 0x9a, 0x60,
	0x23, 0x85, 0xb1, 0x07, 0xf7, 0x4a, 0xa0, 0xe2, 0xea, 0x1f, 0xf6, 0xe9, 0x79, 0x24, 0x05, 0xeb,
	0x82, 0xb8, 0x3e, 0x2e, 0x4f, 0xe0, 0x4e, 0x4e, 0x8d, 0x35, 0xa1, 0x9f, 0xfb, 0xc7, 0xea, 0xb3,
	0x88, 0x87, 0xde, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xda, 0x4d, 0x11, 0xa8, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PeerContractServiceClient is the client API for PeerContractService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PeerContractServiceClient interface {
	Invoke(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*RpcReply, error)
	InitContract(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*RpcReply, error)
}

type peerContractServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerContractServiceClient(cc *grpc.ClientConn) PeerContractServiceClient {
	return &peerContractServiceClient{cc}
}

func (c *peerContractServiceClient) Invoke(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, "/api.PeerContractService/invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerContractServiceClient) InitContract(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*RpcReply, error) {
	out := new(RpcReply)
	err := c.cc.Invoke(ctx, "/api.PeerContractService/initContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerContractServiceServer is the server API for PeerContractService service.
type PeerContractServiceServer interface {
	Invoke(context.Context, *PeerContractRequest) (*RpcReply, error)
	InitContract(context.Context, *PeerContractRequest) (*RpcReply, error)
}

// UnimplementedPeerContractServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPeerContractServiceServer struct {
}

func (*UnimplementedPeerContractServiceServer) Invoke(ctx context.Context, req *PeerContractRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedPeerContractServiceServer) InitContract(ctx context.Context, req *PeerContractRequest) (*RpcReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitContract not implemented")
}

func RegisterPeerContractServiceServer(s *grpc.Server, srv PeerContractServiceServer) {
	s.RegisterService(&_PeerContractService_serviceDesc, srv)
}

func _PeerContractService_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerContractServiceServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PeerContractService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerContractServiceServer).Invoke(ctx, req.(*PeerContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeerContractService_InitContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerContractServiceServer).InitContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PeerContractService/InitContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerContractServiceServer).InitContract(ctx, req.(*PeerContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerContractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PeerContractService",
	HandlerType: (*PeerContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "invoke",
			Handler:    _PeerContractService_Invoke_Handler,
		},
		{
			MethodName: "initContract",
			Handler:    _PeerContractService_InitContract_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer_contract.proto",
}
