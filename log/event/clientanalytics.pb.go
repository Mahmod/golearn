// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clientanalytics.proto

package clientanalytics

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

type LogRequest struct {
	ClientInfo           *ClientInfo `protobuf:"bytes,1,opt,name=client_info,json=clientInfo" json:"client_info,omitempty"`
	LogSource            *int32      `protobuf:"varint,2,opt,name=log_source,json=logSource" json:"log_source,omitempty"`
	RequestTimeMs        *int64      `protobuf:"varint,4,opt,name=request_time_ms,json=requestTimeMs" json:"request_time_ms,omitempty"`
	LogEvent             []*LogEvent `protobuf:"bytes,3,rep,name=log_event,json=logEvent" json:"log_event,omitempty"`
	LogSourceName        *string     `protobuf:"bytes,5,opt,name=log_source_name,json=logSourceName" json:"log_source_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_56bb8612b41df164, []int{0}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

func (m *LogRequest) GetLogSource() int32 {
	if m != nil && m.LogSource != nil {
		return *m.LogSource
	}
	return 0
}

func (m *LogRequest) GetRequestTimeMs() int64 {
	if m != nil && m.RequestTimeMs != nil {
		return *m.RequestTimeMs
	}
	return 0
}

func (m *LogRequest) GetLogEvent() []*LogEvent {
	if m != nil {
		return m.LogEvent
	}
	return nil
}

func (m *LogRequest) GetLogSourceName() string {
	if m != nil && m.LogSourceName != nil {
		return *m.LogSourceName
	}
	return ""
}

type ClientInfo struct {
	ClientType           *int32   `protobuf:"varint,1,opt,name=client_type,json=clientType" json:"client_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInfo) Reset()         { *m = ClientInfo{} }
func (m *ClientInfo) String() string { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()    {}
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_56bb8612b41df164, []int{1}
}

func (m *ClientInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInfo.Unmarshal(m, b)
}
func (m *ClientInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInfo.Marshal(b, m, deterministic)
}
func (m *ClientInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInfo.Merge(m, src)
}
func (m *ClientInfo) XXX_Size() int {
	return xxx_messageInfo_ClientInfo.Size(m)
}
func (m *ClientInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInfo proto.InternalMessageInfo

func (m *ClientInfo) GetClientType() int32 {
	if m != nil && m.ClientType != nil {
		return *m.ClientType
	}
	return 0
}

type LogResponse struct {
	NextRequestWaitMillis *int64   `protobuf:"varint,1,opt,name=next_request_wait_millis,json=nextRequestWaitMillis" json:"next_request_wait_millis,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_56bb8612b41df164, []int{2}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetNextRequestWaitMillis() int64 {
	if m != nil && m.NextRequestWaitMillis != nil {
		return *m.NextRequestWaitMillis
	}
	return 0
}

type LogEvent struct {
	EventTimeMs          *int64   `protobuf:"varint,1,opt,name=event_time_ms,json=eventTimeMs" json:"event_time_ms,omitempty"`
	SourceExtension      []byte   `protobuf:"bytes,6,opt,name=source_extension,json=sourceExtension" json:"source_extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEvent) Reset()         { *m = LogEvent{} }
func (m *LogEvent) String() string { return proto.CompactTextString(m) }
func (*LogEvent) ProtoMessage()    {}
func (*LogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56bb8612b41df164, []int{3}
}

func (m *LogEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEvent.Unmarshal(m, b)
}
func (m *LogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEvent.Marshal(b, m, deterministic)
}
func (m *LogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEvent.Merge(m, src)
}
func (m *LogEvent) XXX_Size() int {
	return xxx_messageInfo_LogEvent.Size(m)
}
func (m *LogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogEvent proto.InternalMessageInfo

func (m *LogEvent) GetEventTimeMs() int64 {
	if m != nil && m.EventTimeMs != nil {
		return *m.EventTimeMs
	}
	return 0
}

func (m *LogEvent) GetSourceExtension() []byte {
	if m != nil {
		return m.SourceExtension
	}
	return nil
}

func init() {
	proto.RegisterType((*LogRequest)(nil), "LogRequest")
	proto.RegisterType((*ClientInfo)(nil), "ClientInfo")
	proto.RegisterType((*LogResponse)(nil), "LogResponse")
	proto.RegisterType((*LogEvent)(nil), "LogEvent")
}

func init() {
	proto.RegisterFile("clientanalytics.proto", fileDescriptor_56bb8612b41df164)
}

var fileDescriptor_56bb8612b41df164 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0xc9, 0xb3, 0x4f, 0xa5, 0x9d, 0xb5, 0x54, 0x02, 0x85, 0x5c, 0xc4, 0x65, 0x0f, 0x65,
	0x05, 0xed, 0xa1, 0x17, 0x3f, 0x80, 0x54, 0x10, 0x5a, 0x0f, 0xb1, 0x20, 0x9e, 0x42, 0x58, 0xa6,
	0x4b, 0x20, 0x2f, 0x6b, 0x93, 0x6a, 0xfb, 0x31, 0xfd, 0x46, 0xb2, 0xd9, 0x97, 0xde, 0xc2, 0x6f,
	0x32, 0x33, 0xff, 0xdf, 0xc0, 0xbc, 0xd4, 0x0a, 0x6d, 0x90, 0x56, 0xea, 0x73, 0x50, 0xa5, 0x5f,
	0xd6, 0x07, 0x17, 0x5c, 0xfe, 0x4b, 0x00, 0x36, 0xae, 0xe2, 0xf8, 0x75, 0x44, 0x1f, 0xe8, 0x03,
	0xa4, 0xed, 0x3f, 0xa1, 0xec, 0xde, 0x31, 0x92, 0x91, 0x22, 0x5d, 0xa5, 0xcb, 0xe7, 0xc8, 0x5e,
	0xed, 0xde, 0x71, 0x28, 0x87, 0x37, 0xbd, 0x05, 0xd0, 0xae, 0x12, 0xde, 0x1d, 0x0f, 0x25, 0xb2,
	0x7f, 0x19, 0x29, 0x46, 0x7c, 0xa2, 0x5d, 0xf5, 0x1e, 0x01, 0x5d, 0xc0, 0xec, 0xd0, 0xce, 0x15,
	0x41, 0x19, 0x14, 0xc6, 0xb3, 0xff, 0x19, 0x29, 0x12, 0x3e, 0xed, 0xf0, 0x4e, 0x19, 0xdc, 0x7a,
	0xba, 0x80, 0xa6, 0x49, 0xe0, 0x37, 0xda, 0xc0, 0x92, 0x2c, 0x29, 0xd2, 0xd5, 0x64, 0xb9, 0x71,
	0xd5, 0xba, 0x01, 0x7c, 0xac, 0xbb, 0x57, 0x33, 0xef, 0xb2, 0x4e, 0x58, 0x69, 0x90, 0x8d, 0x32,
	0x52, 0x4c, 0xf8, 0x74, 0xd8, 0xf9, 0x26, 0x0d, 0xe6, 0x8f, 0x00, 0x97, 0xc0, 0xf4, 0x6e, 0x50,
	0x0a, 0xe7, 0x1a, 0xa3, 0xd2, 0xa8, 0xb7, 0xd8, 0x9d, 0x6b, 0xcc, 0x5f, 0x20, 0x8d, 0x17, 0xf0,
	0xb5, 0xb3, 0x1e, 0xe9, 0x13, 0x30, 0x8b, 0xa7, 0x20, 0xfa, 0xe8, 0x3f, 0x52, 0x05, 0x61, 0x94,
	0xd6, 0xca, 0xc7, 0xe6, 0x84, 0xcf, 0x9b, 0x7a, 0x77, 0xb1, 0x0f, 0xa9, 0xc2, 0x36, 0x16, 0xf3,
	0x4f, 0x18, 0xf7, 0xa1, 0x69, 0x0e, 0xd3, 0xa8, 0x33, 0x88, 0xb7, 0x9d, 0x69, 0x84, 0x9d, 0xf6,
	0x3d, 0xdc, 0x74, 0x2a, 0x78, 0x0a, 0x68, 0xbd, 0x72, 0x96, 0x5d, 0x65, 0xa4, 0xb8, 0xe6, 0xb3,
	0x96, 0xaf, 0x7b, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0x84, 0xe2, 0x94, 0xe2, 0xbd, 0x01, 0x00,
	0x00,
}
