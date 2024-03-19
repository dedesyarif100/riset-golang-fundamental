// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: model/garage.proto

package model

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

type GarageCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GarageCoordinate) Reset() {
	*x = GarageCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_garage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageCoordinate) ProtoMessage() {}

func (x *GarageCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_model_garage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageCoordinate.ProtoReflect.Descriptor instead.
func (*GarageCoordinate) Descriptor() ([]byte, []int) {
	return file_model_garage_proto_rawDescGZIP(), []int{0}
}

func (x *GarageCoordinate) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GarageCoordinate) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Garage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coordinate *GarageCoordinate `protobuf:"bytes,3,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
}

func (x *Garage) Reset() {
	*x = Garage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_garage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Garage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Garage) ProtoMessage() {}

func (x *Garage) ProtoReflect() protoreflect.Message {
	mi := &file_model_garage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Garage.ProtoReflect.Descriptor instead.
func (*Garage) Descriptor() ([]byte, []int) {
	return file_model_garage_proto_rawDescGZIP(), []int{1}
}

func (x *Garage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Garage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Garage) GetCoordinate() *GarageCoordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

type GarageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Garage `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *GarageList) Reset() {
	*x = GarageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_garage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageList) ProtoMessage() {}

func (x *GarageList) ProtoReflect() protoreflect.Message {
	mi := &file_model_garage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageList.ProtoReflect.Descriptor instead.
func (*GarageList) Descriptor() ([]byte, []int) {
	return file_model_garage_proto_rawDescGZIP(), []int{2}
}

func (x *GarageList) GetList() []*Garage {
	if x != nil {
		return x.List
	}
	return nil
}

type GarageListByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List map[string]*GarageList `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GarageListByUser) Reset() {
	*x = GarageListByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_garage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GarageListByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GarageListByUser) ProtoMessage() {}

func (x *GarageListByUser) ProtoReflect() protoreflect.Message {
	mi := &file_model_garage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GarageListByUser.ProtoReflect.Descriptor instead.
func (*GarageListByUser) Descriptor() ([]byte, []int) {
	return file_model_garage_proto_rawDescGZIP(), []int{3}
}

func (x *GarageListByUser) GetList() map[string]*GarageList {
	if x != nil {
		return x.List
	}
	return nil
}

var File_model_garage_proto protoreflect.FileDescriptor

var file_model_garage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x61, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x4c, 0x0a, 0x10, 0x47,
	0x61, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09,
	0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x65, 0x0a, 0x06, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65,
	0x22, 0x2f, 0x0a, 0x0a, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x22, 0x95, 0x01, 0x0a, 0x10, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x4a, 0x0a,
	0x09, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x27, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x47, 0x61, 0x72, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_garage_proto_rawDescOnce sync.Once
	file_model_garage_proto_rawDescData = file_model_garage_proto_rawDesc
)

func file_model_garage_proto_rawDescGZIP() []byte {
	file_model_garage_proto_rawDescOnce.Do(func() {
		file_model_garage_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_garage_proto_rawDescData)
	})
	return file_model_garage_proto_rawDescData
}

var file_model_garage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_model_garage_proto_goTypes = []interface{}{
	(*GarageCoordinate)(nil), // 0: model.GarageCoordinate
	(*Garage)(nil),           // 1: model.Garage
	(*GarageList)(nil),       // 2: model.GarageList
	(*GarageListByUser)(nil), // 3: model.GarageListByUser
	nil,                      // 4: model.GarageListByUser.ListEntry
}
var file_model_garage_proto_depIdxs = []int32{
	0, // 0: model.Garage.coordinate:type_name -> model.GarageCoordinate
	1, // 1: model.GarageList.list:type_name -> model.Garage
	4, // 2: model.GarageListByUser.list:type_name -> model.GarageListByUser.ListEntry
	2, // 3: model.GarageListByUser.ListEntry.value:type_name -> model.GarageList
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_model_garage_proto_init() }
func file_model_garage_proto_init() {
	if File_model_garage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_garage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageCoordinate); i {
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
		file_model_garage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Garage); i {
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
		file_model_garage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageList); i {
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
		file_model_garage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GarageListByUser); i {
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
			RawDescriptor: file_model_garage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_garage_proto_goTypes,
		DependencyIndexes: file_model_garage_proto_depIdxs,
		MessageInfos:      file_model_garage_proto_msgTypes,
	}.Build()
	File_model_garage_proto = out.File
	file_model_garage_proto_rawDesc = nil
	file_model_garage_proto_goTypes = nil
	file_model_garage_proto_depIdxs = nil
}
