// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// UserInfo : 用户基本信息
type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 int64    `protobuf:"varint,4,opt,name=role,proto3" json:"role,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  bool     `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
	Phone                string   `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Email                string   `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	School               string   `protobuf:"bytes,9,opt,name=school,proto3" json:"school,omitempty"`
	LastLogin            int64    `protobuf:"varint,10,opt,name=last_login,json=lastLogin,proto3" json:"last_login,omitempty"`
	Create               int64    `protobuf:"varint,11,opt,name=create,proto3" json:"create,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserInfo) GetRole() int64 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *UserInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserInfo) GetSex() bool {
	if m != nil {
		return m.Sex
	}
	return false
}

func (m *UserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *UserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserInfo) GetSchool() string {
	if m != nil {
		return m.School
	}
	return ""
}

func (m *UserInfo) GetLastLogin() int64 {
	if m != nil {
		return m.LastLogin
	}
	return 0
}

func (m *UserInfo) GetCreate() int64 {
	if m != nil {
		return m.Create
	}
	return 0
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// ProblemExample : 题目输入输出样例
type ProblemExample struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Input                string   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	Output               string   `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProblemExample) Reset()         { *m = ProblemExample{} }
func (m *ProblemExample) String() string { return proto.CompactTextString(m) }
func (*ProblemExample) ProtoMessage()    {}
func (*ProblemExample) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *ProblemExample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProblemExample.Unmarshal(m, b)
}
func (m *ProblemExample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProblemExample.Marshal(b, m, deterministic)
}
func (m *ProblemExample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProblemExample.Merge(m, src)
}
func (m *ProblemExample) XXX_Size() int {
	return xxx_messageInfo_ProblemExample.Size(m)
}
func (m *ProblemExample) XXX_DiscardUnknown() {
	xxx_messageInfo_ProblemExample.DiscardUnknown(m)
}

var xxx_messageInfo_ProblemExample proto.InternalMessageInfo

func (m *ProblemExample) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ProblemExample) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *ProblemExample) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

// Problem : 题目
type Problem struct {
	Id                   int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	In                   string            `protobuf:"bytes,4,opt,name=in,proto3" json:"in,omitempty"`
	Out                  string            `protobuf:"bytes,5,opt,name=out,proto3" json:"out,omitempty"`
	Hint                 string            `protobuf:"bytes,6,opt,name=hint,proto3" json:"hint,omitempty"`
	InOutExamples        []*ProblemExample `protobuf:"bytes,7,rep,name=in_out_examples,json=inOutExamples,proto3" json:"in_out_examples,omitempty"`
	JudgeLimitTime       int64             `protobuf:"varint,8,opt,name=judge_limit_time,json=judgeLimitTime,proto3" json:"judge_limit_time,omitempty"`
	JudgeLimitMem        int64             `protobuf:"varint,9,opt,name=judge_limit_mem,json=judgeLimitMem,proto3" json:"judge_limit_mem,omitempty"`
	Tags                 []int64           `protobuf:"varint,10,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	Difficulty           int64             `protobuf:"varint,11,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Cognition            int64             `protobuf:"varint,12,opt,name=cognition,proto3" json:"cognition,omitempty"`
	SubmitTime           int64             `protobuf:"varint,13,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	AcceptTime           int64             `protobuf:"varint,14,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
	CreateTime           int64             `protobuf:"varint,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Publisher            string            `protobuf:"bytes,16,opt,name=publisher,proto3" json:"publisher,omitempty"`
	JudgeFile            string            `protobuf:"bytes,17,opt,name=judge_file,json=judgeFile,proto3" json:"judge_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Problem) Reset()         { *m = Problem{} }
func (m *Problem) String() string { return proto.CompactTextString(m) }
func (*Problem) ProtoMessage()    {}
func (*Problem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{2}
}

func (m *Problem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Problem.Unmarshal(m, b)
}
func (m *Problem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Problem.Marshal(b, m, deterministic)
}
func (m *Problem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Problem.Merge(m, src)
}
func (m *Problem) XXX_Size() int {
	return xxx_messageInfo_Problem.Size(m)
}
func (m *Problem) XXX_DiscardUnknown() {
	xxx_messageInfo_Problem.DiscardUnknown(m)
}

var xxx_messageInfo_Problem proto.InternalMessageInfo

func (m *Problem) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Problem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Problem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Problem) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Problem) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

