// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: nika_operator.proto

package nika_operator

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

type ExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenant             string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	UserName           string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserID             uint64 `protobuf:"varint,3,opt,name=userID,proto3" json:"userID,omitempty"`
	ApiServer          string `protobuf:"bytes,4,opt,name=ApiServer,proto3" json:"ApiServer,omitempty"`
	Token              string `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`
	CA                 string `protobuf:"bytes,6,opt,name=CA,proto3" json:"CA,omitempty"`
	HelmRepositoryName string `protobuf:"bytes,7,opt,name=HelmRepositoryName,proto3" json:"HelmRepositoryName,omitempty"`
	HelmRepositoryURL  string `protobuf:"bytes,8,opt,name=HelmRepositoryURL,proto3" json:"HelmRepositoryURL,omitempty"`
	Namespace          string `protobuf:"bytes,9,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Release            string `protobuf:"bytes,10,opt,name=Release,proto3" json:"Release,omitempty"`
	Chart              string `protobuf:"bytes,11,opt,name=Chart,proto3" json:"Chart,omitempty"`
	ReleaseValues      string `protobuf:"bytes,12,opt,name=ReleaseValues,proto3" json:"ReleaseValues,omitempty"`
	EnvName            string `protobuf:"bytes,13,opt,name=EnvName,proto3" json:"EnvName,omitempty"`
	EnvValues          string `protobuf:"bytes,14,opt,name=EnvValues,proto3" json:"EnvValues,omitempty"`
	Cmd                string `protobuf:"bytes,15,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Debug              bool   `protobuf:"varint,16,opt,name=Debug,proto3" json:"Debug,omitempty"`
	NickName           string `protobuf:"bytes,17,opt,name=nickName,proto3" json:"nickName,omitempty"`
	ApplicationName    string `protobuf:"bytes,18,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	ApplicationID      uint64 `protobuf:"varint,19,opt,name=applicationID,proto3" json:"applicationID,omitempty"`
	Version            string `protobuf:"bytes,20,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ExecRequest) Reset() {
	*x = ExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecRequest) ProtoMessage() {}

func (x *ExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecRequest.ProtoReflect.Descriptor instead.
func (*ExecRequest) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{0}
}

func (x *ExecRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *ExecRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ExecRequest) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ExecRequest) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *ExecRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ExecRequest) GetCA() string {
	if x != nil {
		return x.CA
	}
	return ""
}

func (x *ExecRequest) GetHelmRepositoryName() string {
	if x != nil {
		return x.HelmRepositoryName
	}
	return ""
}

func (x *ExecRequest) GetHelmRepositoryURL() string {
	if x != nil {
		return x.HelmRepositoryURL
	}
	return ""
}

func (x *ExecRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ExecRequest) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *ExecRequest) GetChart() string {
	if x != nil {
		return x.Chart
	}
	return ""
}

func (x *ExecRequest) GetReleaseValues() string {
	if x != nil {
		return x.ReleaseValues
	}
	return ""
}

func (x *ExecRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *ExecRequest) GetEnvValues() string {
	if x != nil {
		return x.EnvValues
	}
	return ""
}

func (x *ExecRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *ExecRequest) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *ExecRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *ExecRequest) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *ExecRequest) GetApplicationID() uint64 {
	if x != nil {
		return x.ApplicationID
	}
	return 0
}

func (x *ExecRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecRecord *GetExecRecordResponse `protobuf:"bytes,1,opt,name=ExecRecord,proto3" json:"ExecRecord,omitempty"`
}

func (x *ExecResponse) Reset() {
	*x = ExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecResponse) ProtoMessage() {}

func (x *ExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecResponse.ProtoReflect.Descriptor instead.
func (*ExecResponse) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{1}
}

func (x *ExecResponse) GetExecRecord() *GetExecRecordResponse {
	if x != nil {
		return x.ExecRecord
	}
	return nil
}

type GetExecRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetExecRecordRequest) Reset() {
	*x = GetExecRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecRecordRequest) ProtoMessage() {}

func (x *GetExecRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecRecordRequest.ProtoReflect.Descriptor instead.
func (*GetExecRecordRequest) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{2}
}

