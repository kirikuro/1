// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub.proto

package sonm

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

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ListReply struct {
	Info map[string]*ListReply_ListValue `protobuf:"bytes,1,rep,name=info" json:"info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListReply) Reset()                    { *m = ListReply{} }
func (m *ListReply) String() string            { return proto.CompactTextString(m) }
func (*ListReply) ProtoMessage()               {}
func (*ListReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ListReply) GetInfo() map[string]*ListReply_ListValue {
	if m != nil {
		return m.Info
	}
	return nil
}

type ListReply_ListValue struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *ListReply_ListValue) Reset()                    { *m = ListReply_ListValue{} }
func (m *ListReply_ListValue) String() string            { return proto.CompactTextString(m) }
func (*ListReply_ListValue) ProtoMessage()               {}
func (*ListReply_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *ListReply_ListValue) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type HubInfoRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubInfoRequest) Reset()                    { *m = HubInfoRequest{} }
func (m *HubInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*HubInfoRequest) ProtoMessage()               {}
func (*HubInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *HubInfoRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type TaskResourceRequirements struct {
	// Number of CPU cores.
	CPUCores uint64 `protobuf:"varint,1,opt,name=CPUCores" json:"CPUCores,omitempty"`
	// Memory in bytes required.
	MaxMemory uint64 `protobuf:"varint,2,opt,name=maxMemory" json:"maxMemory,omitempty"`
	// Describes whether a task requires GPU supoort.
	GPUSupport bool `protobuf:"varint,3,opt,name=GPUSupport" json:"GPUSupport,omitempty"`
}

func (m *TaskResourceRequirements) Reset()                    { *m = TaskResourceRequirements{} }
func (m *TaskResourceRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskResourceRequirements) ProtoMessage()               {}
func (*TaskResourceRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TaskResourceRequirements) GetCPUCores() uint64 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *TaskResourceRequirements) GetMaxMemory() uint64 {
	if m != nil {
		return m.MaxMemory
	}
	return 0
}

func (m *TaskResourceRequirements) GetGPUSupport() bool {
	if m != nil {
		return m.GPUSupport
	}
	return false
}

type TaskRequirements struct {
	// How much resources consumes a task. This is used both for scheduling and for cgroups configuration.
	Resources *TaskResourceRequirements `protobuf:"bytes,1,opt,name=resources" json:"resources,omitempty"`
	// Optional miner ID restrictions (currently IP:port), that are allowed to start a task.
	Miners []string `protobuf:"bytes,2,rep,name=miners" json:"miners,omitempty"`
}

func (m *TaskRequirements) Reset()                    { *m = TaskRequirements{} }
func (m *TaskRequirements) String() string            { return proto.CompactTextString(m) }
func (*TaskRequirements) ProtoMessage()               {}
func (*TaskRequirements) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *TaskRequirements) GetResources() *TaskResourceRequirements {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *TaskRequirements) GetMiners() []string {
	if m != nil {
		return m.Miners
	}
	return nil
}

type HubStartTaskRequest struct {
	Requirements  *TaskRequirements `protobuf:"bytes,1,opt,name=requirements" json:"requirements,omitempty"`
	Registry      string            `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string            `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string            `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	PublicKeyData string            `protobuf:"bytes,5,opt,name=PublicKeyData" json:"PublicKeyData,omitempty"`
}

