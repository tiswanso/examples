// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ipprovider.proto

/*
Package ipprovider is a generated protocol buffer package.

It is generated from these files:
	ipprovider.proto

It has these top-level messages:
	IpFamily
	SubnetRequest
	Subnet
	IpPrefix
	Identifier
	Empty
*/
package ipprovider

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

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

type IpFamily_Family int32

const (
	IpFamily_IPV4 IpFamily_Family = 0
	IpFamily_IPV6 IpFamily_Family = 1
)

var IpFamily_Family_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}
var IpFamily_Family_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IpFamily_Family) String() string {
	return proto.EnumName(IpFamily_Family_name, int32(x))
}
func (IpFamily_Family) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type IpFamily struct {
	Family IpFamily_Family `protobuf:"varint,1,opt,name=family,enum=ippool.IpFamily_Family" json:"family,omitempty"`
}

func (m *IpFamily) Reset()                    { *m = IpFamily{} }
func (m *IpFamily) String() string            { return proto.CompactTextString(m) }
func (*IpFamily) ProtoMessage()               {}
func (*IpFamily) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IpFamily) GetFamily() IpFamily_Family {
	if m != nil {
		return m.Family
	}
	return IpFamily_IPV4
}

type SubnetRequest struct {
	Identifier *Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	AddrFamily *IpFamily   `protobuf:"bytes,2,opt,name=addr_family,json=addrFamily" json:"addr_family,omitempty"`
	PrefixLen  uint32      `protobuf:"varint,3,opt,name=prefix_len,json=prefixLen" json:"prefix_len,omitempty"`
}

func (m *SubnetRequest) Reset()                    { *m = SubnetRequest{} }
func (m *SubnetRequest) String() string            { return proto.CompactTextString(m) }
func (*SubnetRequest) ProtoMessage()               {}
func (*SubnetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubnetRequest) GetIdentifier() *Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *SubnetRequest) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *SubnetRequest) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

type Subnet struct {
	Identifier   *Identifier `protobuf:"bytes,1,opt,name=identifier" json:"identifier,omitempty"`
	Prefix       *IpPrefix   `protobuf:"bytes,2,opt,name=prefix" json:"prefix,omitempty"`
	LeaseTimeout int32       `protobuf:"varint,3,opt,name=lease_timeout,json=leaseTimeout" json:"lease_timeout,omitempty"`
}

func (m *Subnet) Reset()                    { *m = Subnet{} }
func (m *Subnet) String() string            { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()               {}
func (*Subnet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Subnet) GetIdentifier() *Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Subnet) GetPrefix() *IpPrefix {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Subnet) GetLeaseTimeout() int32 {
	if m != nil {
		return m.LeaseTimeout
	}
	return 0
}

type IpPrefix struct {
	AddrFamily *IpFamily `protobuf:"bytes,1,opt,name=addr_family,json=addrFamily" json:"addr_family,omitempty"`
	Subnet     string    `protobuf:"bytes,2,opt,name=subnet" json:"subnet,omitempty"`
}

func (m *IpPrefix) Reset()                    { *m = IpPrefix{} }
func (m *IpPrefix) String() string            { return proto.CompactTextString(m) }
func (*IpPrefix) ProtoMessage()               {}
func (*IpPrefix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IpPrefix) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *IpPrefix) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

