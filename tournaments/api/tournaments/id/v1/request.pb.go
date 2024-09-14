// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/tournaments/id/v1/request.proto

package idv1

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

type GetTournamentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTournamentRequest) Reset() {
	*x = GetTournamentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tournaments_id_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentRequest) ProtoMessage() {}

func (x *GetTournamentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tournaments_id_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentRequest.ProtoReflect.Descriptor instead.
func (*GetTournamentRequest) Descriptor() ([]byte, []int) {
	return file_api_tournaments_id_v1_request_proto_rawDescGZIP(), []int{0}
}

type DeleteTournamentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTournamentRequest) Reset() {
	*x = DeleteTournamentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tournaments_id_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTournamentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTournamentRequest) ProtoMessage() {}

func (x *DeleteTournamentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tournaments_id_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTournamentRequest.ProtoReflect.Descriptor instead.
func (*DeleteTournamentRequest) Descriptor() ([]byte, []int) {
	return file_api_tournaments_id_v1_request_proto_rawDescGZIP(), []int{1}
}

var File_api_tournaments_id_v1_request_proto protoreflect.FileDescriptor

var file_api_tournaments_id_v1_request_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x16, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0xdd, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x69, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x54, 0x49,
	0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x49, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x41, 0x70, 0x69, 0x5c, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x49, 0x64, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5c, 0x49, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x49, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tournaments_id_v1_request_proto_rawDescOnce sync.Once
	file_api_tournaments_id_v1_request_proto_rawDescData = file_api_tournaments_id_v1_request_proto_rawDesc
)

func file_api_tournaments_id_v1_request_proto_rawDescGZIP() []byte {
	file_api_tournaments_id_v1_request_proto_rawDescOnce.Do(func() {
		file_api_tournaments_id_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tournaments_id_v1_request_proto_rawDescData)
	})
	return file_api_tournaments_id_v1_request_proto_rawDescData
}

var file_api_tournaments_id_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_tournaments_id_v1_request_proto_goTypes = []any{
	(*GetTournamentRequest)(nil),    // 0: api.tournaments.id.v1.GetTournamentRequest
	(*DeleteTournamentRequest)(nil), // 1: api.tournaments.id.v1.DeleteTournamentRequest
}
var file_api_tournaments_id_v1_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_tournaments_id_v1_request_proto_init() }
func file_api_tournaments_id_v1_request_proto_init() {
	if File_api_tournaments_id_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tournaments_id_v1_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTournamentRequest); i {
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
		file_api_tournaments_id_v1_request_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTournamentRequest); i {
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
			RawDescriptor: file_api_tournaments_id_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_tournaments_id_v1_request_proto_goTypes,
		DependencyIndexes: file_api_tournaments_id_v1_request_proto_depIdxs,
		MessageInfos:      file_api_tournaments_id_v1_request_proto_msgTypes,
	}.Build()
	File_api_tournaments_id_v1_request_proto = out.File
	file_api_tournaments_id_v1_request_proto_rawDesc = nil
	file_api_tournaments_id_v1_request_proto_goTypes = nil
	file_api_tournaments_id_v1_request_proto_depIdxs = nil
}
