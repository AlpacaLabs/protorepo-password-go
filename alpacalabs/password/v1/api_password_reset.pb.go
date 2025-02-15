// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/password/v1/api_password_reset.proto

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

type GetDeliveryOptionsRequest struct {
	// email, username, or phone number
	AccountIdentifier    string   `protobuf:"bytes,1,opt,name=account_identifier,json=accountIdentifier,proto3" json:"account_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeliveryOptionsRequest) Reset()         { *m = GetDeliveryOptionsRequest{} }
func (m *GetDeliveryOptionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeliveryOptionsRequest) ProtoMessage()    {}
func (*GetDeliveryOptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{0}
}

func (m *GetDeliveryOptionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeliveryOptionsRequest.Unmarshal(m, b)
}
func (m *GetDeliveryOptionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeliveryOptionsRequest.Marshal(b, m, deterministic)
}
func (m *GetDeliveryOptionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeliveryOptionsRequest.Merge(m, src)
}
func (m *GetDeliveryOptionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeliveryOptionsRequest.Size(m)
}
func (m *GetDeliveryOptionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeliveryOptionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeliveryOptionsRequest proto.InternalMessageInfo

func (m *GetDeliveryOptionsRequest) GetAccountIdentifier() string {
	if m != nil {
		return m.AccountIdentifier
	}
	return ""
}

type GetDeliveryOptionsResponse struct {
	CodeId               string               `protobuf:"bytes,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	CodeOptions          *CodeDeliveryOptions `protobuf:"bytes,2,opt,name=code_options,json=codeOptions,proto3" json:"code_options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetDeliveryOptionsResponse) Reset()         { *m = GetDeliveryOptionsResponse{} }
func (m *GetDeliveryOptionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeliveryOptionsResponse) ProtoMessage()    {}
func (*GetDeliveryOptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{1}
}

