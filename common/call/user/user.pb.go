// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.19.4
// source: application/user-center/interfaces/rpc/user.proto

package user

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

type Id struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Id) Reset() {
	*x = Id{}
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[0]
	if x != nil {
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
	return file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP(), []int{0}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Phone struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Phone         string                 `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Phone) Reset() {
	*x = Phone{}
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP(), []int{1}
}

func (x *Phone) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName      string                 `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Role          string                 `protobuf:"bytes,5,opt,name=Role,proto3" json:"Role,omitempty"`
	Sex           string                 `protobuf:"bytes,6,opt,name=Sex,proto3" json:"Sex,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserInfo) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

type OSSUploadReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	FileName      string                 `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FilePath      string                 `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OSSUploadReq) Reset() {
	*x = OSSUploadReq{}
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OSSUploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSSUploadReq) ProtoMessage() {}

func (x *OSSUploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSSUploadReq.ProtoReflect.Descriptor instead.
func (*OSSUploadReq) Descriptor() ([]byte, []int) {
	return file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP(), []int{3}
}

func (x *OSSUploadReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OSSUploadReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *OSSUploadReq) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type OSSUploadResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OSSUploadResp) Reset() {
	*x = OSSUploadResp{}
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OSSUploadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OSSUploadResp) ProtoMessage() {}

func (x *OSSUploadResp) ProtoReflect() protoreflect.Message {
	mi := &file_application_user_center_interfaces_rpc_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OSSUploadResp.ProtoReflect.Descriptor instead.
func (*OSSUploadResp) Descriptor() ([]byte, []int) {
	return file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP(), []int{4}
}

func (x *OSSUploadResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_application_user_center_interfaces_rpc_user_proto protoreflect.FileDescriptor

var file_application_user_center_interfaces_rpc_user_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2d, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1d, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x88,
	0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x65, 0x78, 0x22, 0x5a, 0x0a, 0x0c, 0x4f, 0x53, 0x53,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x21, 0x0a, 0x0d, 0x4f, 0x53, 0x53, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x9a, 0x01, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x70, 0x63, 0x12, 0x28, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2e, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x1a,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x34, 0x0a, 0x09, 0x4f, 0x73, 0x73, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4f, 0x53, 0x53, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x53, 0x53, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_application_user_center_interfaces_rpc_user_proto_rawDescOnce sync.Once
	file_application_user_center_interfaces_rpc_user_proto_rawDescData = file_application_user_center_interfaces_rpc_user_proto_rawDesc
)

func file_application_user_center_interfaces_rpc_user_proto_rawDescGZIP() []byte {
	file_application_user_center_interfaces_rpc_user_proto_rawDescOnce.Do(func() {
		file_application_user_center_interfaces_rpc_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_application_user_center_interfaces_rpc_user_proto_rawDescData)
	})
	return file_application_user_center_interfaces_rpc_user_proto_rawDescData
}

var file_application_user_center_interfaces_rpc_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_application_user_center_interfaces_rpc_user_proto_goTypes = []any{
	(*Id)(nil),            // 0: user.Id
	(*Phone)(nil),         // 1: user.Phone
	(*UserInfo)(nil),      // 2: user.UserInfo
	(*OSSUploadReq)(nil),  // 3: user.OSSUploadReq
	(*OSSUploadResp)(nil), // 4: user.OSSUploadResp
}
var file_application_user_center_interfaces_rpc_user_proto_depIdxs = []int32{
	0, // 0: user.user_rpc.FindUserById:input_type -> user.Id
	1, // 1: user.user_rpc.FindUserByPhone:input_type -> user.Phone
	3, // 2: user.user_rpc.OssUpload:input_type -> user.OSSUploadReq
	2, // 3: user.user_rpc.FindUserById:output_type -> user.UserInfo
	2, // 4: user.user_rpc.FindUserByPhone:output_type -> user.UserInfo
	4, // 5: user.user_rpc.OssUpload:output_type -> user.OSSUploadResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_application_user_center_interfaces_rpc_user_proto_init() }
func file_application_user_center_interfaces_rpc_user_proto_init() {
	if File_application_user_center_interfaces_rpc_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_application_user_center_interfaces_rpc_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_application_user_center_interfaces_rpc_user_proto_goTypes,
		DependencyIndexes: file_application_user_center_interfaces_rpc_user_proto_depIdxs,
		MessageInfos:      file_application_user_center_interfaces_rpc_user_proto_msgTypes,
	}.Build()
	File_application_user_center_interfaces_rpc_user_proto = out.File
	file_application_user_center_interfaces_rpc_user_proto_rawDesc = nil
	file_application_user_center_interfaces_rpc_user_proto_goTypes = nil
	file_application_user_center_interfaces_rpc_user_proto_depIdxs = nil
}
