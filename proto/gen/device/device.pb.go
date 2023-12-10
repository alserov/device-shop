// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.1
// source: device/device.proto

package device

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID         string  `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Title        string  `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Price        float32 `protobuf:"fixed32,4,opt,name=Price,proto3" json:"Price,omitempty"`
	Manufacturer string  `protobuf:"bytes,5,opt,name=Manufacturer,proto3" json:"Manufacturer,omitempty"`
	Amount       uint32  `protobuf:"varint,6,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{0}
}

func (x *Device) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *Device) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Device) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Device) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Device) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Device) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type IncreaseDeviceAmountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUUID string `protobuf:"bytes,1,opt,name=DeviceUUID,proto3" json:"DeviceUUID,omitempty"`
}

func (x *IncreaseDeviceAmountReq) Reset() {
	*x = IncreaseDeviceAmountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseDeviceAmountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseDeviceAmountReq) ProtoMessage() {}

func (x *IncreaseDeviceAmountReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseDeviceAmountReq.ProtoReflect.Descriptor instead.
func (*IncreaseDeviceAmountReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{1}
}

func (x *IncreaseDeviceAmountReq) GetDeviceUUID() string {
	if x != nil {
		return x.DeviceUUID
	}
	return ""
}

type CreateDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string  `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string  `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Price        float32 `protobuf:"fixed32,3,opt,name=Price,proto3" json:"Price,omitempty"`
	Manufacturer string  `protobuf:"bytes,4,opt,name=Manufacturer,proto3" json:"Manufacturer,omitempty"`
	Amount       uint32  `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *CreateDeviceReq) Reset() {
	*x = CreateDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeviceReq) ProtoMessage() {}

func (x *CreateDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeviceReq.ProtoReflect.Descriptor instead.
func (*CreateDeviceReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{2}
}

func (x *CreateDeviceReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateDeviceReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateDeviceReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateDeviceReq) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *CreateDeviceReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type DeleteDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *DeleteDeviceReq) Reset() {
	*x = DeleteDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDeviceReq) ProtoMessage() {}

func (x *DeleteDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDeviceReq.ProtoReflect.Descriptor instead.
func (*DeleteDeviceReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteDeviceReq) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type UpdateDeviceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=Description,proto3" json:"Description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=Price,proto3" json:"Price,omitempty"`
	UUID        string  `protobuf:"bytes,4,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *UpdateDeviceReq) Reset() {
	*x = UpdateDeviceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeviceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeviceReq) ProtoMessage() {}

func (x *UpdateDeviceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeviceReq.ProtoReflect.Descriptor instead.
func (*UpdateDeviceReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateDeviceReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateDeviceReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateDeviceReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *UpdateDeviceReq) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type GetAllDevicesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index  uint32 `protobuf:"varint,1,opt,name=Index,proto3" json:"Index,omitempty"`
	Amount uint32 `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *GetAllDevicesReq) Reset() {
	*x = GetAllDevicesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDevicesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDevicesReq) ProtoMessage() {}

func (x *GetAllDevicesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDevicesReq.ProtoReflect.Descriptor instead.
func (*GetAllDevicesReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllDevicesReq) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *GetAllDevicesReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetDeviceByTitleReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *GetDeviceByTitleReq) Reset() {
	*x = GetDeviceByTitleReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceByTitleReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceByTitleReq) ProtoMessage() {}