func (m *Problem) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *Problem) GetInOutExamples() []*ProblemExample {
	if m != nil {
		return m.InOutExamples
	}
	return nil
}

func (m *Problem) GetJudgeLimitTime() int64 {
	if m != nil {
		return m.JudgeLimitTime
	}
	return 0
}

func (m *Problem) GetJudgeLimitMem() int64 {
	if m != nil {
		return m.JudgeLimitMem
	}
	return 0
}

func (m *Problem) GetTags() []int64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Problem) GetDifficulty() int64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Problem) GetCognition() int64 {
	if m != nil {
		return m.Cognition
	}
	return 0
}

func (m *Problem) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *Problem) GetAcceptTime() int64 {
	if m != nil {
		return m.AcceptTime
	}
	return 0
}

func (m *Problem) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Problem) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Problem) GetJudgeFile() string {
	if m != nil {
		return m.JudgeFile
	}
	return ""
}

// SubmitRecord : 提交情况
type SubmitRecord struct {
	ProblemId            int64    `protobuf:"varint,1,opt,name=problem_id,json=problemId,proto3" json:"problem_id,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SubmitTime           int64    `protobuf:"varint,3,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	IsPass               bool     `protobuf:"varint,4,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
	RunningTime          int64    `protobuf:"varint,5,opt,name=running_time,json=runningTime,proto3" json:"running_time,omitempty"`
	RunningMem           int64    `protobuf:"varint,6,opt,name=running_mem,json=runningMem,proto3" json:"running_mem,omitempty"`
	Code                 string   `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	RunningLan           int64    `protobuf:"varint,8,opt,name=running_lan,json=runningLan,proto3" json:"running_lan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitRecord) Reset()         { *m = SubmitRecord{} }
func (m *SubmitRecord) String() string { return proto.CompactTextString(m) }
func (*SubmitRecord) ProtoMessage()    {}
func (*SubmitRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{3}
}

func (m *SubmitRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitRecord.Unmarshal(m, b)
}
func (m *SubmitRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitRecord.Marshal(b, m, deterministic)
}
func (m *SubmitRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitRecord.Merge(m, src)
}
func (m *SubmitRecord) XXX_Size() int {
	return xxx_messageInfo_SubmitRecord.Size(m)
}
func (m *SubmitRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitRecord.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitRecord proto.InternalMessageInfo

func (m *SubmitRecord) GetProblemId() int64 {
	if m != nil {
		return m.ProblemId
	}
	return 0
}

func (m *SubmitRecord) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SubmitRecord) GetSubmitTime() int64 {
	if m != nil {
		return m.SubmitTime
	}
	return 0
}

func (m *SubmitRecord) GetIsPass() bool {
	if m != nil {
		return m.IsPass
	}
	return false
}

func (m *SubmitRecord) GetRunningTime() int64 {
	if m != nil {
		return m.RunningTime
	}
	return 0
}

func (m *SubmitRecord) GetRunningMem() int64 {
	if m != nil {
		return m.RunningMem
	}
	return 0
}

func (m *SubmitRecord) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *SubmitRecord) GetRunningLan() int64 {
	if m != nil {
		return m.RunningLan
	}
	return 0
}

// Announcement : 公告，包括班级公告和全局公告
type Announcement struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Publisher            string   `protobuf:"bytes,2,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Detail               string   `protobuf:"bytes,4,opt,name=detail,proto3" json:"detail,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	LastUpdateTime       int64    `protobuf:"varint,6,opt,name=last_update_time,json=lastUpdateTime,proto3" json:"last_update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{4}
}

func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announcement.Unmarshal(m, b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return xxx_messageInfo_Announcement.Size(m)
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Announcement) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *Announcement) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Announcement) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func (m *Announcement) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Announcement) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

