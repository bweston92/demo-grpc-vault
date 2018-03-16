// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	SayHelloToNameRequest
	SayHelloToNameResponse
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SayHelloToNameRequest struct {
	FullName string `protobuf:"bytes,1,opt,name=fullName" json:"fullName,omitempty"`
}

func (m *SayHelloToNameRequest) Reset()                    { *m = SayHelloToNameRequest{} }
func (m *SayHelloToNameRequest) String() string            { return proto.CompactTextString(m) }
func (*SayHelloToNameRequest) ProtoMessage()               {}
func (*SayHelloToNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SayHelloToNameRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type SayHelloToNameResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
}

func (m *SayHelloToNameResponse) Reset()                    { *m = SayHelloToNameResponse{} }
func (m *SayHelloToNameResponse) String() string            { return proto.CompactTextString(m) }
func (*SayHelloToNameResponse) ProtoMessage()               {}
func (*SayHelloToNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SayHelloToNameResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloToNameRequest)(nil), "SayHelloToNameRequest")
	proto.RegisterType((*SayHelloToNameResponse)(nil), "SayHelloToNameResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hello service

type HelloClient interface {
	SayHelloToName(ctx context.Context, in *SayHelloToNameRequest, opts ...grpc.CallOption) (*SayHelloToNameResponse, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) SayHelloToName(ctx context.Context, in *SayHelloToNameRequest, opts ...grpc.CallOption) (*SayHelloToNameResponse, error) {
	out := new(SayHelloToNameResponse)
	err := grpc.Invoke(ctx, "/Hello/SayHelloToName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloServer interface {
	SayHelloToName(context.Context, *SayHelloToNameRequest) (*SayHelloToNameResponse, error)
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_SayHelloToName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloToNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).SayHelloToName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello/SayHelloToName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).SayHelloToName(ctx, req.(*SayHelloToNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHelloToName",
			Handler:    _Hello_SayHelloToName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe6, 0x12, 0x0d, 0x4e, 0xac, 0xf4, 0x00,
	0x89, 0x84, 0xe4, 0xfb, 0x25, 0xe6, 0xa6, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49,
	0x71, 0x71, 0xa4, 0x95, 0xe6, 0xe4, 0x80, 0x84, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0,
	0x7c, 0x25, 0x1d, 0x2e, 0x31, 0x74, 0x4d, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x42, 0x5c,
	0x2c, 0x25, 0xa9, 0x15, 0x25, 0x50, 0x1d, 0x60, 0xb6, 0x91, 0x0f, 0x17, 0x2b, 0x58, 0xa9, 0x90,
	0x33, 0x17, 0x1f, 0xaa, 0x36, 0x21, 0x31, 0x3d, 0xac, 0x96, 0x4b, 0x89, 0xeb, 0x61, 0x37, 0x5f,
	0x89, 0x21, 0x89, 0x0d, 0xec, 0x6e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x22, 0x42,
	0xbe, 0xc6, 0x00, 0x00, 0x00,
}