func (m *HubStartTaskRequest) Reset()                    { *m = HubStartTaskRequest{} }
func (m *HubStartTaskRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskRequest) ProtoMessage()               {}
func (*HubStartTaskRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *HubStartTaskRequest) GetRequirements() *TaskRequirements {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *HubStartTaskRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *HubStartTaskRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *HubStartTaskRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *HubStartTaskRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

type HubStartTaskReply struct {
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Endpoint []string `protobuf:"bytes,2,rep,name=endpoint" json:"endpoint,omitempty"`
}

func (m *HubStartTaskReply) Reset()                    { *m = HubStartTaskReply{} }
func (m *HubStartTaskReply) String() string            { return proto.CompactTextString(m) }
func (*HubStartTaskReply) ProtoMessage()               {}
func (*HubStartTaskReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *HubStartTaskReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *HubStartTaskReply) GetEndpoint() []string {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type HubStatusMapRequest struct {
	Miner string `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
}

func (m *HubStatusMapRequest) Reset()                    { *m = HubStatusMapRequest{} }
func (m *HubStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusMapRequest) ProtoMessage()               {}
func (*HubStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *HubStatusMapRequest) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

type HubStatusRequest struct {
}

func (m *HubStatusRequest) Reset()                    { *m = HubStatusRequest{} }
func (m *HubStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*HubStatusRequest) ProtoMessage()               {}
func (*HubStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type HubStatusReply struct {
	MinerCount uint64 `protobuf:"varint,1,opt,name=minerCount" json:"minerCount,omitempty"`
	Uptime     uint64 `protobuf:"varint,2,opt,name=uptime" json:"uptime,omitempty"`
}

func (m *HubStatusReply) Reset()                    { *m = HubStatusReply{} }
func (m *HubStatusReply) String() string            { return proto.CompactTextString(m) }
func (*HubStatusReply) ProtoMessage()               {}
func (*HubStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *HubStatusReply) GetMinerCount() uint64 {
	if m != nil {
		return m.MinerCount
	}
	return 0
}

func (m *HubStatusReply) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "sonm.ListRequest")
	proto.RegisterType((*ListReply)(nil), "sonm.ListReply")
	proto.RegisterType((*ListReply_ListValue)(nil), "sonm.ListReply.ListValue")
	proto.RegisterType((*HubInfoRequest)(nil), "sonm.HubInfoRequest")
	proto.RegisterType((*TaskResourceRequirements)(nil), "sonm.TaskResourceRequirements")
	proto.RegisterType((*TaskRequirements)(nil), "sonm.TaskRequirements")
	proto.RegisterType((*HubStartTaskRequest)(nil), "sonm.HubStartTaskRequest")
	proto.RegisterType((*HubStartTaskReply)(nil), "sonm.HubStartTaskReply")
	proto.RegisterType((*HubStatusMapRequest)(nil), "sonm.HubStatusMapRequest")
	proto.RegisterType((*HubStatusRequest)(nil), "sonm.HubStatusRequest")
	proto.RegisterType((*HubStatusReply)(nil), "sonm.HubStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hub service

type HubClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error)
	Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error)
	StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error)
}

type hubClient struct {
	cc *grpc.ClientConn
}

func NewHubClient(cc *grpc.ClientConn) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Status(ctx context.Context, in *HubStatusRequest, opts ...grpc.CallOption) (*HubStatusReply, error) {
	out := new(HubStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) Info(ctx context.Context, in *HubInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StartTask(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*HubStartTaskReply, error) {
	out := new(HubStartTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StartTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) StopTask(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/StopTask", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskStatus(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/TaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) MinerStatus(ctx context.Context, in *HubStatusMapRequest, opts ...grpc.CallOption) (*StatusMapReply, error) {
	out := new(StatusMapReply)
	err := grpc.Invoke(ctx, "/sonm.Hub/MinerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Hub_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hub_serviceDesc.Streams[0], c.cc, "/sonm.Hub/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &hubTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hub_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type hubTaskLogsClient struct {
	grpc.ClientStream
}

func (x *hubTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hub service

type HubServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Status(context.Context, *HubStatusRequest) (*HubStatusReply, error)
	List(context.Context, *ListRequest) (*ListReply, error)
	Info(context.Context, *HubInfoRequest) (*InfoReply, error)
	StartTask(context.Context, *HubStartTaskRequest) (*HubStartTaskReply, error)
	StopTask(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TaskStatus(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	MinerStatus(context.Context, *HubStatusMapRequest) (*StatusMapReply, error)
	TaskLogs(*TaskLogsRequest, Hub_TaskLogsServer) error
}

func RegisterHubServer(s *grpc.Server, srv HubServer) {
	s.RegisterService(&_Hub_serviceDesc, srv)
}

func _Hub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Status(ctx, req.(*HubStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).Info(ctx, req.(*HubInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StartTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StartTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StartTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StartTask(ctx, req.(*HubStartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_StopTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).StopTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/StopTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).StopTask(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).TaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/TaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).TaskStatus(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_MinerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStatusMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).MinerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Hub/MinerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).MinerStatus(ctx, req.(*HubStatusMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hub_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HubServer).TaskLogs(m, &hubTaskLogsServer{stream})
}

type Hub_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type hubTaskLogsServer struct {
	grpc.ServerStream
}

func (x *hubTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

var _Hub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Hub_Ping_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Hub_Status_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Hub_List_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Hub_Info_Handler,
		},
		{
			MethodName: "StartTask",
			Handler:    _Hub_StartTask_Handler,
		},
		{
			MethodName: "StopTask",
			Handler:    _Hub_StopTask_Handler,
		},
		{
			MethodName: "TaskStatus",
			Handler:    _Hub_TaskStatus_Handler,
		},
		{
			MethodName: "MinerStatus",
			Handler:    _Hub_MinerStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TaskLogs",
			Handler:       _Hub_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "hub.proto",
}

func init() { proto.RegisterFile("hub.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x13, 0xb7, 0x8a, 0x27, 0xf4, 0xb6, 0xbd, 0xb9, 0x16, 0xaa, 0x22, 0x83, 0x50, 0x24,
	0x4a, 0x8a, 0xca, 0x4b, 0x55, 0x21, 0x2e, 0x0a, 0x88, 0x20, 0x5a, 0x29, 0xda, 0x52, 0xde, 0x9d,
	0x76, 0x9b, 0xac, 0x12, 0xef, 0x1a, 0x7b, 0x17, 0xe1, 0x2f, 0xe0, 0x83, 0x78, 0xe6, 0xdf, 0xd0,
	0x5e, 0xec, 0x6c, 0xd2, 0xc2, 0x53, 0x7c, 0xce, 0xce, 0x99, 0x33, 0x33, 0x3b, 0x1b, 0x08, 0x26,
	0x72, 0xd4, 0xcb, 0x72, 0x2e, 0x38, 0xf2, 0x0b, 0xce, 0xd2, 0x68, 0x93, 0x32, 0xf5, 0xcb, 0x68,
	0x62, 0xe8, 0x78, 0x1d, 0xda, 0x17, 0xb4, 0x10, 0x98, 0x7c, 0x97, 0xa4, 0x10, 0xf1, 0x6f, 0x0f,
	0x02, 0x83, 0xb3, 0x59, 0x89, 0x5e, 0x80, 0x4f, 0xd9, 0x1d, 0x0f, 0xbd, 0x4e, 0xb3, 0xdb, 0x3e,
	0x3d, 0xec, 0x29, 0x69, 0xaf, 0x3e, 0xee, 0x7d, 0x66, 0x77, 0xfc, 0x23, 0x13, 0x79, 0x89, 0x75,
	0x58, 0xf4, 0xc4, 0x68, 0xbf, 0x25, 0x33, 0x49, 0xd0, 0x3e, 0xac, 0xfd, 0x50, 0x1f, 0x85, 0x56,
	0x07, 0xd8, 0xa2, 0x08, 0x43, 0x50, 0xeb, 0xd0, 0x16, 0x34, 0xa7, 0xa4, 0x0c, 0xbd, 0x8e, 0xd7,
	0x0d, 0xb0, 0xfa, 0x44, 0x27, 0xb0, 0xaa, 0x03, 0xc3, 0x46, 0xc7, 0x7b, 0xc8, 0xb3, 0x36, 0xc0,
	0x26, 0xee, 0xbc, 0x71, 0xe6, 0xc5, 0xcf, 0x60, 0x63, 0x20, 0x47, 0x2a, 0xad, 0xed, 0x03, 0xed,
	0xc2, 0x6a, 0x4a, 0x19, 0xc9, 0x6d, 0x6a, 0x03, 0x62, 0x01, 0xe1, 0xd7, 0xa4, 0x98, 0x62, 0x52,
	0x70, 0x99, 0xdf, 0x10, 0x15, 0x4c, 0x73, 0x92, 0x12, 0x26, 0x0a, 0x14, 0x41, 0xab, 0x3f, 0xbc,
	0xee, 0xf3, 0x5c, 0x57, 0xec, 0x75, 0x7d, 0x5c, 0x63, 0xf4, 0x18, 0x82, 0x34, 0xf9, 0x79, 0x49,
	0x52, 0x9e, 0x97, 0xba, 0x30, 0x1f, 0xcf, 0x09, 0x74, 0x04, 0xf0, 0x69, 0x78, 0x7d, 0x25, 0xb3,
	0x8c, 0xe7, 0x22, 0x6c, 0x76, 0xbc, 0x6e, 0x0b, 0x3b, 0x4c, 0x3c, 0x81, 0x2d, 0xe3, 0xea, 0xb8,
	0xbd, 0x86, 0x20, 0xb7, 0x55, 0x18, 0xbb, 0xf6, 0xe9, 0x91, 0x69, 0xf5, 0x5f, 0x05, 0xe2, 0xb9,
	0x40, 0xcd, 0x56, 0x37, 0x54, 0x84, 0x0d, 0x33, 0x5b, 0x83, 0xe2, 0x3f, 0x1e, 0xec, 0x0c, 0xe4,
	0xe8, 0x4a, 0x24, 0xb9, 0xa8, 0x2c, 0xd5, 0x34, 0xce, 0xe1, 0x51, 0xee, 0xa4, 0xb2, 0x86, 0xfb,
	0xae, 0xa1, 0x63, 0xb4, 0x10, 0xab, 0xe6, 0x92, 0x93, 0x31, 0x2d, 0x84, 0x6d, 0x3d, 0xc0, 0x35,
	0x56, 0x53, 0xa6, 0x69, 0x32, 0x26, 0xba, 0xe9, 0x00, 0x1b, 0x80, 0x10, 0xf8, 0x89, 0x14, 0x93,
	0xd0, 0xd7, 0xa4, 0xfe, 0x46, 0x4f, 0x61, 0x7d, 0x28, 0x47, 0x33, 0x7a, 0xf3, 0x85, 0x94, 0x1f,
	0x12, 0x91, 0x84, 0xab, 0xfa, 0x70, 0x91, 0x8c, 0xdf, 0xc2, 0xf6, 0x62, 0xf9, 0x6a, 0x09, 0x37,
	0xa0, 0x41, 0x6f, 0xed, 0x3d, 0x36, 0xe8, 0xad, 0x2a, 0x88, 0xb0, 0xdb, 0x8c, 0x53, 0x26, 0x6c,
	0xfb, 0x35, 0x8e, 0x9f, 0x57, 0xfd, 0x0b, 0x59, 0x5c, 0x26, 0xd9, 0xff, 0xb7, 0x01, 0xc1, 0x56,
	0x1d, 0x5c, 0xed, 0xff, 0x40, 0x6f, 0x52, 0xc5, 0x29, 0xfb, 0x23, 0x00, 0x1d, 0xde, 0xe7, 0x92,
	0x09, 0xbb, 0x19, 0x0e, 0xa3, 0xee, 0x42, 0x66, 0x82, 0xa6, 0xc4, 0x2e, 0x86, 0x45, 0xa7, 0xbf,
	0x7c, 0x68, 0x0e, 0xe4, 0x08, 0x1d, 0x83, 0x3f, 0xa4, 0x6c, 0x8c, 0xb6, 0xcd, 0xb4, 0xd5, 0xb7,
	0x35, 0x8b, 0x36, 0x5d, 0x2a, 0x9b, 0x95, 0xf1, 0x0a, 0x3a, 0x83, 0x35, 0x63, 0x8e, 0xec, 0xed,
	0x2c, 0x57, 0x18, 0xed, 0xde, 0xe3, 0x8d, 0xf2, 0x18, 0x7c, 0xf5, 0x36, 0x2a, 0x1f, 0xe7, 0x51,
	0x57, 0x3e, 0xf5, 0x23, 0x8a, 0x57, 0xd0, 0x09, 0xf8, 0xea, 0xb9, 0xa0, 0x79, 0x36, 0xe7, 0xf5,
	0x54, 0x02, 0x43, 0x19, 0xc1, 0x7b, 0x08, 0xea, 0x7b, 0x41, 0x87, 0x6e, 0x0d, 0x0b, 0xab, 0x16,
	0x1d, 0x3c, 0x74, 0x54, 0xf5, 0xd6, 0xba, 0x12, 0x3c, 0xd3, 0x19, 0xf6, 0x4c, 0x58, 0x85, 0x2b,
	0xf5, 0xce, 0x32, 0x6d, 0x94, 0x6f, 0x00, 0x14, 0xb4, 0x93, 0x39, 0x98, 0xef, 0xed, 0xe2, 0x68,
	0xf6, 0xee, 0x1f, 0x18, 0xfd, 0x3b, 0x68, 0x5f, 0xaa, 0x1b, 0xb3, 0x09, 0x0e, 0x97, 0x46, 0x38,
	0xdf, 0x94, 0x6a, 0xba, 0x0e, 0x6f, 0x32, 0x9c, 0x43, 0x4b, 0xa5, 0xbd, 0xe0, 0xe3, 0x02, 0x39,
	0x36, 0x0a, 0x2f, 0xd5, 0x5e, 0xd1, 0xfd, 0x89, 0x64, 0xd3, 0x78, 0xe5, 0xa5, 0x37, 0x5a, 0xd3,
	0xff, 0xb4, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x08, 0xd5, 0xfc, 0xb3, 0x8d, 0x05, 0x00,
	0x00,
}
