// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/ocp-user-api/ocp-user-api.proto

package ocp_user_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListUsersV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListUsersV1Request) Reset() {
	*x = ListUsersV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersV1Request) ProtoMessage() {}

func (x *ListUsersV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersV1Request.ProtoReflect.Descriptor instead.
func (*ListUsersV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{0}
}

func (x *ListUsersV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUsersV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListUsersV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersV1Response) Reset() {
	*x = ListUsersV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersV1Response) ProtoMessage() {}

func (x *ListUsersV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersV1Response.ProtoReflect.Descriptor instead.
func (*ListUsersV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListUsersV1Response) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type CreateUserV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CalendarId uint64       `protobuf:"varint,2,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	ResumeId   uint64       `protobuf:"varint,3,opt,name=resumeId,proto3" json:"resumeId,omitempty"`
	Profile    *UserProfile `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CreateUserV1Request) Reset() {
	*x = CreateUserV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserV1Request) ProtoMessage() {}

func (x *CreateUserV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserV1Request.ProtoReflect.Descriptor instead.
func (*CreateUserV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserV1Request) GetCalendarId() uint64 {
	if x != nil {
		return x.CalendarId
	}
	return 0
}

func (x *CreateUserV1Request) GetResumeId() uint64 {
	if x != nil {
		return x.ResumeId
	}
	return 0
}

func (x *CreateUserV1Request) GetProfile() *UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

type CreateUserV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *CreateUserV1Response) Reset() {
	*x = CreateUserV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserV1Response) ProtoMessage() {}

func (x *CreateUserV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserV1Response.ProtoReflect.Descriptor instead.
func (*CreateUserV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserV1Response) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveUserV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *RemoveUserV1Request) Reset() {
	*x = RemoveUserV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserV1Request) ProtoMessage() {}

func (x *RemoveUserV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserV1Request.ProtoReflect.Descriptor instead.
func (*RemoveUserV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveUserV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type RemoveUserV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *RemoveUserV1Response) Reset() {
	*x = RemoveUserV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserV1Response) ProtoMessage() {}

func (x *RemoveUserV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserV1Response.ProtoReflect.Descriptor instead.
func (*RemoveUserV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{5}
}

func (x *RemoveUserV1Response) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type DescribeUserV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *DescribeUserV1Request) Reset() {
	*x = DescribeUserV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserV1Request) ProtoMessage() {}

func (x *DescribeUserV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserV1Request.ProtoReflect.Descriptor instead.
func (*DescribeUserV1Request) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{6}
}

func (x *DescribeUserV1Request) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DescribeUserV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *DescribeUserV1Response) Reset() {
	*x = DescribeUserV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeUserV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeUserV1Response) ProtoMessage() {}

func (x *DescribeUserV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeUserV1Response.ProtoReflect.Descriptor instead.
func (*DescribeUserV1Response) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{7}
}

func (x *DescribeUserV1Response) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname    string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Patronymic string `protobuf:"bytes,3,opt,name=patronymic,proto3" json:"patronymic,omitempty"`
	Email      string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{8}
}