// Class : 班级信息
type Class struct {
	Id                   int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Tutor                string          `protobuf:"bytes,2,opt,name=tutor,proto3" json:"tutor,omitempty"`
	Name                 string          `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string          `protobuf:"bytes,4,opt,name=introduction,proto3" json:"introduction,omitempty"`
	Number               int64           `protobuf:"varint,5,opt,name=number,proto3" json:"number,omitempty"`
	IsCheck              bool            `protobuf:"varint,6,opt,name=is_check,json=isCheck,proto3" json:"is_check,omitempty"`
	CreateTime           int64           `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Announcements        []*Announcement `protobuf:"bytes,8,rep,name=announcements,proto3" json:"announcements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{5}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Class) GetTutor() string {
	if m != nil {
		return m.Tutor
	}
	return ""
}

func (m *Class) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Class) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *Class) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Class) GetIsCheck() bool {
	if m != nil {
		return m.IsCheck
	}
	return false
}

func (m *Class) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Class) GetAnnouncements() []*Announcement {
	if m != nil {
		return m.Announcements
	}
	return nil
}

// RankItem : 排名item
type RankItem struct {
	Ranking              int64    `protobuf:"varint,1,opt,name=ranking,proto3" json:"ranking,omitempty"`
	UserId               int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PassNum              int64    `protobuf:"varint,4,opt,name=pass_num,json=passNum,proto3" json:"pass_num,omitempty"`
	SubmitNum            int64    `protobuf:"varint,5,opt,name=submit_num,json=submitNum,proto3" json:"submit_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RankItem) Reset()         { *m = RankItem{} }
func (m *RankItem) String() string { return proto.CompactTextString(m) }
func (*RankItem) ProtoMessage()    {}
func (*RankItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{6}
}

func (m *RankItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RankItem.Unmarshal(m, b)
}
func (m *RankItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RankItem.Marshal(b, m, deterministic)
}
func (m *RankItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RankItem.Merge(m, src)
}
func (m *RankItem) XXX_Size() int {
	return xxx_messageInfo_RankItem.Size(m)
}
func (m *RankItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RankItem.DiscardUnknown(m)
}

var xxx_messageInfo_RankItem proto.InternalMessageInfo

func (m *RankItem) GetRanking() int64 {
	if m != nil {
		return m.Ranking
	}
	return 0
}

func (m *RankItem) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RankItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RankItem) GetPassNum() int64 {
	if m != nil {
		return m.PassNum
	}
	return 0
}

func (m *RankItem) GetSubmitNum() int64 {
	if m != nil {
		return m.SubmitNum
	}
	return 0
}

// Paper : 试卷
type Paper struct {
	Id       int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Problems []*Problem `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
	// 组卷的参数
	Difficulty           int64    `protobuf:"varint,3,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	ProblemNum           int64    `protobuf:"varint,4,opt,name=problem_num,json=problemNum,proto3" json:"problem_num,omitempty"`
	Tags                 []int64  `protobuf:"varint,5,rep,packed,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Paper) Reset()         { *m = Paper{} }
func (m *Paper) String() string { return proto.CompactTextString(m) }
func (*Paper) ProtoMessage()    {}
func (*Paper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{7}
}

func (m *Paper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Paper.Unmarshal(m, b)
}
func (m *Paper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Paper.Marshal(b, m, deterministic)
}
func (m *Paper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Paper.Merge(m, src)
}
func (m *Paper) XXX_Size() int {
	return xxx_messageInfo_Paper.Size(m)
}
func (m *Paper) XXX_DiscardUnknown() {
	xxx_messageInfo_Paper.DiscardUnknown(m)
}

var xxx_messageInfo_Paper proto.InternalMessageInfo

func (m *Paper) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Paper) GetProblems() []*Problem {
	if m != nil {
		return m.Problems
	}
	return nil
}

func (m *Paper) GetDifficulty() int64 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *Paper) GetProblemNum() int64 {
	if m != nil {
		return m.ProblemNum
	}
	return 0
}

func (m *Paper) GetTags() []int64 {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Match : 比赛
type Match struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsPublic             bool     `protobuf:"varint,2,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	IsOver               bool     `protobuf:"varint,5,opt,name=is_over,json=isOver,proto3" json:"is_over,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Introduction         string   `protobuf:"bytes,7,opt,name=introduction,proto3" json:"introduction,omitempty"`
	PaperId              int64    `protobuf:"varint,8,opt,name=paper_id,json=paperId,proto3" json:"paper_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{8}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Match) GetIsPublic() bool {
	if m != nil {
		return m.IsPublic
	}
	return false
}

func (m *Match) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Match) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Match) GetIsOver() bool {
	if m != nil {
		return m.IsOver
	}
	return false
}

func (m *Match) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Match) GetIntroduction() string {
	if m != nil {
		return m.Introduction
	}
	return ""
}

func (m *Match) GetPaperId() int64 {
	if m != nil {
		return m.PaperId
	}
	return 0
}

type ClassMember struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IsAdministrator      bool     `protobuf:"varint,3,opt,name=is_administrator,json=isAdministrator,proto3" json:"is_administrator,omitempty"`
	IsChecked            bool     `protobuf:"varint,4,opt,name=is_checked,json=isChecked,proto3" json:"is_checked,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassMember) Reset()         { *m = ClassMember{} }
func (m *ClassMember) String() string { return proto.CompactTextString(m) }
func (*ClassMember) ProtoMessage()    {}
func (*ClassMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{9}
}

func (m *ClassMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassMember.Unmarshal(m, b)
}
func (m *ClassMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassMember.Marshal(b, m, deterministic)
}
func (m *ClassMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassMember.Merge(m, src)
}
func (m *ClassMember) XXX_Size() int {
	return xxx_messageInfo_ClassMember.Size(m)
}
func (m *ClassMember) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassMember.DiscardUnknown(m)
}

