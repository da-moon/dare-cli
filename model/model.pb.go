// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model.proto

package model

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Hash ...
type Hash struct {
	Md5                  string   `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Sha256               string   `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *Hash) Reset() { *m = Hash{} }

// String ...
func (m *Hash) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*Hash) ProtoMessage() {}

// Descriptor ...
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{0}
}

// XXX_Unmarshal ...
func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}

// XXX_Size ...
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}

// XXX_DiscardUnknown ...
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

// GetMd5 ...
func (m *Hash) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

// GetSha256 ...
func (m *Hash) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

// EncryptRequest ...
type EncryptRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *EncryptRequest) Reset() { *m = EncryptRequest{} }

// String ...
func (m *EncryptRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*EncryptRequest) ProtoMessage() {}

// Descriptor ...
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{1}
}

// XXX_Unmarshal ...
func (m *EncryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *EncryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *EncryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptRequest.Merge(m, src)
}

// XXX_Size ...
func (m *EncryptRequest) XXX_Size() int {
	return xxx_messageInfo_EncryptRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *EncryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptRequest proto.InternalMessageInfo

// GetSource ...
func (m *EncryptRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// GetDestination ...
func (m *EncryptRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

// GetKey ...
func (m *EncryptRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// EncryptResponse ...
type EncryptResponse struct {
	OutputHash           *Hash    `protobuf:"bytes,1,opt,name=output_hash,json=outputHash,proto3" json:"output_hash,omitempty"`
	RandomNonce          string   `protobuf:"bytes,2,opt,name=random_nonce,json=randomNonce,proto3" json:"random_nonce,omitempty"`
	RandomKey            string   `protobuf:"bytes,3,opt,name=random_key,json=randomKey,proto3" json:"random_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *EncryptResponse) Reset() { *m = EncryptResponse{} }

// String ...
func (m *EncryptResponse) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*EncryptResponse) ProtoMessage() {}

// Descriptor ...
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{2}
}

// XXX_Unmarshal ...
func (m *EncryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptResponse.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *EncryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptResponse.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *EncryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptResponse.Merge(m, src)
}

// XXX_Size ...
func (m *EncryptResponse) XXX_Size() int {
	return xxx_messageInfo_EncryptResponse.Size(m)
}

// XXX_DiscardUnknown ...
func (m *EncryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptResponse proto.InternalMessageInfo

// GetOutputHash ...
func (m *EncryptResponse) GetOutputHash() *Hash {
	if m != nil {
		return m.OutputHash
	}
	return nil
}

// GetRandomNonce ...
func (m *EncryptResponse) GetRandomNonce() string {
	if m != nil {
		return m.RandomNonce
	}
	return ""
}

// GetRandomKey ...
func (m *EncryptResponse) GetRandomKey() string {
	if m != nil {
		return m.RandomKey
	}
	return ""
}

// DecryptRequest ...
type DecryptRequest struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	Nonce                string   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Key                  string   `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *DecryptRequest) Reset() { *m = DecryptRequest{} }

// String ...
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*DecryptRequest) ProtoMessage() {}

// Descriptor ...
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{3}
}

// XXX_Unmarshal ...
func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}

// XXX_Size ...
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}

// XXX_DiscardUnknown ...
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

