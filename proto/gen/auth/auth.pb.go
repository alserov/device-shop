// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.1
// source: auth/auth.proto

package auth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckIfAdminReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *CheckIfAdminReq) Reset() {
	*x = CheckIfAdminReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfAdminReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfAdminReq) ProtoMessage() {}

func (x *CheckIfAdminReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfAdminReq.ProtoReflect.Descriptor instead.
func (*CheckIfAdminReq) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CheckIfAdminReq) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type CheckIfAdminRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin bool `protobuf:"varint,1,opt,name=IsAdmin,proto3" json:"IsAdmin,omitempty"`
}

func (x *CheckIfAdminRes) Reset() {
	*x = CheckIfAdminRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckIfAdminRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckIfAdminRes) ProtoMessage() {}

func (x *CheckIfAdminRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckIfAdminRes.ProtoReflect.Descriptor instead.
func (*CheckIfAdminRes) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CheckIfAdminRes) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUUID string `protobuf:"bytes,1,opt,name=UserUUID,proto3" json:"UserUUID,omitempty"`
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserInfoReq) GetUserUUID() string {
	if x != nil {
		return x.UserUUID
	}
	return ""
}

type GetUserInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string  `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email    string  `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	UUID     string  `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Cash     float32 `protobuf:"fixed32,4,opt,name=Cash,proto3" json:"Cash,omitempty"`
}

func (x *GetUserInfoRes) Reset() {
	*x = GetUserInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoRes) ProtoMessage() {}

func (x *GetUserInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoRes.ProtoReflect.Descriptor instead.
func (*GetUserInfoRes) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserInfoRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserInfoRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserInfoRes) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *GetUserInfoRes) GetCash() float32 {
	if x != nil {
		return x.Cash
	}
	return 0
}

type SignupRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username     string  `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Email        string  `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	UUID         string  `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Cash         float32 `protobuf:"fixed32,4,opt,name=Cash,proto3" json:"Cash,omitempty"`
	RefreshToken string  `protobuf:"bytes,5,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Token        string  `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *SignupRes) Reset() {
	*x = SignupRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRes) ProtoMessage() {}

func (x *SignupRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRes.ProtoReflect.Descriptor instead.
func (*SignupRes) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SignupRes) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignupRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignupRes) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *SignupRes) GetCash() float32 {
	if x != nil {
		return x.Cash
	}
	return 0
}

func (x *SignupRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *SignupRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SignupReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *SignupReq) Reset() {
	*x = SignupReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupReq) ProtoMessage() {}

func (x *SignupReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupReq.ProtoReflect.Descriptor instead.
func (*SignupReq) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{5}
}

func (x *SignupReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SignupReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignupReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	UUID         string `protobuf:"bytes,3,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *LoginRes) Reset() {
	*x = LoginRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRes) ProtoMessage() {}

func (x *LoginRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRes.ProtoReflect.Descriptor instead.
func (*LoginRes) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRes) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginRes) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginRes) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_protos_auth_proto_rawDescGZIP(), []int{7}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_protos_auth_proto protoreflect.FileDescriptor

var file_protos_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x2d, 0x0a, 0x0f, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x49, 0x66, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x55, 0x55, 0x49, 0x44, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x49, 0x66, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x49,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x49, 0x73,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x55,
	0x55, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x43,
	0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x43, 0x61, 0x73, 0x68, 0x22,
	0x9f, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x43, 0x61, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x59, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x58, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x42, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xd4, 0x01, 0x0a, 0x04, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x0f, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x12,
	0x27, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x49, 0x66, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x66, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x61, 0x6c, 0x73, 0x65, 0x72, 0x6f, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_auth_proto_rawDescOnce sync.Once
	file_protos_auth_proto_rawDescData = file_protos_auth_proto_rawDesc
)

func file_protos_auth_proto_rawDescGZIP() []byte {
	file_protos_auth_proto_rawDescOnce.Do(func() {
		file_protos_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_auth_proto_rawDescData)
	})
	return file_protos_auth_proto_rawDescData
}

var file_protos_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_auth_proto_goTypes = []interface{}{
	(*CheckIfAdminReq)(nil), // 0: auth.CheckIfAdminReq
	(*CheckIfAdminRes)(nil), // 1: auth.CheckIfAdminRes
	(*GetUserInfoReq)(nil),  // 2: auth.GetUserInfoReq
	(*GetUserInfoRes)(nil),  // 3: auth.GetUserInfoRes
	(*SignupRes)(nil),       // 4: auth.SignupRes
	(*SignupReq)(nil),       // 5: auth.SignupReq
	(*LoginRes)(nil),        // 6: auth.LoginRes
	(*LoginReq)(nil),        // 7: auth.LoginReq
}
var file_protos_auth_proto_depIdxs = []int32{
	5, // 0: auth.Auth.Signup:input_type -> auth.SignupReq
	7, // 1: auth.Auth.Login:input_type -> auth.LoginReq
	2, // 2: auth.Auth.GetUserInfo:input_type -> auth.GetUserInfoReq
	0, // 3: auth.Auth.CheckIfAdmin:input_type -> auth.CheckIfAdminReq
	4, // 4: auth.Auth.Signup:output_type -> auth.SignupRes
	6, // 5: auth.Auth.Login:output_type -> auth.LoginRes
	3, // 6: auth.Auth.GetUserInfo:output_type -> auth.GetUserInfoRes
	1, // 7: auth.Auth.CheckIfAdmin:output_type -> auth.CheckIfAdminRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_auth_proto_init() }
func file_protos_auth_proto_init() {
	if File_protos_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfAdminReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckIfAdminRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_auth_proto_goTypes,
		DependencyIndexes: file_protos_auth_proto_depIdxs,
		MessageInfos:      file_protos_auth_proto_msgTypes,
	}.Build()
	File_protos_auth_proto = out.File
	file_protos_auth_proto_rawDesc = nil
	file_protos_auth_proto_goTypes = nil
	file_protos_auth_proto_depIdxs = nil
}
