// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/judge.proto

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

// 判题目结果
type JudgeResult_Result int32

const (
	JudgeResult_SUCCESS                  JudgeResult_Result = 0
	JudgeResult_WRONG_ANSWER             JudgeResult_Result = -1
	JudgeResult_CPU_TIME_LIMIT_EXCEEDED  JudgeResult_Result = 1
	JudgeResult_REAL_TIME_LIMIT_EXCEEDED JudgeResult_Result = 2
	JudgeResult_MEMORY_LIMIT_EXCEEDED    JudgeResult_Result = 3
	JudgeResult_RUNTIME_ERROR            JudgeResult_Result = 4
	JudgeResult_SYSTEM_ERROR             JudgeResult_Result = 5
)

var JudgeResult_Result_name = map[int32]string{
	0:  "SUCCESS",
	-1: "WRONG_ANSWER",
	1:  "CPU_TIME_LIMIT_EXCEEDED",
	2:  "REAL_TIME_LIMIT_EXCEEDED",
	3:  "MEMORY_LIMIT_EXCEEDED",
	4:  "RUNTIME_ERROR",
	5:  "SYSTEM_ERROR",
}

var JudgeResult_Result_value = map[string]int32{
	"SUCCESS":                  0,
	"WRONG_ANSWER":             -1,
	"CPU_TIME_LIMIT_EXCEEDED":  1,
	"REAL_TIME_LIMIT_EXCEEDED": 2,
	"MEMORY_LIMIT_EXCEEDED":    3,
	"RUNTIME_ERROR":            4,
	"SYSTEM_ERROR":             5,
}

func (x JudgeResult_Result) String() string {
	return proto.EnumName(JudgeResult_Result_name, int32(x))
}

func (JudgeResult_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fc3ba8201c1ea690, []int{0, 0}
}

