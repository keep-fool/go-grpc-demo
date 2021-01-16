// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: server/protos/person.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Gender int32

const (
	Gender_NOT_SPECIFIED Gender = 0
	Gender_FAMAL         Gender = 1
	Gender_MALE          Gender = 2
	Gender_WOMAN         Gender = 1
	Gender_MAN           Gender = 2
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "NOT_SPECIFIED",
		1: "FAMAL",
		2: "MALE",
		// Duplicate value: 1: "WOMAN",
		// Duplicate value: 2: "MAN",
	}
	Gender_value = map[string]int32{
		"NOT_SPECIFIED": 0,
		"FAMAL":         1,
		"MALE":          2,
		"WOMAN":         1,
		"MAN":           2,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_server_protos_person_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_server_protos_person_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_server_protos_person_proto_rawDescGZIP(), []int{0}
}

type EmployeeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32                       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Height        float32                     `protobuf:"fixed32,3,opt,name=height,proto3" json:"height,omitempty"`
	Weight        float32                     `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
	Avatar        []byte                      `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Email         string                      `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	EmailVerified bool                        `protobuf:"varint,7,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	PhoneNumbers  []string                    `protobuf:"bytes,8,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	Gender        Gender                      `protobuf:"varint,11,opt,name=gender,proto3,enum=pb.Gender" json:"gender,omitempty"`
	Addresses     []*EmployeeResponse_Address `protobuf:"bytes,13,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *EmployeeResponse) Reset() {
	*x = EmployeeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_protos_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse) ProtoMessage() {}

func (x *EmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_protos_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse.ProtoReflect.Descriptor instead.
func (*EmployeeResponse) Descriptor() ([]byte, []int) {
	return file_server_protos_person_proto_rawDescGZIP(), []int{0}
}

func (x *EmployeeResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EmployeeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EmployeeResponse) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *EmployeeResponse) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *EmployeeResponse) GetAvatar() []byte {
	if x != nil {
		return x.Avatar
	}
	return nil
}

func (x *EmployeeResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmployeeResponse) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *EmployeeResponse) GetPhoneNumbers() []string {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *EmployeeResponse) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_NOT_SPECIFIED
}

func (x *EmployeeResponse) GetAddresses() []*EmployeeResponse_Address {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type EmployeeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *EmployeeRequest) Reset() {
	*x = EmployeeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_protos_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeRequest) ProtoMessage() {}

func (x *EmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_protos_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeRequest.ProtoReflect.Descriptor instead.
func (*EmployeeRequest) Descriptor() ([]byte, []int) {
	return file_server_protos_person_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type EmployeeResponse_Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Province string `protobuf:"bytes,1,opt,name=province,proto3" json:"province,omitempty"`
	City     string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	ZipCode  string `protobuf:"bytes,3,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	Street   string `protobuf:"bytes,4,opt,name=street,proto3" json:"street,omitempty"`
	Number   string `protobuf:"bytes,5,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *EmployeeResponse_Address) Reset() {
	*x = EmployeeResponse_Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_protos_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeResponse_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeResponse_Address) ProtoMessage() {}

func (x *EmployeeResponse_Address) ProtoReflect() protoreflect.Message {
	mi := &file_server_protos_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeResponse_Address.ProtoReflect.Descriptor instead.
func (*EmployeeResponse_Address) Descriptor() ([]byte, []int) {
	return file_server_protos_person_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EmployeeResponse_Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *EmployeeResponse_Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *EmployeeResponse_Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *EmployeeResponse_Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *EmployeeResponse_Address) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

var File_server_protos_person_proto protoreflect.FileDescriptor

var file_server_protos_person_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xee, 0x03, 0x0a, 0x10, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x1a, 0x84, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x7a, 0x69, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a,
	0x4a, 0x04, 0x08, 0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x14, 0x10, 0x65, 0x4a, 0x09, 0x08, 0xc8,
	0x01, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x52, 0x03, 0x62, 0x61,
	0x72, 0x22, 0x21, 0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x2a, 0x48, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x11,
	0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x41, 0x4e, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x4e, 0x10, 0x02, 0x1a, 0x02, 0x10, 0x01, 0x32, 0x4b,
	0x0a, 0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_protos_person_proto_rawDescOnce sync.Once
	file_server_protos_person_proto_rawDescData = file_server_protos_person_proto_rawDesc
)

func file_server_protos_person_proto_rawDescGZIP() []byte {
	file_server_protos_person_proto_rawDescOnce.Do(func() {
		file_server_protos_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_protos_person_proto_rawDescData)
	})
	return file_server_protos_person_proto_rawDescData
}

var file_server_protos_person_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_server_protos_person_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_server_protos_person_proto_goTypes = []interface{}{
	(Gender)(0),                      // 0: pb.Gender
	(*EmployeeResponse)(nil),         // 1: pb.EmployeeResponse
	(*EmployeeRequest)(nil),          // 2: pb.EmployeeRequest
	(*EmployeeResponse_Address)(nil), // 3: pb.EmployeeResponse.Address
}
var file_server_protos_person_proto_depIdxs = []int32{
	0, // 0: pb.EmployeeResponse.gender:type_name -> pb.Gender
	3, // 1: pb.EmployeeResponse.addresses:type_name -> pb.EmployeeResponse.Address
	2, // 2: pb.EmployeeService.GetEmployee:input_type -> pb.EmployeeRequest
	1, // 3: pb.EmployeeService.GetEmployee:output_type -> pb.EmployeeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_server_protos_person_proto_init() }
func file_server_protos_person_proto_init() {
	if File_server_protos_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_protos_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse); i {
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
		file_server_protos_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeRequest); i {
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
		file_server_protos_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeResponse_Address); i {
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
			RawDescriptor: file_server_protos_person_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_protos_person_proto_goTypes,
		DependencyIndexes: file_server_protos_person_proto_depIdxs,
		EnumInfos:         file_server_protos_person_proto_enumTypes,
		MessageInfos:      file_server_protos_person_proto_msgTypes,
	}.Build()
	File_server_protos_person_proto = out.File
	file_server_protos_person_proto_rawDesc = nil
	file_server_protos_person_proto_goTypes = nil
	file_server_protos_person_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetEmployee(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/pb.EmployeeService/GetEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
type EmployeeServiceServer interface {
	GetEmployee(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
}

// UnimplementedEmployeeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (*UnimplementedEmployeeServiceServer) GetEmployee(context.Context, *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmployee not implemented")
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EmployeeService/GetEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetEmployee(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmployee",
			Handler:    _EmployeeService_GetEmployee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/protos/person.proto",
}