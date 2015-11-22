// Code generated by protoc-gen-go.
// source: proto/auth/auth.proto
// DO NOT EDIT!

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	proto/auth/auth.proto

It has these top-level messages:
	Request
	Result
	Customer
*/
package auth

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

type Request struct {
	AuthToken string `protobuf:"bytes,1,opt,name=authToken" json:"authToken,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Customer *Customer `protobuf:"bytes,1,opt,name=customer" json:"customer,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type Customer struct {
	Id        int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	AuthToken string `protobuf:"bytes,2,opt,name=authToken" json:"authToken,omitempty"`
}

func (m *Customer) Reset()                    { *m = Customer{} }
func (m *Customer) String() string            { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()               {}
func (*Customer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Result)(nil), "auth.Result")
	proto.RegisterType((*Customer)(nil), "auth.Customer")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Auth service

type AuthClient interface {
	VerifyToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) VerifyToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/auth.Auth/VerifyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	VerifyToken(context.Context, *Request) (*Result, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthServer).VerifyToken(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _Auth_VerifyToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2c, 0x2d, 0xc9, 0x00, 0x13, 0x7a, 0x60, 0xbe, 0x10, 0x0b, 0x88, 0xad, 0x24,
	0xc3, 0xc5, 0x1e, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc8, 0xc5, 0x09, 0x12, 0x0a,
	0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x54, 0xd2, 0xe2, 0x62, 0x0b, 0x4a,
	0x2d, 0x2e, 0xcd, 0x29, 0x11, 0x52, 0xe0, 0xe2, 0x48, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2d,
	0x02, 0xcb, 0x71, 0x1b, 0xf1, 0xe9, 0x81, 0x0d, 0x73, 0x86, 0x8a, 0x2a, 0x69, 0x72, 0x71, 0xc0,
	0xd8, 0x42, 0x5c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0x75, 0xac, 0xa8, 0xc6, 0x32, 0x81, 0x8c, 0x35,
	0x32, 0xe1, 0x62, 0x71, 0x04, 0x0a, 0x09, 0xe9, 0x70, 0x71, 0x87, 0xa5, 0x16, 0x65, 0xa6, 0x55,
	0x82, 0x25, 0x85, 0x78, 0x21, 0x26, 0x42, 0xdd, 0x23, 0xc5, 0x03, 0xe3, 0x82, 0x1c, 0xa0, 0xc4,
	0x90, 0xc4, 0x06, 0x76, 0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xef, 0xe8, 0x2d, 0xd0,
	0x00, 0x00, 0x00,
}
