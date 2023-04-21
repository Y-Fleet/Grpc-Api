// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: fleet.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FuelConsumption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date             *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	DistanceTraveled int32                  `protobuf:"varint,2,opt,name=distance_traveled,json=distanceTraveled,proto3" json:"distance_traveled,omitempty"`
	FuelUsed         int32                  `protobuf:"varint,3,opt,name=fuel_used,json=fuelUsed,proto3" json:"fuel_used,omitempty"`
	Consumption      int32                  `protobuf:"varint,4,opt,name=consumption,proto3" json:"consumption,omitempty"`
}

func (x *FuelConsumption) Reset() {
	*x = FuelConsumption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FuelConsumption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuelConsumption) ProtoMessage() {}

func (x *FuelConsumption) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuelConsumption.ProtoReflect.Descriptor instead.
func (*FuelConsumption) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{0}
}

func (x *FuelConsumption) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *FuelConsumption) GetDistanceTraveled() int32 {
	if x != nil {
		return x.DistanceTraveled
	}
	return 0
}

func (x *FuelConsumption) GetFuelUsed() int32 {
	if x != nil {
		return x.FuelUsed
	}
	return 0
}

func (x *FuelConsumption) GetConsumption() int32 {
	if x != nil {
		return x.Consumption
	}
	return 0
}

type MaintenanceHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Description string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Cost        int32                  `protobuf:"varint,3,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *MaintenanceHistory) Reset() {
	*x = MaintenanceHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaintenanceHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceHistory) ProtoMessage() {}

func (x *MaintenanceHistory) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceHistory.ProtoReflect.Descriptor instead.
func (*MaintenanceHistory) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{1}
}

func (x *MaintenanceHistory) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *MaintenanceHistory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MaintenanceHistory) GetCost() int32 {
	if x != nil {
		return x.Cost
	}
	return 0
}

type AssignedDriver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Certifications []string `protobuf:"bytes,2,rep,name=certifications,proto3" json:"certifications,omitempty"`
}

func (x *AssignedDriver) Reset() {
	*x = AssignedDriver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssignedDriver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignedDriver) ProtoMessage() {}

func (x *AssignedDriver) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignedDriver.ProtoReflect.Descriptor instead.
func (*AssignedDriver) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{2}
}

func (x *AssignedDriver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssignedDriver) GetCertifications() []string {
	if x != nil {
		return x.Certifications
	}
	return nil
}

type Vehicle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Make               string                `protobuf:"bytes,2,opt,name=make,proto3" json:"make,omitempty"`
	Model              string                `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Year               int32                 `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	LicensePlateNumber string                `protobuf:"bytes,5,opt,name=license_plate_number,json=licensePlateNumber,proto3" json:"license_plate_number,omitempty"`
	VehicleIdNumber    string                `protobuf:"bytes,6,opt,name=vehicle_id_number,json=vehicleIdNumber,proto3" json:"vehicle_id_number,omitempty"`
	WarehouseId        string                `protobuf:"bytes,7,opt,name=warehouse_id,json=warehouseId,proto3" json:"warehouse_id,omitempty"`
	Status             string                `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	FuelConsumption    []*FuelConsumption    `protobuf:"bytes,9,rep,name=fuel_consumption,json=fuelConsumption,proto3" json:"fuel_consumption,omitempty"`
	MaintenanceHistory []*MaintenanceHistory `protobuf:"bytes,10,rep,name=maintenance_history,json=maintenanceHistory,proto3" json:"maintenance_history,omitempty"`
	AssignedDriver     *AssignedDriver       `protobuf:"bytes,11,opt,name=assigned_driver,json=assignedDriver,proto3" json:"assigned_driver,omitempty"`
}

func (x *Vehicle) Reset() {
	*x = Vehicle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vehicle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vehicle) ProtoMessage() {}

func (x *Vehicle) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vehicle.ProtoReflect.Descriptor instead.
func (*Vehicle) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{3}
}

func (x *Vehicle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vehicle) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *Vehicle) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Vehicle) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Vehicle) GetLicensePlateNumber() string {
	if x != nil {
		return x.LicensePlateNumber
	}
	return ""
}

func (x *Vehicle) GetVehicleIdNumber() string {
	if x != nil {
		return x.VehicleIdNumber
	}
	return ""
}

func (x *Vehicle) GetWarehouseId() string {
	if x != nil {
		return x.WarehouseId
	}
	return ""
}

