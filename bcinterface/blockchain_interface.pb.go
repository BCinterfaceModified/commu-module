// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: blockchain_interface.proto

package bcinterface

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

type EnrollAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Pubkey    []byte `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *EnrollAccountRequest) Reset() {
	*x = EnrollAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_interface_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollAccountRequest) ProtoMessage() {}

func (x *EnrollAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_interface_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollAccountRequest.ProtoReflect.Descriptor instead.
func (*EnrollAccountRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_interface_proto_rawDescGZIP(), []int{0}
}

func (x *EnrollAccountRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EnrollAccountRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EnrollAccountRequest) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *EnrollAccountRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EnrollAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *EnrollAccountResponse) Reset() {
	*x = EnrollAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_interface_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrollAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrollAccountResponse) ProtoMessage() {}

func (x *EnrollAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_interface_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrollAccountResponse.ProtoReflect.Descriptor instead.
func (*EnrollAccountResponse) Descriptor() ([]byte, []int) {
	return file_blockchain_interface_proto_rawDescGZIP(), []int{1}
}

func (x *EnrollAccountResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SetupCommitteeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Round int32 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
}

func (x *SetupCommitteeRequest) Reset() {
	*x = SetupCommitteeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_interface_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupCommitteeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupCommitteeRequest) ProtoMessage() {}

func (x *SetupCommitteeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_interface_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupCommitteeRequest.ProtoReflect.Descriptor instead.
func (*SetupCommitteeRequest) Descriptor() ([]byte, []int) {
	return file_blockchain_interface_proto_rawDescGZIP(), []int{2}
}

func (x *SetupCommitteeRequest) GetRound() int32 {
	if x != nil {
		return x.Round
	}
	return 0
}

type SetupCommitteeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SetupCommitteeResponse) Reset() {
	*x = SetupCommitteeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_interface_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupCommitteeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupCommitteeResponse) ProtoMessage() {}

func (x *SetupCommitteeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_interface_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupCommitteeResponse.ProtoReflect.Descriptor instead.
func (*SetupCommitteeResponse) Descriptor() ([]byte, []int) {
	return file_blockchain_interface_proto_rawDescGZIP(), []int{3}
}

func (x *SetupCommitteeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_blockchain_interface_proto protoreflect.FileDescriptor

var file_blockchain_interface_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x22, 0x7a, 0x0a, 0x14, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x2b, 0x0a, 0x15, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x22, 0x2c, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x32, 0xcc, 0x01, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x45, 0x6e, 0x72, 0x6f,
	0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x5f, 0x6d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x75, 0x70, 0x43, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x62, 0x63, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blockchain_interface_proto_rawDescOnce sync.Once
	file_blockchain_interface_proto_rawDescData = file_blockchain_interface_proto_rawDesc
)

func file_blockchain_interface_proto_rawDescGZIP() []byte {
	file_blockchain_interface_proto_rawDescOnce.Do(func() {
		file_blockchain_interface_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_interface_proto_rawDescData)
	})
	return file_blockchain_interface_proto_rawDescData
}

var file_blockchain_interface_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_blockchain_interface_proto_goTypes = []interface{}{
	(*EnrollAccountRequest)(nil),   // 0: commu_module.EnrollAccountRequest
	(*EnrollAccountResponse)(nil),  // 1: commu_module.EnrollAccountResponse
	(*SetupCommitteeRequest)(nil),  // 2: commu_module.SetupCommitteeRequest
	(*SetupCommitteeResponse)(nil), // 3: commu_module.SetupCommitteeResponse
}
var file_blockchain_interface_proto_depIdxs = []int32{
	0, // 0: commu_module.BlockchainInterface.EnrollAccount:input_type -> commu_module.EnrollAccountRequest
	2, // 1: commu_module.BlockchainInterface.SetupCommittee:input_type -> commu_module.SetupCommitteeRequest
	1, // 2: commu_module.BlockchainInterface.EnrollAccount:output_type -> commu_module.EnrollAccountResponse
	3, // 3: commu_module.BlockchainInterface.SetupCommittee:output_type -> commu_module.SetupCommitteeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_blockchain_interface_proto_init() }
func file_blockchain_interface_proto_init() {
	if File_blockchain_interface_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_interface_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollAccountRequest); i {
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
		file_blockchain_interface_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrollAccountResponse); i {
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
		file_blockchain_interface_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupCommitteeRequest); i {
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
		file_blockchain_interface_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupCommitteeResponse); i {
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
			RawDescriptor: file_blockchain_interface_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blockchain_interface_proto_goTypes,
		DependencyIndexes: file_blockchain_interface_proto_depIdxs,
		MessageInfos:      file_blockchain_interface_proto_msgTypes,
	}.Build()
	File_blockchain_interface_proto = out.File
	file_blockchain_interface_proto_rawDesc = nil
	file_blockchain_interface_proto_goTypes = nil
	file_blockchain_interface_proto_depIdxs = nil
}