// GetSource ...
func (m *DecryptRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// GetDestination ...
func (m *DecryptRequest) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

// GetNonce ...
func (m *DecryptRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

// GetKey ...
func (m *DecryptRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// DecryptResponse ...
type DecryptResponse struct {
	OutputHash           *Hash    `protobuf:"bytes,1,opt,name=output_hash,json=outputHash,proto3" json:"output_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

// Reset ...
func (m *DecryptResponse) Reset() { *m = DecryptResponse{} }

// String ...
func (m *DecryptResponse) String() string { return proto.CompactTextString(m) }

// ProtoMessage ...
func (*DecryptResponse) ProtoMessage() {}

// Descriptor ...
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c16552f9fdb66d8, []int{4}
}

// XXX_Unmarshal ...
func (m *DecryptResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptResponse.Unmarshal(m, b)
}

// XXX_Marshal ...
func (m *DecryptResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptResponse.Marshal(b, m, deterministic)
}

// XXX_Merge ...
func (m *DecryptResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptResponse.Merge(m, src)
}

// XXX_Size ...
func (m *DecryptResponse) XXX_Size() int {
	return xxx_messageInfo_DecryptResponse.Size(m)
}

// XXX_DiscardUnknown ...
func (m *DecryptResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptResponse proto.InternalMessageInfo

// GetOutputHash ...
func (m *DecryptResponse) GetOutputHash() *Hash {
	if m != nil {
		return m.OutputHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Hash)(nil), "model.Hash")
	proto.RegisterType((*EncryptRequest)(nil), "model.EncryptRequest")
	proto.RegisterType((*EncryptResponse)(nil), "model.EncryptResponse")
	proto.RegisterType((*DecryptRequest)(nil), "model.DecryptRequest")
	proto.RegisterType((*DecryptResponse)(nil), "model.DecryptResponse")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_4c16552f9fdb66d8) }

var fileDescriptor_4c16552f9fdb66d8 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0xfb, 0x47, 0x32, 0x91, 0x56, 0x16, 0x2d, 0x41, 0x10, 0xea, 0x9e, 0x3c, 0x48,
	0x91, 0x48, 0x3d, 0x78, 0xf1, 0xa2, 0x28, 0x08, 0x1e, 0x72, 0x16, 0x4a, 0x4c, 0x06, 0x52, 0x6a,
	0x76, 0xe3, 0xee, 0xe6, 0x90, 0xa3, 0xdf, 0x5c, 0x76, 0x77, 0x9a, 0xda, 0x1e, 0xa5, 0xb7, 0xbc,
	0x97, 0xdd, 0xf7, 0x9b, 0x37, 0x09, 0x44, 0x95, 0x2c, 0xf0, 0x6b, 0x5e, 0x2b, 0x69, 0x24, 0x1b,
	0x3a, 0xc1, 0x6f, 0x61, 0xf0, 0x9a, 0xe9, 0x92, 0x9d, 0x42, 0xbf, 0x2a, 0x16, 0x71, 0x30, 0x0b,
	0xae, 0xc3, 0xd4, 0x3e, 0xb2, 0x29, 0x8c, 0x74, 0x99, 0x25, 0x8b, 0xfb, 0xf8, 0xc8, 0x99, 0xa4,
	0xf8, 0x07, 0x8c, 0x9f, 0x45, 0xae, 0xda, 0xda, 0xa4, 0xf8, 0xdd, 0xa0, 0x36, 0xee, 0xa4, 0x6c,
	0x54, 0x8e, 0x74, 0x9d, 0x14, 0x9b, 0x41, 0x54, 0xa0, 0x36, 0x2b, 0x91, 0x99, 0x95, 0x14, 0x14,
	0xf3, 0xd7, 0xb2, 0xd4, 0x35, 0xb6, 0x71, 0xdf, 0x53, 0xd7, 0xd8, 0xf2, 0x9f, 0x00, 0x26, 0x5d,
	0xbc, 0xae, 0xa5, 0xd0, 0xc8, 0x6e, 0x20, 0x92, 0x8d, 0xa9, 0x1b, 0xb3, 0x2c, 0x33, 0x5d, 0x3a,
	0x48, 0x94, 0x44, 0x73, 0xdf, 0xc6, 0x4e, 0x9f, 0x82, 0x7f, 0xef, 0x9a, 0x5c, 0xc1, 0x89, 0xca,
	0x44, 0x21, 0xab, 0xa5, 0x90, 0x22, 0xc7, 0x0d, 0xd6, 0x7b, 0xef, 0xd6, 0x62, 0x97, 0x00, 0x74,
	0x64, 0x4b, 0x0f, 0xbd, 0xf3, 0x86, 0x2d, 0x57, 0x30, 0x7e, 0xc2, 0x03, 0x35, 0x3c, 0x83, 0xa1,
	0x1f, 0xc3, 0x53, 0xbc, 0xd8, 0xf4, 0x1e, 0x6c, 0x7b, 0x3f, 0xc2, 0xa4, 0x63, 0xfe, 0xa7, 0x76,
	0xf2, 0x02, 0x21, 0xed, 0x4d, 0x2a, 0xf6, 0x00, 0xc7, 0x24, 0xd8, 0x39, 0x5d, 0xd8, 0xfd, 0x66,
	0x17, 0xd3, 0x7d, 0xdb, 0x43, 0x79, 0xcf, 0x06, 0xd1, 0x24, 0x3e, 0x88, 0x44, 0x17, 0xb4, 0xbb,
	0x9a, 0x2e, 0x68, 0x6f, 0x7a, 0xde, 0xfb, 0x1c, 0xb9, 0x1f, 0xed, 0xee, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xc2, 0xbd, 0x2d, 0x1a, 0x77, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EncryptorClient is the client API for Encryptor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EncryptorClient interface {
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
}

type encryptorClient struct {
	cc *grpc.ClientConn
}

// NewEncryptorClient ...
func NewEncryptorClient(cc *grpc.ClientConn) EncryptorClient {
	return &encryptorClient{cc}
}

// Encrypt ...
func (c *encryptorClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/model.Encryptor/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncryptorServer is the server API for Encryptor service.
type EncryptorServer interface {
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
}

// UnimplementedEncryptorServer can be embedded to have forward compatible implementations.
type UnimplementedEncryptorServer struct {
}

// Encrypt ...
func (*UnimplementedEncryptorServer) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}

// RegisterEncryptorServer ...
func RegisterEncryptorServer(s *grpc.Server, srv EncryptorServer) {
	s.RegisterService(&_Encryptor_serviceDesc, srv)
}

func _Encryptor_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncryptorServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Encryptor/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncryptorServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Encryptor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Encryptor",
	HandlerType: (*EncryptorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _Encryptor_Encrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}

// DecryptorClient is the client API for Decryptor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DecryptorClient interface {
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type decryptorClient struct {
	cc *grpc.ClientConn
}

// NewDecryptorClient ...
func NewDecryptorClient(cc *grpc.ClientConn) DecryptorClient {
	return &decryptorClient{cc}
}

// Decrypt ...
func (c *decryptorClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/model.Decryptor/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecryptorServer is the server API for Decryptor service.
type DecryptorServer interface {
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

// UnimplementedDecryptorServer can be embedded to have forward compatible implementations.
type UnimplementedDecryptorServer struct {
}

// Decrypt ...
func (*UnimplementedDecryptorServer) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

// RegisterDecryptorServer ...
func RegisterDecryptorServer(s *grpc.Server, srv DecryptorServer) {
	s.RegisterService(&_Decryptor_serviceDesc, srv)
}

func _Decryptor_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptorServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Decryptor/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptorServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Decryptor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "model.Decryptor",
	HandlerType: (*DecryptorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decrypt",
			Handler:    _Decryptor_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "model.proto",
}
