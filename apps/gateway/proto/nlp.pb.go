// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: nlp.proto

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

type TokenizeParagraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Para string `protobuf:"bytes,1,opt,name=para,proto3" json:"para,omitempty"`
}

func (x *TokenizeParagraph) Reset() {
	*x = TokenizeParagraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeParagraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeParagraph) ProtoMessage() {}

func (x *TokenizeParagraph) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeParagraph.ProtoReflect.Descriptor instead.
func (*TokenizeParagraph) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{0}
}

func (x *TokenizeParagraph) GetPara() string {
	if x != nil {
		return x.Para
	}
	return ""
}

type EmbedSentence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentence string `protobuf:"bytes,1,opt,name=sentence,proto3" json:"sentence,omitempty"`
}

func (x *EmbedSentence) Reset() {
	*x = EmbedSentence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmbedSentence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmbedSentence) ProtoMessage() {}

func (x *EmbedSentence) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmbedSentence.ProtoReflect.Descriptor instead.
func (*EmbedSentence) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{1}
}

func (x *EmbedSentence) GetSentence() string {
	if x != nil {
		return x.Sentence
	}
	return ""
}

type MeasureVectors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector1 []int32 `protobuf:"varint,1,rep,packed,name=vector1,proto3" json:"vector1,omitempty"`
	Vector2 []int32 `protobuf:"varint,2,rep,packed,name=vector2,proto3" json:"vector2,omitempty"`
}

func (x *MeasureVectors) Reset() {
	*x = MeasureVectors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasureVectors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasureVectors) ProtoMessage() {}

func (x *MeasureVectors) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasureVectors.ProtoReflect.Descriptor instead.
func (*MeasureVectors) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{2}
}

func (x *MeasureVectors) GetVector1() []int32 {
	if x != nil {
		return x.Vector1
	}
	return nil
}

func (x *MeasureVectors) GetVector2() []int32 {
	if x != nil {
		return x.Vector2
	}
	return nil
}

type TokenizeSentences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sentences []string `protobuf:"bytes,1,rep,name=sentences,proto3" json:"sentences,omitempty"`
}

func (x *TokenizeSentences) Reset() {
	*x = TokenizeSentences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenizeSentences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenizeSentences) ProtoMessage() {}

func (x *TokenizeSentences) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenizeSentences.ProtoReflect.Descriptor instead.
func (*TokenizeSentences) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{3}
}

func (x *TokenizeSentences) GetSentences() []string {
	if x != nil {
		return x.Sentences
	}
	return nil
}

type EncodedSentence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector []int32 `protobuf:"varint,1,rep,packed,name=vector,proto3" json:"vector,omitempty"`
}

func (x *EncodedSentence) Reset() {
	*x = EncodedSentence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncodedSentence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncodedSentence) ProtoMessage() {}

func (x *EncodedSentence) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncodedSentence.ProtoReflect.Descriptor instead.
func (*EncodedSentence) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{4}
}

func (x *EncodedSentence) GetVector() []int32 {
	if x != nil {
		return x.Vector
	}
	return nil
}

type MeasureSimilarity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Similarity float32 `protobuf:"fixed32,1,opt,name=similarity,proto3" json:"similarity,omitempty"`
}

func (x *MeasureSimilarity) Reset() {
	*x = MeasureSimilarity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nlp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeasureSimilarity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeasureSimilarity) ProtoMessage() {}

func (x *MeasureSimilarity) ProtoReflect() protoreflect.Message {
	mi := &file_nlp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeasureSimilarity.ProtoReflect.Descriptor instead.
func (*MeasureSimilarity) Descriptor() ([]byte, []int) {
	return file_nlp_proto_rawDescGZIP(), []int{5}
}

func (x *MeasureSimilarity) GetSimilarity() float32 {
	if x != nil {
		return x.Similarity
	}
	return 0
}

var File_nlp_proto protoreflect.FileDescriptor

var file_nlp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x6c, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6e, 0x6c, 0x70,
	0x22, 0x27, 0x0a, 0x11, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x72, 0x61, 0x22, 0x2b, 0x0a, 0x0d, 0x45, 0x6d, 0x62,
	0x65, 0x64, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x0e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72,
	0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x31, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x32, 0x22, 0x31, 0x0a, 0x11,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x29, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0x33, 0x0a, 0x11, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x32,
	0xb2, 0x01, 0x0a, 0x03, 0x4e, 0x6c, 0x70, 0x12, 0x3c, 0x0a, 0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x6e, 0x6c, 0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69,
	0x7a, 0x65, 0x50, 0x61, 0x72, 0x61, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x16, 0x2e, 0x6e, 0x6c,
	0x70, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x7a, 0x65, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x12,
	0x2e, 0x6e, 0x6c, 0x70, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e,
	0x63, 0x65, 0x1a, 0x14, 0x2e, 0x6e, 0x6c, 0x70, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x53, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x4d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x12, 0x13, 0x2e, 0x6e, 0x6c, 0x70, 0x2e, 0x4d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x16, 0x2e, 0x6e, 0x6c, 0x70,
	0x2e, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x53, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nlp_proto_rawDescOnce sync.Once
	file_nlp_proto_rawDescData = file_nlp_proto_rawDesc
)

func file_nlp_proto_rawDescGZIP() []byte {
	file_nlp_proto_rawDescOnce.Do(func() {
		file_nlp_proto_rawDescData = protoimpl.X.CompressGZIP(file_nlp_proto_rawDescData)
	})
	return file_nlp_proto_rawDescData
}

var file_nlp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_nlp_proto_goTypes = []interface{}{
	(*TokenizeParagraph)(nil), // 0: nlp.TokenizeParagraph
	(*EmbedSentence)(nil),     // 1: nlp.EmbedSentence
	(*MeasureVectors)(nil),    // 2: nlp.MeasureVectors
	(*TokenizeSentences)(nil), // 3: nlp.TokenizeSentences
	(*EncodedSentence)(nil),   // 4: nlp.EncodedSentence
	(*MeasureSimilarity)(nil), // 5: nlp.MeasureSimilarity
}
var file_nlp_proto_depIdxs = []int32{
	0, // 0: nlp.Nlp.Tokenize:input_type -> nlp.TokenizeParagraph
	1, // 1: nlp.Nlp.Embed:input_type -> nlp.EmbedSentence
	2, // 2: nlp.Nlp.Measure:input_type -> nlp.MeasureVectors
	3, // 3: nlp.Nlp.Tokenize:output_type -> nlp.TokenizeSentences
	4, // 4: nlp.Nlp.Embed:output_type -> nlp.EncodedSentence
	5, // 5: nlp.Nlp.Measure:output_type -> nlp.MeasureSimilarity
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_nlp_proto_init() }
func file_nlp_proto_init() {
	if File_nlp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nlp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeParagraph); i {
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
		file_nlp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmbedSentence); i {
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
		file_nlp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasureVectors); i {
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
		file_nlp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenizeSentences); i {
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
		file_nlp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncodedSentence); i {
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
		file_nlp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeasureSimilarity); i {
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
			RawDescriptor: file_nlp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nlp_proto_goTypes,
		DependencyIndexes: file_nlp_proto_depIdxs,
		MessageInfos:      file_nlp_proto_msgTypes,
	}.Build()
	File_nlp_proto = out.File
	file_nlp_proto_rawDesc = nil
	file_nlp_proto_goTypes = nil
	file_nlp_proto_depIdxs = nil
}