// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: demo.proto

package pbgen

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

type ZooLevel int32

const (
	ZooLevel_Green  ZooLevel = 0
	ZooLevel_Blue   ZooLevel = 1
	ZooLevel_Orange ZooLevel = 2
	ZooLevel_Red    ZooLevel = 3
)

// Enum value maps for ZooLevel.
var (
	ZooLevel_name = map[int32]string{
		0: "Green",
		1: "Blue",
		2: "Orange",
		3: "Red",
	}
	ZooLevel_value = map[string]int32{
		"Green":  0,
		"Blue":   1,
		"Orange": 2,
		"Red":    3,
	}
)

func (x ZooLevel) Enum() *ZooLevel {
	p := new(ZooLevel)
	*p = x
	return p
}

func (x ZooLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ZooLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_demo_proto_enumTypes[0].Descriptor()
}

func (ZooLevel) Type() protoreflect.EnumType {
	return &file_demo_proto_enumTypes[0]
}

func (x ZooLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ZooLevel.Descriptor instead.
func (ZooLevel) EnumDescriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

type FoodType int32

const (
	FoodType_Water     FoodType = 0
	FoodType_Vegetable FoodType = 1
	FoodType_Carbon    FoodType = 2
)

// Enum value maps for FoodType.
var (
	FoodType_name = map[int32]string{
		0: "Water",
		1: "Vegetable",
		2: "Carbon",
	}
	FoodType_value = map[string]int32{
		"Water":     0,
		"Vegetable": 1,
		"Carbon":    2,
	}
)

func (x FoodType) Enum() *FoodType {
	p := new(FoodType)
	*p = x
	return p
}

func (x FoodType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FoodType) Descriptor() protoreflect.EnumDescriptor {
	return file_demo_proto_enumTypes[1].Descriptor()
}

func (FoodType) Type() protoreflect.EnumType {
	return &file_demo_proto_enumTypes[1]
}

func (x FoodType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FoodType.Descriptor instead.
func (FoodType) EnumDescriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

type Assistant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Level     *int32  `protobuf:"varint,2,opt,name=level,proto3,oneof" json:"level,omitempty"`
	Direction *string `protobuf:"bytes,3,opt,name=direction,proto3,oneof" json:"direction,omitempty"`
}

func (x *Assistant) Reset() {
	*x = Assistant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assistant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assistant) ProtoMessage() {}

func (x *Assistant) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assistant.ProtoReflect.Descriptor instead.
func (*Assistant) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{0}
}

func (x *Assistant) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Assistant) GetLevel() int32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *Assistant) GetDirection() string {
	if x != nil && x.Direction != nil {
		return *x.Direction
	}
	return ""
}

type Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age        *int32         `protobuf:"varint,1,opt,name=age,proto3,oneof" json:"age,omitempty"`
	Name       *string        `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Addr       *AddressDetail `protobuf:"bytes,3,opt,name=addr,proto3,oneof" json:"addr,omitempty"`
	Assistants []*Assistant   `protobuf:"bytes,4,rep,name=assistants,proto3" json:"assistants,omitempty"`
}

func (x *Manager) Reset() {
	*x = Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Manager) ProtoMessage() {}

func (x *Manager) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Manager.ProtoReflect.Descriptor instead.
func (*Manager) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{1}
}

func (x *Manager) GetAge() int32 {
	if x != nil && x.Age != nil {
		return *x.Age
	}
	return 0
}

func (x *Manager) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Manager) GetAddr() *AddressDetail {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *Manager) GetAssistants() []*Assistant {
	if x != nil {
		return x.Assistants
	}
	return nil
}

type AddressDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name *string  `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	X    *float64 `protobuf:"fixed64,2,opt,name=x,proto3,oneof" json:"x,omitempty"`
	Y    *float64 `protobuf:"fixed64,3,opt,name=y,proto3,oneof" json:"y,omitempty"`
}

func (x *AddressDetail) Reset() {
	*x = AddressDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressDetail) ProtoMessage() {}

func (x *AddressDetail) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressDetail.ProtoReflect.Descriptor instead.
func (*AddressDetail) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{2}
}

func (x *AddressDetail) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *AddressDetail) GetX() float64 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *AddressDetail) GetY() float64 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

type Institute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Years  *int32  `protobuf:"varint,2,opt,name=years,proto3,oneof" json:"years,omitempty"`
	Normal *bool   `protobuf:"varint,3,opt,name=normal,proto3,oneof" json:"normal,omitempty"`
}

func (x *Institute) Reset() {
	*x = Institute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Institute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Institute) ProtoMessage() {}

func (x *Institute) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Institute.ProtoReflect.Descriptor instead.
func (*Institute) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{3}
}

func (x *Institute) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Institute) GetYears() int32 {
	if x != nil && x.Years != nil {
		return *x.Years
	}
	return 0
}