func (x *GetDeviceByTitleReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceByTitleReq.ProtoReflect.Descriptor instead.
func (*GetDeviceByTitleReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{6}
}

func (x *GetDeviceByTitleReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type IncreaseDeviceAmountByUUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceUUID string `protobuf:"bytes,1,opt,name=DeviceUUID,proto3" json:"DeviceUUID,omitempty"`
	Amount     uint32 `protobuf:"varint,2,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *IncreaseDeviceAmountByUUIDReq) Reset() {
	*x = IncreaseDeviceAmountByUUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseDeviceAmountByUUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseDeviceAmountByUUIDReq) ProtoMessage() {}

func (x *IncreaseDeviceAmountByUUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseDeviceAmountByUUIDReq.ProtoReflect.Descriptor instead.
func (*IncreaseDeviceAmountByUUIDReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{7}
}

func (x *IncreaseDeviceAmountByUUIDReq) GetDeviceUUID() string {
	if x != nil {
		return x.DeviceUUID
	}
	return ""
}

func (x *IncreaseDeviceAmountByUUIDReq) GetAmount() uint32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetByPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min float32 `protobuf:"fixed32,1,opt,name=Min,proto3" json:"Min,omitempty"`
	Max float32 `protobuf:"fixed32,2,opt,name=Max,proto3" json:"Max,omitempty"`
}

func (x *GetByPrice) Reset() {
	*x = GetByPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPrice) ProtoMessage() {}

func (x *GetByPrice) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPrice.ProtoReflect.Descriptor instead.
func (*GetByPrice) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{8}
}

func (x *GetByPrice) GetMin() float32 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *GetByPrice) GetMax() float32 {
	if x != nil {
		return x.Max
	}
	return 0
}

type GetByManufacturer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Manufacturer string `protobuf:"bytes,1,opt,name=Manufacturer,proto3" json:"Manufacturer,omitempty"`
}

func (x *GetByManufacturer) Reset() {
	*x = GetByManufacturer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByManufacturer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByManufacturer) ProtoMessage() {}

func (x *GetByManufacturer) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByManufacturer.ProtoReflect.Descriptor instead.
func (*GetByManufacturer) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{9}
}

func (x *GetByManufacturer) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

type GetDeviceByUUIDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *GetDeviceByUUIDReq) Reset() {
	*x = GetDeviceByUUIDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceByUUIDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceByUUIDReq) ProtoMessage() {}

func (x *GetDeviceByUUIDReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceByUUIDReq.ProtoReflect.Descriptor instead.
func (*GetDeviceByUUIDReq) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{10}
}

func (x *GetDeviceByUUIDReq) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type DevicesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=Devices,proto3" json:"Devices,omitempty"`
}

func (x *DevicesRes) Reset() {
	*x = DevicesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DevicesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DevicesRes) ProtoMessage() {}

func (x *DevicesRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DevicesRes.ProtoReflect.Descriptor instead.
func (*DevicesRes) Descriptor() ([]byte, []int) {
	return file_protos_device_proto_rawDescGZIP(), []int{11}
}

func (x *DevicesRes) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

var File_protos_device_proto protoreflect.FileDescriptor

var file_protos_device_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x39, 0x0a, 0x17, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x22, 0x9b,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x22, 0x73, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x57, 0x0a, 0x1d, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x55, 0x55, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x30, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x4d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x4d, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x4d,
	0x61, 0x78, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x75, 0x66,
	0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xef, 0x04,
	0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x2e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x42, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x49,
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x4d, 0x61,
	0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12,
	0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x1a, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x14, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x25, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a,
	0x0e, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x42,
	0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c,
	0x73, 0x65, 0x72, 0x6f, 0x76, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x68, 0x6f,
	0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_device_proto_rawDescOnce sync.Once
	file_protos_device_proto_rawDescData = file_protos_device_proto_rawDesc
)

func file_protos_device_proto_rawDescGZIP() []byte {
	file_protos_device_proto_rawDescOnce.Do(func() {
		file_protos_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_device_proto_rawDescData)
	})
	return file_protos_device_proto_rawDescData
}

