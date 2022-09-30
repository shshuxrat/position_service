// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: attribute.proto

package position_service

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

type Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Attribute) Reset() {
	*x = Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Attribute) ProtoMessage() {}

func (x *Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Attribute.ProtoReflect.Descriptor instead.
func (*Attribute) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{0}
}

func (x *Attribute) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Attribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Attribute) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Attribute) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Attribute) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CreateAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateAttribute) Reset() {
	*x = CreateAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAttribute) ProtoMessage() {}

func (x *CreateAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAttribute.ProtoReflect.Descriptor instead.
func (*CreateAttribute) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAttribute) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAttribute) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type AttributeId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AttributeId) Reset() {
	*x = AttributeId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeId) ProtoMessage() {}

func (x *AttributeId) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeId.ProtoReflect.Descriptor instead.
func (*AttributeId) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{2}
}

func (x *AttributeId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAllAttributeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetAllAttributeRequest) Reset() {
	*x = GetAllAttributeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeRequest) ProtoMessage() {}

func (x *GetAllAttributeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeRequest.ProtoReflect.Descriptor instead.
func (*GetAllAttributeRequest) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllAttributeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllAttributeRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetAllAttributeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetAllAttributeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attributes []*Attribute `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	Count      int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllAttributeResponse) Reset() {
	*x = GetAllAttributeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAttributeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAttributeResponse) ProtoMessage() {}

func (x *GetAllAttributeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAttributeResponse.ProtoReflect.Descriptor instead.
func (*GetAllAttributeResponse) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllAttributeResponse) GetAttributes() []*Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *GetAllAttributeResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AttributeAfterUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Old *Attribute `protobuf:"bytes,1,opt,name=old,proto3" json:"old,omitempty"`
	New *Attribute `protobuf:"bytes,2,opt,name=new,proto3" json:"new,omitempty"`
}

func (x *AttributeAfterUpdate) Reset() {
	*x = AttributeAfterUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeAfterUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeAfterUpdate) ProtoMessage() {}

func (x *AttributeAfterUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeAfterUpdate.ProtoReflect.Descriptor instead.
func (*AttributeAfterUpdate) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{5}
}

func (x *AttributeAfterUpdate) GetOld() *Attribute {
	if x != nil {
		return x.Old
	}
	return nil
}

func (x *AttributeAfterUpdate) GetNew() *Attribute {
	if x != nil {
		return x.New
	}
	return nil
}

type AttributeRowsAffected struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RowsAffected int64 `protobuf:"varint,1,opt,name=rowsAffected,proto3" json:"rowsAffected,omitempty"`
}

func (x *AttributeRowsAffected) Reset() {
	*x = AttributeRowsAffected{}
	if protoimpl.UnsafeEnabled {
		mi := &file_attribute_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeRowsAffected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeRowsAffected) ProtoMessage() {}

func (x *AttributeRowsAffected) ProtoReflect() protoreflect.Message {
	mi := &file_attribute_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeRowsAffected.ProtoReflect.Descriptor instead.
func (*AttributeRowsAffected) Descriptor() ([]byte, []int) {
	return file_attribute_proto_rawDescGZIP(), []int{6}
}

func (x *AttributeRowsAffected) GetRowsAffected() int64 {
	if x != nil {
		return x.RowsAffected
	}
	return 0
}

var File_attribute_proto protoreflect.FileDescriptor

var file_attribute_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x09,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x39, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5a, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x64, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x33, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x64, 0x0a, 0x14, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x66, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x03, 0x6f, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x03, 0x6f, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x6e, 0x65,
	0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x03, 0x6e, 0x65,
	0x77, 0x22, 0x3b, 0x0a, 0x15, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x6f,
	0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0c, 0x72, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0xdf,
	0x02, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x22,
	0x00, 0x12, 0x4f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x49, 0x64, 0x1a, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x1e, 0x2e, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x64, 0x1a, 0x1f, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x52, 0x6f, 0x77, 0x73, 0x41, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x00,
	0x42, 0x1b, 0x5a, 0x19, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_attribute_proto_rawDescOnce sync.Once
	file_attribute_proto_rawDescData = file_attribute_proto_rawDesc
)

func file_attribute_proto_rawDescGZIP() []byte {
	file_attribute_proto_rawDescOnce.Do(func() {
		file_attribute_proto_rawDescData = protoimpl.X.CompressGZIP(file_attribute_proto_rawDescData)
	})
	return file_attribute_proto_rawDescData
}

var file_attribute_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_attribute_proto_goTypes = []interface{}{
	(*Attribute)(nil),               // 0: genproto.Attribute
	(*CreateAttribute)(nil),         // 1: genproto.CreateAttribute
	(*AttributeId)(nil),             // 2: genproto.AttributeId
	(*GetAllAttributeRequest)(nil),  // 3: genproto.GetAllAttributeRequest
	(*GetAllAttributeResponse)(nil), // 4: genproto.GetAllAttributeResponse
	(*AttributeAfterUpdate)(nil),    // 5: genproto.AttributeAfterUpdate
	(*AttributeRowsAffected)(nil),   // 6: genproto.AttributeRowsAffected
}
var file_attribute_proto_depIdxs = []int32{
	0, // 0: genproto.GetAllAttributeResponse.attributes:type_name -> genproto.Attribute
	0, // 1: genproto.AttributeAfterUpdate.old:type_name -> genproto.Attribute
	0, // 2: genproto.AttributeAfterUpdate.new:type_name -> genproto.Attribute
	1, // 3: genproto.AttributeService.Create:input_type -> genproto.CreateAttribute
	3, // 4: genproto.AttributeService.GetAll:input_type -> genproto.GetAllAttributeRequest
	2, // 5: genproto.AttributeService.GetById:input_type -> genproto.AttributeId
	0, // 6: genproto.AttributeService.Update:input_type -> genproto.Attribute
	2, // 7: genproto.AttributeService.Delete:input_type -> genproto.AttributeId
	2, // 8: genproto.AttributeService.Create:output_type -> genproto.AttributeId
	4, // 9: genproto.AttributeService.GetAll:output_type -> genproto.GetAllAttributeResponse
	0, // 10: genproto.AttributeService.GetById:output_type -> genproto.Attribute
	5, // 11: genproto.AttributeService.Update:output_type -> genproto.AttributeAfterUpdate
	6, // 12: genproto.AttributeService.Delete:output_type -> genproto.AttributeRowsAffected
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_attribute_proto_init() }
func file_attribute_proto_init() {
	if File_attribute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_attribute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Attribute); i {
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
		file_attribute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAttribute); i {
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
		file_attribute_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeId); i {
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
		file_attribute_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeRequest); i {
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
		file_attribute_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAttributeResponse); i {
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
		file_attribute_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeAfterUpdate); i {
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
		file_attribute_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeRowsAffected); i {
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
			RawDescriptor: file_attribute_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_attribute_proto_goTypes,
		DependencyIndexes: file_attribute_proto_depIdxs,
		MessageInfos:      file_attribute_proto_msgTypes,
	}.Build()
	File_attribute_proto = out.File
	file_attribute_proto_rawDesc = nil
	file_attribute_proto_goTypes = nil
	file_attribute_proto_depIdxs = nil
}