func (x *Institute) GetNormal() bool {
	if x != nil && x.Normal != nil {
		return *x.Normal
	}
	return false
}

type Government struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    *string `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Level *uint32 `protobuf:"varint,2,opt,name=level,proto3,oneof" json:"level,omitempty"`
}

func (x *Government) Reset() {
	*x = Government{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Government) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Government) ProtoMessage() {}

func (x *Government) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Government.ProtoReflect.Descriptor instead.
func (*Government) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{4}
}

func (x *Government) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *Government) GetLevel() uint32 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

type Paper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   *int32  `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Desc *string `protobuf:"bytes,2,opt,name=desc,proto3,oneof" json:"desc,omitempty"`
}

func (x *Paper) Reset() {
	*x = Paper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paper) ProtoMessage() {}

func (x *Paper) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paper.ProtoReflect.Descriptor instead.
func (*Paper) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{5}
}

func (x *Paper) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Paper) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

type BorrowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string  `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Count  *uint32  `protobuf:"varint,2,opt,name=count,proto3,oneof" json:"count,omitempty"`
	Reason *string  `protobuf:"bytes,3,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	Out    *bool    `protobuf:"varint,4,opt,name=out,proto3,oneof" json:"out,omitempty"`
	Paper  []*Paper `protobuf:"bytes,5,rep,name=paper,proto3" json:"paper,omitempty"`
}

func (x *BorrowInfo) Reset() {
	*x = BorrowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BorrowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BorrowInfo) ProtoMessage() {}

func (x *BorrowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BorrowInfo.ProtoReflect.Descriptor instead.
func (*BorrowInfo) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{6}
}

func (x *BorrowInfo) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *BorrowInfo) GetCount() uint32 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *BorrowInfo) GetReason() string {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return ""
}

func (x *BorrowInfo) GetOut() bool {
	if x != nil && x.Out != nil {
		return *x.Out
	}
	return false
}

func (x *BorrowInfo) GetPaper() []*Paper {
	if x != nil {
		return x.Paper
	}
	return nil
}

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string   `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Ft     *FoodType `protobuf:"varint,2,opt,name=ft,proto3,enum=FoodType,oneof" json:"ft,omitempty"`
	Weight *float32  `protobuf:"fixed32,3,opt,name=weight,proto3,oneof" json:"weight,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{7}
}

func (x *Food) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Food) GetFt() FoodType {
	if x != nil && x.Ft != nil {
		return *x.Ft
	}
	return FoodType_Water
}

func (x *Food) GetWeight() float32 {
	if x != nil && x.Weight != nil {
		return *x.Weight
	}
	return 0
}

type Animal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Detail *string `protobuf:"bytes,2,opt,name=detail,proto3,oneof" json:"detail,omitempty"`
	Count  *uint64 `protobuf:"varint,3,opt,name=count,proto3,oneof" json:"count,omitempty"`
	Foods  []*Food `protobuf:"bytes,4,rep,name=foods,proto3" json:"foods,omitempty"`
}

func (x *Animal) Reset() {
	*x = Animal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Animal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Animal) ProtoMessage() {}

func (x *Animal) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Animal.ProtoReflect.Descriptor instead.
func (*Animal) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{8}
}

func (x *Animal) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Animal) GetDetail() string {
	if x != nil && x.Detail != nil {
		return *x.Detail
	}
	return ""
}

func (x *Animal) GetCount() uint64 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *Animal) GetFoods() []*Food {
	if x != nil {
		return x.Foods
	}
	return nil
}

type Zoo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *int32   `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"` //test //maka
	Desc    *string  `protobuf:"bytes,2,opt,name=desc,proto3,oneof" json:"desc,omitempty"`
	Manager *Manager `protobuf:"bytes,3,opt,name=manager,proto3,oneof" json:"manager,omitempty"`
	// Types that are assignable to Belong:
	//
	//	*Zoo_Institute
	//	*Zoo_Government
	//	*Zoo_Other
	Belong   isZoo_Belong           `protobuf_oneof:"belong"`
	Borrows  map[uint64]*BorrowInfo `protobuf:"bytes,7,rep,name=borrows,proto3" json:"borrows,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Animals  []*Animal              `protobuf:"bytes,8,rep,name=animals,proto3" json:"animals,omitempty"`
	Logs     map[int32]string       `protobuf:"bytes,9,rep,name=logs,proto3" json:"logs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Level    *ZooLevel              `protobuf:"varint,10,opt,name=level,proto3,enum=ZooLevel,oneof" json:"level,omitempty"`
	OpenDays []int32                `protobuf:"varint,11,rep,packed,name=openDays,proto3" json:"openDays,omitempty"`
}

func (x *Zoo) Reset() {
	*x = Zoo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zoo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zoo) ProtoMessage() {}