var xxx_messageInfo_ClassMember proto.InternalMessageInfo

func (m *ClassMember) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ClassMember) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClassMember) GetIsAdministrator() bool {
	if m != nil {
		return m.IsAdministrator
	}
	return false
}

func (m *ClassMember) GetIsChecked() bool {
	if m != nil {
		return m.IsChecked
	}
	return false
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "protocol.UserInfo")
	proto.RegisterType((*ProblemExample)(nil), "protocol.ProblemExample")
	proto.RegisterType((*Problem)(nil), "protocol.Problem")
	proto.RegisterType((*SubmitRecord)(nil), "protocol.SubmitRecord")
	proto.RegisterType((*Announcement)(nil), "protocol.Announcement")
	proto.RegisterType((*Class)(nil), "protocol.Class")
	proto.RegisterType((*RankItem)(nil), "protocol.RankItem")
	proto.RegisterType((*Paper)(nil), "protocol.Paper")
	proto.RegisterType((*Match)(nil), "protocol.Match")
	proto.RegisterType((*ClassMember)(nil), "protocol.ClassMember")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 1008 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xcd, 0x6e, 0xe4, 0x44,
	0x10, 0x96, 0xc7, 0x99, 0xb1, 0x5d, 0x93, 0xdf, 0x66, 0x95, 0x75, 0x60, 0x21, 0x83, 0x0f, 0x68,
	0x38, 0x10, 0x24, 0x38, 0xc2, 0x81, 0x65, 0x05, 0x52, 0xa4, 0x64, 0x37, 0x32, 0xbb, 0x17, 0x2e,
	0x56, 0xc7, 0xee, 0xcc, 0x34, 0xb1, 0xbb, 0x2d, 0x77, 0xf7, 0x12, 0xee, 0x70, 0x42, 0x3c, 0x01,
	0xcf, 0xc1, 0xcb, 0xf0, 0x22, 0x88, 0x1b, 0xea, 0xea, 0xf6, 0x8c, 0x67, 0x26, 0xec, 0x29, 0xfd,
	0x7d, 0xae, 0xee, 0x54, 0xd5, 0xf7, 0x55, 0x0d, 0x90, 0xb6, 0x93, 0x5a, 0x7e, 0x5e, 0xca, 0xa6,
	0x91, 0xe2, 0x02, 0x01, 0x89, 0xf1, 0x4f, 0x29, 0xeb, 0xec, 0xb7, 0x11, 0xc4, 0x6f, 0x14, 0xeb,
	0x2e, 0xc5, 0x9d, 0x24, 0x87, 0x30, 0xe2, 0x55, 0x1a, 0xcc, 0x82, 0x79, 0x98, 0x8f, 0x78, 0x45,
	0x08, 0xec, 0x75, 0xb2, 0x66, 0xe9, 0x1e, 0x32, 0x78, 0xb6, 0x9c, 0xa0, 0x0d, 0x4b, 0xc7, 0xb3,
	0x60, 0x9e, 0xe4, 0x78, 0x26, 0xc7, 0x10, 0x2a, 0xf6, 0x90, 0x4e, 0x66, 0xc1, 0x3c, 0xce, 0xed,
	0x91, 0x3c, 0x81, 0x71, 0xbb, 0x94, 0x82, 0xa5, 0x11, 0x86, 0x39, 0x60, 0x59, 0xd6, 0x50, 0x5e,
	0xa7, 0xb1, 0x63, 0x11, 0x90, 0x53, 0x98, 0xa8, 0x72, 0x29, 0x65, 0x9d, 0x26, 0x48, 0x7b, 0x44,
	0x3e, 0x04, 0xa8, 0xa9, 0xd2, 0x45, 0x2d, 0x17, 0x5c, 0xa4, 0x80, 0x39, 0x24, 0x96, 0xb9, 0xb2,
	0x84, 0xbd, 0x56, 0x76, 0x8c, 0x6a, 0x96, 0x4e, 0xf1, 0x93, 0x47, 0x24, 0x85, 0x88, 0x96, 0xa5,
	0x34, 0x42, 0xa7, 0x23, 0x7c, 0xaf, 0x87, 0xe4, 0x7d, 0x88, 0x5b, 0xaa, 0xd4, 0xcf, 0xb2, 0xab,
	0xd2, 0x10, 0x3f, 0xad, 0x70, 0xf6, 0x1a, 0x0e, 0x6f, 0x3a, 0x79, 0x5b, 0xb3, 0xe6, 0xbb, 0x07,
	0xda, 0xb4, 0x35, 0x26, 0xcb, 0x45, 0xc5, 0x1e, 0x7c, 0x3f, 0x1c, 0x70, 0x6c, 0x6b, 0xfa, 0xb7,
	0x1d, 0xb0, 0xb9, 0x48, 0xa3, 0x2d, 0xed, 0xde, 0xf5, 0x28, 0xfb, 0x63, 0x0f, 0x22, 0xff, 0xec,
	0x4e, 0x73, 0x9f, 0xc0, 0x58, 0x73, 0x5d, 0xb3, 0xfe, 0x25, 0x04, 0x64, 0x06, 0xd3, 0x8a, 0xa9,
	0xb2, 0xe3, 0xad, 0xe6, 0x52, 0xf8, 0xe7, 0x86, 0x14, 0xbe, 0x23, 0x50, 0x92, 0x24, 0x1f, 0x71,
	0x61, 0x9b, 0x2f, 0x8d, 0xf6, 0x7a, 0xd8, 0xa3, 0x95, 0x68, 0xc9, 0x85, 0x46, 0x3d, 0x92, 0x1c,
	0xcf, 0xe4, 0x1b, 0x38, 0xe2, 0xa2, 0x90, 0x46, 0x17, 0xcc, 0xd5, 0xa7, 0xd2, 0x68, 0x16, 0xce,
	0xa7, 0x5f, 0xa4, 0x17, 0xbd, 0x17, 0x2e, 0x36, 0x1b, 0x90, 0x1f, 0x70, 0xf1, 0xca, 0x68, 0x8f,
	0x14, 0x99, 0xc3, 0xf1, 0x4f, 0xa6, 0x5a, 0xb0, 0xa2, 0xe6, 0x0d, 0xd7, 0x85, 0xe6, 0x0d, 0x43,
	0x1d, 0xc3, 0xfc, 0x10, 0xf9, 0x2b, 0x4b, 0xbf, 0xe6, 0x0d, 0x23, 0x9f, 0xc0, 0xd1, 0x30, 0xb2,
	0x61, 0x0d, 0x2a, 0x1b, 0xe6, 0x07, 0xeb, 0xc0, 0x6b, 0xd6, 0xd8, 0x3c, 0x35, 0x5d, 0xa8, 0x14,
	0x66, 0xa1, 0xb5, 0x97, 0x3d, 0x93, 0x8f, 0x00, 0x2a, 0x7e, 0x77, 0xc7, 0x4b, 0x53, 0xeb, 0x5f,
	0xbc, 0xb2, 0x03, 0x86, 0x3c, 0x83, 0xa4, 0x94, 0x0b, 0xc1, 0xb1, 0x3b, 0xfb, 0xce, 0x13, 0x2b,
	0x82, 0x9c, 0xc3, 0x54, 0x99, 0xdb, 0x55, 0x7a, 0x07, 0xee, 0xba, 0xa3, 0x30, 0xb5, 0x73, 0x98,
	0xd2, 0xb2, 0x64, 0xad, 0x0f, 0x38, 0x74, 0x01, 0x8e, 0xea, 0x03, 0x9c, 0x8f, 0x5c, 0xc0, 0x91,
	0x0b, 0x70, 0x14, 0x06, 0x3c, 0x83, 0xa4, 0x35, 0xb7, 0x35, 0x57, 0x4b, 0xd6, 0xa5, 0xc7, 0xd8,
	0xe1, 0x35, 0x61, 0x3d, 0xeb, 0x4a, 0xbf, 0xe3, 0x35, 0x4b, 0x4f, 0xdc, 0x67, 0x64, 0xbe, 0xe7,
	0x35, 0xcb, 0xfe, 0x0d, 0x60, 0xff, 0x07, 0xcc, 0x26, 0x67, 0xa5, 0xec, 0x2a, 0x1b, 0xdf, 0xba,
	0xae, 0x17, 0x2b, 0x73, 0x24, 0x9e, 0xb9, 0xac, 0xc8, 0x53, 0x88, 0x8c, 0x62, 0x9d, 0xfd, 0x36,
	0x72, 0x26, 0xb7, 0xf0, 0xb2, 0xda, 0x2e, 0x34, 0xdc, 0x29, 0xf4, 0x29, 0x44, 0x5c, 0x15, 0xd6,
	0xde, 0x68, 0x95, 0x38, 0x9f, 0x70, 0x75, 0x43, 0x95, 0x22, 0x1f, 0xc3, 0x7e, 0x67, 0x84, 0xe0,
	0x62, 0xe1, 0xae, 0x8e, 0xf1, 0xea, 0xd4, 0x73, 0x7d, 0x0f, 0xfa, 0x10, 0xab, 0xdd, 0xc4, 0x3d,
	0xee, 0x29, 0x2f, 0x5c, 0x29, 0xab, 0x7e, 0xb8, 0xf1, 0x3c, 0xbc, 0x54, 0x53, 0xe1, 0x9d, 0xd1,
	0x5f, 0xba, 0xa2, 0x22, 0xfb, 0x2b, 0x80, 0xfd, 0xe7, 0x42, 0x48, 0x23, 0x4a, 0xd6, 0x30, 0xa1,
	0x77, 0x06, 0x62, 0xa3, 0xb3, 0xa3, 0xed, 0xce, 0xae, 0xc6, 0x25, 0x1c, 0x8e, 0xcb, 0x29, 0x4c,
	0x2a, 0xa6, 0xed, 0x4a, 0x71, 0x03, 0xe1, 0xd1, 0xb6, 0x8c, 0xe3, 0x1d, 0x19, 0xe7, 0x70, 0x8c,
	0xcb, 0xc5, 0xb4, 0xd5, 0x2a, 0xca, 0x15, 0x7a, 0x68, 0xf9, 0x37, 0x48, 0xdb, 0xc8, 0xec, 0x9f,
	0x00, 0xc6, 0x2f, 0x6a, 0xdb, 0xba, 0xc7, 0x26, 0xd8, 0x68, 0xd9, 0xad, 0x26, 0xd8, 0x82, 0xd5,
	0x82, 0x0c, 0x07, 0x0b, 0x32, 0x83, 0x7d, 0x2e, 0x74, 0x27, 0x2b, 0x53, 0xa2, 0x71, 0x5d, 0xb2,
	0x1b, 0x9c, 0x2d, 0x45, 0x98, 0xe6, 0x96, 0x75, 0x3e, 0x5b, 0x8f, 0xc8, 0x19, 0xc4, 0x5c, 0x15,
	0xe5, 0x92, 0x95, 0xf7, 0x7e, 0xc3, 0x46, 0x5c, 0xbd, 0xb0, 0x70, 0xbb, 0xca, 0x68, 0xa7, 0xca,
	0xaf, 0xe1, 0x80, 0x0e, 0x5a, 0xae, 0xd2, 0x18, 0x67, 0xfe, 0x74, 0x3d, 0xf3, 0x43, 0x45, 0xf2,
	0xcd, 0xe0, 0xec, 0xf7, 0x00, 0xe2, 0x9c, 0x8a, 0xfb, 0x4b, 0xcd, 0x1a, 0xbb, 0x56, 0x3b, 0x2a,
	0xee, 0xb9, 0x58, 0xf8, 0x0e, 0xf4, 0xf0, 0xff, 0x4d, 0xfa, 0x58, 0x27, 0xce, 0xdc, 0x0e, 0x2e,
	0x84, 0x69, 0xfc, 0xcf, 0x4a, 0x64, 0xf1, 0x4b, 0xd3, 0xd8, 0x59, 0xf0, 0x9e, 0xb6, 0x1f, 0x5d,
	0x13, 0x12, 0xc7, 0xbc, 0x34, 0x4d, 0xf6, 0x67, 0x00, 0xe3, 0x1b, 0xda, 0xb2, 0x6e, 0x47, 0x87,
	0xcf, 0x20, 0xf6, 0x23, 0xa3, 0xd2, 0x11, 0x16, 0x78, 0xb2, 0xb3, 0xd4, 0xf2, 0x55, 0xc8, 0xd6,
	0x8a, 0x09, 0x77, 0x56, 0xcc, 0x39, 0x4c, 0xfb, 0x99, 0x5c, 0x67, 0xd9, 0x8f, 0xa9, 0x4d, 0xb4,
	0xdf, 0x5b, 0xe3, 0xf5, 0xde, 0xca, 0xfe, 0x0e, 0x60, 0x7c, 0x4d, 0x75, 0xb9, 0xdc, 0xc9, 0xee,
	0x03, 0x48, 0xec, 0x24, 0x5a, 0x23, 0x97, 0xd8, 0xa0, 0x38, 0x8f, 0xb9, 0xba, 0x41, 0x8c, 0x35,
	0x6b, 0xda, 0x6d, 0x8c, 0x71, 0x82, 0x0c, 0xea, 0x77, 0x06, 0x31, 0x13, 0x95, 0xfb, 0xe8, 0xbb,
	0xc5, 0x44, 0x35, 0x18, 0x70, 0xf9, 0xd6, 0xfb, 0x05, 0x07, 0xfc, 0xd5, 0x5b, 0xb6, 0xf6, 0xdf,
	0xe4, 0x1d, 0xfe, 0x8b, 0x1e, 0xf1, 0x1f, 0x2a, 0xd3, 0x3a, 0x1d, 0xe3, 0x5e, 0x99, 0xd6, 0x0a,
	0x99, 0xfd, 0x1a, 0xc0, 0x14, 0x47, 0xe0, 0x9a, 0xa1, 0x25, 0x07, 0x8a, 0x07, 0x8f, 0x2a, 0x3e,
	0x1a, 0xfc, 0xef, 0x4f, 0xe1, 0x98, 0xab, 0x82, 0x56, 0x0d, 0x17, 0x5c, 0xe9, 0x8e, 0xda, 0x81,
	0x09, 0x31, 0xe3, 0x23, 0xae, 0x9e, 0x0f, 0x69, 0xdb, 0x8d, 0xde, 0xea, 0xac, 0xf2, 0x7b, 0x2b,
	0xf1, 0x66, 0x67, 0xd5, 0xb7, 0xef, 0xfd, 0x78, 0xd2, 0xcb, 0xfa, 0x55, 0x7f, 0xb8, 0x9d, 0xe0,
	0xe9, 0xcb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x22, 0xe9, 0x8c, 0xe7, 0x08, 0x00, 0x00,
}
