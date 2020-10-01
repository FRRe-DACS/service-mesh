// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package foo

import (
	context "context"
	fmt "fmt"
	bar "github.com/FRRe-DACS/service-mesh/pkg/dacs/bar"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

// Foo message
type Foo struct {
	// Id of the bar
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bar                  *bar.Bar `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Foo) GetBar() *bar.Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

// CreateFooRequest is a request for creating a Foo
type CreateFooRequest struct {
	// Foo to create
	Foo                  *Foo     `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFooRequest) Reset()         { *m = CreateFooRequest{} }
func (m *CreateFooRequest) String() string { return proto.CompactTextString(m) }
func (*CreateFooRequest) ProtoMessage()    {}
func (*CreateFooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}

func (m *CreateFooRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFooRequest.Unmarshal(m, b)
}
func (m *CreateFooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFooRequest.Marshal(b, m, deterministic)
}
func (m *CreateFooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFooRequest.Merge(m, src)
}
func (m *CreateFooRequest) XXX_Size() int {
	return xxx_messageInfo_CreateFooRequest.Size(m)
}
func (m *CreateFooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFooRequest proto.InternalMessageInfo

func (m *CreateFooRequest) GetFoo() *Foo {
	if m != nil {
		return m.Foo
	}
	return nil
}

// CreateFooResponse is the response after creating a Foo
type CreateFooResponse struct {
	// Created bar
	Foo                  *Foo     `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateFooResponse) Reset()         { *m = CreateFooResponse{} }
func (m *CreateFooResponse) String() string { return proto.CompactTextString(m) }
func (*CreateFooResponse) ProtoMessage()    {}
func (*CreateFooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{2}
}

func (m *CreateFooResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateFooResponse.Unmarshal(m, b)
}
func (m *CreateFooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateFooResponse.Marshal(b, m, deterministic)
}
func (m *CreateFooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateFooResponse.Merge(m, src)
}
func (m *CreateFooResponse) XXX_Size() int {
	return xxx_messageInfo_CreateFooResponse.Size(m)
}
func (m *CreateFooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateFooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateFooResponse proto.InternalMessageInfo

func (m *CreateFooResponse) GetFoo() *Foo {
	if m != nil {
		return m.Foo
	}
	return nil
}

func init() {
	proto.RegisterType((*Foo)(nil), "ar.edu.utn.frre.dacs.foo.Foo")
	proto.RegisterType((*CreateFooRequest)(nil), "ar.edu.utn.frre.dacs.foo.CreateFooRequest")
	proto.RegisterType((*CreateFooResponse)(nil), "ar.edu.utn.frre.dacs.foo.CreateFooResponse")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0x0a, 0xd2, 0x5c, 0xc1, 0x9f, 0x80, 0x10, 0x82, 0x42, 0xc9, 0xaa, 0x28, 0xcc,
	0x60, 0xdd, 0xb9, 0x6c, 0x25, 0x0f, 0x10, 0x77, 0xee, 0x26, 0xcd, 0x9d, 0x32, 0x20, 0x73, 0xe2,
	0xcc, 0xa4, 0x1b, 0x77, 0xee, 0x5c, 0xfb, 0x68, 0xbe, 0x82, 0x0f, 0x22, 0xd3, 0x80, 0x88, 0x58,
	0xd0, 0xdd, 0x30, 0x9c, 0xef, 0xbb, 0xe7, 0x5e, 0xca, 0x34, 0x20, 0x7a, 0x87, 0x80, 0xbc, 0x50,
	0x4e, 0x70, 0x37, 0x88, 0x21, 0x58, 0xa1, 0x9d, 0x63, 0xd1, 0xa9, 0xb5, 0x17, 0x1a, 0x28, 0xcf,
	0x37, 0xc0, 0xe6, 0x91, 0xa5, 0xea, 0x8d, 0x54, 0xd6, 0x22, 0xa8, 0x60, 0x60, 0xfd, 0xc8, 0x95,
	0x59, 0xab, 0xdc, 0xf8, 0xac, 0x6a, 0x9a, 0xd4, 0x40, 0x7e, 0x44, 0xa9, 0xe9, 0x8a, 0x64, 0x96,
	0xcc, 0x27, 0x4d, 0x6a, 0xba, 0x5c, 0xd2, 0xa4, 0x55, 0xae, 0x48, 0x67, 0xc9, 0xfc, 0x70, 0x71,
	0x21, 0x7e, 0x9d, 0x13, 0x25, 0x4b, 0xe5, 0x9a, 0x98, 0xac, 0x56, 0x74, 0xb2, 0x72, 0xac, 0x02,
	0xd7, 0x40, 0xc3, 0x4f, 0x03, 0xfb, 0x10, 0x25, 0x1a, 0xd8, 0x59, 0xf7, 0x4a, 0xe2, 0x32, 0x11,
	0x89, 0xc9, 0xea, 0x8e, 0x4e, 0xbf, 0x49, 0x7c, 0x0f, 0xeb, 0xf9, 0xdf, 0x96, 0xc5, 0x6b, 0x42,
	0x54, 0x03, 0xf7, 0xec, 0xb6, 0x66, 0xcd, 0xf9, 0x33, 0x65, 0x5f, 0xd2, 0xfc, 0x72, 0x3f, 0xff,
	0xb3, 0x7e, 0x79, 0xf5, 0xa7, 0xec, 0xd8, 0xb2, 0x3a, 0x7b, 0x79, 0xff, 0x78, 0x4b, 0x8f, 0xab,
	0xa9, 0xdc, 0x5e, 0x4b, 0x0d, 0xf8, 0xdb, 0xd8, 0x65, 0x49, 0x0f, 0xd3, 0x08, 0xc5, 0x9f, 0xf6,
	0x60, 0x77, 0xf1, 0x9b, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xc9, 0xfc, 0x83, 0xc1, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FooServiceClient is the client API for FooService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FooServiceClient interface {
	// Creates a new Foo in the system
	CreateFoo(ctx context.Context, in *CreateFooRequest, opts ...grpc.CallOption) (*CreateFooResponse, error)
}

type fooServiceClient struct {
	cc *grpc.ClientConn
}

func NewFooServiceClient(cc *grpc.ClientConn) FooServiceClient {
	return &fooServiceClient{cc}
}

func (c *fooServiceClient) CreateFoo(ctx context.Context, in *CreateFooRequest, opts ...grpc.CallOption) (*CreateFooResponse, error) {
	out := new(CreateFooResponse)
	err := c.cc.Invoke(ctx, "/ar.edu.utn.frre.dacs.foo.FooService/CreateFoo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FooServiceServer is the server API for FooService service.
type FooServiceServer interface {
	// Creates a new Foo in the system
	CreateFoo(context.Context, *CreateFooRequest) (*CreateFooResponse, error)
}

func RegisterFooServiceServer(s *grpc.Server, srv FooServiceServer) {
	s.RegisterService(&_FooService_serviceDesc, srv)
}

func _FooService_CreateFoo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFooRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FooServiceServer).CreateFoo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ar.edu.utn.frre.dacs.foo.FooService/CreateFoo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FooServiceServer).CreateFoo(ctx, req.(*CreateFooRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FooService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ar.edu.utn.frre.dacs.foo.FooService",
	HandlerType: (*FooServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFoo",
			Handler:    _FooService_CreateFoo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}
