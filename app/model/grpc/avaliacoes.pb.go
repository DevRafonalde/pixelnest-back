// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.1
// source: avaliacoes.proto

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

type ListaAvaliacoes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Avaliacoes []*Avaliacao `protobuf:"bytes,1,rep,name=avaliacoes,proto3" json:"avaliacoes,omitempty"`
}

func (x *ListaAvaliacoes) Reset() {
	*x = ListaAvaliacoes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_avaliacoes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListaAvaliacoes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListaAvaliacoes) ProtoMessage() {}

func (x *ListaAvaliacoes) ProtoReflect() protoreflect.Message {
	mi := &file_avaliacoes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListaAvaliacoes.ProtoReflect.Descriptor instead.
func (*ListaAvaliacoes) Descriptor() ([]byte, []int) {
	return file_avaliacoes_proto_rawDescGZIP(), []int{0}
}

func (x *ListaAvaliacoes) GetAvaliacoes() []*Avaliacao {
	if x != nil {
		return x.Avaliacoes
	}
	return nil
}

var File_avaliacoes_proto protoreflect.FileDescriptor

var file_avaliacoes_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x6f,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x61, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x12,
	0x2f, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x61, 0x63, 0x61, 0x6f, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73,
	0x32, 0xb4, 0x06, 0x0a, 0x0a, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x12,
	0x61, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61,
	0x63, 0x6f, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x56, 0x61, 0x7a, 0x69, 0x6f, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x61, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x22,
	0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e,
	0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f,
	0x65, 0x73, 0x12, 0x5d, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61,
	0x63, 0x61, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x12, 0x1e, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x12, 0x70, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63,
	0x61, 0x6f, 0x42, 0x79, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63,
	0x6f, 0x65, 0x73, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76, 0x61, 0x6c,
	0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x2f, 0x7b,
	0x49, 0x44, 0x7d, 0x12, 0x70, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x61, 0x63, 0x61, 0x6f, 0x42, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x12, 0x0f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x15,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x61, 0x63, 0x6f, 0x65, 0x73, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f,
	0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76,
	0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f,
	0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x6a, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x76, 0x61,
	0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f, 0x42, 0x79, 0x4a, 0x6f, 0x67, 0x6f, 0x12, 0x0f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x15, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x61, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61,
	0x63, 0x6f, 0x65, 0x73, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x70,
	0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76, 0x61,
	0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x2f, 0x6a, 0x6f, 0x67, 0x6f, 0x2f, 0x7b, 0x49, 0x44,
	0x7d, 0x12, 0x59, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x6c, 0x69,
	0x61, 0x63, 0x61, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x76, 0x61, 0x6c,
	0x69, 0x61, 0x63, 0x61, 0x6f, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x76, 0x61,
	0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01,
	0x2a, 0x22, 0x19, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f, 0x12,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f,
	0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x61,
	0x6f, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x1a, 0x19, 0x2f, 0x70,
	0x69, 0x78, 0x65, 0x6c, 0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76, 0x61,
	0x6c, 0x69, 0x61, 0x63, 0x6f, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x61, 0x6f, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x22,
	0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x2a, 0x1e, 0x2f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x6e,
	0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x61, 0x63, 0x6f,
	0x65, 0x73, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x69, 0x78, 0x65, 0x6c,
	0x6e, 0x65, 0x73, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_avaliacoes_proto_rawDescOnce sync.Once
	file_avaliacoes_proto_rawDescData = file_avaliacoes_proto_rawDesc
)

func file_avaliacoes_proto_rawDescGZIP() []byte {
	file_avaliacoes_proto_rawDescOnce.Do(func() {
		file_avaliacoes_proto_rawDescData = protoimpl.X.CompressGZIP(file_avaliacoes_proto_rawDescData)
	})
	return file_avaliacoes_proto_rawDescData
}

var file_avaliacoes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_avaliacoes_proto_goTypes = []any{
	(*ListaAvaliacoes)(nil), // 0: grpc.ListaAvaliacoes
	(*Avaliacao)(nil),       // 1: grpc.Avaliacao
	(*RequestVazio)(nil),    // 2: grpc.RequestVazio
	(*RequestId)(nil),       // 3: grpc.RequestId
	(*ResponseBool)(nil),    // 4: grpc.ResponseBool
}
var file_avaliacoes_proto_depIdxs = []int32{
	1, // 0: grpc.ListaAvaliacoes.avaliacoes:type_name -> grpc.Avaliacao
	2, // 1: grpc.Avaliacoes.FindAllAvaliacoes:input_type -> grpc.RequestVazio
	3, // 2: grpc.Avaliacoes.FindAvaliacaoById:input_type -> grpc.RequestId
	3, // 3: grpc.Avaliacoes.FindAvaliacaoByUsuario:input_type -> grpc.RequestId
	3, // 4: grpc.Avaliacoes.FindAvaliacaoByProduto:input_type -> grpc.RequestId
	3, // 5: grpc.Avaliacoes.FindAvaliacaoByJogo:input_type -> grpc.RequestId
	1, // 6: grpc.Avaliacoes.CreateAvaliacao:input_type -> grpc.Avaliacao
	1, // 7: grpc.Avaliacoes.UpdateAvaliacao:input_type -> grpc.Avaliacao
	3, // 8: grpc.Avaliacoes.DeleteAvaliacao:input_type -> grpc.RequestId
	0, // 9: grpc.Avaliacoes.FindAllAvaliacoes:output_type -> grpc.ListaAvaliacoes
	1, // 10: grpc.Avaliacoes.FindAvaliacaoById:output_type -> grpc.Avaliacao
	0, // 11: grpc.Avaliacoes.FindAvaliacaoByUsuario:output_type -> grpc.ListaAvaliacoes
	0, // 12: grpc.Avaliacoes.FindAvaliacaoByProduto:output_type -> grpc.ListaAvaliacoes
	0, // 13: grpc.Avaliacoes.FindAvaliacaoByJogo:output_type -> grpc.ListaAvaliacoes
	1, // 14: grpc.Avaliacoes.CreateAvaliacao:output_type -> grpc.Avaliacao
	1, // 15: grpc.Avaliacoes.UpdateAvaliacao:output_type -> grpc.Avaliacao
	4, // 16: grpc.Avaliacoes.DeleteAvaliacao:output_type -> grpc.ResponseBool
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_avaliacoes_proto_init() }
func file_avaliacoes_proto_init() {
	if File_avaliacoes_proto != nil {
		return
	}
	file_modelos_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_avaliacoes_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListaAvaliacoes); i {
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
			RawDescriptor: file_avaliacoes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_avaliacoes_proto_goTypes,
		DependencyIndexes: file_avaliacoes_proto_depIdxs,
		MessageInfos:      file_avaliacoes_proto_msgTypes,
	}.Build()
	File_avaliacoes_proto = out.File
	file_avaliacoes_proto_rawDesc = nil
	file_avaliacoes_proto_goTypes = nil
	file_avaliacoes_proto_depIdxs = nil
}
