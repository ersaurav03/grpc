// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Num1 int64 `protobuf:"varint,1,opt,name=num1" json:"num1,omitempty"`
	Num2 int64 `protobuf:"varint,2,opt,name=num2" json:"num2,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetNum1() int64 {
	if m != nil {
		return m.Num1
	}
	return 0
}

func (m *Request) GetNum2() int64 {
	if m != nil {
		return m.Num2
	}
	return 0
}

type Response struct {
	Result int64 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto1.RegisterType((*Request)(nil), "proto.Request")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AddService service

type AddServiceClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Mul(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Div(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Mul(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/Mul", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) Div(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.AddService/Div", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddService service

type AddServiceServer interface {
	Add(context.Context, *Request) (*Response, error)
	Mul(context.Context, *Request) (*Response, error)
	Div(context.Context, *Request) (*Response, error)
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Mul_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Mul(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Mul",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Mul(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_Div_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Div(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddService/Div",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Div(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
		{
			MethodName: "Mul",
			Handler:    _AddService_Mul_Handler,
		},
		{
			MethodName: "Div",
			Handler:    _AddService_Div_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x86, 0x5c,
	0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0xa5, 0xb9, 0x86,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x60, 0x36, 0x54, 0xcc, 0x48, 0x82, 0x09, 0x2e, 0x66,
	0xa4, 0xa4, 0xc4, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5,
	0x56, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0x02, 0xd5, 0x05, 0xe5, 0x19, 0x75, 0x31, 0x72, 0x71, 0x39,
	0xa6, 0xa4, 0x04, 0x43, 0xac, 0x14, 0xd2, 0xe0, 0x62, 0x76, 0x4c, 0x49, 0x11, 0xe2, 0x83, 0xd8,
	0xad, 0x07, 0xb5, 0x51, 0x8a, 0x1f, 0xce, 0x87, 0x18, 0xa7, 0xc4, 0x00, 0x52, 0xe9, 0x5b, 0x9a,
	0x43, 0xa4, 0x4a, 0x97, 0xcc, 0x32, 0x22, 0x54, 0x26, 0xb1, 0x81, 0x45, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x02, 0xdb, 0xa5, 0x02, 0x01, 0x00, 0x00,
}