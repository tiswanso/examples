// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipreceiver.proto

/*
Package ipreceiver is a generated protocol buffer package.

It is generated from these files:
	ipreceiver.proto

It has these top-level messages:
	IpRange
	RangeIdentifier
*/
package ipreceiver

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	ippool "github.com/networkservicemesh/examples/api/ipam/ipprovider"

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

type IpRange struct {
	Identifier *RangeIdentifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Prefix     *ippool.IpPrefix `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
}

func (m *IpRange) Reset()                    { *m = IpRange{} }
func (m *IpRange) String() string            { return proto.CompactTextString(m) }
func (*IpRange) ProtoMessage()               {}
func (*IpRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IpRange) GetIdentifier() *RangeIdentifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *IpRange) GetPrefix() *ippool.IpPrefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

type RangeIdentifier struct {
	Fqdn               string `protobuf:"bytes,1,opt,name=fqdn" json:"fqdn,omitempty"`
	ConnectivityDomain string `protobuf:"bytes,3,opt,name=connectivity_domain,json=connectivityDomain" json:"connectivity_domain,omitempty"`
}

func (m *RangeIdentifier) Reset()                    { *m = RangeIdentifier{} }
func (m *RangeIdentifier) String() string            { return proto.CompactTextString(m) }
func (*RangeIdentifier) ProtoMessage()               {}
func (*RangeIdentifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RangeIdentifier) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *RangeIdentifier) GetConnectivityDomain() string {
	if m != nil {
		return m.ConnectivityDomain
	}
	return ""
}

func init() {
	proto.RegisterType((*IpRange)(nil), "ippool.IpRange")
	proto.RegisterType((*RangeIdentifier)(nil), "ippool.RangeIdentifier")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PrefixRangeAllocator service

type PrefixRangeAllocatorClient interface {
	AssignRange(ctx context.Context, in *IpRange, opts ...grpc.CallOption) (*ippool.Empty, error)
	UnassignRange(ctx context.Context, in *IpRange, opts ...grpc.CallOption) (*ippool.Empty, error)
}

type prefixRangeAllocatorClient struct {
	cc *grpc.ClientConn
}

func NewPrefixRangeAllocatorClient(cc *grpc.ClientConn) PrefixRangeAllocatorClient {
	return &prefixRangeAllocatorClient{cc}
}

func (c *prefixRangeAllocatorClient) AssignRange(ctx context.Context, in *IpRange, opts ...grpc.CallOption) (*ippool.Empty, error) {
	out := new(ippool.Empty)
	err := grpc.Invoke(ctx, "/ippool.PrefixRangeAllocator/AssignRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prefixRangeAllocatorClient) UnassignRange(ctx context.Context, in *IpRange, opts ...grpc.CallOption) (*ippool.Empty, error) {
	out := new(ippool.Empty)
	err := grpc.Invoke(ctx, "/ippool.PrefixRangeAllocator/UnassignRange", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PrefixRangeAllocator service

type PrefixRangeAllocatorServer interface {
	AssignRange(context.Context, *IpRange) (*ippool.Empty, error)
	UnassignRange(context.Context, *IpRange) (*ippool.Empty, error)
}

func RegisterPrefixRangeAllocatorServer(s *grpc.Server, srv PrefixRangeAllocatorServer) {
	s.RegisterService(&_PrefixRangeAllocator_serviceDesc, srv)
}

func _PrefixRangeAllocator_AssignRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixRangeAllocatorServer).AssignRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.PrefixRangeAllocator/AssignRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixRangeAllocatorServer).AssignRange(ctx, req.(*IpRange))
	}
	return interceptor(ctx, in, info, handler)
}

func _PrefixRangeAllocator_UnassignRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IpRange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrefixRangeAllocatorServer).UnassignRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.PrefixRangeAllocator/UnassignRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrefixRangeAllocatorServer).UnassignRange(ctx, req.(*IpRange))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrefixRangeAllocator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ippool.PrefixRangeAllocator",
	HandlerType: (*PrefixRangeAllocatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssignRange",
			Handler:    _PrefixRangeAllocator_AssignRange_Handler,
		},
		{
			MethodName: "UnassignRange",
			Handler:    _PrefixRangeAllocator_UnassignRange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipreceiver.proto",
}

func init() { proto.RegisterFile("ipreceiver.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0xbd, 0x4b, 0xc3, 0x40,
	0x14, 0xb7, 0x2a, 0x11, 0x5f, 0x2d, 0x2d, 0xa7, 0x60, 0xe8, 0x24, 0x99, 0xba, 0x98, 0x60, 0x1d,
	0x5c, 0x5c, 0x2a, 0x3a, 0x64, 0x93, 0x80, 0x0e, 0x2e, 0x72, 0xb9, 0x5c, 0xe2, 0x83, 0xe4, 0xdd,
	0xf3, 0x72, 0x06, 0xeb, 0x5f, 0x2f, 0x5e, 0x5a, 0x1b, 0xdc, 0xdc, 0x7e, 0xfc, 0x3e, 0xef, 0x03,
	0x66, 0xc8, 0x56, 0x2b, 0x8d, 0x9d, 0xb6, 0x31, 0x5b, 0xe3, 0x8c, 0x08, 0x90, 0xd9, 0x98, 0x7a,
	0x7e, 0x5b, 0xa1, 0x7b, 0xfb, 0xc8, 0x63, 0x65, 0x9a, 0x44, 0x11, 0xb5, 0x97, 0xd4, 0xda, 0x1e,
	0x20, 0xcb, 0x26, 0xe1, 0x3c, 0x41, 0x66, 0x6b, 0x3a, 0x2c, 0xb4, 0x1d, 0xc0, 0xbe, 0x25, 0xaa,
	0xe1, 0x28, 0xe5, 0x4c, 0x52, 0xa5, 0xc5, 0x0d, 0x00, 0x16, 0x9a, 0x1c, 0x96, 0xa8, 0x6d, 0x38,
	0xba, 0x18, 0x2d, 0xc6, 0xcb, 0xf3, 0xb8, 0x5f, 0x89, 0xbd, 0x25, 0xfd, 0x95, 0xb3, 0x81, 0x55,
	0x2c, 0x20, 0x60, 0xab, 0x4b, 0xfc, 0x0c, 0xf7, 0x7d, 0x68, 0xb6, 0x0d, 0xa5, 0xfc, 0xe8, 0xf9,
	0x6c, 0xa3, 0x47, 0xcf, 0x30, 0xfd, 0x53, 0x24, 0x04, 0x1c, 0x96, 0xef, 0x05, 0xf9, 0xbd, 0xe3,
	0xcc, 0x63, 0x91, 0xc0, 0xa9, 0x32, 0x44, 0x5a, 0x39, 0xec, 0xd0, 0xad, 0x5f, 0x0b, 0xd3, 0x48,
	0xa4, 0xf0, 0xc0, 0x5b, 0xc4, 0x50, 0xba, 0xf7, 0xca, 0xf2, 0x0b, 0xce, 0x36, 0x4b, 0x3f, 0xed,
	0xab, 0xba, 0x36, 0x4a, 0x3a, 0x63, 0x45, 0x02, 0xe3, 0x55, 0xdb, 0x62, 0x45, 0xfd, 0x0d, 0xa7,
	0xbb, 0x83, 0x79, 0x62, 0x3e, 0xd9, 0x12, 0x0f, 0x0d, 0xbb, 0x75, 0xb4, 0x27, 0xae, 0x60, 0xf2,
	0x44, 0xf2, 0x3f, 0x91, 0xbb, 0x93, 0x17, 0xd8, 0xfd, 0x4d, 0x1e, 0xf8, 0x67, 0xbd, 0xfe, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0x78, 0x2b, 0xd5, 0xb0, 0x01, 0x00, 0x00,
}
