// Code generated by protoc-gen-go.
// source: proto/geo/geo.proto
// DO NOT EDIT!

/*
Package geo is a generated protocol buffer package.

It is generated from these files:
	proto/geo/geo.proto

It has these top-level messages:
	Request
	Point
	Result
*/
package geo

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

// A latitude-longitude bounding box, represented as two diagonally opposite
// points "lo" and "hi".
type Request struct {
	Lo *Point `protobuf:"bytes,2,opt,name=lo" json:"lo,omitempty"`
	Hi *Point `protobuf:"bytes,3,opt,name=hi" json:"hi,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetLo() *Point {
	if m != nil {
		return m.Lo
	}
	return nil
}

func (m *Request) GetHi() *Point {
	if m != nil {
		return m.Hi
	}
	return nil
}

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	Latitude  int32 `protobuf:"varint,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Result struct {
	HotelIds []int32 `protobuf:"varint,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Request)(nil), "geo.Request")
	proto.RegisterType((*Point)(nil), "geo.Point")
	proto.RegisterType((*Result)(nil), "geo.Result")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Geo service

type GeoClient interface {
	// Obtains the Locations contained within the given Rectangle.
	BoundedBox(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type geoClient struct {
	cc *grpc.ClientConn
}

func NewGeoClient(cc *grpc.ClientConn) GeoClient {
	return &geoClient{cc}
}

func (c *geoClient) BoundedBox(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/geo.Geo/BoundedBox", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Geo service

type GeoServer interface {
	// Obtains the Locations contained within the given Rectangle.
	BoundedBox(context.Context, *Request) (*Result, error)
}

func RegisterGeoServer(s *grpc.Server, srv GeoServer) {
	s.RegisterService(&_Geo_serviceDesc, srv)
}

func _Geo_BoundedBox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GeoServer).BoundedBox(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Geo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "geo.Geo",
	HandlerType: (*GeoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BoundedBox",
			Handler:    _Geo_BoundedBox_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4f, 0x05, 0x63, 0x3d, 0x30, 0x4f, 0x88, 0x19, 0xc8, 0x54, 0xb2, 0xe4, 0x62,
	0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0xca, 0xc9, 0x97, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x36, 0xe2, 0xd2, 0x03, 0xa9, 0x0b, 0xc8, 0xcf, 0xcc, 0x03, 0x8b, 0x67, 0x64,
	0x4a, 0x30, 0xa3, 0x8b, 0x2b, 0xe9, 0x70, 0xb1, 0x42, 0x14, 0x08, 0x70, 0x71, 0xe4, 0x24, 0x96,
	0x64, 0x96, 0x94, 0xa6, 0xa4, 0x4a, 0x30, 0x02, 0x95, 0xb1, 0x0a, 0x09, 0x72, 0x71, 0xe6, 0xe4,
	0xe7, 0xa5, 0x43, 0x84, 0x40, 0x26, 0xb2, 0x2a, 0x49, 0x71, 0xb1, 0x05, 0xa5, 0x16, 0x97, 0xe6,
	0x80, 0x95, 0x67, 0xe4, 0x97, 0xa4, 0xe6, 0x78, 0xa6, 0x14, 0x03, 0x95, 0x33, 0x6b, 0xb0, 0x1a,
	0xe9, 0x71, 0x31, 0xbb, 0xa7, 0xe6, 0x0b, 0xa9, 0x73, 0x71, 0x39, 0xe5, 0x97, 0xe6, 0xa5, 0xa4,
	0xa6, 0x38, 0xe5, 0x57, 0x08, 0xf1, 0x80, 0xad, 0x82, 0x3a, 0x4e, 0x8a, 0x1b, 0xca, 0x03, 0x99,
	0x90, 0xc4, 0x06, 0xf6, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x91, 0x1f, 0x2e, 0x26, 0xd7,
	0x00, 0x00, 0x00,
}
