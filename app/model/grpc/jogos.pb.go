// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: jogos.proto

package grpc

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListaJogos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jogos []*Jogo `protobuf:"bytes,1,rep,name=jogos,proto3" json:"jogos,omitempty"`
}

func (x *ListaJogos) Reset() {
	*x = ListaJogos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jogos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaJogos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaJogos) ProtoMessage() {}

func (x *ListaJogos) ProtoReflect() protoreflect.Message {
	mi := &file_jogos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaJogos.ProtoReflect.Descriptor instead.
func (*ListaJogos) Descriptor() ([]byte, []int) {
	return file_jogos_proto_rawDescGZIP(), []int{0}
}

func (x *ListaJogos) GetJogos() []*Jogo {
	if x != nil {
		return x.Jogos
	}
	return nil
}

var File_jogos_proto protoreflect.FileDescriptor

var file_jogos_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x61,
	0x4a, 0x6f, 0x67, 0x6f, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x67, 0x6f,
	0x52, 0x05, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x32, 0xad, 0x06, 0x0a, 0x05, 0x4a, 0x6f, 0x67, 0x6f,
	0x73, 0x12, 0x52, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x4a, 0x6f, 0x67, 0x6f,
	0x73, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x56, 0x61, 0x7a, 0x69, 0x6f, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x61, 0x4a, 0x6f, 0x67, 0x6f, 0x73, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x12, 0x4e, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x67,
	0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f,
	0x67, 0x6f, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x73,
	0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x5f, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x67,
	0x6f, 0x42, 0x79, 0x4e, 0x6f, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x6d, 0x65, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4a, 0x6f, 0x67, 0x6f, 0x73, 0x22, 0x28, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x2f, 0x6e, 0x6f, 0x6d, 0x65, 0x2f,
	0x7b, 0x4e, 0x6f, 0x6d, 0x65, 0x7d, 0x12, 0x63, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f,
	0x67, 0x6f, 0x42, 0x79, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x6f, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x6f, 0x6d, 0x65, 0x1a, 0x10, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4a, 0x6f, 0x67, 0x6f, 0x73, 0x22,
	0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x24, 0x12, 0x22, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e,
	0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x6f, 0x2f, 0x7b, 0x4e, 0x6f, 0x6d, 0x65, 0x7d, 0x12, 0x61, 0x0a, 0x11, 0x46,
	0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x67, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f,
	0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x1a, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4a, 0x6f,
	0x67, 0x6f, 0x73, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f,
	0x73, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x73,
	0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x4a, 0x6f, 0x67, 0x6f, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x4a, 0x6f, 0x67, 0x6f, 0x73, 0x22, 0x33,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x2b, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65,
	0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x6f, 0x73, 0x2f, 0x7b,
	0x49, 0x44, 0x7d, 0x12, 0x45, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x67,
	0x6f, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x67, 0x6f, 0x1a, 0x0a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x67, 0x6f, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x73, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x67, 0x6f, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4a, 0x6f, 0x67, 0x6f, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4a, 0x6f, 0x67, 0x6f,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x1a, 0x14, 0x2f, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67, 0x6f,
	0x73, 0x12, 0x54, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x67, 0x6f, 0x12,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x1a, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x6f, 0x6f, 0x6c, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x70,
	0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x67,
	0x6f, 0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x69, 0x78, 0x65, 0x6c,
	0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jogos_proto_rawDescOnce sync.Once
	file_jogos_proto_rawDescData = file_jogos_proto_rawDesc
)

func file_jogos_proto_rawDescGZIP() []byte {
	file_jogos_proto_rawDescOnce.Do(func() {
		file_jogos_proto_rawDescData = protoimpl.X.CompressGZIP(file_jogos_proto_rawDescData)
	})
	return file_jogos_proto_rawDescData
}

var file_jogos_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_jogos_proto_goTypes = []any{
	(*ListaJogos)(nil),   // 0: grpc.ListaJogos
	(*Jogo)(nil),         // 1: grpc.Jogo
	(*RequestVazio)(nil), // 2: grpc.RequestVazio
	(*RequestId)(nil),    // 3: grpc.RequestId
	(*RequestNome)(nil),  // 4: grpc.RequestNome
	(*ResponseBool)(nil), // 5: grpc.ResponseBool
}
var file_jogos_proto_depIdxs = []int32{
	1,  // 0: grpc.ListaJogos.jogos:type_name -> grpc.Jogo
	2,  // 1: grpc.Jogos.FindAllJogos:input_type -> grpc.RequestVazio
	3,  // 2: grpc.Jogos.FindJogoById:input_type -> grpc.RequestId
	4,  // 3: grpc.Jogos.FindJogoByNome:input_type -> grpc.RequestNome
	4,  // 4: grpc.Jogos.FindJogoByGenero:input_type -> grpc.RequestNome
	3,  // 5: grpc.Jogos.FindJogoByUsuario:input_type -> grpc.RequestId
	3,  // 6: grpc.Jogos.FindJogoFavoritoByUsuario:input_type -> grpc.RequestId
	1,  // 7: grpc.Jogos.CreateJogo:input_type -> grpc.Jogo
	1,  // 8: grpc.Jogos.UpdateJogo:input_type -> grpc.Jogo
	3,  // 9: grpc.Jogos.DeleteJogo:input_type -> grpc.RequestId
	0,  // 10: grpc.Jogos.FindAllJogos:output_type -> grpc.ListaJogos
	1,  // 11: grpc.Jogos.FindJogoById:output_type -> grpc.Jogo
	0,  // 12: grpc.Jogos.FindJogoByNome:output_type -> grpc.ListaJogos
	0,  // 13: grpc.Jogos.FindJogoByGenero:output_type -> grpc.ListaJogos
	0,  // 14: grpc.Jogos.FindJogoByUsuario:output_type -> grpc.ListaJogos
	0,  // 15: grpc.Jogos.FindJogoFavoritoByUsuario:output_type -> grpc.ListaJogos
	1,  // 16: grpc.Jogos.CreateJogo:output_type -> grpc.Jogo
	1,  // 17: grpc.Jogos.UpdateJogo:output_type -> grpc.Jogo
	5,  // 18: grpc.Jogos.DeleteJogo:output_type -> grpc.ResponseBool
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_jogos_proto_init() }
func file_jogos_proto_init() {
	if File_jogos_proto != nil {
		return
	}
	file_modelos_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_jogos_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListaJogos); i {
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
			RawDescriptor: file_jogos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jogos_proto_goTypes,
		DependencyIndexes: file_jogos_proto_depIdxs,
		MessageInfos:      file_jogos_proto_msgTypes,
	}.Build()
	File_jogos_proto = out.File
	file_jogos_proto_rawDesc = nil
	file_jogos_proto_goTypes = nil
	file_jogos_proto_depIdxs = nil
}