func (x *GetExecRecordRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetExecRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status             uint32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	ApiServer          string `protobuf:"bytes,3,opt,name=ApiServer,proto3" json:"ApiServer,omitempty"`
	Token              string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	CA                 string `protobuf:"bytes,5,opt,name=CA,proto3" json:"CA,omitempty"`
	HelmRepositoryName string `protobuf:"bytes,6,opt,name=HelmRepositoryName,proto3" json:"HelmRepositoryName,omitempty"`
	HelmRepositoryURL  string `protobuf:"bytes,7,opt,name=HelmRepositoryURL,proto3" json:"HelmRepositoryURL,omitempty"`
	Namespace          string `protobuf:"bytes,8,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
	Release            string `protobuf:"bytes,9,opt,name=Release,proto3" json:"Release,omitempty"`
	Chart              string `protobuf:"bytes,10,opt,name=Chart,proto3" json:"Chart,omitempty"`
	Version            string `protobuf:"bytes,11,opt,name=Version,proto3" json:"Version,omitempty"`
	ReleaseValues      string `protobuf:"bytes,12,opt,name=ReleaseValues,proto3" json:"ReleaseValues,omitempty"`
	EnvName            string `protobuf:"bytes,13,opt,name=EnvName,proto3" json:"EnvName,omitempty"`
	Cmd                string `protobuf:"bytes,14,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Log                string `protobuf:"bytes,15,opt,name=Log,proto3" json:"Log,omitempty"`
	Tenant             string `protobuf:"bytes,16,opt,name=Tenant,proto3" json:"Tenant,omitempty"`
	UserName           string `protobuf:"bytes,17,opt,name=UserName,proto3" json:"UserName,omitempty"`
	NickName           string `protobuf:"bytes,18,opt,name=NickName,proto3" json:"NickName,omitempty"`
	UserID             uint64 `protobuf:"varint,19,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ApplicationName    string `protobuf:"bytes,20,opt,name=ApplicationName,proto3" json:"ApplicationName,omitempty"`
	Debug              bool   `protobuf:"varint,21,opt,name=Debug,proto3" json:"Debug,omitempty"`
	ApplicationID      uint64 `protobuf:"varint,22,opt,name=ApplicationID,proto3" json:"ApplicationID,omitempty"`
	EnvValues          string `protobuf:"bytes,23,opt,name=EnvValues,proto3" json:"EnvValues,omitempty"`
}

func (x *GetExecRecordResponse) Reset() {
	*x = GetExecRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExecRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExecRecordResponse) ProtoMessage() {}

func (x *GetExecRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExecRecordResponse.ProtoReflect.Descriptor instead.
func (*GetExecRecordResponse) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{3}
}

func (x *GetExecRecordResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetExecRecordResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetExecRecordResponse) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *GetExecRecordResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetExecRecordResponse) GetCA() string {
	if x != nil {
		return x.CA
	}
	return ""
}

func (x *GetExecRecordResponse) GetHelmRepositoryName() string {
	if x != nil {
		return x.HelmRepositoryName
	}
	return ""
}

func (x *GetExecRecordResponse) GetHelmRepositoryURL() string {
	if x != nil {
		return x.HelmRepositoryURL
	}
	return ""
}

func (x *GetExecRecordResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetExecRecordResponse) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *GetExecRecordResponse) GetChart() string {
	if x != nil {
		return x.Chart
	}
	return ""
}

func (x *GetExecRecordResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetExecRecordResponse) GetReleaseValues() string {
	if x != nil {
		return x.ReleaseValues
	}
	return ""
}

func (x *GetExecRecordResponse) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *GetExecRecordResponse) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *GetExecRecordResponse) GetLog() string {
	if x != nil {
		return x.Log
	}
	return ""
}

func (x *GetExecRecordResponse) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *GetExecRecordResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *GetExecRecordResponse) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *GetExecRecordResponse) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetExecRecordResponse) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *GetExecRecordResponse) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *GetExecRecordResponse) GetApplicationID() uint64 {
	if x != nil {
		return x.ApplicationID
	}
	return 0
}

func (x *GetExecRecordResponse) GetEnvValues() string {
	if x != nil {
		return x.EnvValues
	}
	return ""
}

type ListExecRecordByAppIDAndEnvNameAndCmdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	EnvName string `protobuf:"bytes,2,opt,name=EnvName,proto3" json:"EnvName,omitempty"`
	Cmd     string `protobuf:"bytes,3,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	Offset  int64  `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit   int64  `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) Reset() {
	*x = ListExecRecordByAppIDAndEnvNameAndCmdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExecRecordByAppIDAndEnvNameAndCmdRequest) ProtoMessage() {}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExecRecordByAppIDAndEnvNameAndCmdRequest.ProtoReflect.Descriptor instead.
