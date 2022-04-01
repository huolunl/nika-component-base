// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: nika_cluster.proto

package nika_cluster

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

type CreateClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ApiServer   string `protobuf:"bytes,2,opt,name=apiServer,proto3" json:"apiServer,omitempty"`
	Tenant      string `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CA          string `protobuf:"bytes,5,opt,name=CA,proto3" json:"CA,omitempty"`
	Status      uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Name        string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateClusterRequest) Reset() {
	*x = CreateClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterRequest) ProtoMessage() {}

func (x *CreateClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterRequest.ProtoReflect.Descriptor instead.
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return file_nika_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *CreateClusterRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateClusterRequest) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *CreateClusterRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CreateClusterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateClusterRequest) GetCA() string {
	if x != nil {
		return x.CA
	}
	return ""
}

func (x *CreateClusterRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateClusterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID uint64 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *CreateClusterResponse) Reset() {
	*x = CreateClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateClusterResponse) ProtoMessage() {}

func (x *CreateClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateClusterResponse.ProtoReflect.Descriptor instead.
func (*CreateClusterResponse) Descriptor() ([]byte, []int) {
	return file_nika_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *CreateClusterResponse) GetClusterID() uint64 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

type GetClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterID uint64 `protobuf:"varint,1,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
}

func (x *GetClusterRequest) Reset() {
	*x = GetClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterRequest) ProtoMessage() {}

func (x *GetClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nika_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterRequest.ProtoReflect.Descriptor instead.
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return file_nika_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *GetClusterRequest) GetClusterID() uint64 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

type GetClusterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ApiServer   string `protobuf:"bytes,2,opt,name=apiServer,proto3" json:"apiServer,omitempty"`
	Tenant      string `protobuf:"bytes,3,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CA          string `protobuf:"bytes,5,opt,name=CA,proto3" json:"CA,omitempty"`
	Status      uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Name        string `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	CreateAt    string `protobuf:"bytes,8,opt,name=createAt,proto3" json:"createAt,omitempty"`
	UpdatedAt   string `protobuf:"bytes,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
}

func (x *GetClusterResponse) Reset() {
	*x = GetClusterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nika_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterResponse) ProtoMessage() {}

func (x *GetClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nika_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterResponse.ProtoReflect.Descriptor instead.
func (*GetClusterResponse) Descriptor() ([]byte, []int) {
	return file_nika_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *GetClusterResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetClusterResponse) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *GetClusterResponse) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *GetClusterResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetClusterResponse) GetCA() string {
	if x != nil {
		return x.CA
	}
	return ""
}

func (x *GetClusterResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetClusterResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetClusterResponse) GetCreateAt() string {
	if x != nil {
		return x.CreateAt
	}
	return ""
}

func (x *GetClusterResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_nika_cluster_proto protoreflect.FileDescriptor

var file_nika_cluster_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x43,
	0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x43, 0x41, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x22, 0x31,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x22, 0xf8, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x41, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x43, 0x41, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x8b, 0x01, 0x0a,
	0x0b, 0x4e, 0x69, 0x6b, 0x61, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6f, 0x6c, 0x75, 0x6e, 0x6c,
	0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2d,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x61, 0x70,
	0x69, 0x2f, 0x6e, 0x69, 0x6b, 0x61, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x69, 0x6b, 0x61,
	0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nika_cluster_proto_rawDescOnce sync.Once
	file_nika_cluster_proto_rawDescData = file_nika_cluster_proto_rawDesc
)

func file_nika_cluster_proto_rawDescGZIP() []byte {
	file_nika_cluster_proto_rawDescOnce.Do(func() {
		file_nika_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_nika_cluster_proto_rawDescData)
	})
	return file_nika_cluster_proto_rawDescData
}

var file_nika_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nika_cluster_proto_goTypes = []interface{}{
	(*CreateClusterRequest)(nil),  // 0: CreateClusterRequest
	(*CreateClusterResponse)(nil), // 1: CreateClusterResponse
	(*GetClusterRequest)(nil),     // 2: GetClusterRequest
	(*GetClusterResponse)(nil),    // 3: GetClusterResponse
}
var file_nika_cluster_proto_depIdxs = []int32{
	0, // 0: NikaCluster.CreateCluster:input_type -> CreateClusterRequest
	0, // 1: NikaCluster.GetCluster:input_type -> CreateClusterRequest
	1, // 2: NikaCluster.CreateCluster:output_type -> CreateClusterResponse
	3, // 3: NikaCluster.GetCluster:output_type -> GetClusterResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nika_cluster_proto_init() }
func file_nika_cluster_proto_init() {
	if File_nika_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nika_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterRequest); i {
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
		file_nika_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateClusterResponse); i {
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
		file_nika_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterRequest); i {
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
		file_nika_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClusterResponse); i {
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
			RawDescriptor: file_nika_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nika_cluster_proto_goTypes,
		DependencyIndexes: file_nika_cluster_proto_depIdxs,
		MessageInfos:      file_nika_cluster_proto_msgTypes,
	}.Build()
	File_nika_cluster_proto = out.File
	file_nika_cluster_proto_rawDesc = nil
	file_nika_cluster_proto_goTypes = nil
	file_nika_cluster_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.2.1. DO NOT EDIT.

type NikaCluster interface {
	CreateCluster(ctx context.Context, req *CreateClusterRequest) (res *CreateClusterResponse, err error)
	GetCluster(ctx context.Context, req *CreateClusterRequest) (res *GetClusterResponse, err error)
}
