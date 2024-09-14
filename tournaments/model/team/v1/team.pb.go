// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: model/team/v1/team.proto

package teamv1

import (
	v1 "github.com/bufbuild/buf-tour/gen/model/player/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Team struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Players   []*v1.Player           `protobuf:"bytes,5,rep,name=players,proto3" json:"players,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Team) Reset() {
	*x = Team{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_team_v1_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Team) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Team) ProtoMessage() {}

func (x *Team) ProtoReflect() protoreflect.Message {
	mi := &file_model_team_v1_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Team.ProtoReflect.Descriptor instead.
func (*Team) Descriptor() ([]byte, []int) {
	return file_model_team_v1_team_proto_rawDescGZIP(), []int{0}
}

func (x *Team) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Team) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Team) GetPlayers() []*v1.Player {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *Team) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Team) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_model_team_v1_team_proto protoreflect.FileDescriptor

var file_model_team_v1_team_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x04, 0x54, 0x65, 0x61,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xab,
	0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75,
	0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75, 0x72, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x54, 0x58, 0xaa, 0x02,
	0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0d, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x54, 0x65, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x19, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x5c, 0x54, 0x65, 0x61, 0x6d, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x3a, 0x3a, 0x54, 0x65, 0x61, 0x6d, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_team_v1_team_proto_rawDescOnce sync.Once
	file_model_team_v1_team_proto_rawDescData = file_model_team_v1_team_proto_rawDesc
)

func file_model_team_v1_team_proto_rawDescGZIP() []byte {
	file_model_team_v1_team_proto_rawDescOnce.Do(func() {
		file_model_team_v1_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_team_v1_team_proto_rawDescData)
	})
	return file_model_team_v1_team_proto_rawDescData
}

var file_model_team_v1_team_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_team_v1_team_proto_goTypes = []any{
	(*Team)(nil),                  // 0: model.team.v1.Team
	(*v1.Player)(nil),             // 1: model.player.v1.Player
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_model_team_v1_team_proto_depIdxs = []int32{
	1, // 0: model.team.v1.Team.players:type_name -> model.player.v1.Player
	2, // 1: model.team.v1.Team.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: model.team.v1.Team.updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_team_v1_team_proto_init() }
func file_model_team_v1_team_proto_init() {
	if File_model_team_v1_team_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_team_v1_team_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Team); i {
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
			RawDescriptor: file_model_team_v1_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_team_v1_team_proto_goTypes,
		DependencyIndexes: file_model_team_v1_team_proto_depIdxs,
		MessageInfos:      file_model_team_v1_team_proto_msgTypes,
	}.Build()
	File_model_team_v1_team_proto = out.File
	file_model_team_v1_team_proto_rawDesc = nil
	file_model_team_v1_team_proto_goTypes = nil
	file_model_team_v1_team_proto_depIdxs = nil
}
