// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: nika_application.proto

package nika_application

import (
	context "context"
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

type CreateApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Git           string         `protobuf:"bytes,1,opt,name=git,proto3" json:"git,omitempty"`
	Type          uint32         `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Status        uint32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	UserIDList    []uint64       `protobuf:"varint,4,rep,packed,name=userIDList,proto3" json:"userIDList,omitempty"`
	ProjectIDList []uint64       `protobuf:"varint,5,rep,packed,name=projectIDList,proto3" json:"projectIDList,omitempty"`
	Role          string         `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	TenantName    string         `protobuf:"bytes,7,opt,name=tenantName,proto3" json:"tenantName,omitempty"`
	Description   string         `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Environments  []*Environment `protobuf:"bytes,9,rep,name=Environments,proto3" json:"Environments,omitempty"`
}

func (x *CreateApplicationRequest) Reset() {
	*x = CreateApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationRequest) ProtoMessage() {}

func (x *CreateApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationRequest.ProtoReflect.Descriptor instead.
func (*CreateApplicationRequest) Descriptor() ([]byte, []int) {
	return file_nika_application_proto_rawDescGZIP(), []int{0}
}

func (x *CreateApplicationRequest) GetGit() string {
	if x != nil {
		return x.Git
	}
	return ""
}

func (x *CreateApplicationRequest) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateApplicationRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateApplicationRequest) GetUserIDList() []uint64 {
	if x != nil {
		return x.UserIDList
	}
	return nil
}

func (x *CreateApplicationRequest) GetProjectIDList() []uint64 {
	if x != nil {
		return x.ProjectIDList
	}
	return nil
}

func (x *CreateApplicationRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *CreateApplicationRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *CreateApplicationRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateApplicationRequest) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

type CreateApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppID uint64 `protobuf:"varint,1,opt,name=appID,proto3" json:"appID,omitempty"`
}

func (x *CreateApplicationResponse) Reset() {
	*x = CreateApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApplicationResponse) ProtoMessage() {}

func (x *CreateApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApplicationResponse.ProtoReflect.Descriptor instead.
func (*CreateApplicationResponse) Descriptor() ([]byte, []int) {
	return file_nika_application_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApplicationResponse) GetAppID() uint64 {
	if x != nil {
		return x.AppID
	}
	return 0
}

type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repository string `protobuf:"bytes,1,opt,name=Repository,proto3" json:"Repository,omitempty"`
	URL        string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Version    string `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *Chart) Reset() {
	*x = Chart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_nika_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_nika_application_proto_rawDescGZIP(), []int{2}
}

func (x *Chart) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *Chart) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Chart) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Chart) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Environment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value         []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Namespace     string `protobuf:"bytes,3,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Release       string `protobuf:"bytes,4,opt,name=Release,proto3" json:"Release,omitempty"`
	ApiServer     string `protobuf:"bytes,5,opt,name=ApiServer,proto3" json:"ApiServer,omitempty"`
	Token         string `protobuf:"bytes,6,opt,name=Token,proto3" json:"Token,omitempty"`
	CA            string `protobuf:"bytes,7,opt,name=CA,proto3" json:"CA,omitempty"`
	Chart         *Chart `protobuf:"bytes,8,opt,name=Chart,proto3" json:"Chart,omitempty"`
	ValueTemplate []byte `protobuf:"bytes,9,opt,name=ValueTemplate,proto3" json:"ValueTemplate,omitempty"`
}

func (x *Environment) Reset() {
	*x = Environment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Environment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Environment) ProtoMessage() {}

func (x *Environment) ProtoReflect() protoreflect.Message {
	mi := &file_nika_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Environment.ProtoReflect.Descriptor instead.
func (*Environment) Descriptor() ([]byte, []int) {
	return file_nika_application_proto_rawDescGZIP(), []int{3}
}

func (x *Environment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Environment) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Environment) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Environment) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *Environment) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *Environment) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Environment) GetCA() string {
	if x != nil {
		return x.CA
	}
	return ""
}

func (x *Environment) GetChart() *Chart {
	if x != nil {
		return x.Chart
	}
	return nil
}

func (x *Environment) GetValueTemplate() []byte {
	if x != nil {
		return x.ValueTemplate
	}
	return nil
}

var File_nika_application_proto protoreflect.FileDescriptor

var file_nika_application_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x18, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c, 0x69, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x0c, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x31, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x70, 0x70, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61,
	0x70, 0x70, 0x49, 0x44, 0x22, 0x67, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xf7, 0x01,
	0x0a, 0x0b, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x41, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x43, 0x41, 0x12, 0x1c, 0x0a, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x06, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x05, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x32, 0x5f, 0x0a, 0x0f, 0x4e, 0x69, 0x6b, 0x61, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x64, 0x5a, 0x62, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6f, 0x6c, 0x75, 0x6e, 0x6c, 0x2f, 0x6e,
	0x69, 0x6b, 0x61, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2d, 0x62, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x69,
	0x6b, 0x61, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nika_application_proto_rawDescOnce sync.Once
	file_nika_application_proto_rawDescData = file_nika_application_proto_rawDesc
)

func file_nika_application_proto_rawDescGZIP() []byte {
	file_nika_application_proto_rawDescOnce.Do(func() {
		file_nika_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_nika_application_proto_rawDescData)
	})
	return file_nika_application_proto_rawDescData
}

var file_nika_application_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nika_application_proto_goTypes = []interface{}{
	(*CreateApplicationRequest)(nil),  // 0: CreateApplicationRequest
	(*CreateApplicationResponse)(nil), // 1: CreateApplicationResponse
	(*Chart)(nil),                     // 2: Chart
	(*Environment)(nil),               // 3: Environment
}
var file_nika_application_proto_depIdxs = []int32{
	3, // 0: CreateApplicationRequest.Environments:type_name -> Environment
	2, // 1: Environment.Chart:type_name -> Chart
	0, // 2: NikaApplication.CreateApplication:input_type -> CreateApplicationRequest
	1, // 3: NikaApplication.CreateApplication:output_type -> CreateApplicationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nika_application_proto_init() }
func file_nika_application_proto_init() {
	if File_nika_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nika_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationRequest); i {
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
		file_nika_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApplicationResponse); i {
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
		file_nika_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chart); i {
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
		file_nika_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Environment); i {
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
			RawDescriptor: file_nika_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nika_application_proto_goTypes,
		DependencyIndexes: file_nika_application_proto_depIdxs,
		MessageInfos:      file_nika_application_proto_msgTypes,
	}.Build()
	File_nika_application_proto = out.File
	file_nika_application_proto_rawDesc = nil
	file_nika_application_proto_goTypes = nil
	file_nika_application_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.2.1. DO NOT EDIT.

type NikaApplication interface {
	CreateApplication(ctx context.Context, req *CreateApplicationRequest) (res *CreateApplicationResponse, err error)
}
