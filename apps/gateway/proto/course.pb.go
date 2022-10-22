// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: course.proto

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

type GetAllCourseSummaryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCourseSummaryRequest) Reset() {
	*x = GetAllCourseSummaryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCourseSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCourseSummaryRequest) ProtoMessage() {}

func (x *GetAllCourseSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCourseSummaryRequest.ProtoReflect.Descriptor instead.
func (*GetAllCourseSummaryRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{0}
}

type GetAllCourseSummaryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*CourseSummary `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *GetAllCourseSummaryResponse) Reset() {
	*x = GetAllCourseSummaryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCourseSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCourseSummaryResponse) ProtoMessage() {}

func (x *GetAllCourseSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCourseSummaryResponse.ProtoReflect.Descriptor instead.
func (*GetAllCourseSummaryResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllCourseSummaryResponse) GetCourses() []*CourseSummary {
	if x != nil {
		return x.Courses
	}
	return nil
}

type CourseSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId    string `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	CourseName  string `protobuf:"bytes,2,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	ThreadCount int32  `protobuf:"varint,3,opt,name=thread_count,json=threadCount,proto3" json:"thread_count,omitempty"`
}

func (x *CourseSummary) Reset() {
	*x = CourseSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CourseSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CourseSummary) ProtoMessage() {}

func (x *CourseSummary) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CourseSummary.ProtoReflect.Descriptor instead.
func (*CourseSummary) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{2}
}

func (x *CourseSummary) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

func (x *CourseSummary) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *CourseSummary) GetThreadCount() int32 {
	if x != nil {
		return x.ThreadCount
	}
	return 0
}

type GetCourseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId string `protobuf:"bytes,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *GetCourseRequest) Reset() {
	*x = GetCourseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseRequest) ProtoMessage() {}

func (x *GetCourseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseRequest.ProtoReflect.Descriptor instead.
func (*GetCourseRequest) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{3}
}

func (x *GetCourseRequest) GetCourseId() string {
	if x != nil {
		return x.CourseId
	}
	return ""
}

type GetCourseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseName string         `protobuf:"bytes,1,opt,name=course_name,json=courseName,proto3" json:"course_name,omitempty"`
	Exams      []*ExamSummary `protobuf:"bytes,2,rep,name=exams,proto3" json:"exams,omitempty"`
}

func (x *GetCourseResponse) Reset() {
	*x = GetCourseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCourseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCourseResponse) ProtoMessage() {}

func (x *GetCourseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_course_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCourseResponse.ProtoReflect.Descriptor instead.
func (*GetCourseResponse) Descriptor() ([]byte, []int) {
	return file_course_proto_rawDescGZIP(), []int{4}
}

func (x *GetCourseResponse) GetCourseName() string {
	if x != nil {
		return x.CourseName
	}
	return ""
}

func (x *GetCourseResponse) GetExams() []*ExamSummary {
	if x != nil {
		return x.Exams
	}
	return nil
}

var File_course_proto protoreflect.FileDescriptor

var file_course_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0a, 0x65, 0x78, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x4e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x22, 0x70, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x65, 0x78, 0x61,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x2e,
	0x45, 0x78, 0x61, 0x6d, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x05, 0x65, 0x78, 0x61,
	0x6d, 0x73, 0x32, 0xae, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x60, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x18, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_proto_rawDescOnce sync.Once
	file_course_proto_rawDescData = file_course_proto_rawDesc
)

func file_course_proto_rawDescGZIP() []byte {
	file_course_proto_rawDescOnce.Do(func() {
		file_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_proto_rawDescData)
	})
	return file_course_proto_rawDescData
}

var file_course_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_course_proto_goTypes = []interface{}{
	(*GetAllCourseSummaryRequest)(nil),  // 0: course.GetAllCourseSummaryRequest
	(*GetAllCourseSummaryResponse)(nil), // 1: course.GetAllCourseSummaryResponse
	(*CourseSummary)(nil),               // 2: course.CourseSummary
	(*GetCourseRequest)(nil),            // 3: course.GetCourseRequest
	(*GetCourseResponse)(nil),           // 4: course.GetCourseResponse
	(*ExamSummary)(nil),                 // 5: exam.ExamSummary
}
var file_course_proto_depIdxs = []int32{
	2, // 0: course.GetAllCourseSummaryResponse.courses:type_name -> course.CourseSummary
	5, // 1: course.GetCourseResponse.exams:type_name -> exam.ExamSummary
	0, // 2: course.Course.GetAllCourseSummary:input_type -> course.GetAllCourseSummaryRequest
	3, // 3: course.Course.GetCourse:input_type -> course.GetCourseRequest
	1, // 4: course.Course.GetAllCourseSummary:output_type -> course.GetAllCourseSummaryResponse
	4, // 5: course.Course.GetCourse:output_type -> course.GetCourseResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_course_proto_init() }
func file_course_proto_init() {
	if File_course_proto != nil {
		return
	}
	file_exam_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCourseSummaryRequest); i {
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
		file_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCourseSummaryResponse); i {
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
		file_course_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CourseSummary); i {
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
		file_course_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseRequest); i {
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
		file_course_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCourseResponse); i {
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
			RawDescriptor: file_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_proto_goTypes,
		DependencyIndexes: file_course_proto_depIdxs,
		MessageInfos:      file_course_proto_msgTypes,
	}.Build()
	File_course_proto = out.File
	file_course_proto_rawDesc = nil
	file_course_proto_goTypes = nil
	file_course_proto_depIdxs = nil
}
