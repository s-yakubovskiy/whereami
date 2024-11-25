// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: whrmi.proto

package whrmi

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

type ShowLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ShowLocationRequest) Reset() {
	*x = ShowLocationRequest{}
	mi := &file_whrmi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShowLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowLocationRequest) ProtoMessage() {}

func (x *ShowLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowLocationRequest.ProtoReflect.Descriptor instead.
func (*ShowLocationRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{0}
}

func (x *ShowLocationRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ShowLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ShowLocationResponse) Reset() {
	*x = ShowLocationResponse{}
	mi := &file_whrmi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShowLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShowLocationResponse) ProtoMessage() {}

func (x *ShowLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShowLocationResponse.ProtoReflect.Descriptor instead.
func (*ShowLocationResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{1}
}

func (x *ShowLocationResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	mi := &file_whrmi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{2}
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	mi := &file_whrmi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{3}
}

type AddVpnInterfaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vpninterface string `protobuf:"bytes,1,opt,name=vpninterface,proto3" json:"vpninterface,omitempty"`
}

func (x *AddVpnInterfaceRequest) Reset() {
	*x = AddVpnInterfaceRequest{}
	mi := &file_whrmi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddVpnInterfaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVpnInterfaceRequest) ProtoMessage() {}

func (x *AddVpnInterfaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVpnInterfaceRequest.ProtoReflect.Descriptor instead.
func (*AddVpnInterfaceRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{4}
}

func (x *AddVpnInterfaceRequest) GetVpninterface() string {
	if x != nil {
		return x.Vpninterface
	}
	return ""
}

type AddVpnInterfaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddVpnInterfaceResponse) Reset() {
	*x = AddVpnInterfaceResponse{}
	mi := &file_whrmi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddVpnInterfaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVpnInterfaceResponse) ProtoMessage() {}

func (x *AddVpnInterfaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVpnInterfaceResponse.ProtoReflect.Descriptor instead.
func (*AddVpnInterfaceResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{5}
}

type ListVpnInterfacesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVpnInterfacesRequest) Reset() {
	*x = ListVpnInterfacesRequest{}
	mi := &file_whrmi_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListVpnInterfacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVpnInterfacesRequest) ProtoMessage() {}

func (x *ListVpnInterfacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVpnInterfacesRequest.ProtoReflect.Descriptor instead.
func (*ListVpnInterfacesRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{6}
}

type ListVpnInterfacesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vpninterfaces []string `protobuf:"bytes,1,rep,name=vpninterfaces,proto3" json:"vpninterfaces,omitempty"`
}

func (x *ListVpnInterfacesResponse) Reset() {
	*x = ListVpnInterfacesResponse{}
	mi := &file_whrmi_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListVpnInterfacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVpnInterfacesResponse) ProtoMessage() {}

func (x *ListVpnInterfacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVpnInterfacesResponse.ProtoReflect.Descriptor instead.
func (*ListVpnInterfacesResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{7}
}

func (x *ListVpnInterfacesResponse) GetVpninterfaces() []string {
	if x != nil {
		return x.Vpninterfaces
	}
	return nil
}

type ExportLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exportpath string `protobuf:"bytes,1,opt,name=exportpath,proto3" json:"exportpath,omitempty"`
}

func (x *ExportLocationsRequest) Reset() {
	*x = ExportLocationsRequest{}
	mi := &file_whrmi_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportLocationsRequest) ProtoMessage() {}

func (x *ExportLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportLocationsRequest.ProtoReflect.Descriptor instead.
func (*ExportLocationsRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{8}
}

func (x *ExportLocationsRequest) GetExportpath() string {
	if x != nil {
		return x.Exportpath
	}
	return ""
}

type ExportLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ExportLocationsResponse) Reset() {
	*x = ExportLocationsResponse{}
	mi := &file_whrmi_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportLocationsResponse) ProtoMessage() {}

func (x *ExportLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportLocationsResponse.ProtoReflect.Descriptor instead.
func (*ExportLocationsResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{9}
}

type ImportLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Importpath string `protobuf:"bytes,1,opt,name=importpath,proto3" json:"importpath,omitempty"`
}

func (x *ImportLocationsRequest) Reset() {
	*x = ImportLocationsRequest{}
	mi := &file_whrmi_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportLocationsRequest) ProtoMessage() {}

func (x *ImportLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportLocationsRequest.ProtoReflect.Descriptor instead.
func (*ImportLocationsRequest) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{10}
}

func (x *ImportLocationsRequest) GetImportpath() string {
	if x != nil {
		return x.Importpath
	}
	return ""
}

type ImportLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ImportLocationsResponse) Reset() {
	*x = ImportLocationsResponse{}
	mi := &file_whrmi_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportLocationsResponse) ProtoMessage() {}