func (x *UserProfile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserProfile) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *UserProfile) GetPatronymic() string {
	if x != nil {
		return x.Patronymic
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CalendarId uint64       `protobuf:"varint,2,opt,name=calendarId,proto3" json:"calendarId,omitempty"`
	ResumeId   uint64       `protobuf:"varint,3,opt,name=resumeId,proto3" json:"resumeId,omitempty"`
	Profile    *UserProfile `protobuf:"bytes,4,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_ocp_user_api_ocp_user_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetCalendarId() uint64 {
	if x != nil {
		return x.CalendarId
	}
	return 0
}

func (x *User) GetResumeId() uint64 {
	if x != nil {
		return x.ResumeId
	}
	return 0
}

func (x *User) GetProfile() *UserProfile {
	if x != nil {
		return x.Profile
	}
	return nil
}

var File_api_ocp_user_api_ocp_user_api_proto protoreflect.FileDescriptor

var file_api_ocp_user_api_ocp_user_api_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x36, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x38, 0x0a, 0x15,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x16, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f, 0x6e, 0x79,
	0x6d, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x72, 0x6f,
	0x6e, 0x79, 0x6d, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x90, 0x01, 0x0a, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x32, 0xc9,
	0x03, 0x0a, 0x0a, 0x4f, 0x63, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x70, 0x69, 0x12, 0x65, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x56, 0x31, 0x12, 0x20, 0x2e, 0x6f,
	0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x12, 0x77, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f, 0x63,
	0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x12, 0x68, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x21, 0x2e,
	0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x71, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x7d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x68, 0x74, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63,
	0x70, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_ocp_user_api_ocp_user_api_proto_rawDescOnce sync.Once
	file_api_ocp_user_api_ocp_user_api_proto_rawDescData = file_api_ocp_user_api_ocp_user_api_proto_rawDesc
)

func file_api_ocp_user_api_ocp_user_api_proto_rawDescGZIP() []byte {
	file_api_ocp_user_api_ocp_user_api_proto_rawDescOnce.Do(func() {
		file_api_ocp_user_api_ocp_user_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ocp_user_api_ocp_user_api_proto_rawDescData)
	})
	return file_api_ocp_user_api_ocp_user_api_proto_rawDescData
}

var file_api_ocp_user_api_ocp_user_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_ocp_user_api_ocp_user_api_proto_goTypes = []interface{}{
	(*ListUsersV1Request)(nil),     // 0: ocp.user.api.ListUsersV1Request
	(*ListUsersV1Response)(nil),    // 1: ocp.user.api.ListUsersV1Response
	(*CreateUserV1Request)(nil),    // 2: ocp.user.api.CreateUserV1Request
	(*CreateUserV1Response)(nil),   // 3: ocp.user.api.CreateUserV1Response
	(*RemoveUserV1Request)(nil),    // 4: ocp.user.api.RemoveUserV1Request
	(*RemoveUserV1Response)(nil),   // 5: ocp.user.api.RemoveUserV1Response
	(*DescribeUserV1Request)(nil),  // 6: ocp.user.api.DescribeUserV1Request
	(*DescribeUserV1Response)(nil), // 7: ocp.user.api.DescribeUserV1Response
	(*UserProfile)(nil),            // 8: ocp.user.api.UserProfile
	(*User)(nil),                   // 9: ocp.user.api.User
}
var file_api_ocp_user_api_ocp_user_api_proto_depIdxs = []int32{
	9, // 0: ocp.user.api.ListUsersV1Response.users:type_name -> ocp.user.api.User
	8, // 1: ocp.user.api.CreateUserV1Request.profile:type_name -> ocp.user.api.UserProfile
	9, // 2: ocp.user.api.DescribeUserV1Response.user:type_name -> ocp.user.api.User
	8, // 3: ocp.user.api.User.profile:type_name -> ocp.user.api.UserProfile
	0, // 4: ocp.user.api.OcpUserApi.ListUsersV1:input_type -> ocp.user.api.ListUsersV1Request
	6, // 5: ocp.user.api.OcpUserApi.DescribeUserV1:input_type -> ocp.user.api.DescribeUserV1Request
	2, // 6: ocp.user.api.OcpUserApi.CreateUserV1:input_type -> ocp.user.api.CreateUserV1Request
	4, // 7: ocp.user.api.OcpUserApi.RemoveUserV1:input_type -> ocp.user.api.RemoveUserV1Request
	1, // 8: ocp.user.api.OcpUserApi.ListUsersV1:output_type -> ocp.user.api.ListUsersV1Response
	7, // 9: ocp.user.api.OcpUserApi.DescribeUserV1:output_type -> ocp.user.api.DescribeUserV1Response
	3, // 10: ocp.user.api.OcpUserApi.CreateUserV1:output_type -> ocp.user.api.CreateUserV1Response
	5, // 11: ocp.user.api.OcpUserApi.RemoveUserV1:output_type -> ocp.user.api.RemoveUserV1Response
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_ocp_user_api_ocp_user_api_proto_init() }
func file_api_ocp_user_api_ocp_user_api_proto_init() {
	if File_api_ocp_user_api_ocp_user_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersV1Request); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersV1Response); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserV1Request); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserV1Response); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserV1Request); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserV1Response); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserV1Request); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeUserV1Response); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_api_ocp_user_api_ocp_user_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_api_ocp_user_api_ocp_user_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ocp_user_api_ocp_user_api_proto_goTypes,
		DependencyIndexes: file_api_ocp_user_api_ocp_user_api_proto_depIdxs,
		MessageInfos:      file_api_ocp_user_api_ocp_user_api_proto_msgTypes,
	}.Build()
	File_api_ocp_user_api_ocp_user_api_proto = out.File
	file_api_ocp_user_api_ocp_user_api_proto_rawDesc = nil
	file_api_ocp_user_api_ocp_user_api_proto_goTypes = nil
	file_api_ocp_user_api_ocp_user_api_proto_depIdxs = nil
}