func (m *GetDeliveryOptionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeliveryOptionsResponse.Unmarshal(m, b)
}
func (m *GetDeliveryOptionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeliveryOptionsResponse.Marshal(b, m, deterministic)
}
func (m *GetDeliveryOptionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeliveryOptionsResponse.Merge(m, src)
}
func (m *GetDeliveryOptionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeliveryOptionsResponse.Size(m)
}
func (m *GetDeliveryOptionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeliveryOptionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeliveryOptionsResponse proto.InternalMessageInfo

func (m *GetDeliveryOptionsResponse) GetCodeId() string {
	if m != nil {
		return m.CodeId
	}
	return ""
}

func (m *GetDeliveryOptionsResponse) GetCodeOptions() *CodeDeliveryOptions {
	if m != nil {
		return m.CodeOptions
	}
	return nil
}

type DeliverCodeRequest struct {
	CodeId               string   `protobuf:"bytes,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	EmailAddressId       string   `protobuf:"bytes,2,opt,name=email_address_id,json=emailAddressId,proto3" json:"email_address_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverCodeRequest) Reset()         { *m = DeliverCodeRequest{} }
func (m *DeliverCodeRequest) String() string { return proto.CompactTextString(m) }
func (*DeliverCodeRequest) ProtoMessage()    {}
func (*DeliverCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{2}
}

func (m *DeliverCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverCodeRequest.Unmarshal(m, b)
}
func (m *DeliverCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverCodeRequest.Marshal(b, m, deterministic)
}
func (m *DeliverCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverCodeRequest.Merge(m, src)
}
func (m *DeliverCodeRequest) XXX_Size() int {
	return xxx_messageInfo_DeliverCodeRequest.Size(m)
}
func (m *DeliverCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverCodeRequest proto.InternalMessageInfo

func (m *DeliverCodeRequest) GetCodeId() string {
	if m != nil {
		return m.CodeId
	}
	return ""
}

func (m *DeliverCodeRequest) GetEmailAddressId() string {
	if m != nil {
		return m.EmailAddressId
	}
	return ""
}

type DeliverCodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeliverCodeResponse) Reset()         { *m = DeliverCodeResponse{} }
func (m *DeliverCodeResponse) String() string { return proto.CompactTextString(m) }
func (*DeliverCodeResponse) ProtoMessage()    {}
func (*DeliverCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{3}
}

func (m *DeliverCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeliverCodeResponse.Unmarshal(m, b)
}
func (m *DeliverCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeliverCodeResponse.Marshal(b, m, deterministic)
}
func (m *DeliverCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeliverCodeResponse.Merge(m, src)
}
func (m *DeliverCodeResponse) XXX_Size() int {
	return xxx_messageInfo_DeliverCodeResponse.Size(m)
}
func (m *DeliverCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeliverCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeliverCodeResponse proto.InternalMessageInfo

type VerifyCodeRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	AccountId            string   `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	NewPassword          string   `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeRequest) Reset()         { *m = VerifyCodeRequest{} }
func (m *VerifyCodeRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeRequest) ProtoMessage()    {}
func (*VerifyCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{4}
}

func (m *VerifyCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeRequest.Unmarshal(m, b)
}
func (m *VerifyCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeRequest.Marshal(b, m, deterministic)
}
func (m *VerifyCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeRequest.Merge(m, src)
}
func (m *VerifyCodeRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeRequest.Size(m)
}
func (m *VerifyCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeRequest proto.InternalMessageInfo

func (m *VerifyCodeRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *VerifyCodeRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *VerifyCodeRequest) GetNewPassword() string {
	if m != nil {
		return m.NewPassword
	}
	return ""
}

type VerifyCodeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyCodeResponse) Reset()         { *m = VerifyCodeResponse{} }
func (m *VerifyCodeResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyCodeResponse) ProtoMessage()    {}
func (*VerifyCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_820da2458cb0ca4e, []int{5}
}

func (m *VerifyCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyCodeResponse.Unmarshal(m, b)
}
func (m *VerifyCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyCodeResponse.Marshal(b, m, deterministic)
}
func (m *VerifyCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyCodeResponse.Merge(m, src)
}
func (m *VerifyCodeResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyCodeResponse.Size(m)
}
func (m *VerifyCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyCodeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetDeliveryOptionsRequest)(nil), "alpacalabs.password.v1.GetDeliveryOptionsRequest")
	proto.RegisterType((*GetDeliveryOptionsResponse)(nil), "alpacalabs.password.v1.GetDeliveryOptionsResponse")
	proto.RegisterType((*DeliverCodeRequest)(nil), "alpacalabs.password.v1.DeliverCodeRequest")
	proto.RegisterType((*DeliverCodeResponse)(nil), "alpacalabs.password.v1.DeliverCodeResponse")
	proto.RegisterType((*VerifyCodeRequest)(nil), "alpacalabs.password.v1.VerifyCodeRequest")
	proto.RegisterType((*VerifyCodeResponse)(nil), "alpacalabs.password.v1.VerifyCodeResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/password/v1/api_password_reset.proto", fileDescriptor_820da2458cb0ca4e)
}

