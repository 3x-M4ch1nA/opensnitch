// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ui.proto

/*
Package ui is a generated protocol buffer package.

It is generated from these files:
	ui.proto

It has these top-level messages:
	PingRequest
	PingReply
	RuleRequest
	RuleReply
*/
package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PingRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PingReply struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingReply) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type RuleRequest struct {
	Protocol    string   `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	SrcIp       string   `protobuf:"bytes,2,opt,name=src_ip,json=srcIp" json:"src_ip,omitempty"`
	SrcPort     uint32   `protobuf:"varint,3,opt,name=src_port,json=srcPort" json:"src_port,omitempty"`
	DstIp       string   `protobuf:"bytes,4,opt,name=dst_ip,json=dstIp" json:"dst_ip,omitempty"`
	DstHost     string   `protobuf:"bytes,5,opt,name=dst_host,json=dstHost" json:"dst_host,omitempty"`
	DstPort     uint32   `protobuf:"varint,6,opt,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	ProcessId   uint32   `protobuf:"varint,7,opt,name=process_id,json=processId" json:"process_id,omitempty"`
	ProcessPath string   `protobuf:"bytes,8,opt,name=process_path,json=processPath" json:"process_path,omitempty"`
	ProcessArgs []string `protobuf:"bytes,9,rep,name=process_args,json=processArgs" json:"process_args,omitempty"`
}

func (m *RuleRequest) Reset()                    { *m = RuleRequest{} }
func (m *RuleRequest) String() string            { return proto.CompactTextString(m) }
func (*RuleRequest) ProtoMessage()               {}
func (*RuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RuleRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *RuleRequest) GetSrcIp() string {
	if m != nil {
		return m.SrcIp
	}
	return ""
}

func (m *RuleRequest) GetSrcPort() uint32 {
	if m != nil {
		return m.SrcPort
	}
	return 0
}

func (m *RuleRequest) GetDstIp() string {
	if m != nil {
		return m.DstIp
	}
	return ""
}

func (m *RuleRequest) GetDstHost() string {
	if m != nil {
		return m.DstHost
	}
	return ""
}

func (m *RuleRequest) GetDstPort() uint32 {
	if m != nil {
		return m.DstPort
	}
	return 0
}

func (m *RuleRequest) GetProcessId() uint32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *RuleRequest) GetProcessPath() string {
	if m != nil {
		return m.ProcessPath
	}
	return ""
}

func (m *RuleRequest) GetProcessArgs() []string {
	if m != nil {
		return m.ProcessArgs
	}
	return nil
}

type RuleReply struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Duration string `protobuf:"bytes,3,opt,name=duration" json:"duration,omitempty"`
	What     string `protobuf:"bytes,4,opt,name=what" json:"what,omitempty"`
	Value    string `protobuf:"bytes,5,opt,name=value" json:"value,omitempty"`
}

func (m *RuleReply) Reset()                    { *m = RuleReply{} }
func (m *RuleReply) String() string            { return proto.CompactTextString(m) }
func (*RuleReply) ProtoMessage()               {}
func (*RuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RuleReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RuleReply) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RuleReply) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *RuleReply) GetWhat() string {
	if m != nil {
		return m.What
	}
	return ""
}