type JudgeResult struct {
	JudgeResult          JudgeResult_Result `protobuf:"varint,2,opt,name=judge_result,json=judgeResult,proto3,enum=protocol.JudgeResult_Result" json:"judge_result,omitempty"`
	CpuTime              int64              `protobuf:"varint,3,opt,name=cpu_time,json=cpuTime,proto3" json:"cpu_time,omitempty"`
	RealTime             int64              `protobuf:"varint,4,opt,name=real_time,json=realTime,proto3" json:"real_time,omitempty"`
	Memory               int64              `protobuf:"varint,5,opt,name=memory,proto3" json:"memory,omitempty"`
	Signal               int64              `protobuf:"varint,6,opt,name=signal,proto3" json:"signal,omitempty"`
	ExitCode             int64              `protobuf:"varint,7,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *JudgeResult) Reset()         { *m = JudgeResult{} }
func (m *JudgeResult) String() string { return proto.CompactTextString(m) }
func (*JudgeResult) ProtoMessage()    {}
func (*JudgeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc3ba8201c1ea690, []int{0}
}

func (m *JudgeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeResult.Unmarshal(m, b)
}
func (m *JudgeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeResult.Marshal(b, m, deterministic)
}
func (m *JudgeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeResult.Merge(m, src)
}
func (m *JudgeResult) XXX_Size() int {
	return xxx_messageInfo_JudgeResult.Size(m)
}
func (m *JudgeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeResult.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeResult proto.InternalMessageInfo

func (m *JudgeResult) GetJudgeResult() JudgeResult_Result {
	if m != nil {
		return m.JudgeResult
	}
	return JudgeResult_SUCCESS
}

func (m *JudgeResult) GetCpuTime() int64 {
	if m != nil {
		return m.CpuTime
	}
	return 0
}

func (m *JudgeResult) GetRealTime() int64 {
	if m != nil {
		return m.RealTime
	}
	return 0
}

func (m *JudgeResult) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *JudgeResult) GetSignal() int64 {
	if m != nil {
		return m.Signal
	}
	return 0
}

func (m *JudgeResult) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

type JudgeRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Src                  string   `protobuf:"bytes,2,opt,name=src,proto3" json:"src,omitempty"`
	Language             int64    `protobuf:"varint,3,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JudgeRequest) Reset()         { *m = JudgeRequest{} }
func (m *JudgeRequest) String() string { return proto.CompactTextString(m) }
func (*JudgeRequest) ProtoMessage()    {}
func (*JudgeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc3ba8201c1ea690, []int{1}
}

func (m *JudgeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeRequest.Unmarshal(m, b)
}
func (m *JudgeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeRequest.Marshal(b, m, deterministic)
}
func (m *JudgeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeRequest.Merge(m, src)
}
func (m *JudgeRequest) XXX_Size() int {
	return xxx_messageInfo_JudgeRequest.Size(m)
}
func (m *JudgeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeRequest proto.InternalMessageInfo

func (m *JudgeRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *JudgeRequest) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *JudgeRequest) GetLanguage() int64 {
	if m != nil {
		return m.Language
	}
	return 0
}

type JudgeResponse struct {
	Status               *Status        `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Result               int64          `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	Results              []*JudgeResult `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *JudgeResponse) Reset()         { *m = JudgeResponse{} }
func (m *JudgeResponse) String() string { return proto.CompactTextString(m) }
func (*JudgeResponse) ProtoMessage()    {}
func (*JudgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc3ba8201c1ea690, []int{2}
}

func (m *JudgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JudgeResponse.Unmarshal(m, b)
}
func (m *JudgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JudgeResponse.Marshal(b, m, deterministic)
}
func (m *JudgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JudgeResponse.Merge(m, src)
}
func (m *JudgeResponse) XXX_Size() int {
	return xxx_messageInfo_JudgeResponse.Size(m)
}
func (m *JudgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JudgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JudgeResponse proto.InternalMessageInfo

func (m *JudgeResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JudgeResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *JudgeResponse) GetResults() []*JudgeResult {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.JudgeResult_Result", JudgeResult_Result_name, JudgeResult_Result_value)
	proto.RegisterType((*JudgeResult)(nil), "protocol.JudgeResult")
	proto.RegisterType((*JudgeRequest)(nil), "protocol.JudgeRequest")
	proto.RegisterType((*JudgeResponse)(nil), "protocol.JudgeResponse")
}

func init() { proto.RegisterFile("proto/judge.proto", fileDescriptor_fc3ba8201c1ea690) }

var fileDescriptor_fc3ba8201c1ea690 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x6f, 0x9b, 0x30,
	0x14, 0x1e, 0x71, 0x4b, 0xd2, 0x47, 0x5a, 0x11, 0x4f, 0xdd, 0x9c, 0xb6, 0x87, 0x88, 0x13, 0x27,
	0x2a, 0x65, 0xc7, 0x1d, 0xa6, 0x8e, 0x5a, 0x53, 0xa6, 0x90, 0x4c, 0x86, 0xa8, 0xeb, 0x2e, 0x88,
	0x81, 0x85, 0xa8, 0x20, 0x30, 0x0c, 0xd2, 0x76, 0xdc, 0xfe, 0x9d, 0xfd, 0x91, 0x9b, 0xb0, 0x61,
	0x8b, 0xaa, 0x70, 0xf1, 0xf7, 0xe3, 0xf9, 0x13, 0xbc, 0x0f, 0x98, 0x55, 0x75, 0xd9, 0x94, 0xb7,
	0x4f, 0x6d, 0x92, 0x72, 0x47, 0x62, 0x3c, 0x91, 0x47, 0x5c, 0xe6, 0x57, 0x58, 0x99, 0xa2, 0x89,
	0x9a, 0x56, 0x28, 0xd7, 0xfa, 0x89, 0xc0, 0xf8, 0xd8, 0x4d, 0x33, 0x2e, 0xda, 0xbc, 0xc1, 0xef,
	0x60, 0x2a, 0x2f, 0x87, 0xb5, 0xe4, 0x64, 0xb4, 0xd0, 0xec, 0x8b, 0xe5, 0x8d, 0x33, 0x84, 0x38,
	0x07, 0xc3, 0x8e, 0x3a, 0x98, 0xf1, 0x74, 0x10, 0x30, 0x87, 0x49, 0x5c, 0xb5, 0x61, 0x93, 0x15,
	0x9c, 0xa0, 0x85, 0x66, 0x23, 0x36, 0x8e, 0xab, 0x36, 0xc8, 0x0a, 0x8e, 0xaf, 0xe1, 0xac, 0xe6,
	0x51, 0xae, 0xbc, 0x13, 0xe9, 0x4d, 0x3a, 0x41, 0x9a, 0xaf, 0x40, 0x2f, 0x78, 0x51, 0xd6, 0x3f,
	0xc8, 0xa9, 0x74, 0x7a, 0xd6, 0xe9, 0x22, 0x4b, 0xf7, 0x51, 0x4e, 0x74, 0xa5, 0x2b, 0xd6, 0x85,
	0xf1, 0xef, 0x59, 0x13, 0xc6, 0x65, 0xc2, 0xc9, 0x58, 0x85, 0x75, 0x82, 0x5b, 0x26, 0xdc, 0xfa,
	0xad, 0x81, 0xde, 0xbf, 0x8f, 0x01, 0x63, 0x7f, 0xe7, 0xba, 0xd4, 0xf7, 0xcd, 0x17, 0x78, 0x0e,
	0xd3, 0x07, 0xb6, 0xdd, 0x7c, 0x08, 0xef, 0x36, 0xfe, 0x03, 0x65, 0xe6, 0x9f, 0xe1, 0xd1, 0xf0,
	0x35, 0xbc, 0x76, 0x3f, 0xed, 0xc2, 0x60, 0xe5, 0xd1, 0x70, 0xbd, 0xf2, 0x56, 0x41, 0x48, 0x3f,
	0xbb, 0x94, 0xde, 0xd3, 0x7b, 0x53, 0xc3, 0x37, 0x40, 0x18, 0xbd, 0x5b, 0x1f, 0x75, 0x47, 0x78,
	0x0e, 0x97, 0x1e, 0xf5, 0xb6, 0xec, 0xf1, 0xb9, 0x85, 0xf0, 0x0c, 0xce, 0xd9, 0x6e, 0x23, 0xaf,
	0x51, 0xc6, 0xb6, 0xcc, 0x3c, 0xc1, 0x26, 0x4c, 0xfd, 0x47, 0x3f, 0xa0, 0x5e, 0xaf, 0x9c, 0x5a,
	0x6b, 0x98, 0xf6, 0x5b, 0xfd, 0xd6, 0x72, 0xd1, 0xe0, 0x0b, 0x18, 0x65, 0x09, 0xd1, 0xe4, 0x37,
	0x8d, 0xb2, 0x04, 0x9b, 0x80, 0x44, 0x1d, 0xcb, 0x2a, 0xce, 0x58, 0x07, 0xf1, 0x15, 0x4c, 0xf2,
	0x68, 0x9f, 0xb6, 0x51, 0x3a, 0x2c, 0xf9, 0x1f, 0xb7, 0x7e, 0x69, 0x70, 0x3e, 0x94, 0x54, 0x95,
	0x7b, 0xc1, 0xb1, 0x0d, 0xba, 0xea, 0x5c, 0x66, 0x1a, 0x4b, 0xf3, 0x7f, 0x9b, 0xbe, 0xd4, 0x59,
	0xef, 0x77, 0xcb, 0x3e, 0xe8, 0x1d, 0xb1, 0x9e, 0xe1, 0x5b, 0x18, 0x2b, 0x24, 0x08, 0x5a, 0x20,
	0xdb, 0x58, 0x5e, 0x1e, 0xfd, 0x21, 0xd8, 0x30, 0xf5, 0xfe, 0xe5, 0x97, 0xd9, 0x30, 0xf0, 0x76,
	0x00, 0x5f, 0x75, 0x89, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa3, 0x9e, 0x23, 0xa5,
	0x02, 0x00, 0x00,
}
