// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.3
// source: user-service/proto/user_service.proto

package proto

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PathParamsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PathParamsRequest) Reset() {
	*x = PathParamsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathParamsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathParamsRequest) ProtoMessage() {}

func (x *PathParamsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathParamsRequest.ProtoReflect.Descriptor instead.
func (*PathParamsRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *PathParamsRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PathParamsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PathParamsAndBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Number  int32  `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PathParamsAndBodyRequest) Reset() {
	*x = PathParamsAndBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathParamsAndBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathParamsAndBodyRequest) ProtoMessage() {}

func (x *PathParamsAndBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathParamsAndBodyRequest.ProtoReflect.Descriptor instead.
func (*PathParamsAndBodyRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *PathParamsAndBodyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PathParamsAndBodyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PathParamsAndBodyRequest) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *PathParamsAndBodyRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PathParamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PathParamRequest) Reset() {
	*x = PathParamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathParamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathParamRequest) ProtoMessage() {}

func (x *PathParamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathParamRequest.ProtoReflect.Descriptor instead.
func (*PathParamRequest) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *PathParamRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PathParamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PathParamResponse) Reset() {
	*x = PathParamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_service_proto_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathParamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathParamResponse) ProtoMessage() {}

func (x *PathParamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_service_proto_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathParamResponse.ProtoReflect.Descriptor instead.
func (*PathParamResponse) Descriptor() ([]byte, []int) {
	return file_user_service_proto_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *PathParamResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_user_service_proto_user_service_proto protoreflect.FileDescriptor

var file_user_service_proto_user_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x50, 0x61,
	0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x18, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x41, 0x6e, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x11, 0x50, 0x61, 0x74,
	0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x86, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x54,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x0d, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x54, 0x65, 0x73, 0x74, 0x12, 0x11,
	0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x15, 0x50, 0x61, 0x74, 0x68, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x41, 0x6e, 0x64, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x41, 0x6e, 0x64,
	0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x50, 0x61,
	0x74, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x14, 0x5a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_service_proto_user_service_proto_rawDescOnce sync.Once
	file_user_service_proto_user_service_proto_rawDescData = file_user_service_proto_user_service_proto_rawDesc
)

func file_user_service_proto_user_service_proto_rawDescGZIP() []byte {
	file_user_service_proto_user_service_proto_rawDescOnce.Do(func() {
		file_user_service_proto_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_service_proto_user_service_proto_rawDescData)
	})
	return file_user_service_proto_user_service_proto_rawDescData
}

var file_user_service_proto_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_service_proto_user_service_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),        // 0: CreateUserRequest
	(*CreateUserResponse)(nil),       // 1: CreateUserResponse
	(*PathParamsRequest)(nil),        // 2: PathParamsRequest
	(*PathParamsAndBodyRequest)(nil), // 3: PathParamsAndBodyRequest
	(*PathParamRequest)(nil),         // 4: PathParamRequest
	(*PathParamResponse)(nil),        // 5: PathParamResponse
}
var file_user_service_proto_user_service_proto_depIdxs = []int32{
	0, // 0: UserService.CreateUser:input_type -> CreateUserRequest
	2, // 1: UserService.PathParamsTest:input_type -> PathParamsRequest
	4, // 2: UserService.PathParamTest:input_type -> PathParamRequest
	3, // 3: UserService.PathParamsAndBodyTest:input_type -> PathParamsAndBodyRequest
	1, // 4: UserService.CreateUser:output_type -> CreateUserResponse
	5, // 5: UserService.PathParamsTest:output_type -> PathParamResponse
	5, // 6: UserService.PathParamTest:output_type -> PathParamResponse
	5, // 7: UserService.PathParamsAndBodyTest:output_type -> PathParamResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_service_proto_user_service_proto_init() }
func file_user_service_proto_user_service_proto_init() {
	if File_user_service_proto_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_service_proto_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_service_proto_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_user_service_proto_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathParamsRequest); i {
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
		file_user_service_proto_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathParamsAndBodyRequest); i {
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
		file_user_service_proto_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathParamRequest); i {
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
		file_user_service_proto_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathParamResponse); i {
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
			RawDescriptor: file_user_service_proto_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_service_proto_user_service_proto_goTypes,
		DependencyIndexes: file_user_service_proto_user_service_proto_depIdxs,
		MessageInfos:      file_user_service_proto_user_service_proto_msgTypes,
	}.Build()
	File_user_service_proto_user_service_proto = out.File
	file_user_service_proto_user_service_proto_rawDesc = nil
	file_user_service_proto_user_service_proto_goTypes = nil
	file_user_service_proto_user_service_proto_depIdxs = nil
}
