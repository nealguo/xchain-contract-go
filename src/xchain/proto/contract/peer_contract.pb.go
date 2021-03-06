// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer_contract.proto

package contract

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PeerContractReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Payload              string   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Is_Original          bool     `protobuf:"varint,4,opt,name=is_Original,json=isOriginal,proto3" json:"is_Original,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerContractReply) Reset()         { *m = PeerContractReply{} }
func (m *PeerContractReply) String() string { return proto.CompactTextString(m) }
func (*PeerContractReply) ProtoMessage()    {}
func (*PeerContractReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd926e208b226158, []int{0}
}

func (m *PeerContractReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerContractReply.Unmarshal(m, b)
}
func (m *PeerContractReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerContractReply.Marshal(b, m, deterministic)
}
func (m *PeerContractReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerContractReply.Merge(m, src)
}
func (m *PeerContractReply) XXX_Size() int {
	return xxx_messageInfo_PeerContractReply.Size(m)
}
func (m *PeerContractReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerContractReply.DiscardUnknown(m)
}

var xxx_messageInfo_PeerContractReply proto.InternalMessageInfo

func (m *PeerContractReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PeerContractReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PeerContractReply) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PeerContractReply) GetIs_Original() bool {
	if m != nil {
		return m.Is_Original
	}
	return false
}

type PeerContractRequest struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Payload              string   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	ChannelName          string   `protobuf:"bytes,3,opt,name=channel_name,json=channelName,proto3" json:"channel_name,omitempty"`
	Sign                 string   `protobuf:"bytes,4,opt,name=sign,proto3" json:"sign,omitempty"`
	NodeId               string   `protobuf:"bytes,5,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PeerIp               string   `protobuf:"bytes,6,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PeerContractRequest) Reset()         { *m = PeerContractRequest{} }
func (m *PeerContractRequest) String() string { return proto.CompactTextString(m) }
func (*PeerContractRequest) ProtoMessage()    {}
func (*PeerContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd926e208b226158, []int{1}
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

func (m *PeerContractRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *PeerContractRequest) GetPeerIp() string {
	if m != nil {
		return m.PeerIp
	}
	return ""
}

func init() {
	proto.RegisterType((*PeerContractReply)(nil), "contract.PeerContractReply")
	proto.RegisterType((*PeerContractRequest)(nil), "contract.PeerContractRequest")
}

func init() { proto.RegisterFile("peer_contract.proto", fileDescriptor_cd926e208b226158) }

var fileDescriptor_cd926e208b226158 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x6d, 0xd3, 0x76, 0xea, 0xc5, 0x2d, 0xe8, 0xa2, 0xa8, 0xb5, 0xa7, 0x9e, 0x72,
	0x50, 0xfc, 0x03, 0x15, 0x84, 0x5e, 0xb4, 0xc4, 0xb3, 0x84, 0xe9, 0x66, 0x48, 0x17, 0x93, 0xdd,
	0x75, 0x37, 0x16, 0x02, 0xfe, 0x28, 0x7f, 0xa2, 0x64, 0x9b, 0x14, 0x03, 0x7a, 0xdb, 0x6f, 0x1e,
	0xcb, 0x7b, 0x33, 0x0f, 0xa6, 0x86, 0xc8, 0x26, 0x42, 0xab, 0xd2, 0xa2, 0x28, 0x23, 0x63, 0x75,
	0xa9, 0xd9, 0xa8, 0xe5, 0xf9, 0x17, 0x9c, 0xae, 0x89, 0xec, 0x63, 0xc3, 0x31, 0x99, 0xbc, 0x62,
	0x0c, 0xfa, 0x42, 0xa7, 0xc4, 0x83, 0x59, 0xb0, 0x18, 0xc4, 0xfe, 0xcd, 0x38, 0x0c, 0x0b, 0x72,
	0x0e, 0x33, 0xe2, 0xbd, 0x59, 0xb0, 0x18, 0xc7, 0x2d, 0xd6, 0x8a, 0xc1, 0x2a, 0xd7, 0x98, 0xf2,
	0xe3, 0xbd, 0xd2, 0x20, 0xbb, 0x81, 0x89, 0x74, 0xc9, 0x8b, 0x95, 0x99, 0x54, 0x98, 0xf3, 0xfe,
	0x2c, 0x58, 0x8c, 0x62, 0x90, 0xae, 0x9d, 0xcc, 0xbf, 0x03, 0x98, 0x76, 0xed, 0x3f, 0x3e, 0xc9,
	0x95, 0xec, 0x0c, 0xc2, 0x82, 0xca, 0xad, 0x4e, 0x7d, 0x84, 0x71, 0xdc, 0xd0, 0x6f, 0xab, 0x5e,
	0xd7, 0xea, 0x16, 0x4e, 0xc4, 0x16, 0x95, 0xa2, 0x3c, 0x51, 0x58, 0x50, 0x93, 0x64, 0xd2, 0xcc,
	0x9e, 0xb1, 0xa0, 0x7a, 0x2b, 0x27, 0x33, 0xe5, 0x63, 0x8c, 0x63, 0xff, 0x66, 0xe7, 0x30, 0x54,
	0x3a, 0xa5, 0x44, 0xa6, 0x7c, 0xb0, 0x77, 0xaa, 0x71, 0x95, 0xd6, 0x82, 0x3f, 0x9c, 0x34, 0x3c,
	0xdc, 0x0b, 0x35, 0xae, 0xcc, 0xdd, 0x5b, 0x37, 0xf1, 0x2b, 0xd9, 0x9d, 0x14, 0xc4, 0x9e, 0x20,
	0x94, 0x6a, 0xa7, 0xdf, 0x89, 0x5d, 0x45, 0x87, 0x63, 0xff, 0xb1, 0xda, 0xc5, 0xe5, 0x7f, 0xb2,
	0xc9, 0xab, 0xf9, 0xd1, 0xf2, 0x01, 0xae, 0x85, 0x8a, 0x0a, 0x6b, 0xb1, 0x8a, 0x2c, 0x56, 0x1b,
	0x44, 0x17, 0x65, 0xd6, 0x88, 0xc3, 0xaf, 0x65, 0xa7, 0xaf, 0x75, 0x5d, 0xe7, 0x3a, 0xd8, 0x84,
	0xbe, 0xd7, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x12, 0x7f, 0xe2, 0x0c, 0xee, 0x01, 0x00,
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
	Invoke(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*PeerContractReply, error)
}

type peerContractServiceClient struct {
	cc *grpc.ClientConn
}

func NewPeerContractServiceClient(cc *grpc.ClientConn) PeerContractServiceClient {
	return &peerContractServiceClient{cc}
}

func (c *peerContractServiceClient) Invoke(ctx context.Context, in *PeerContractRequest, opts ...grpc.CallOption) (*PeerContractReply, error) {
	out := new(PeerContractReply)
	err := c.cc.Invoke(ctx, "/contract.PeerContractService/invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerContractServiceServer is the server API for PeerContractService service.
type PeerContractServiceServer interface {
	Invoke(context.Context, *PeerContractRequest) (*PeerContractReply, error)
}

// UnimplementedPeerContractServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPeerContractServiceServer struct {
}

func (*UnimplementedPeerContractServiceServer) Invoke(ctx context.Context, req *PeerContractRequest) (*PeerContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
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
		FullMethod: "/contract.PeerContractService/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerContractServiceServer).Invoke(ctx, req.(*PeerContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PeerContractService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "contract.PeerContractService",
	HandlerType: (*PeerContractServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "invoke",
			Handler:    _PeerContractService_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer_contract.proto",
}
