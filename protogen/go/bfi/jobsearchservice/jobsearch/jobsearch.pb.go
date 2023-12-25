// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: proto/jobsearch/jobsearch.proto

package jobsearch

import (
	first "github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/firstService/first"
	second "github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/secondservice/second"
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

type JobSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCandidateId uint32              `protobuf:"varint,1,opt,name=job_candidate_id,proto3" json:"job_candidate_id,omitempty"`
	Application    *second.Application `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *JobSearch) Reset() {
	*x = JobSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSearch) ProtoMessage() {}

func (x *JobSearch) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSearch.ProtoReflect.Descriptor instead.
func (*JobSearch) Descriptor() ([]byte, []int) {
	return file_proto_jobsearch_jobsearch_proto_rawDescGZIP(), []int{0}
}

func (x *JobSearch) GetJobCandidateId() uint32 {
	if x != nil {
		return x.JobCandidateId
	}
	return 0
}

func (x *JobSearch) GetApplication() *second.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

type JobSofware struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobCandidateId uint32             `protobuf:"varint,1,opt,name=job_candidate_id,proto3" json:"job_candidate_id,omitempty"`
	Application    *first.Application `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
}

func (x *JobSofware) Reset() {
	*x = JobSofware{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobSofware) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobSofware) ProtoMessage() {}

func (x *JobSofware) ProtoReflect() protoreflect.Message {
	mi := &file_proto_jobsearch_jobsearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobSofware.ProtoReflect.Descriptor instead.
func (*JobSofware) Descriptor() ([]byte, []int) {
	return file_proto_jobsearch_jobsearch_proto_rawDescGZIP(), []int{1}
}

func (x *JobSofware) GetJobCandidateId() uint32 {
	if x != nil {
		return x.JobCandidateId
	}
	return 0
}

func (x *JobSofware) GetApplication() *first.Application {
	if x != nil {
		return x.Application
	}
	return nil
}

var File_proto_jobsearch_jobsearch_proto protoreflect.FileDescriptor

var file_proto_jobsearch_jobsearch_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x09, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x2a, 0x0a,
	0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x0b, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x6e, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x6f, 0x66, 0x77, 0x61, 0x72, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6a, 0x6f, 0x62, 0x5f, 0x63, 0x61,
	0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x69, 0x7a, 0x6b, 0x69, 0x72, 0x61, 0x6d, 0x61, 0x64, 0x68, 0x61, 0x6e, 0x32, 0x30, 0x31, 0x39,
	0x38, 0x37, 0x2f, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x66, 0x69, 0x2f, 0x6a, 0x6f,
	0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_jobsearch_jobsearch_proto_rawDescOnce sync.Once
	file_proto_jobsearch_jobsearch_proto_rawDescData = file_proto_jobsearch_jobsearch_proto_rawDesc
)

func file_proto_jobsearch_jobsearch_proto_rawDescGZIP() []byte {
	file_proto_jobsearch_jobsearch_proto_rawDescOnce.Do(func() {
		file_proto_jobsearch_jobsearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_jobsearch_jobsearch_proto_rawDescData)
	})
	return file_proto_jobsearch_jobsearch_proto_rawDescData
}

var file_proto_jobsearch_jobsearch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_jobsearch_jobsearch_proto_goTypes = []interface{}{
	(*JobSearch)(nil),          // 0: JobSearch
	(*JobSofware)(nil),         // 1: JobSofware
	(*second.Application)(nil), // 2: second.Application
	(*first.Application)(nil),  // 3: first.Application
}
var file_proto_jobsearch_jobsearch_proto_depIdxs = []int32{
	2, // 0: JobSearch.application:type_name -> second.Application
	3, // 1: JobSofware.application:type_name -> first.Application
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_jobsearch_jobsearch_proto_init() }
func file_proto_jobsearch_jobsearch_proto_init() {
	if File_proto_jobsearch_jobsearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_jobsearch_jobsearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSearch); i {
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
		file_proto_jobsearch_jobsearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobSofware); i {
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
			RawDescriptor: file_proto_jobsearch_jobsearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_jobsearch_jobsearch_proto_goTypes,
		DependencyIndexes: file_proto_jobsearch_jobsearch_proto_depIdxs,
		MessageInfos:      file_proto_jobsearch_jobsearch_proto_msgTypes,
	}.Build()
	File_proto_jobsearch_jobsearch_proto = out.File
	file_proto_jobsearch_jobsearch_proto_rawDesc = nil
	file_proto_jobsearch_jobsearch_proto_goTypes = nil
	file_proto_jobsearch_jobsearch_proto_depIdxs = nil
}