func (x *Zoo) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zoo.ProtoReflect.Descriptor instead.
func (*Zoo) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{9}
}

func (x *Zoo) GetId() int32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Zoo) GetDesc() string {
	if x != nil && x.Desc != nil {
		return *x.Desc
	}
	return ""
}

func (x *Zoo) GetManager() *Manager {
	if x != nil {
		return x.Manager
	}
	return nil
}

func (m *Zoo) GetBelong() isZoo_Belong {
	if m != nil {
		return m.Belong
	}
	return nil
}

func (x *Zoo) GetInstitute() *Institute {
	if x, ok := x.GetBelong().(*Zoo_Institute); ok {
		return x.Institute
	}
	return nil
}

func (x *Zoo) GetGovernment() *Government {
	if x, ok := x.GetBelong().(*Zoo_Government); ok {
		return x.Government
	}
	return nil
}

func (x *Zoo) GetOther() string {
	if x, ok := x.GetBelong().(*Zoo_Other); ok {
		return x.Other
	}
	return ""
}

func (x *Zoo) GetBorrows() map[uint64]*BorrowInfo {
	if x != nil {
		return x.Borrows
	}
	return nil
}

func (x *Zoo) GetAnimals() []*Animal {
	if x != nil {
		return x.Animals
	}
	return nil
}

func (x *Zoo) GetLogs() map[int32]string {
	if x != nil {
		return x.Logs
	}
	return nil
}

func (x *Zoo) GetLevel() ZooLevel {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return ZooLevel_Green
}

func (x *Zoo) GetOpenDays() []int32 {
	if x != nil {
		return x.OpenDays
	}
	return nil
}

type isZoo_Belong interface {
	isZoo_Belong()
}

type Zoo_Institute struct {
	Institute *Institute `protobuf:"bytes,4,opt,name=institute,proto3,oneof"`
}

type Zoo_Government struct {
	Government *Government `protobuf:"bytes,5,opt,name=government,proto3,oneof"`
}

type Zoo_Other struct {
	Other string `protobuf:"bytes,6,opt,name=other,proto3,oneof"`
}

func (*Zoo_Institute) isZoo_Belong() {}

func (*Zoo_Government) isZoo_Belong() {}

func (*Zoo_Other) isZoo_Belong() {}

type ZooTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zoos []*Zoo `protobuf:"bytes,1,rep,name=zoos,proto3" json:"zoos,omitempty"`
}

func (x *ZooTable) Reset() {
	*x = ZooTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZooTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZooTable) ProtoMessage() {}

func (x *ZooTable) ProtoReflect() protoreflect.Message {
	mi := &file_demo_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZooTable.ProtoReflect.Descriptor instead.
func (*ZooTable) Descriptor() ([]byte, []int) {
	return file_demo_proto_rawDescGZIP(), []int{10}
}

func (x *ZooTable) GetZoos() []*Zoo {
	if x != nil {
		return x.Zoos
	}
	return nil
}

var File_demo_proto protoreflect.FileDescriptor

var file_demo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a,
	0x09, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x21,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x15,
	0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x48, 0x02, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x61, 0x67, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x22, 0x63, 0x0a,
	0x0d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x01, 0x52, 0x01, 0x78, 0x88, 0x01, 0x01, 0x12, 0x11, 0x0a, 0x01, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x01, 0x79, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x04, 0x0a, 0x02, 0x5f, 0x78, 0x42, 0x04, 0x0a, 0x02,
	0x5f, 0x79, 0x22, 0x7a, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x79, 0x65, 0x61, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x05, 0x79, 0x65, 0x61, 0x72, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x06, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x88, 0x01, 0x01,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x79, 0x65,
	0x61, 0x72, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x22, 0x4d,
	0x0a, 0x0a, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x48, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x45, 0x0a,
	0x05, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x03, 0x52, 0x03, 0x6f, 0x75, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x05, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x70, 0x65, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6f, 0x75, 0x74, 0x22,
	0x77, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x02, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x46,
	0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x01, 0x52, 0x02, 0x66, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x1b, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x48, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a,
	0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x66, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x41, 0x6e, 0x69,
	0x6d, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x06,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x05, 0x66, 0x6f, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x05, 0x66, 0x6f, 0x6f, 0x64,
	0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0xb7, 0x04, 0x0a, 0x03, 0x5a, 0x6f, 0x6f, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04,
	0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x48, 0x03, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x09, 0x69, 0x6e, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x67, 0x6f,
	0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x67,
	0x6f, 0x76, 0x65, 0x72, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x2b, 0x0a, 0x07, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x5a, 0x6f, 0x6f, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x62, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x21,
	0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x61, 0x6c, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x61, 0x6c,
	0x73, 0x12, 0x22, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x5a, 0x6f, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x5a, 0x6f, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x48,
	0x04, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x1a, 0x47, 0x0a, 0x0c, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x6f, 0x72, 0x72, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x37, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x62, 0x65, 0x6c,
	0x6f, 0x6e, 0x67, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x24, 0x0a, 0x08, 0x5a, 0x6f, 0x6f,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x7a, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x5a, 0x6f, 0x6f, 0x52, 0x04, 0x7a, 0x6f, 0x6f, 0x73, 0x2a,
	0x34, 0x0a, 0x08, 0x5a, 0x6f, 0x6f, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x72, 0x65, 0x65, 0x6e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x6c, 0x75, 0x65, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x52, 0x65, 0x64, 0x10, 0x03, 0x2a, 0x30, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x56, 0x65, 0x67, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43,
	0x61, 0x72, 0x62, 0x6f, 0x6e, 0x10, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x62,
	0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_proto_rawDescOnce sync.Once
	file_demo_proto_rawDescData = file_demo_proto_rawDesc
)

