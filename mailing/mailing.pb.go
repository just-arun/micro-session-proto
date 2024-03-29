// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: mailing/mailing.proto

package __

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

type MailType int32

const (
	MailType_ForgotPassword    MailType = 0
	MailType_TwoFactorAuth     MailType = 1
	MailType_VerificationEmail MailType = 2
)

// Enum value maps for MailType.
var (
	MailType_name = map[int32]string{
		0: "ForgotPassword",
		1: "TwoFactorAuth",
		2: "VerificationEmail",
	}
	MailType_value = map[string]int32{
		"ForgotPassword":    0,
		"TwoFactorAuth":     1,
		"VerificationEmail": 2,
	}
)

func (x MailType) Enum() *MailType {
	p := new(MailType)
	*p = x
	return p
}

func (x MailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MailType) Descriptor() protoreflect.EnumDescriptor {
	return file_mailing_mailing_proto_enumTypes[0].Descriptor()
}

func (MailType) Type() protoreflect.EnumType {
	return &file_mailing_mailing_proto_enumTypes[0]
}

func (x MailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MailType.Descriptor instead.
func (MailType) EnumDescriptor() ([]byte, []int) {
	return file_mailing_mailing_proto_rawDescGZIP(), []int{0}
}

type MailSendTwoFactorAuthOtpPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromName  string   `protobuf:"bytes,1,opt,name=fromName,proto3" json:"fromName,omitempty"`
	FromEmail string   `protobuf:"bytes,2,opt,name=fromEmail,proto3" json:"fromEmail,omitempty"`
	ToName    string   `protobuf:"bytes,3,opt,name=toName,proto3" json:"toName,omitempty"`
	ToEmail   string   `protobuf:"bytes,4,opt,name=toEmail,proto3" json:"toEmail,omitempty"`
	Otp       string   `protobuf:"bytes,5,opt,name=otp,proto3" json:"otp,omitempty"`
	Link      string   `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	MailType  MailType `protobuf:"varint,7,opt,name=mailType,proto3,enum=mailing_notification_service.MailType" json:"mailType,omitempty"`
}

func (x *MailSendTwoFactorAuthOtpPayload) Reset() {
	*x = MailSendTwoFactorAuthOtpPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailing_mailing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailSendTwoFactorAuthOtpPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailSendTwoFactorAuthOtpPayload) ProtoMessage() {}

func (x *MailSendTwoFactorAuthOtpPayload) ProtoReflect() protoreflect.Message {
	mi := &file_mailing_mailing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailSendTwoFactorAuthOtpPayload.ProtoReflect.Descriptor instead.
func (*MailSendTwoFactorAuthOtpPayload) Descriptor() ([]byte, []int) {
	return file_mailing_mailing_proto_rawDescGZIP(), []int{0}
}

func (x *MailSendTwoFactorAuthOtpPayload) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetFromEmail() string {
	if x != nil {
		return x.FromEmail
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetToName() string {
	if x != nil {
		return x.ToName
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetToEmail() string {
	if x != nil {
		return x.ToEmail
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *MailSendTwoFactorAuthOtpPayload) GetMailType() MailType {
	if x != nil {
		return x.MailType
	}
	return MailType_ForgotPassword
}

type OkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *OkResponse) Reset() {
	*x = OkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailing_mailing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkResponse) ProtoMessage() {}

func (x *OkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailing_mailing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkResponse.ProtoReflect.Descriptor instead.
func (*OkResponse) Descriptor() ([]byte, []int) {
	return file_mailing_mailing_proto_rawDescGZIP(), []int{1}
}

func (x *OkResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_mailing_mailing_proto protoreflect.FileDescriptor

var file_mailing_mailing_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xf7, 0x01, 0x0a, 0x1f, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65,
	0x6e, 0x64, 0x54, 0x77, 0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x4f,
	0x74, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74,
	0x6f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x42, 0x0a, 0x08, 0x6d,
	0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e,
	0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x61, 0x69,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x1c, 0x0a, 0x0a, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x2a, 0x48, 0x0a,
	0x08, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x6f, 0x72,
	0x67, 0x6f, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x54, 0x77, 0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x10, 0x01,
	0x12, 0x15, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x10, 0x02, 0x32, 0x91, 0x01, 0x0a, 0x1a, 0x4d, 0x61, 0x69, 0x6c,
	0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x3d, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x77, 0x6f, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x74, 0x70, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x1a, 0x28, 0x2e, 0x6d, 0x61, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mailing_mailing_proto_rawDescOnce sync.Once
	file_mailing_mailing_proto_rawDescData = file_mailing_mailing_proto_rawDesc
)

func file_mailing_mailing_proto_rawDescGZIP() []byte {
	file_mailing_mailing_proto_rawDescOnce.Do(func() {
		file_mailing_mailing_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailing_mailing_proto_rawDescData)
	})
	return file_mailing_mailing_proto_rawDescData
}

var file_mailing_mailing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mailing_mailing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mailing_mailing_proto_goTypes = []interface{}{
	(MailType)(0),                           // 0: mailing_notification_service.MailType
	(*MailSendTwoFactorAuthOtpPayload)(nil), // 1: mailing_notification_service.MailSendTwoFactorAuthOtpPayload
	(*OkResponse)(nil),                      // 2: mailing_notification_service.OkResponse
}
var file_mailing_mailing_proto_depIdxs = []int32{
	0, // 0: mailing_notification_service.MailSendTwoFactorAuthOtpPayload.mailType:type_name -> mailing_notification_service.MailType
	1, // 1: mailing_notification_service.MailingNotificationService.SendMail:input_type -> mailing_notification_service.MailSendTwoFactorAuthOtpPayload
	2, // 2: mailing_notification_service.MailingNotificationService.SendMail:output_type -> mailing_notification_service.OkResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_mailing_mailing_proto_init() }
func file_mailing_mailing_proto_init() {
	if File_mailing_mailing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailing_mailing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailSendTwoFactorAuthOtpPayload); i {
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
		file_mailing_mailing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OkResponse); i {
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
			RawDescriptor: file_mailing_mailing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mailing_mailing_proto_goTypes,
		DependencyIndexes: file_mailing_mailing_proto_depIdxs,
		EnumInfos:         file_mailing_mailing_proto_enumTypes,
		MessageInfos:      file_mailing_mailing_proto_msgTypes,
	}.Build()
	File_mailing_mailing_proto = out.File
	file_mailing_mailing_proto_rawDesc = nil
	file_mailing_mailing_proto_goTypes = nil
	file_mailing_mailing_proto_depIdxs = nil
}
