// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/tournaments/id/v1/response.proto

package idv1

import (
	v1 "github.com/finsco-re/pgo/model/tournament/v1"
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

type GetTournamentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tournament *v1.Tournament `protobuf:"bytes,1,opt,name=tournament,proto3" json:"tournament,omitempty"`
}

func (x *GetTournamentResponse) Reset() {
	*x = GetTournamentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tournaments_id_v1_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTournamentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTournamentResponse) ProtoMessage() {}

func (x *GetTournamentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tournaments_id_v1_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTournamentResponse.ProtoReflect.Descriptor instead.
func (*GetTournamentResponse) Descriptor() ([]byte, []int) {
	return file_api_tournaments_id_v1_response_proto_rawDescGZIP(), []int{0}
}

func (x *GetTournamentResponse) GetTournament() *v1.Tournament {
	if x != nil {
		return x.Tournament
	}
	return nil
}

type DeleteTournamentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTournamentResponse) Reset() {
	*x = DeleteTournamentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_tournaments_id_v1_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTournamentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTournamentResponse) ProtoMessage() {}

func (x *DeleteTournamentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tournaments_id_v1_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTournamentResponse.ProtoReflect.Descriptor instead.
func (*DeleteTournamentResponse) Descriptor() ([]byte, []int) {
	return file_api_tournaments_id_v1_response_proto_rawDescGZIP(), []int{1}
}

var File_api_tournaments_id_v1_response_proto protoreflect.FileDescriptor

var file_api_tournaments_id_v1_response_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x24, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0a, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x1a, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd6, 0x01, 0x0a, 0x19, 0x63, 0x6f,
	0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x69, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x73, 0x63, 0x6f, 0x2d, 0x72, 0x65, 0x2f, 0x70,
	0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x69, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x41, 0x54, 0x49, 0xaa, 0x02, 0x15, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x41, 0x70,
	0x69, 0x5c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x49, 0x64,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x41, 0x70, 0x69, 0x5c, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x49, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x54,
	0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x49, 0x64, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_tournaments_id_v1_response_proto_rawDescOnce sync.Once
	file_api_tournaments_id_v1_response_proto_rawDescData = file_api_tournaments_id_v1_response_proto_rawDesc
)

func file_api_tournaments_id_v1_response_proto_rawDescGZIP() []byte {
	file_api_tournaments_id_v1_response_proto_rawDescOnce.Do(func() {
		file_api_tournaments_id_v1_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_tournaments_id_v1_response_proto_rawDescData)
	})
	return file_api_tournaments_id_v1_response_proto_rawDescData
}

var file_api_tournaments_id_v1_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_tournaments_id_v1_response_proto_goTypes = []any{
	(*GetTournamentResponse)(nil),    // 0: api.tournaments.id.v1.GetTournamentResponse
	(*DeleteTournamentResponse)(nil), // 1: api.tournaments.id.v1.DeleteTournamentResponse
	(*v1.Tournament)(nil),            // 2: model.tournament.v1.Tournament
}
var file_api_tournaments_id_v1_response_proto_depIdxs = []int32{
	2, // 0: api.tournaments.id.v1.GetTournamentResponse.tournament:type_name -> model.tournament.v1.Tournament
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_tournaments_id_v1_response_proto_init() }
func file_api_tournaments_id_v1_response_proto_init() {
	if File_api_tournaments_id_v1_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_tournaments_id_v1_response_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetTournamentResponse); i {
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
		file_api_tournaments_id_v1_response_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTournamentResponse); i {
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
			RawDescriptor: file_api_tournaments_id_v1_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_tournaments_id_v1_response_proto_goTypes,
		DependencyIndexes: file_api_tournaments_id_v1_response_proto_depIdxs,
		MessageInfos:      file_api_tournaments_id_v1_response_proto_msgTypes,
	}.Build()
	File_api_tournaments_id_v1_response_proto = out.File
	file_api_tournaments_id_v1_response_proto_rawDesc = nil
	file_api_tournaments_id_v1_response_proto_goTypes = nil
	file_api_tournaments_id_v1_response_proto_depIdxs = nil
}