func file_demo_proto_rawDescGZIP() []byte {
	file_demo_proto_rawDescOnce.Do(func() {
		file_demo_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_proto_rawDescData)
	})
	return file_demo_proto_rawDescData
}

var file_demo_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_demo_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_demo_proto_goTypes = []any{
	(ZooLevel)(0),         // 0: ZooLevel
	(FoodType)(0),         // 1: FoodType
	(*Assistant)(nil),     // 2: Assistant
	(*Manager)(nil),       // 3: Manager
	(*AddressDetail)(nil), // 4: AddressDetail
	(*Institute)(nil),     // 5: Institute
	(*Government)(nil),    // 6: Government
	(*Paper)(nil),         // 7: Paper
	(*BorrowInfo)(nil),    // 8: BorrowInfo
	(*Food)(nil),          // 9: Food
	(*Animal)(nil),        // 10: Animal
	(*Zoo)(nil),           // 11: Zoo
	(*ZooTable)(nil),      // 12: ZooTable
	nil,                   // 13: Zoo.BorrowsEntry
	nil,                   // 14: Zoo.LogsEntry
}
var file_demo_proto_depIdxs = []int32{
	4,  // 0: Manager.addr:type_name -> AddressDetail
	2,  // 1: Manager.assistants:type_name -> Assistant
	7,  // 2: BorrowInfo.paper:type_name -> Paper
	1,  // 3: Food.ft:type_name -> FoodType
	9,  // 4: Animal.foods:type_name -> Food
	3,  // 5: Zoo.manager:type_name -> Manager
	5,  // 6: Zoo.institute:type_name -> Institute
	6,  // 7: Zoo.government:type_name -> Government
	13, // 8: Zoo.borrows:type_name -> Zoo.BorrowsEntry
	10, // 9: Zoo.animals:type_name -> Animal
	14, // 10: Zoo.logs:type_name -> Zoo.LogsEntry
	0,  // 11: Zoo.level:type_name -> ZooLevel
	11, // 12: ZooTable.zoos:type_name -> Zoo
	8,  // 13: Zoo.BorrowsEntry.value:type_name -> BorrowInfo
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_demo_proto_init() }
func file_demo_proto_init() {
	if File_demo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Assistant); i {
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
		file_demo_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Manager); i {
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
		file_demo_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddressDetail); i {
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
		file_demo_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Institute); i {
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
		file_demo_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Government); i {
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
		file_demo_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Paper); i {
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
		file_demo_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*BorrowInfo); i {
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
		file_demo_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Food); i {
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
		file_demo_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Animal); i {
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
		file_demo_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Zoo); i {
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
		file_demo_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ZooTable); i {
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
	file_demo_proto_msgTypes[0].OneofWrappers = []any{}
	file_demo_proto_msgTypes[1].OneofWrappers = []any{}
	file_demo_proto_msgTypes[2].OneofWrappers = []any{}
	file_demo_proto_msgTypes[3].OneofWrappers = []any{}
	file_demo_proto_msgTypes[4].OneofWrappers = []any{}
	file_demo_proto_msgTypes[5].OneofWrappers = []any{}
	file_demo_proto_msgTypes[6].OneofWrappers = []any{}
	file_demo_proto_msgTypes[7].OneofWrappers = []any{}
	file_demo_proto_msgTypes[8].OneofWrappers = []any{}
	file_demo_proto_msgTypes[9].OneofWrappers = []any{
		(*Zoo_Institute)(nil),
		(*Zoo_Government)(nil),
		(*Zoo_Other)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_demo_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_demo_proto_goTypes,
		DependencyIndexes: file_demo_proto_depIdxs,
		EnumInfos:         file_demo_proto_enumTypes,
		MessageInfos:      file_demo_proto_msgTypes,
	}.Build()
	File_demo_proto = out.File
	file_demo_proto_rawDesc = nil
	file_demo_proto_goTypes = nil
	file_demo_proto_depIdxs = nil
}