func (x *ImportLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_whrmi_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportLocationsResponse.ProtoReflect.Descriptor instead.
func (*ImportLocationsResponse) Descriptor() ([]byte, []int) {
	return file_whrmi_proto_rawDescGZIP(), []int{11}
}

var File_whrmi_proto protoreflect.FileDescriptor

var file_whrmi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61,
	0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x2d, 0x0a, 0x13, 0x53,
	0x68, 0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x68,
	0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x49, 0x6e, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x41, 0x64, 0x64,
	0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x76, 0x70, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x70, 0x6e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x56, 0x70,
	0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41,
	0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x76,
	0x70, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x76, 0x70, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x22, 0x38, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x74, 0x68, 0x22, 0x19, 0x0a, 0x17, 0x45,
	0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x0a, 0x16, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x61, 0x74, 0x68,
	0x22, 0x19, 0x0a, 0x17, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x68, 0x0a, 0x0f, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55,
	0x0a, 0x0c, 0x53, 0x68, 0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68,
	0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xdc, 0x03, 0x0a, 0x15, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3d, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68,
	0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68,
	0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x70, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68,
	0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5e, 0x0a, 0x0f, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68,
	0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x77, 0x68, 0x72, 0x6d,
	0x69, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x2d, 0x79, 0x61, 0x6b, 0x75, 0x62, 0x6f, 0x76, 0x73, 0x6b, 0x69, 0x79,
	0x2f, 0x77, 0x68, 0x65, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x68,
	0x72, 0x6d, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x68, 0x72, 0x6d, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_whrmi_proto_rawDescOnce sync.Once
	file_whrmi_proto_rawDescData = file_whrmi_proto_rawDesc
)

func file_whrmi_proto_rawDescGZIP() []byte {
	file_whrmi_proto_rawDescOnce.Do(func() {
		file_whrmi_proto_rawDescData = protoimpl.X.CompressGZIP(file_whrmi_proto_rawDescData)
	})
	return file_whrmi_proto_rawDescData
}

var file_whrmi_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_whrmi_proto_goTypes = []any{
	(*ShowLocationRequest)(nil),       // 0: api.whrmi.v1.ShowLocationRequest
	(*ShowLocationResponse)(nil),      // 1: api.whrmi.v1.ShowLocationResponse
	(*InitRequest)(nil),               // 2: api.whrmi.v1.InitRequest
	(*InitResponse)(nil),              // 3: api.whrmi.v1.InitResponse
	(*AddVpnInterfaceRequest)(nil),    // 4: api.whrmi.v1.AddVpnInterfaceRequest
	(*AddVpnInterfaceResponse)(nil),   // 5: api.whrmi.v1.AddVpnInterfaceResponse
	(*ListVpnInterfacesRequest)(nil),  // 6: api.whrmi.v1.ListVpnInterfacesRequest
	(*ListVpnInterfacesResponse)(nil), // 7: api.whrmi.v1.ListVpnInterfacesResponse
	(*ExportLocationsRequest)(nil),    // 8: api.whrmi.v1.ExportLocationsRequest
	(*ExportLocationsResponse)(nil),   // 9: api.whrmi.v1.ExportLocationsResponse
	(*ImportLocationsRequest)(nil),    // 10: api.whrmi.v1.ImportLocationsRequest
	(*ImportLocationsResponse)(nil),   // 11: api.whrmi.v1.ImportLocationsResponse
}
var file_whrmi_proto_depIdxs = []int32{
	0,  // 0: api.whrmi.v1.LocationService.ShowLocation:input_type -> api.whrmi.v1.ShowLocationRequest
	2,  // 1: api.whrmi.v1.LocationKeeperService.Init:input_type -> api.whrmi.v1.InitRequest
	4,  // 2: api.whrmi.v1.LocationKeeperService.AddVpnInterface:input_type -> api.whrmi.v1.AddVpnInterfaceRequest
	6,  // 3: api.whrmi.v1.LocationKeeperService.ListVpnInterfaces:input_type -> api.whrmi.v1.ListVpnInterfacesRequest
	8,  // 4: api.whrmi.v1.LocationKeeperService.ExportLocations:input_type -> api.whrmi.v1.ExportLocationsRequest
	10, // 5: api.whrmi.v1.LocationKeeperService.ImportLocations:input_type -> api.whrmi.v1.ImportLocationsRequest
	1,  // 6: api.whrmi.v1.LocationService.ShowLocation:output_type -> api.whrmi.v1.ShowLocationResponse
	3,  // 7: api.whrmi.v1.LocationKeeperService.Init:output_type -> api.whrmi.v1.InitResponse
	5,  // 8: api.whrmi.v1.LocationKeeperService.AddVpnInterface:output_type -> api.whrmi.v1.AddVpnInterfaceResponse
	7,  // 9: api.whrmi.v1.LocationKeeperService.ListVpnInterfaces:output_type -> api.whrmi.v1.ListVpnInterfacesResponse
	9,  // 10: api.whrmi.v1.LocationKeeperService.ExportLocations:output_type -> api.whrmi.v1.ExportLocationsResponse
	11, // 11: api.whrmi.v1.LocationKeeperService.ImportLocations:output_type -> api.whrmi.v1.ImportLocationsResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_whrmi_proto_init() }
func file_whrmi_proto_init() {
	if File_whrmi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_whrmi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_whrmi_proto_goTypes,
		DependencyIndexes: file_whrmi_proto_depIdxs,
		MessageInfos:      file_whrmi_proto_msgTypes,
	}.Build()
	File_whrmi_proto = out.File
	file_whrmi_proto_rawDesc = nil
	file_whrmi_proto_goTypes = nil
	file_whrmi_proto_depIdxs = nil
}