// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: protos/driversapp.proto

package protos

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driversapp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driversapp_proto_msgTypes[0]
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
	return file_protos_driversapp_proto_rawDescGZIP(), []int{0}
}

type DriverInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tele  string `protobuf:"bytes,3,opt,name=tele,proto3" json:"tele,omitempty"`
	Point *Point `protobuf:"bytes,4,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *DriverInfo) Reset() {
	*x = DriverInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driversapp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverInfo) ProtoMessage() {}

func (x *DriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driversapp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverInfo.ProtoReflect.Descriptor instead.
func (*DriverInfo) Descriptor() ([]byte, []int) {
	return file_protos_driversapp_proto_rawDescGZIP(), []int{1}
}

func (x *DriverInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DriverInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DriverInfo) GetTele() string {
	if x != nil {
		return x.Tele
	}
	return ""
}

func (x *DriverInfo) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driversapp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driversapp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_protos_driversapp_proto_rawDescGZIP(), []int{2}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driversapp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driversapp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_protos_driversapp_proto_rawDescGZIP(), []int{3}
}

func (x *Id) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_driversapp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_protos_driversapp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_protos_driversapp_proto_rawDescGZIP(), []int{4}
}

func (x *Status) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_protos_driversapp_proto protoreflect.FileDescriptor

var file_protos_driversapp_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x70, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6d,
	0x0a, 0x0a, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x41, 0x0a,
	0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x22, 0x1a, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xa0, 0x02, 0x0a,
	0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x11, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x30, 0x01, 0x12, 0x33, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x0e, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x12,
	0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x12, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x32, 0x0a, 0x0c, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x22, 0x5a, 0x1a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x70, 0x70, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03,
	0x4d, 0x41, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_driversapp_proto_rawDescOnce sync.Once
	file_protos_driversapp_proto_rawDescData = file_protos_driversapp_proto_rawDesc
)

func file_protos_driversapp_proto_rawDescGZIP() []byte {
	file_protos_driversapp_proto_rawDescOnce.Do(func() {
		file_protos_driversapp_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_driversapp_proto_rawDescData)
	})
	return file_protos_driversapp_proto_rawDescData
}

var file_protos_driversapp_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_driversapp_proto_goTypes = []interface{}{
	(*Empty)(nil),      // 0: driversapp.Empty
	(*DriverInfo)(nil), // 1: driversapp.DriverInfo
	(*Point)(nil),      // 2: driversapp.Point
	(*Id)(nil),         // 3: driversapp.Id
	(*Status)(nil),     // 4: driversapp.Status
}
var file_protos_driversapp_proto_depIdxs = []int32{
	2, // 0: driversapp.DriverInfo.point:type_name -> driversapp.Point
	0, // 1: driversapp.Driver.GetDrivers:input_type -> driversapp.Empty
	3, // 2: driversapp.Driver.GetDriver:input_type -> driversapp.Id
	1, // 3: driversapp.Driver.CreateDriver:input_type -> driversapp.DriverInfo
	1, // 4: driversapp.Driver.UpdateDriver:input_type -> driversapp.DriverInfo
	3, // 5: driversapp.Driver.DeleteDriver:input_type -> driversapp.Id
	1, // 6: driversapp.Driver.GetDrivers:output_type -> driversapp.DriverInfo
	1, // 7: driversapp.Driver.GetDriver:output_type -> driversapp.DriverInfo
	3, // 8: driversapp.Driver.CreateDriver:output_type -> driversapp.Id
	4, // 9: driversapp.Driver.UpdateDriver:output_type -> driversapp.Status
	4, // 10: driversapp.Driver.DeleteDriver:output_type -> driversapp.Status
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_driversapp_proto_init() }
func file_protos_driversapp_proto_init() {
	if File_protos_driversapp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_driversapp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_driversapp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverInfo); i {
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
		file_protos_driversapp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_protos_driversapp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_protos_driversapp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_protos_driversapp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_driversapp_proto_goTypes,
		DependencyIndexes: file_protos_driversapp_proto_depIdxs,
		MessageInfos:      file_protos_driversapp_proto_msgTypes,
	}.Build()
	File_protos_driversapp_proto = out.File
	file_protos_driversapp_proto_rawDesc = nil
	file_protos_driversapp_proto_goTypes = nil
	file_protos_driversapp_proto_depIdxs = nil
}
