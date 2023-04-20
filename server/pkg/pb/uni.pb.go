// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: proto/uni.proto

package pb

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

type Student struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
}

func (x *Student) Reset() {
	*x = Student{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Student) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Student) ProtoMessage() {}

func (x *Student) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Student.ProtoReflect.Descriptor instead.
func (*Student) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{0}
}

func (x *Student) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Student) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Student) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

type Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number   int32      `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Students []*Student `protobuf:"bytes,2,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *Group) Reset() {
	*x = Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Group.ProtoReflect.Descriptor instead.
func (*Group) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{1}
}

func (x *Group) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Group) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type Professor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *Professor) Reset() {
	*x = Professor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Professor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Professor) ProtoMessage() {}

func (x *Professor) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Professor.ProtoReflect.Descriptor instead.
func (*Professor) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{2}
}

func (x *Professor) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Professor) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Professor) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *Professor) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Groups     []*Group     `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	Professors []*Professor `protobuf:"bytes,3,rep,name=professors,proto3" json:"professors,omitempty"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{3}
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *Department) GetProfessors() []*Professor {
	if x != nil {
		return x.Professors
	}
	return nil
}

type GetStudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{4}
}

func (x *GetStudentRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetStudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Student *Student `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"`
}

func (x *GetStudentResponse) Reset() {
	*x = GetStudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentResponse) ProtoMessage() {}

func (x *GetStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentResponse.ProtoReflect.Descriptor instead.
func (*GetStudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{5}
}

func (x *GetStudentResponse) GetStudent() *Student {
	if x != nil {
		return x.Student
	}
	return nil
}

type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type GetGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetGroupResponse) Reset() {
	*x = GetGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupResponse) ProtoMessage() {}

func (x *GetGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupResponse.ProtoReflect.Descriptor instead.
func (*GetGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{7}
}

func (x *GetGroupResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

type GetProfessorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProfessorRequest) Reset() {
	*x = GetProfessorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfessorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfessorRequest) ProtoMessage() {}

func (x *GetProfessorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfessorRequest.ProtoReflect.Descriptor instead.
func (*GetProfessorRequest) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{8}
}

func (x *GetProfessorRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProfessorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Professor *Professor `protobuf:"bytes,1,opt,name=professor,proto3" json:"professor,omitempty"`
}

func (x *GetProfessorResponse) Reset() {
	*x = GetProfessorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfessorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfessorResponse) ProtoMessage() {}

func (x *GetProfessorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfessorResponse.ProtoReflect.Descriptor instead.
func (*GetProfessorResponse) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{9}
}

func (x *GetProfessorResponse) GetProfessor() *Professor {
	if x != nil {
		return x.Professor
	}
	return nil
}

type GetDepartmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetDepartmentRequest) Reset() {
	*x = GetDepartmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepartmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentRequest) ProtoMessage() {}

func (x *GetDepartmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentRequest.ProtoReflect.Descriptor instead.
func (*GetDepartmentRequest) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{10}
}

func (x *GetDepartmentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetDepartmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups     []*Group     `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	Professors []*Professor `protobuf:"bytes,2,rep,name=professors,proto3" json:"professors,omitempty"`
}

func (x *GetDepartmentResponse) Reset() {
	*x = GetDepartmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDepartmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDepartmentResponse) ProtoMessage() {}

func (x *GetDepartmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDepartmentResponse.ProtoReflect.Descriptor instead.
func (*GetDepartmentResponse) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{11}
}

func (x *GetDepartmentResponse) GetGroups() []*Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetDepartmentResponse) GetProfessors() []*Professor {
	if x != nil {
		return x.Professors
	}
	return nil
}

type GetStudentAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStudentAllRequest) Reset() {
	*x = GetStudentAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentAllRequest) ProtoMessage() {}

func (x *GetStudentAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentAllRequest.ProtoReflect.Descriptor instead.
func (*GetStudentAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{12}
}

type GetStudentAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students []*Student `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (x *GetStudentAllResponse) Reset() {
	*x = GetStudentAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_uni_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStudentAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentAllResponse) ProtoMessage() {}

func (x *GetStudentAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_uni_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentAllResponse.ProtoReflect.Descriptor instead.
func (*GetStudentAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_uni_proto_rawDescGZIP(), []int{13}
}

func (x *GetStudentAllResponse) GetStudents() []*Student {
	if x != nil {
		return x.Students
	}
	return nil
}

var File_proto_uni_proto protoreflect.FileDescriptor

var file_proto_uni_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x47, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x05, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x08, 0x73,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x63, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x6c, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66,
	0x65, 0x73, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x38,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x65,
	0x73, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x63, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73,
	0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f,
	0x72, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x48, 0x0a, 0x0d, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x32, 0x40, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65,
	0x6e, 0x74, 0x73, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x54, 0x0a, 0x10, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x40, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_uni_proto_rawDescOnce sync.Once
	file_proto_uni_proto_rawDescData = file_proto_uni_proto_rawDesc
)

func file_proto_uni_proto_rawDescGZIP() []byte {
	file_proto_uni_proto_rawDescOnce.Do(func() {
		file_proto_uni_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_uni_proto_rawDescData)
	})
	return file_proto_uni_proto_rawDescData
}