func (x *Vehicle) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Vehicle) GetFuelConsumption() []*FuelConsumption {
	if x != nil {
		return x.FuelConsumption
	}
	return nil
}

func (x *Vehicle) GetMaintenanceHistory() []*MaintenanceHistory {
	if x != nil {
		return x.MaintenanceHistory
	}
	return nil
}

func (x *Vehicle) GetAssignedDriver() *AssignedDriver {
	if x != nil {
		return x.AssignedDriver
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{4}
}

type VehicleListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vehicles []*Vehicle `protobuf:"bytes,1,rep,name=vehicles,proto3" json:"vehicles,omitempty"`
}

func (x *VehicleListResponse) Reset() {
	*x = VehicleListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleListResponse) ProtoMessage() {}

func (x *VehicleListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_fleet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleListResponse.ProtoReflect.Descriptor instead.
func (*VehicleListResponse) Descriptor() ([]byte, []int) {
	return file_fleet_proto_rawDescGZIP(), []int{5}
}

func (x *VehicleListResponse) GetVehicles() []*Vehicle {
	if x != nil {
		return x.Vehicles
	}
	return nil
}

var File_fleet_proto protoreflect.FileDescriptor

var file_fleet_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x46,
	0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x01, 0x0a, 0x0f, 0x46, 0x75, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x55, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x7a, 0x0a, 0x12, 0x4d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x0e, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xd4, 0x03, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12,
	0x30, 0x0a, 0x14, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2a, 0x0a, 0x11, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x76, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x10, 0x66, 0x75, 0x65, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x46, 0x75, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0f, 0x66, 0x75, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x13, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d,
	0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x12, 0x6d, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x45, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x0e, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x48, 0x0a, 0x13, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x08,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x08, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x32,
	0x68, 0x0a, 0x0c, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x58, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x46, 0x6c,
	0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x21, 0x2e, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09,
	0x2f, 0x47, 0x65, 0x74, 0x46, 0x6c, 0x65, 0x65, 0x74, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleet_proto_rawDescOnce sync.Once
	file_fleet_proto_rawDescData = file_fleet_proto_rawDesc
)

func file_fleet_proto_rawDescGZIP() []byte {
	file_fleet_proto_rawDescOnce.Do(func() {
		file_fleet_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleet_proto_rawDescData)
	})
	return file_fleet_proto_rawDescData
}

var file_fleet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fleet_proto_goTypes = []interface{}{
	(*FuelConsumption)(nil),       // 0: FleetService.FuelConsumption
	(*MaintenanceHistory)(nil),    // 1: FleetService.MaintenanceHistory
	(*AssignedDriver)(nil),        // 2: FleetService.AssignedDriver
	(*Vehicle)(nil),               // 3: FleetService.Vehicle
	(*Empty)(nil),                 // 4: FleetService.Empty
	(*VehicleListResponse)(nil),   // 5: FleetService.VehicleListResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_fleet_proto_depIdxs = []int32{
	6, // 0: FleetService.FuelConsumption.date:type_name -> google.protobuf.Timestamp
	6, // 1: FleetService.MaintenanceHistory.date:type_name -> google.protobuf.Timestamp
	0, // 2: FleetService.Vehicle.fuel_consumption:type_name -> FleetService.FuelConsumption
	1, // 3: FleetService.Vehicle.maintenance_history:type_name -> FleetService.MaintenanceHistory
	2, // 4: FleetService.Vehicle.assigned_driver:type_name -> FleetService.AssignedDriver
	3, // 5: FleetService.VehicleListResponse.vehicles:type_name -> FleetService.Vehicle
	4, // 6: FleetService.fleetService.GetFleet:input_type -> FleetService.Empty
	5, // 7: FleetService.fleetService.GetFleet:output_type -> FleetService.VehicleListResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_fleet_proto_init() }
func file_fleet_proto_init() {
	if File_fleet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FuelConsumption); i {
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
		file_fleet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaintenanceHistory); i {
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
		file_fleet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssignedDriver); i {
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
		file_fleet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vehicle); i {
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
		file_fleet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_fleet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleListResponse); i {
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
			RawDescriptor: file_fleet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_fleet_proto_goTypes,
		DependencyIndexes: file_fleet_proto_depIdxs,
		MessageInfos:      file_fleet_proto_msgTypes,
	}.Build()
	File_fleet_proto = out.File
	file_fleet_proto_rawDesc = nil
	file_fleet_proto_goTypes = nil
	file_fleet_proto_depIdxs = nil
}