type Identifier struct {
	Fqdn               string `protobuf:"bytes,1,opt,name=fqdn" json:"fqdn,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ConnectivityDomain string `protobuf:"bytes,3,opt,name=connectivity_domain,json=connectivityDomain" json:"connectivity_domain,omitempty"`
}

func (m *Identifier) Reset()                    { *m = Identifier{} }
func (m *Identifier) String() string            { return proto.CompactTextString(m) }
func (*Identifier) ProtoMessage()               {}
func (*Identifier) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Identifier) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Identifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identifier) GetConnectivityDomain() string {
	if m != nil {
		return m.ConnectivityDomain
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*IpFamily)(nil), "ippool.IpFamily")
	proto.RegisterType((*SubnetRequest)(nil), "ippool.SubnetRequest")
	proto.RegisterType((*Subnet)(nil), "ippool.Subnet")
	proto.RegisterType((*IpPrefix)(nil), "ippool.IpPrefix")
	proto.RegisterType((*Identifier)(nil), "ippool.Identifier")
	proto.RegisterType((*Empty)(nil), "ippool.Empty")
	proto.RegisterEnum("ippool.IpFamily_Family", IpFamily_Family_name, IpFamily_Family_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Allocator service

type AllocatorClient interface {
	AllocateSubnet(ctx context.Context, in *SubnetRequest, opts ...grpc.CallOption) (*Subnet, error)
	FreeSubnet(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Empty, error)
	RenewSubnetLease(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Subnet, error)
}

type allocatorClient struct {
	cc *grpc.ClientConn
}

func NewAllocatorClient(cc *grpc.ClientConn) AllocatorClient {
	return &allocatorClient{cc}
}

func (c *allocatorClient) AllocateSubnet(ctx context.Context, in *SubnetRequest, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := grpc.Invoke(ctx, "/ippool.Allocator/AllocateSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocatorClient) FreeSubnet(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/ippool.Allocator/FreeSubnet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocatorClient) RenewSubnetLease(ctx context.Context, in *Subnet, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := grpc.Invoke(ctx, "/ippool.Allocator/RenewSubnetLease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Allocator service

type AllocatorServer interface {
	AllocateSubnet(context.Context, *SubnetRequest) (*Subnet, error)
	FreeSubnet(context.Context, *Subnet) (*Empty, error)
	RenewSubnetLease(context.Context, *Subnet) (*Subnet, error)
}

func RegisterAllocatorServer(s *grpc.Server, srv AllocatorServer) {
	s.RegisterService(&_Allocator_serviceDesc, srv)
}

func _Allocator_AllocateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocatorServer).AllocateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.Allocator/AllocateSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocatorServer).AllocateSubnet(ctx, req.(*SubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Allocator_FreeSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocatorServer).FreeSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.Allocator/FreeSubnet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocatorServer).FreeSubnet(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Allocator_RenewSubnetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Subnet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocatorServer).RenewSubnetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ippool.Allocator/RenewSubnetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocatorServer).RenewSubnetLease(ctx, req.(*Subnet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Allocator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ippool.Allocator",
	HandlerType: (*AllocatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocateSubnet",
			Handler:    _Allocator_AllocateSubnet_Handler,
		},
		{
			MethodName: "FreeSubnet",
			Handler:    _Allocator_FreeSubnet_Handler,
		},
		{
			MethodName: "RenewSubnetLease",
			Handler:    _Allocator_RenewSubnetLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipprovider.proto",
}

func init() { proto.RegisterFile("ipprovider.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x5d, 0x43, 0x1b, 0x9a, 0x69, 0xb3, 0x8a, 0x8c, 0x80, 0xaa, 0x02, 0xa9, 0x32, 0x97, 0xbd,
	0x90, 0x8a, 0x50, 0x21, 0x71, 0x04, 0x41, 0xa5, 0x4a, 0x3d, 0xac, 0xcc, 0x87, 0x04, 0x97, 0x28,
	0xbb, 0x99, 0x48, 0x96, 0x12, 0xdb, 0xeb, 0xf5, 0x2e, 0xec, 0x6f, 0xe0, 0xcc, 0xaf, 0xe0, 0x4f,
	0xa2, 0xd8, 0xce, 0x7e, 0x5e, 0x50, 0x4f, 0x99, 0x79, 0x6f, 0x9e, 0xe7, 0xcd, 0xc4, 0x86, 0x54,
	0x68, 0x6d, 0xd4, 0x52, 0x54, 0x68, 0x32, 0x6d, 0x94, 0x55, 0x34, 0x12, 0x5a, 0x2b, 0xd5, 0xb0,
	0xef, 0x70, 0x72, 0xab, 0x6f, 0xca, 0x56, 0x34, 0x2b, 0x7a, 0x05, 0x51, 0xed, 0xa2, 0x73, 0x72,
	0x49, 0x46, 0xc3, 0xfc, 0x59, 0xe6, 0x8b, 0xb2, 0xbe, 0x22, 0xf3, 0x1f, 0x1e, 0xca, 0xd8, 0x73,
	0x88, 0x82, 0xf4, 0x04, 0x8e, 0x6e, 0xc7, 0xdf, 0xae, 0xd3, 0x41, 0x88, 0xde, 0xa6, 0x84, 0xfd,
	0x21, 0x90, 0x7c, 0x5e, 0x4c, 0x24, 0x5a, 0x8e, 0xb3, 0x05, 0xce, 0x2d, 0xcd, 0x01, 0x44, 0x85,
	0xd2, 0x8a, 0x5a, 0xa0, 0x71, 0x4d, 0x4e, 0x73, 0xba, 0x6e, 0xb2, 0x66, 0xf8, 0x56, 0x15, 0x7d,
	0x0d, 0xa7, 0x65, 0x55, 0x99, 0x22, 0x38, 0x7b, 0xe0, 0x44, 0xe9, 0xbe, 0x33, 0x0e, 0x5d, 0x51,
	0x30, 0xf3, 0x02, 0x40, 0x1b, 0xac, 0xc5, 0xaf, 0xa2, 0x41, 0x79, 0xfe, 0xf0, 0x92, 0x8c, 0x12,
	0x1e, 0x7b, 0xe4, 0x0e, 0x25, 0xfb, 0x4d, 0x20, 0xf2, 0xbe, 0xee, 0x65, 0x68, 0x04, 0x91, 0x3f,
	0xeb, 0xd0, 0xcb, 0xd8, 0xe1, 0x3c, 0xf0, 0xf4, 0x25, 0x24, 0x0d, 0x96, 0x73, 0x2c, 0xac, 0x68,
	0x51, 0x2d, 0xac, 0xb3, 0x72, 0xcc, 0xcf, 0x1c, 0xf8, 0xc5, 0x63, 0xec, 0x6b, 0xf7, 0x03, 0xbc,
	0x70, 0x7f, 0x56, 0xf2, 0x1f, 0xb3, 0x3e, 0x85, 0x68, 0xee, 0x66, 0x71, 0x6e, 0x62, 0x1e, 0x32,
	0x86, 0x00, 0x1b, 0xff, 0x94, 0xc2, 0x51, 0x3d, 0xab, 0xa4, 0x3b, 0x31, 0xe6, 0x2e, 0xee, 0x30,
	0x59, 0xb6, 0x18, 0x74, 0x2e, 0xa6, 0x57, 0xf0, 0x78, 0xaa, 0xa4, 0xc4, 0xa9, 0x15, 0x4b, 0x61,
	0x57, 0x45, 0xa5, 0xda, 0x52, 0xf8, 0x15, 0xc6, 0x9c, 0x6e, 0x53, 0x1f, 0x1d, 0xc3, 0x1e, 0xc1,
	0xf1, 0xa7, 0x56, 0xdb, 0x55, 0xfe, 0x97, 0x40, 0xfc, 0xbe, 0x69, 0xd4, 0xb4, 0xb4, 0xca, 0xd0,
	0x77, 0x30, 0x0c, 0x09, 0x86, 0x4d, 0x3f, 0xe9, 0xa7, 0xd8, 0xb9, 0x11, 0x17, 0xc3, 0x5d, 0x98,
	0x0d, 0xe8, 0x2b, 0x80, 0x1b, 0x83, 0xbd, 0x6c, 0x8f, 0xbf, 0x48, 0xfa, 0xdc, 0x75, 0x65, 0x03,
	0x7a, 0x0d, 0x29, 0x47, 0x89, 0x3f, 0x3d, 0x7f, 0xd7, 0x6d, 0xf6, 0x40, 0x74, 0xd0, 0xe4, 0xc3,
	0xd9, 0x0f, 0xd8, 0xbc, 0x88, 0x49, 0xe4, 0x9e, 0xc4, 0x9b, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xed, 0x5d, 0xe6, 0xf5, 0x26, 0x03, 0x00, 0x00,
}