func (*ListExecRecordByAppIDAndEnvNameAndCmdRequest) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{4}
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) GetEnvName() string {
	if x != nil {
		return x.EnvName
	}
	return ""
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListExecRecordByAppIDAndEnvNameAndCmdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*GetExecRecordResponse `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdResponse) Reset() {
	*x = ListExecRecordByAppIDAndEnvNameAndCmdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_operator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListExecRecordByAppIDAndEnvNameAndCmdResponse) ProtoMessage() {}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_operator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListExecRecordByAppIDAndEnvNameAndCmdResponse.ProtoReflect.Descriptor instead.
func (*ListExecRecordByAppIDAndEnvNameAndCmdResponse) Descriptor() ([]byte, []int) {
	return file_nika_operator_proto_rawDescGZIP(), []int{5}
}

func (x *ListExecRecordByAppIDAndEnvNameAndCmdResponse) GetRecords() []*GetExecRecordResponse {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_nika_operator_proto protoreflect.FileDescriptor

var file_nika_operator_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x04, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x43, 0x41, 0x12, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x55, 0x52, 0x4c, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x6d,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a,
	0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a,
	0x0c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a,
	0x0a, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x99, 0x05,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x43, 0x41, 0x12, 0x2e, 0x0a, 0x12, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x52, 0x4c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x48, 0x65, 0x6c, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x52,
	0x4c, 0x12, 0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x54,
	0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x44, 0x65, 0x62, 0x75, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x16, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x45,
	0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x2c, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64,
	0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e,
	0x76, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x76,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x61, 0x0a, 0x2d, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44, 0x41, 0x6e, 0x64,
	0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x32, 0x82, 0x02, 0x0a, 0x0c, 0x4e, 0x69, 0x6b, 0x61,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63,
	0x12, 0x0c, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x65,
	0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x88, 0x01, 0x0a, 0x25, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x41, 0x70, 0x70, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x45, 0x6e,
	0x76, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x43, 0x6d, 0x64, 0x12, 0x2d, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x41, 0x70,
	0x70, 0x49, 0x44, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64,
	0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x79, 0x41, 0x70, 0x70,
	0x49, 0x44, 0x41, 0x6e, 0x64, 0x45, 0x6e, 0x76, 0x4e, 0x61, 0x6d, 0x65, 0x41, 0x6e, 0x64, 0x43,
	0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5e, 0x5a, 0x5c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6f, 0x6c, 0x75,
	0x6e, 0x6c, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2d, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6e,
	0x69, 0x6b, 0x61, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nika_operator_proto_rawDescOnce sync.Once
	file_nika_operator_proto_rawDescData = file_nika_operator_proto_rawDesc
)

func file_nika_operator_proto_rawDescGZIP() []byte {
	file_nika_operator_proto_rawDescOnce.Do(func() {
		file_nika_operator_proto_rawDescData = protoimpl.X.CompressGZIP(file_nika_operator_proto_rawDescData)
	})
	return file_nika_operator_proto_rawDescData
}

var file_nika_operator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nika_operator_proto_goTypes = []interface{}{
	(*ExecRequest)(nil),                                   // 0: ExecRequest
	(*ExecResponse)(nil),                                  // 1: ExecResponse
	(*GetExecRecordRequest)(nil),                          // 2: GetExecRecordRequest
	(*GetExecRecordResponse)(nil),                         // 3: GetExecRecordResponse
	(*ListExecRecordByAppIDAndEnvNameAndCmdRequest)(nil),  // 4: ListExecRecordByAppIDAndEnvNameAndCmdRequest
	(*ListExecRecordByAppIDAndEnvNameAndCmdResponse)(nil), // 5: ListExecRecordByAppIDAndEnvNameAndCmdResponse
}
var file_nika_operator_proto_depIdxs = []int32{
	3, // 0: ExecResponse.ExecRecord:type_name -> GetExecRecordResponse
	3, // 1: ListExecRecordByAppIDAndEnvNameAndCmdResponse.records:type_name -> GetExecRecordResponse
	0, // 2: NikaOperator.Exec:input_type -> ExecRequest
	2, // 3: NikaOperator.GetExecRecord:input_type -> GetExecRecordRequest
	4, // 4: NikaOperator.ListExecRecordByAppIDAndEnvNameAndCmd:input_type -> ListExecRecordByAppIDAndEnvNameAndCmdRequest
	1, // 5: NikaOperator.Exec:output_type -> ExecResponse
	3, // 6: NikaOperator.GetExecRecord:output_type -> GetExecRecordResponse
	5, // 7: NikaOperator.ListExecRecordByAppIDAndEnvNameAndCmd:output_type -> ListExecRecordByAppIDAndEnvNameAndCmdResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_nika_operator_proto_init() }
func file_nika_operator_proto_init() {
	if File_nika_operator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nika_operator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecRequest); i {
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
		file_nika_operator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecResponse); i {
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
		file_nika_operator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecRecordRequest); i {
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
		file_nika_operator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExecRecordResponse); i {
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
		file_nika_operator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExecRecordByAppIDAndEnvNameAndCmdRequest); i {
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
		file_nika_operator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListExecRecordByAppIDAndEnvNameAndCmdResponse); i {
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
			RawDescriptor: file_nika_operator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nika_operator_proto_goTypes,
		DependencyIndexes: file_nika_operator_proto_depIdxs,
		MessageInfos:      file_nika_operator_proto_msgTypes,
	}.Build()
	File_nika_operator_proto = out.File
	file_nika_operator_proto_rawDesc = nil
	file_nika_operator_proto_goTypes = nil
	file_nika_operator_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.2.1. DO NOT EDIT.

type NikaOperator interface {
	Exec(ctx context.Context, req *ExecRequest) (res *ExecResponse, err error)
	GetExecRecord(ctx context.Context, req *GetExecRecordRequest) (res *GetExecRecordResponse, err error)
	ListExecRecordByAppIDAndEnvNameAndCmd(ctx context.Context, req *ListExecRecordByAppIDAndEnvNameAndCmdRequest) (res *ListExecRecordByAppIDAndEnvNameAndCmdResponse, err error)
}