var file_protos_device_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_protos_device_proto_goTypes = []interface{}{
	(*Device)(nil),                        // 0: device.Device
	(*IncreaseDeviceAmountReq)(nil),       // 1: device.IncreaseDeviceAmountReq
	(*CreateDeviceReq)(nil),               // 2: device.CreateDeviceReq
	(*DeleteDeviceReq)(nil),               // 3: device.DeleteDeviceReq
	(*UpdateDeviceReq)(nil),               // 4: device.UpdateDeviceReq
	(*GetAllDevicesReq)(nil),              // 5: device.GetAllDevicesReq
	(*GetDeviceByTitleReq)(nil),           // 6: device.GetDeviceByTitleReq
	(*IncreaseDeviceAmountByUUIDReq)(nil), // 7: device.IncreaseDeviceAmountByUUIDReq
	(*GetByPrice)(nil),                    // 8: device.GetByPrice
	(*GetByManufacturer)(nil),             // 9: device.GetByManufacturer
	(*GetDeviceByUUIDReq)(nil),            // 10: device.GetDeviceByUUIDReq
	(*DevicesRes)(nil),                    // 11: device.DevicesRes
	(*emptypb.Empty)(nil),                 // 12: google.protobuf.Empty
}
var file_protos_device_proto_depIdxs = []int32{
	0,  // 0: device.DevicesRes.Devices:type_name -> device.Device
	5,  // 1: device.Devices.GetAllDevices:input_type -> device.GetAllDevicesReq
	6,  // 2: device.Devices.GetDevicesByTitle:input_type -> device.GetDeviceByTitleReq
	9,  // 3: device.Devices.GetDevicesByManufacturer:input_type -> device.GetByManufacturer
	8,  // 4: device.Devices.GetDevicesByPrice:input_type -> device.GetByPrice
	2,  // 5: device.Devices.CreateDevice:input_type -> device.CreateDeviceReq
	3,  // 6: device.Devices.DeleteDevice:input_type -> device.DeleteDeviceReq
	4,  // 7: device.Devices.UpdateDevice:input_type -> device.UpdateDeviceReq
	7,  // 8: device.Devices.IncreaseDeviceAmount:input_type -> device.IncreaseDeviceAmountByUUIDReq
	10, // 9: device.Devices.GetDeviceByUUID:input_type -> device.GetDeviceByUUIDReq
	11, // 10: device.Devices.GetAllDevices:output_type -> device.DevicesRes
	11, // 11: device.Devices.GetDevicesByTitle:output_type -> device.DevicesRes
	11, // 12: device.Devices.GetDevicesByManufacturer:output_type -> device.DevicesRes
	11, // 13: device.Devices.GetDevicesByPrice:output_type -> device.DevicesRes
	12, // 14: device.Devices.CreateDevice:output_type -> google.protobuf.Empty
	12, // 15: device.Devices.DeleteDevice:output_type -> google.protobuf.Empty
	12, // 16: device.Devices.UpdateDevice:output_type -> google.protobuf.Empty
	12, // 17: device.Devices.IncreaseDeviceAmount:output_type -> google.protobuf.Empty
	0,  // 18: device.Devices.GetDeviceByUUID:output_type -> device.Device
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_protos_device_proto_init() }
func file_protos_device_proto_init() {
	if File_protos_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_device_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_protos_device_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseDeviceAmountReq); i {
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
		file_protos_device_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeviceReq); i {
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
		file_protos_device_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDeviceReq); i {
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
		file_protos_device_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeviceReq); i {
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
		file_protos_device_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDevicesReq); i {
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
		file_protos_device_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceByTitleReq); i {
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
		file_protos_device_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseDeviceAmountByUUIDReq); i {
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
		file_protos_device_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPrice); i {
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
		file_protos_device_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByManufacturer); i {
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
		file_protos_device_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceByUUIDReq); i {
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
		file_protos_device_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DevicesRes); i {
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
			RawDescriptor: file_protos_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_device_proto_goTypes,
		DependencyIndexes: file_protos_device_proto_depIdxs,
		MessageInfos:      file_protos_device_proto_msgTypes,
	}.Build()
	File_protos_device_proto = out.File
	file_protos_device_proto_rawDesc = nil
	file_protos_device_proto_goTypes = nil
	file_protos_device_proto_depIdxs = nil
}