func (m *RuleReply) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "ui.PingRequest")
	proto.RegisterType((*PingReply)(nil), "ui.PingReply")
	proto.RegisterType((*RuleRequest)(nil), "ui.RuleRequest")
	proto.RegisterType((*RuleReply)(nil), "ui.RuleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UI service

type UIClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	AskRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleReply, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/ui.UI/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) AskRule(ctx context.Context, in *RuleRequest, opts ...grpc.CallOption) (*RuleReply, error) {
	out := new(RuleReply)
	err := grpc.Invoke(ctx, "/ui.UI/AskRule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UI service

type UIServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	AskRule(context.Context, *RuleRequest) (*RuleReply, error)
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.UI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_AskRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).AskRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.UI/AskRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).AskRule(ctx, req.(*RuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ui.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _UI_Ping_Handler,
		},
		{
			MethodName: "AskRule",
			Handler:    _UI_AskRule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ui.proto",
}

func init() { proto.RegisterFile("ui.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x6d, 0xd3, 0x34, 0x8f, 0x5b, 0xab, 0x30, 0xa8, 0xc4, 0x4a, 0xa1, 0x66, 0x55, 0x10, 0xba,
	0xd0, 0x2f, 0xe8, 0xce, 0xee, 0x4a, 0xc0, 0x95, 0x8b, 0x32, 0x66, 0x42, 0x33, 0x18, 0x3b, 0xe3,
	0xdc, 0x89, 0xd2, 0x85, 0x9f, 0xe0, 0x3f, 0xcb, 0x3c, 0x52, 0x03, 0xee, 0xe6, 0x9e, 0x17, 0xc9,
	0x39, 0x90, 0xb4, 0x7c, 0x25, 0x95, 0xd0, 0x82, 0x04, 0x2d, 0xcf, 0xe7, 0x30, 0xd9, 0xf2, 0xc3,
	0xbe, 0xa8, 0x3e, 0xda, 0x0a, 0x35, 0x39, 0x87, 0x80, 0xb3, 0x6c, 0xb8, 0x18, 0x2e, 0xc3, 0x22,
	0xe0, 0x2c, 0xbf, 0x85, 0xd4, 0xd1, 0xb2, 0x39, 0xfe, 0x23, 0x7f, 0x02, 0x98, 0x14, 0x6d, 0x53,
	0x75, 0xe6, 0x19, 0x24, 0x36, 0xb8, 0x14, 0x8d, 0x55, 0xa5, 0xc5, 0xe9, 0x26, 0x57, 0x10, 0xa1,
	0x2a, 0x77, 0x5c, 0x66, 0x81, 0x65, 0xc6, 0xa8, 0xca, 0x8d, 0x24, 0x37, 0x90, 0x18, 0x58, 0x0a,
	0xa5, 0xb3, 0xd1, 0x62, 0xb8, 0x9c, 0x16, 0x31, 0xaa, 0x72, 0x2b, 0x94, 0x36, 0x0e, 0x86, 0xda,
	0x38, 0x42, 0xe7, 0x60, 0xa8, 0x9d, 0xc3, 0xc0, 0xb5, 0x40, 0x9d, 0x8d, 0x2d, 0x11, 0x33, 0xd4,
	0x4f, 0x02, 0x75, 0x47, 0xd9, 0xb0, 0xc8, 0x85, 0x31, 0xd4, 0x36, 0x6c, 0x0e, 0x20, 0x95, 0x28,
	0x2b, 0xc4, 0x1d, 0x67, 0x59, 0x6c, 0xc9, 0xd4, 0x23, 0x1b, 0x46, 0xee, 0xe0, 0xac, 0xa3, 0x25,
	0xd5, 0x75, 0x96, 0xd8, 0xe0, 0x89, 0xc7, 0xb6, 0x54, 0xd7, 0x7d, 0x09, 0x55, 0x7b, 0xcc, 0xd2,
	0xc5, 0xa8, 0x27, 0x59, 0xab, 0x3d, 0xe6, 0xdf, 0x90, 0xba, 0x3a, 0x4c, 0x59, 0x04, 0xc2, 0x03,
	0x7d, 0xaf, 0x7c, 0x11, 0xf6, 0x4d, 0xae, 0x21, 0xa2, 0xa5, 0xe6, 0xe2, 0xe0, 0x4b, 0xf0, 0x97,
	0x29, 0x8e, 0xb5, 0x8a, 0x5a, 0x66, 0xe4, 0x8a, 0xeb, 0x6e, 0x93, 0xf3, 0x55, 0x53, 0xed, 0x4b,
	0xb0, 0x6f, 0x72, 0x09, 0xe3, 0x4f, 0xda, 0xb4, 0x95, 0x2f, 0xc0, 0x1d, 0x0f, 0x2f, 0x10, 0x3c,
	0x6f, 0xc8, 0x12, 0x42, 0xb3, 0x18, 0xb9, 0x58, 0xb5, 0x7c, 0xd5, 0x9b, 0x76, 0x36, 0xfd, 0x03,
	0x64, 0x73, 0xcc, 0x07, 0xe4, 0x1e, 0xe2, 0x35, 0xbe, 0x99, 0x2f, 0x76, 0xe2, 0xde, 0x94, 0x4e,
	0x7c, 0xfa, 0x99, 0x7c, 0xf0, 0x1a, 0xd9, 0x25, 0x1f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xa3, 0x76, 0xb5, 0x3e, 0x02, 0x00, 0x00,
}
