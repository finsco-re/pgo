// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/organisations/id/v1/response.proto

package idv1

import (
	v1 "github.com/finsco-re/pgo/tournaments/model/organisation/v1"
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

type GetOrganisationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organisation *v1.Organisation `protobuf:"bytes,1,opt,name=organisation,proto3" json:"organisation,omitempty"`
}

func (x *GetOrganisationResponse) Reset() {
	*x = GetOrganisationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_organisations_id_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrganisationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrganisationResponse) ProtoMessage() {}

func (x *GetOrganisationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_organisations_id_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrganisationResponse.ProtoReflect.Descriptor instead.
func (*GetOrganisationResponse) Descriptor() ([]byte, []int) {
	return file_api_organisations_id_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrganisationResponse) GetOrganisation() *v1.Organisation {
	if x != nil {
		return x.Organisation
	}
	return nil
}

type PatchOrganisationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organisation *v1.Organisation `protobuf:"bytes,1,opt,name=organisation,proto3" json:"organisation,omitempty"`
}

func (x *PatchOrganisationResponse) Reset() {
	*x = PatchOrganisationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_organisations_id_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchOrganisationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchOrganisationResponse) ProtoMessage() {}

func (x *PatchOrganisationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_organisations_id_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchOrganisationResponse.ProtoReflect.Descriptor instead.
func (*PatchOrganisationResponse) Descriptor() ([]byte, []int) {
	return file_api_organisations_id_v1_response_proto_rawDescGZIP(), []int{1}
}

func (x *PatchOrganisationResponse) GetOrganisation() *v1.Organisation {
	if x != nil {
		return x.Organisation
	}
	return nil
}

type DeleteOrganisationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteOrganisationResponse) Reset() {
	*x = DeleteOrganisationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_organisations_id_v1_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrganisationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrganisationResponse) ProtoMessage() {}

func (x *DeleteOrganisationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_organisations_id_v1_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrganisationResponse.ProtoReflect.Descriptor instead.
func (*DeleteOrganisationResponse) Descriptor() ([]byte, []int) {
	return file_api_organisations_id_v1_response_proto_rawDescGZIP(), []int{2}
}

var File_api_organisations_id_v1_response_proto protoreflect.FileDescriptor

var file_api_organisations_id_v1_response_proto_rawDesc = []byte{
	0x0a, 0x26, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x64, 0x2e, 0x76,
	0x31, 0x1a, 0x28, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x64, 0x0a, 0x19, 0x50, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xee, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x64,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x69, 0x6e, 0x73, 0x63, 0x6f, 0x2d, 0x72, 0x65, 0x2f, 0x70, 0x67, 0x6f, 0x2f, 0x74,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x64, 0x2f,
	0x76, 0x31, 0x3b, 0x69, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x4f, 0x49, 0xaa, 0x02, 0x17,
	0x41, 0x70, 0x69, 0x2e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x49, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x49, 0x64, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x23, 0x41, 0x70, 0x69, 0x5c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5c, 0x49, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x3a, 0x49, 0x64,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_organisations_id_v1_response_proto_rawDescOnce sync.Once
	file_api_organisations_id_v1_response_proto_rawDescData = file_api_organisations_id_v1_response_proto_rawDesc
)

func file_api_organisations_id_v1_response_proto_rawDescGZIP() []byte {
	file_api_organisations_id_v1_response_proto_rawDescOnce.Do(func() {
		file_api_organisations_id_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_organisations_id_v1_response_proto_rawDescData)
	})
	return file_api_organisations_id_v1_response_proto_rawDescData
}

var file_api_organisations_id_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_organisations_id_v1_response_proto_goTypes = []any{
	(*GetOrganisationResponse)(nil),    // 0: api.organisations.id.v1.GetOrganisationResponse
	(*PatchOrganisationResponse)(nil),  // 1: api.organisations.id.v1.PatchOrganisationResponse
	(*DeleteOrganisationResponse)(nil), // 2: api.organisations.id.v1.DeleteOrganisationResponse
	(*v1.Organisation)(nil),            // 3: model.organisation.v1.Organisation
}
var file_api_organisations_id_v1_response_proto_depIdxs = []int32{
	3, // 0: api.organisations.id.v1.GetOrganisationResponse.organisation:type_name -> model.organisation.v1.Organisation
	3, // 1: api.organisations.id.v1.PatchOrganisationResponse.organisation:type_name -> model.organisation.v1.Organisation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_organisations_id_v1_response_proto_init() }
func file_api_organisations_id_v1_response_proto_init() {
	if File_api_organisations_id_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_organisations_id_v1_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrganisationResponse); i {
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
		file_api_organisations_id_v1_response_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PatchOrganisationResponse); i {
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
		file_api_organisations_id_v1_response_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOrganisationResponse); i {
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
			RawDescriptor: file_api_organisations_id_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_organisations_id_v1_response_proto_goTypes,
		DependencyIndexes: file_api_organisations_id_v1_response_proto_depIdxs,
		MessageInfos:      file_api_organisations_id_v1_response_proto_msgTypes,
	}.Build()
	File_api_organisations_id_v1_response_proto = out.File
	file_api_organisations_id_v1_response_proto_rawDesc = nil
	file_api_organisations_id_v1_response_proto_goTypes = nil
	file_api_organisations_id_v1_response_proto_depIdxs = nil
}
