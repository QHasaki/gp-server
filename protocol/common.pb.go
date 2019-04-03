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

// Role : 用户角色（学生/老师...）
type Role int32

const (
	Role_STUDENT Role = 0
	Role_TEACHER Role = 1
	Role_MANAGER Role = 2
)

var Role_name = map[int32]string{
	0: "STUDENT",
	1: "TEACHER",
	2: "MANAGER",
}

var Role_value = map[string]int32{
	"STUDENT": 0,
	"TEACHER": 1,
	"MANAGER": 2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

// ProblemDifficluty : 题目难度
type ProblemDifficluty int32

const (
	ProblemDifficluty_EASY   ProblemDifficluty = 0
	ProblemDifficluty_MEDIUM ProblemDifficluty = 1
	ProblemDifficluty_HARD   ProblemDifficluty = 2
)

var ProblemDifficluty_name = map[int32]string{
	0: "EASY",
	1: "MEDIUM",
	2: "HARD",
}

var ProblemDifficluty_value = map[string]int32{
	"EASY":   0,
	"MEDIUM": 1,
	"HARD":   2,
}

func (x ProblemDifficluty) String() string {
	return proto.EnumName(ProblemDifficluty_name, int32(x))
}

func (ProblemDifficluty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

// UserInfo : 用户基本信息
type UserInfo struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 Role     `protobuf:"varint,4,opt,name=role,proto3,enum=protocol.Role" json:"role,omitempty"`
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

func (m *UserInfo) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_STUDENT
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
	Input                string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output               string   `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
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
	Difficluty           ProblemDifficluty `protobuf:"varint,11,opt,name=difficluty,proto3,enum=protocol.ProblemDifficluty" json:"difficluty,omitempty"`
	SubmitTime           int64             `protobuf:"varint,12,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	AcceptTime           int64             `protobuf:"varint,13,opt,name=accept_time,json=acceptTime,proto3" json:"accept_time,omitempty"`
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

func (m *Problem) GetDifficluty() ProblemDifficluty {
	if m != nil {
		return m.Difficluty
	}
	return ProblemDifficluty_EASY
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

// SubmitRecord : 提交情况
type SubmitRecord struct {
	Problem              *Problem `protobuf:"bytes,1,opt,name=problem,proto3" json:"problem,omitempty"`
	SubmitTime           int64    `protobuf:"varint,2,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	IsPass               bool     `protobuf:"varint,3,opt,name=is_pass,json=isPass,proto3" json:"is_pass,omitempty"`
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

func (m *SubmitRecord) GetProblem() *Problem {
	if m != nil {
		return m.Problem
	}
	return nil
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
	KnowledgePoints      []int64  `protobuf:"varint,5,rep,packed,name=knowledge_points,json=knowledgePoints,proto3" json:"knowledge_points,omitempty"`
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

func (m *Paper) GetKnowledgePoints() []int64 {
	if m != nil {
		return m.KnowledgePoints
	}
	return nil
}

// Match : 比赛
type Match struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsPublic             bool     `protobuf:"varint,2,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	StartTime            int64    `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration             int64    `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	IsOver               bool     `protobuf:"varint,5,opt,name=is_over,json=isOver,proto3" json:"is_over,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Intriduction         string   `protobuf:"bytes,7,opt,name=intriduction,proto3" json:"intriduction,omitempty"`
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

func (m *Match) GetDuration() int64 {
	if m != nil {
		return m.Duration
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

func (m *Match) GetIntriduction() string {
	if m != nil {
		return m.Intriduction
	}
	return ""
}

func (m *Match) GetPaperId() int64 {
	if m != nil {
		return m.PaperId
	}
	return 0
}

func init() {
	proto.RegisterEnum("protocol.Role", Role_name, Role_value)
	proto.RegisterEnum("protocol.ProblemDifficluty", ProblemDifficluty_name, ProblemDifficluty_value)
	proto.RegisterType((*UserInfo)(nil), "protocol.UserInfo")
	proto.RegisterType((*ProblemExample)(nil), "protocol.ProblemExample")
	proto.RegisterType((*Problem)(nil), "protocol.Problem")
	proto.RegisterType((*SubmitRecord)(nil), "protocol.SubmitRecord")
	proto.RegisterType((*Announcement)(nil), "protocol.Announcement")
	proto.RegisterType((*Class)(nil), "protocol.Class")
	proto.RegisterType((*RankItem)(nil), "protocol.RankItem")
	proto.RegisterType((*Paper)(nil), "protocol.Paper")
	proto.RegisterType((*Match)(nil), "protocol.Match")
}

func init() { proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0) }

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 985 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0x5d, 0x6f, 0xdb, 0x36,
	0x14, 0xad, 0xac, 0xd8, 0x92, 0xaf, 0x13, 0xd7, 0x25, 0x8a, 0x4c, 0x6d, 0xb7, 0xd5, 0xf0, 0xc3,
	0xe0, 0x75, 0x68, 0x06, 0xa4, 0x8f, 0x1d, 0x86, 0x79, 0x89, 0xb1, 0x06, 0x68, 0xd2, 0x80, 0x49,
	0x1e, 0xb6, 0x17, 0x43, 0x91, 0x18, 0x9b, 0x8b, 0x44, 0x0a, 0x22, 0xd9, 0x8f, 0xdf, 0xb0, 0xe7,
	0xfd, 0x8d, 0x3d, 0x0c, 0xd8, 0xaf, 0xd9, 0x0f, 0xd9, 0xeb, 0xc0, 0x4b, 0x5a, 0x76, 0xaa, 0x3c,
	0x99, 0xe7, 0xe8, 0x4a, 0x3c, 0xbc, 0xe7, 0x5c, 0x1a, 0x48, 0x55, 0x4b, 0x2d, 0xbf, 0xcf, 0x64,
	0x59, 0x4a, 0x71, 0x80, 0x80, 0xc4, 0xf8, 0x93, 0xc9, 0x62, 0xf2, 0x67, 0x07, 0xe2, 0x2b, 0xc5,
	0xea, 0x13, 0x71, 0x23, 0xc9, 0x10, 0x3a, 0x3c, 0x4f, 0x82, 0x71, 0x30, 0x0d, 0x69, 0x87, 0xe7,
	0x64, 0x02, 0x3b, 0xb5, 0x2c, 0x58, 0xb2, 0x33, 0x0e, 0xa6, 0xc3, 0xc3, 0xe1, 0xc1, 0xfa, 0xad,
	0x03, 0x2a, 0x0b, 0x46, 0xf1, 0x19, 0x21, 0xb0, 0x23, 0xd2, 0x92, 0x25, 0xdd, 0x71, 0x30, 0xed,
	0x53, 0x5c, 0x93, 0x11, 0x84, 0x8a, 0x7d, 0x4c, 0x7a, 0xe3, 0x60, 0x1a, 0x53, 0xbb, 0x24, 0x8f,
	0xa1, 0x5b, 0xad, 0xa4, 0x60, 0x49, 0x84, 0x65, 0x0e, 0x58, 0x96, 0x95, 0x29, 0x2f, 0x92, 0xd8,
	0xb1, 0x08, 0xc8, 0x3e, 0xf4, 0x54, 0xb6, 0x92, 0xb2, 0x48, 0xfa, 0x48, 0x7b, 0x44, 0xbe, 0x02,
	0x28, 0x52, 0xa5, 0x17, 0x85, 0x5c, 0x72, 0x91, 0x00, 0xaa, 0xec, 0x5b, 0xe6, 0xad, 0x25, 0xec,
	0x6b, 0x59, 0xcd, 0x52, 0xcd, 0x92, 0x01, 0x3e, 0xf2, 0x88, 0x24, 0x10, 0xa5, 0x59, 0x26, 0x8d,
	0xd0, 0x49, 0x07, 0xbf, 0xb7, 0x86, 0xe4, 0x29, 0xc4, 0x55, 0xaa, 0xd4, 0x07, 0x59, 0xe7, 0x49,
	0x88, 0x8f, 0x1a, 0x3c, 0xf9, 0x11, 0x86, 0xe7, 0xb5, 0xbc, 0x2e, 0x58, 0x39, 0xff, 0x98, 0x96,
	0x55, 0x81, 0x62, 0xb9, 0xa8, 0x8c, 0xc6, 0xfe, 0xf4, 0xa9, 0x03, 0x76, 0x57, 0x69, 0xb4, 0xa5,
	0xdd, 0xc7, 0x3d, 0x9a, 0xfc, 0x15, 0x42, 0xe4, 0x3f, 0xd0, 0x6a, 0xeb, 0x63, 0xe8, 0x6a, 0xae,
	0x0b, 0xe6, 0x5f, 0x71, 0x80, 0x8c, 0x61, 0x90, 0x33, 0x95, 0xd5, 0xbc, 0xd2, 0x5c, 0x0a, 0x2f,
	0x68, 0x9b, 0xc2, 0xef, 0x08, 0x34, 0xa3, 0x4f, 0x3b, 0x5c, 0xd8, 0x36, 0x4b, 0xa3, 0x7d, 0xe7,
	0xed, 0xd2, 0x9a, 0xb1, 0xe2, 0x42, 0x63, 0xe7, 0xfb, 0x14, 0xd7, 0xe4, 0x27, 0x78, 0xc8, 0xc5,
	0x42, 0x1a, 0xbd, 0x60, 0xee, 0x24, 0x2a, 0x89, 0xc6, 0xe1, 0x74, 0x70, 0x98, 0x6c, 0xfc, 0xbc,
	0x7b, 0x54, 0xba, 0xc7, 0xc5, 0x3b, 0xa3, 0x3d, 0x52, 0x64, 0x0a, 0xa3, 0xdf, 0x4d, 0xbe, 0x64,
	0x8b, 0x82, 0x97, 0x5c, 0x2f, 0x34, 0x2f, 0x19, 0x3a, 0x16, 0xd2, 0x21, 0xf2, 0x6f, 0x2d, 0x7d,
	0xc9, 0x4b, 0x46, 0xbe, 0x81, 0x87, 0xdb, 0x95, 0x25, 0x2b, 0xd1, 0xc3, 0x90, 0xee, 0x6d, 0x0a,
	0x4f, 0x59, 0x69, 0x75, 0xea, 0x74, 0xa9, 0x12, 0x18, 0x87, 0xd3, 0x90, 0xe2, 0x9a, 0xbc, 0x06,
	0xc8, 0xf9, 0xcd, 0x0d, 0xcf, 0x0a, 0xa3, 0x3f, 0xa1, 0x87, 0xc3, 0xc3, 0x67, 0x2d, 0x89, 0xc7,
	0x4d, 0x09, 0xdd, 0x2a, 0x27, 0xcf, 0x61, 0xa0, 0xcc, 0x75, 0xa3, 0x6e, 0x17, 0x37, 0x05, 0x47,
	0xa1, 0xb2, 0xe7, 0x30, 0x48, 0xb3, 0x8c, 0x55, 0xbe, 0x60, 0xcf, 0x15, 0x38, 0xca, 0x16, 0x4c,
	0x0c, 0xec, 0x5e, 0x60, 0x39, 0x65, 0x99, 0xac, 0x73, 0xf2, 0x1d, 0x44, 0x95, 0xdb, 0x12, 0x9d,
	0x1b, 0x1c, 0x3e, 0x6a, 0x69, 0xa1, 0xeb, 0x8a, 0xcf, 0xb7, 0xef, 0xb4, 0xb6, 0xff, 0x02, 0x22,
	0xae, 0x16, 0x36, 0x5d, 0x68, 0x6c, 0x4c, 0x7b, 0x5c, 0x9d, 0xa7, 0x4a, 0x4d, 0xfe, 0x09, 0x60,
	0x77, 0x26, 0x84, 0x34, 0x22, 0x63, 0x25, 0x13, 0xba, 0x15, 0x96, 0x2f, 0xa1, 0x5f, 0x99, 0xeb,
	0x82, 0xab, 0x15, 0xab, 0x7d, 0x60, 0x36, 0xc4, 0x26, 0x4a, 0xe1, 0x76, 0x94, 0xf6, 0xa1, 0x97,
	0x33, 0x6d, 0x07, 0xcb, 0x85, 0xc5, 0x23, 0x2b, 0xd3, 0x0d, 0x85, 0x93, 0xd9, 0x75, 0x32, 0x1d,
	0x85, 0x32, 0xa7, 0x30, 0xc2, 0x11, 0x33, 0x55, 0xde, 0x54, 0xf5, 0x9c, 0xd3, 0x96, 0xbf, 0x42,
	0x1a, 0xdb, 0xf5, 0x5f, 0x00, 0xdd, 0xa3, 0x22, 0x55, 0xea, 0xde, 0x74, 0x1b, 0x2d, 0xeb, 0x26,
	0xdd, 0x16, 0x34, 0xd7, 0x44, 0xb8, 0x75, 0x4d, 0x4c, 0x60, 0x97, 0x0b, 0x5d, 0xcb, 0xdc, 0x64,
	0x18, 0x79, 0x27, 0xf6, 0x0e, 0x67, 0x8f, 0x22, 0x4c, 0x79, 0xcd, 0x6a, 0xaf, 0xd6, 0x23, 0xf2,
	0x04, 0x62, 0xae, 0x16, 0xd9, 0x8a, 0x65, 0xb7, 0xfe, 0x9e, 0x89, 0xb8, 0x3a, 0xb2, 0xf0, 0xf3,
	0x53, 0x46, 0xad, 0x53, 0xfe, 0x00, 0x7b, 0xe9, 0x56, 0xcb, 0x55, 0x12, 0xe3, 0x3c, 0xec, 0x6f,
	0x0c, 0xde, 0x76, 0x84, 0xde, 0x2d, 0x9e, 0xfc, 0x11, 0x40, 0x4c, 0x53, 0x71, 0x7b, 0xa2, 0x59,
	0x69, 0x2f, 0x97, 0x3a, 0x15, 0xb7, 0x5c, 0x2c, 0x7d, 0x07, 0xd6, 0xd0, 0x3a, 0x6e, 0x14, 0xab,
	0x17, 0x3c, 0xf7, 0x71, 0xe8, 0x59, 0x78, 0x92, 0xdf, 0xdb, 0x89, 0x27, 0xee, 0x26, 0x5a, 0x08,
	0x53, 0x62, 0x17, 0x42, 0x1a, 0x59, 0x7c, 0x66, 0x4a, 0x7b, 0xeb, 0xf9, 0x68, 0xd9, 0x87, 0xae,
	0x09, 0x7d, 0xc7, 0x9c, 0x99, 0x72, 0xf2, 0x77, 0x00, 0xdd, 0xf3, 0xb4, 0x62, 0x75, 0xcb, 0x87,
	0x97, 0x10, 0xfb, 0x78, 0xaa, 0xa4, 0x83, 0x07, 0xbc, 0x27, 0xc1, 0x4d, 0x09, 0xf9, 0x7a, 0x3d,
	0x7e, 0xa6, 0xd0, 0x9f, 0x50, 0x5c, 0x48, 0xb7, 0x18, 0xdb, 0x55, 0x5f, 0xbb, 0xa5, 0x12, 0x3c,
	0x65, 0x85, 0x7e, 0x0b, 0xa3, 0x5b, 0x21, 0x3f, 0x14, 0xcc, 0xce, 0x7f, 0x25, 0xb9, 0x6d, 0x6c,
	0x17, 0xe7, 0xfb, 0x61, 0xc3, 0x9f, 0x23, 0x3d, 0xf9, 0x37, 0x80, 0xee, 0x69, 0xaa, 0xb3, 0x55,
	0x4b, 0xf4, 0x33, 0xe8, 0xdb, 0x39, 0xb1, 0xf9, 0xce, 0xb0, 0x6f, 0x31, 0x8d, 0xb9, 0x3a, 0x47,
	0x8c, 0xad, 0xd0, 0x69, 0xed, 0x87, 0x2c, 0xf4, 0xad, 0xb0, 0x0c, 0xda, 0xfa, 0x14, 0xe2, 0xdc,
	0xd4, 0x69, 0x13, 0xa5, 0x90, 0x36, 0xd8, 0xcf, 0x9f, 0x7c, 0xef, 0x73, 0x84, 0xf3, 0xf7, 0xee,
	0x3d, 0xdb, 0xe4, 0xb2, 0xd7, 0xce, 0x25, 0x5f, 0xe7, 0x32, 0xda, 0xe4, 0x72, 0xcd, 0x39, 0xc7,
	0x2a, 0xe7, 0x6f, 0xbc, 0x76, 0xac, 0xb2, 0x06, 0xbf, 0x78, 0x09, 0x3b, 0xf6, 0xff, 0x91, 0x0c,
	0x20, 0xba, 0xb8, 0xbc, 0x3a, 0x9e, 0x9f, 0x5d, 0x8e, 0x1e, 0x58, 0x70, 0x39, 0x9f, 0x1d, 0xbd,
	0x99, 0xd3, 0x51, 0x60, 0xc1, 0xe9, 0xec, 0x6c, 0xf6, 0xcb, 0x9c, 0x8e, 0x3a, 0x2f, 0x5e, 0xc1,
	0xa3, 0xd6, 0xdd, 0x46, 0x62, 0xd8, 0x99, 0xcf, 0x2e, 0x7e, 0x1d, 0x3d, 0x20, 0x00, 0xbd, 0xd3,
	0xf9, 0xf1, 0xc9, 0xd5, 0xe9, 0x28, 0xb0, 0xec, 0x9b, 0x19, 0x3d, 0x1e, 0x75, 0x7e, 0xde, 0xfb,
	0x6d, 0xb0, 0x94, 0xaf, 0xd7, 0x76, 0x5e, 0xf7, 0x70, 0xf5, 0xea, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0x22, 0xf5, 0x27, 0xec, 0x07, 0x00, 0x00,
}