var fileDescriptor_820da2458cb0ca4e = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xdd, 0xaa, 0xd3, 0x40,
	0x10, 0x26, 0x39, 0x72, 0xe4, 0x4c, 0x8f, 0x62, 0x57, 0xad, 0xb5, 0x20, 0xd4, 0x5c, 0x15, 0xa4,
	0x1b, 0x52, 0x9f, 0xa0, 0x55, 0x90, 0x4a, 0xd1, 0xd0, 0x8b, 0x0a, 0x52, 0x08, 0xdb, 0xec, 0xb4,
	0x2e, 0xa4, 0xd9, 0xb8, 0xbb, 0x6d, 0xe9, 0x03, 0xf8, 0x22, 0xde, 0xe9, 0xa3, 0xf8, 0x54, 0x92,
	0xcd, 0xa6, 0xad, 0xd2, 0xdc, 0x65, 0xe6, 0xfb, 0x9b, 0x99, 0x2c, 0x84, 0x2c, 0x2b, 0x58, 0xca,
	0x32, 0xb6, 0xd2, 0x61, 0xc1, 0xb4, 0x3e, 0x48, 0xc5, 0xc3, 0x7d, 0x14, 0xb2, 0x42, 0x24, 0x75,
	0x9d, 0x28, 0xd4, 0x68, 0x68, 0xa1, 0xa4, 0x91, 0xa4, 0x73, 0x16, 0xd0, 0x9a, 0x40, 0xf7, 0x51,
	0x6f, 0xd4, 0x60, 0x94, 0x4a, 0x8e, 0x09, 0xc7, 0x4c, 0xec, 0x51, 0x1d, 0x13, 0x59, 0x18, 0x21,
	0x73, 0x5d, 0x79, 0x05, 0x1f, 0xe1, 0xe5, 0x07, 0x34, 0xef, 0x1d, 0xf8, 0xb9, 0xc2, 0xe6, 0xf8,
	0x7d, 0x87, 0xda, 0x90, 0x21, 0x10, 0x96, 0xa6, 0x72, 0x97, 0x9b, 0x44, 0x70, 0xcc, 0x8d, 0x58,
	0x0b, 0x54, 0x5d, 0xaf, 0xef, 0x0d, 0xee, 0xe6, 0x6d, 0x87, 0x4c, 0x4f, 0x40, 0xf0, 0xc3, 0x83,
	0xde, 0x35, 0x33, 0x5d, 0xc8, 0x5c, 0x23, 0x79, 0x01, 0x0f, 0xed, 0x24, 0x82, 0x3b, 0x8b, 0xdb,
	0xb2, 0x9c, 0x72, 0xf2, 0x09, 0xee, 0x2d, 0xe0, 0x26, 0xeb, 0xfa, 0x7d, 0x6f, 0xd0, 0x1a, 0xbd,
	0xa1, 0xd7, 0xd7, 0xa4, 0xef, 0x24, 0xc7, 0xff, 0x33, 0x5a, 0xa5, 0x81, 0x2b, 0x82, 0x2f, 0x40,
	0x1c, 0x5e, 0x52, 0xeb, 0x65, 0x1a, 0xe3, 0x07, 0xf0, 0x04, 0xb7, 0x4c, 0x64, 0x09, 0xe3, 0x5c,
	0xa1, 0xd6, 0x25, 0xc3, 0xb7, 0x8c, 0xc7, 0xb6, 0x3f, 0xae, 0xda, 0x53, 0x1e, 0x3c, 0x87, 0xa7,
	0xff, 0x18, 0x57, 0x8b, 0x05, 0x02, 0xda, 0x0b, 0x54, 0x62, 0x7d, 0xbc, 0x8c, 0x23, 0xf0, 0xa0,
	0xf4, 0x77, 0x59, 0xf6, 0x9b, 0xbc, 0x02, 0x38, 0xdf, 0xd3, 0x65, 0xdc, 0x9d, 0xee, 0x48, 0x5e,
	0xc3, 0x7d, 0x8e, 0x87, 0xd3, 0x3f, 0xef, 0xde, 0x58, 0x42, 0x2b, 0xc7, 0x43, 0xec, 0x5a, 0xc1,
	0x33, 0x20, 0x97, 0x51, 0xd5, 0x00, 0x93, 0x5f, 0x1e, 0xf4, 0x52, 0xb9, 0x6d, 0x38, 0xd8, 0xe4,
	0x51, 0x2d, 0x8f, 0xcb, 0x5f, 0x1e, 0x7b, 0x5f, 0x27, 0x1b, 0x61, 0xbe, 0xed, 0x56, 0x34, 0x95,
	0xdb, 0x70, 0x6c, 0x35, 0x33, 0xfb, 0x66, 0x4a, 0x58, 0x61, 0x21, 0x87, 0xb5, 0x7a, 0xb8, 0x91,
	0x0d, 0xaf, 0xf3, 0xa7, 0x7f, 0x33, 0x9e, 0xc5, 0xbf, 0xfd, 0xce, 0x59, 0x4f, 0xeb, 0x18, 0xba,
	0x88, 0xfe, 0x5c, 0x02, 0xcb, 0x1a, 0x58, 0x2e, 0xa2, 0xd5, 0xad, 0x4d, 0x79, 0xfb, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x21, 0xaf, 0x90, 0x0a, 0xf6, 0x02, 0x00, 0x00,
}
