// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/password/v1/api_start_reset_flow.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StartResetFlowRequest struct {
	// email, username, or phone number
	AccountIdentifier    string   `protobuf:"bytes,1,opt,name=account_identifier,json=accountIdentifier,proto3" json:"account_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResetFlowRequest) Reset()         { *m = StartResetFlowRequest{} }
func (m *StartResetFlowRequest) String() string { return proto.CompactTextString(m) }
func (*StartResetFlowRequest) ProtoMessage()    {}
func (*StartResetFlowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5944c0cd52b85c23, []int{0}
}

func (m *StartResetFlowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResetFlowRequest.Unmarshal(m, b)
}
func (m *StartResetFlowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResetFlowRequest.Marshal(b, m, deterministic)
}
func (m *StartResetFlowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResetFlowRequest.Merge(m, src)
}
func (m *StartResetFlowRequest) XXX_Size() int {
	return xxx_messageInfo_StartResetFlowRequest.Size(m)
}
func (m *StartResetFlowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResetFlowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StartResetFlowRequest proto.InternalMessageInfo

func (m *StartResetFlowRequest) GetAccountIdentifier() string {
	if m != nil {
		return m.AccountIdentifier
	}
	return ""
}

type StartResetFlowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartResetFlowResponse) Reset()         { *m = StartResetFlowResponse{} }
func (m *StartResetFlowResponse) String() string { return proto.CompactTextString(m) }
func (*StartResetFlowResponse) ProtoMessage()    {}
func (*StartResetFlowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5944c0cd52b85c23, []int{1}
}

func (m *StartResetFlowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartResetFlowResponse.Unmarshal(m, b)
}
func (m *StartResetFlowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartResetFlowResponse.Marshal(b, m, deterministic)
}
func (m *StartResetFlowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartResetFlowResponse.Merge(m, src)
}
func (m *StartResetFlowResponse) XXX_Size() int {
	return xxx_messageInfo_StartResetFlowResponse.Size(m)
}
func (m *StartResetFlowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StartResetFlowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StartResetFlowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StartResetFlowRequest)(nil), "alpacalabs.password.v1.StartResetFlowRequest")
	proto.RegisterType((*StartResetFlowResponse)(nil), "alpacalabs.password.v1.StartResetFlowResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/password/v1/api_start_reset_flow.proto", fileDescriptor_5944c0cd52b85c23)
}

var fileDescriptor_5944c0cd52b85c23 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcc, 0x29, 0x48,
	0x4c, 0x4e, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0x2f, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0xd1,
	0x2f, 0x33, 0xd4, 0x4f, 0x2c, 0xc8, 0x8c, 0x2f, 0x2e, 0x49, 0x2c, 0x2a, 0x89, 0x2f, 0x4a, 0x2d,
	0x4e, 0x2d, 0x89, 0x4f, 0xcb, 0xc9, 0x2f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x43,
	0x68, 0xd1, 0x83, 0x69, 0xd1, 0x2b, 0x33, 0x54, 0x72, 0xe3, 0x12, 0x0d, 0x06, 0xe9, 0x08, 0x02,
	0x69, 0x70, 0xcb, 0xc9, 0x2f, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xd2, 0xe5, 0x12,
	0x4a, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0x49, 0xcd, 0x2b, 0xc9, 0x4c, 0xcb,
	0x4c, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x12, 0x84, 0xca, 0x78, 0xc2, 0x25, 0x94,
	0x24, 0xb8, 0xc4, 0xd0, 0xcd, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x75, 0x5a, 0xc9, 0xc8, 0x25,
	0x95, 0x9c, 0x9f, 0xab, 0x87, 0xdd, 0x01, 0x4e, 0xbc, 0x01, 0x50, 0x4e, 0x00, 0xc8, 0x9d, 0x01,
	0x8c, 0x51, 0x4e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x8e, 0x60,
	0x3d, 0x3e, 0x60, 0x7f, 0x82, 0xa4, 0x8b, 0x52, 0x0b, 0xf2, 0x75, 0x61, 0xba, 0x75, 0xd3, 0xf3,
	0xf5, 0xb1, 0x07, 0xc4, 0x22, 0x26, 0x66, 0x47, 0x9f, 0x80, 0x55, 0x4c, 0x62, 0x08, 0xfd, 0x7a,
	0x30, 0x6b, 0xf4, 0xc2, 0x0c, 0x4f, 0x21, 0x4b, 0xc4, 0xc0, 0x24, 0x62, 0xc2, 0x0c, 0x93, 0xd8,
	0xc0, 0xb6, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x56, 0x3e, 0x75, 0xd1, 0x61, 0x01, 0x00,
	0x00,
}