var file_proto_uni_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_proto_uni_proto_goTypes = []interface{}{
	(*Student)(nil),               // 0: Student
	(*Group)(nil),                 // 1: Group
	(*Professor)(nil),             // 2: Professor
	(*Department)(nil),            // 3: Department
	(*GetStudentRequest)(nil),     // 4: GetStudentRequest
	(*GetStudentResponse)(nil),    // 5: GetStudentResponse
	(*GetGroupRequest)(nil),       // 6: GetGroupRequest
	(*GetGroupResponse)(nil),      // 7: GetGroupResponse
	(*GetProfessorRequest)(nil),   // 8: GetProfessorRequest
	(*GetProfessorResponse)(nil),  // 9: GetProfessorResponse
	(*GetDepartmentRequest)(nil),  // 10: GetDepartmentRequest
	(*GetDepartmentResponse)(nil), // 11: GetDepartmentResponse
	(*GetStudentAllRequest)(nil),  // 12: GetStudentAllRequest
	(*GetStudentAllResponse)(nil), // 13: GetStudentAllResponse
}
var file_proto_uni_proto_depIdxs = []int32{
	0,  // 0: Group.students:type_name -> Student
	1,  // 1: Department.groups:type_name -> Group
	2,  // 2: Department.professors:type_name -> Professor
	0,  // 3: GetStudentResponse.student:type_name -> Student
	0,  // 4: GetGroupResponse.students:type_name -> Student
	2,  // 5: GetProfessorResponse.professor:type_name -> Professor
	1,  // 6: GetDepartmentResponse.groups:type_name -> Group
	2,  // 7: GetDepartmentResponse.professors:type_name -> Professor
	0,  // 8: GetStudentAllResponse.students:type_name -> Student
	4,  // 9: StudentSearch.GetStudent:input_type -> GetStudentRequest
	6,  // 10: GroupSearch.GetGroup:input_type -> GetGroupRequest
	8,  // 11: ProfessorSearch.GetProfessor:input_type -> GetProfessorRequest
	12, // 12: StudentsSearch.GetStudents:input_type -> GetStudentAllRequest
	10, // 13: DepartmentSearch.GetDepartment:input_type -> GetDepartmentRequest
	5,  // 14: StudentSearch.GetStudent:output_type -> GetStudentResponse
	7,  // 15: GroupSearch.GetGroup:output_type -> GetGroupResponse
	9,  // 16: ProfessorSearch.GetProfessor:output_type -> GetProfessorResponse
	13, // 17: StudentsSearch.GetStudents:output_type -> GetStudentAllResponse
	11, // 18: DepartmentSearch.GetDepartment:output_type -> GetDepartmentResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_proto_uni_proto_init() }
func file_proto_uni_proto_init() {
	if File_proto_uni_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_uni_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Student); i {
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
		file_proto_uni_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Group); i {
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
		file_proto_uni_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Professor); i {
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
		file_proto_uni_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
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
		file_proto_uni_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentRequest); i {
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
		file_proto_uni_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentResponse); i {
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
		file_proto_uni_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_proto_uni_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupResponse); i {
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
		file_proto_uni_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfessorRequest); i {
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
		file_proto_uni_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfessorResponse); i {
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
		file_proto_uni_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepartmentRequest); i {
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
		file_proto_uni_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDepartmentResponse); i {
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
		file_proto_uni_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentAllRequest); i {
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
		file_proto_uni_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStudentAllResponse); i {
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
			RawDescriptor: file_proto_uni_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   5,
		},
		GoTypes:           file_proto_uni_proto_goTypes,
		DependencyIndexes: file_proto_uni_proto_depIdxs,
		MessageInfos:      file_proto_uni_proto_msgTypes,
	}.Build()
	File_proto_uni_proto = out.File
	file_proto_uni_proto_rawDesc = nil
	file_proto_uni_proto_goTypes = nil
	file_proto_uni_proto_depIdxs = nil
}
