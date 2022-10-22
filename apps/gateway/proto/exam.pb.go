// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: exam.proto

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

type GetExamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamId string `protobuf:"bytes,1,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
}

func (x *GetExamRequest) Reset() {
	*x = GetExamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamRequest) ProtoMessage() {}

func (x *GetExamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_exam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamRequest.ProtoReflect.Descriptor instead.
func (*GetExamRequest) Descriptor() ([]byte, []int) {
	return file_exam_proto_rawDescGZIP(), []int{0}
}

func (x *GetExamRequest) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

type GetExamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetExamResponse) Reset() {
	*x = GetExamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExamResponse) ProtoMessage() {}

func (x *GetExamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_exam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExamResponse.ProtoReflect.Descriptor instead.
func (*GetExamResponse) Descriptor() ([]byte, []int) {
	return file_exam_proto_rawDescGZIP(), []int{1}
}

type ExamSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExamId      string `protobuf:"bytes,1,opt,name=exam_id,json=examId,proto3" json:"exam_id,omitempty"`
	ExamName    string `protobuf:"bytes,2,opt,name=exam_name,json=examName,proto3" json:"exam_name,omitempty"`
	ThreadCount int32  `protobuf:"varint,3,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
}

func (x *ExamSummary) Reset() {
	*x = ExamSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_exam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExamSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExamSummary) ProtoMessage() {}

func (x *ExamSummary) ProtoReflect() protoreflect.Message {
	mi := &file_exam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExamSummary.ProtoReflect.Descriptor instead.
func (*ExamSummary) Descriptor() ([]byte, []int) {
	return file_exam_proto_rawDescGZIP(), []int{2}
}

func (x *ExamSummary) GetExamId() string {
	if x != nil {
		return x.ExamId
	}
	return ""
}

func (x *ExamSummary) GetExamName() string {
	if x != nil {
		return x.ExamName
	}
	return ""
}

func (x *ExamSummary) GetThreadCount() int32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

var File_exam_proto protoreflect.FileDescriptor

var file_exam_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x78,
	0x61, 0x6d, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x22, 0x11, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x66, 0x0a, 0x0b, 0x45, 0x78, 0x61, 0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x17, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x78, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x61, 0x6d,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x61,
	0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x40, 0x0a, 0x04, 0x45, 0x78, 0x61, 0x6d,
	0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x12, 0x14, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_exam_proto_rawDescOnce sync.Once
	file_exam_proto_rawDescData = file_exam_proto_rawDesc
)

func file_exam_proto_rawDescGZIP() []byte {
	file_exam_proto_rawDescOnce.Do(func() {
		file_exam_proto_rawDescData = protoimpl.X.CompressGZIP(file_exam_proto_rawDescData)
	})
	return file_exam_proto_rawDescData
}

var file_exam_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_exam_proto_goTypes = []interface{}{
	(*GetExamRequest)(nil),  // 0: exam.GetExamRequest
	(*GetExamResponse)(nil), // 1: exam.GetExamResponse
	(*ExamSummary)(nil),     // 2: exam.ExamSummary
}
var file_exam_proto_depIdxs = []int32{
	0, // 0: exam.Exam.GetExam:input_type -> exam.GetExamRequest
	1, // 1: exam.Exam.GetExam:output_type -> exam.GetExamResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_exam_proto_init() }
func file_exam_proto_init() {
	if File_exam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_exam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamRequest); i {
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
		file_exam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExamResponse); i {
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
		file_exam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExamSummary); i {
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
			RawDescriptor: file_exam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_exam_proto_goTypes,
		DependencyIndexes: file_exam_proto_depIdxs,
		MessageInfos:      file_exam_proto_msgTypes,
	}.Build()
	File_exam_proto = out.File
	file_exam_proto_rawDesc = nil
	file_exam_proto_goTypes = nil
	file_exam_proto_depIdxs = nil
}