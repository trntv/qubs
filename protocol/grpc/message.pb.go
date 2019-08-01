// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Queue                string   `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *Message) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type Ack struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ack) Reset()         { *m = Ack{} }
func (m *Ack) String() string { return proto.CompactTextString(m) }
func (*Ack) ProtoMessage()    {}
func (*Ack) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Ack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ack.Unmarshal(m, b)
}
func (m *Ack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ack.Marshal(b, m, deterministic)
}
func (m *Ack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ack.Merge(m, src)
}
func (m *Ack) XXX_Size() int {
	return xxx_messageInfo_Ack.Size(m)
}
func (m *Ack) XXX_DiscardUnknown() {
	xxx_messageInfo_Ack.DiscardUnknown(m)
}

var xxx_messageInfo_Ack proto.InternalMessageInfo

func (m *Ack) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ConsumerConnect struct {
	Queue                string   `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Tags                 []string `protobuf:"bytes,2,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerConnect) Reset()         { *m = ConsumerConnect{} }
func (m *ConsumerConnect) String() string { return proto.CompactTextString(m) }
func (*ConsumerConnect) ProtoMessage()    {}
func (*ConsumerConnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *ConsumerConnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerConnect.Unmarshal(m, b)
}
func (m *ConsumerConnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerConnect.Marshal(b, m, deterministic)
}
func (m *ConsumerConnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerConnect.Merge(m, src)
}
func (m *ConsumerConnect) XXX_Size() int {
	return xxx_messageInfo_ConsumerConnect.Size(m)
}
func (m *ConsumerConnect) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerConnect.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerConnect proto.InternalMessageInfo

func (m *ConsumerConnect) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *ConsumerConnect) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "protocol.grpc.Message")
	proto.RegisterType((*Ack)(nil), "protocol.grpc.Ack")
	proto.RegisterType((*ConsumerConnect)(nil), "protocol.grpc.ConsumerConnect")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56,
	0x8a, 0xe5, 0x62, 0xf7, 0x85, 0x08, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x16, 0x96, 0xa6, 0x96, 0xa6, 0x4a,
	0x30, 0x81, 0x85, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0x92, 0xc4, 0xf4, 0x62, 0x09, 0x66, 0x05,
	0x66, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x48, 0x82, 0x8b, 0xbd, 0x20, 0xb1, 0x32, 0x27, 0x3f, 0x31,
	0x45, 0x82, 0x45, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc6, 0x55, 0x12, 0xe5, 0x62, 0x76, 0x4c, 0xce,
	0x46, 0x37, 0x5a, 0xc9, 0x9a, 0x8b, 0xdf, 0x39, 0x3f, 0xaf, 0xb8, 0x34, 0x37, 0xb5, 0xc8, 0x39,
	0x3f, 0x2f, 0x2f, 0x35, 0xb9, 0x04, 0x61, 0x1b, 0x23, 0x36, 0xdb, 0x98, 0x10, 0xb6, 0x19, 0x4d,
	0x66, 0x84, 0x2a, 0x15, 0x32, 0xe6, 0x62, 0x87, 0x1a, 0x23, 0x24, 0xaa, 0x07, 0xf2, 0x8e, 0x1e,
	0x9a, 0xa9, 0x52, 0xbc, 0x10, 0x61, 0xa8, 0x17, 0x95, 0x18, 0x0c, 0x18, 0x85, 0x54, 0xb9, 0xd8,
	0x5d, 0xf3, 0x20, 0xfa, 0x51, 0x65, 0xa5, 0x38, 0x21, 0x5c, 0xc7, 0xe4, 0x6c, 0x25, 0x06, 0x21,
	0x3d, 0x2e, 0x1e, 0xa8, 0x32, 0xa7, 0xc4, 0x92, 0xe4, 0x0c, 0x7c, 0x6a, 0x35, 0x18, 0x0d, 0x18,
	0x93, 0xd8, 0xc0, 0xa1, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x8a, 0x6f, 0x9b, 0x66,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protocol.grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueueClient is the client API for Queue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueueClient interface {
	Consume(ctx context.Context, in *ConsumerConnect, opts ...grpc.CallOption) (Queue_ConsumeClient, error)
	Enqueue(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Ack, error)
	EnqueueBatch(ctx context.Context, opts ...grpc.CallOption) (Queue_EnqueueBatchClient, error)
}

type queueClient struct {
	cc *grpc.ClientConn
}

func NewQueueClient(cc *grpc.ClientConn) QueueClient {
	return &queueClient{cc}
}

func (c *queueClient) Consume(ctx context.Context, in *ConsumerConnect, opts ...grpc.CallOption) (Queue_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Queue_serviceDesc.Streams[0], "/protocol.grpc.queue/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &queueConsumeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Queue_ConsumeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type queueConsumeClient struct {
	grpc.ClientStream
}

func (x *queueConsumeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *queueClient) Enqueue(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/protocol.grpc.queue/Enqueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueClient) EnqueueBatch(ctx context.Context, opts ...grpc.CallOption) (Queue_EnqueueBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Queue_serviceDesc.Streams[1], "/protocol.grpc.queue/EnqueueBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &queueEnqueueBatchClient{stream}
	return x, nil
}

type Queue_EnqueueBatchClient interface {
	Send(*Message) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type queueEnqueueBatchClient struct {
	grpc.ClientStream
}

func (x *queueEnqueueBatchClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *queueEnqueueBatchClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueueServer is the server API for Queue service.
type QueueServer interface {
	Consume(*ConsumerConnect, Queue_ConsumeServer) error
	Enqueue(context.Context, *Message) (*Ack, error)
	EnqueueBatch(Queue_EnqueueBatchServer) error
}

// UnimplementedQueueServer can be embedded to have forward compatible implementations.
type UnimplementedQueueServer struct {
}

func (*UnimplementedQueueServer) Consume(req *ConsumerConnect, srv Queue_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}
func (*UnimplementedQueueServer) Enqueue(ctx context.Context, req *Message) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enqueue not implemented")
}
func (*UnimplementedQueueServer) EnqueueBatch(srv Queue_EnqueueBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method EnqueueBatch not implemented")
}

func RegisterQueueServer(s *grpc.Server, srv QueueServer) {
	s.RegisterService(&_Queue_serviceDesc, srv)
}

func _Queue_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumerConnect)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueueServer).Consume(m, &queueConsumeServer{stream})
}

type Queue_ConsumeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type queueConsumeServer struct {
	grpc.ServerStream
}

func (x *queueConsumeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Queue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.grpc.queue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServer).Enqueue(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _Queue_EnqueueBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(QueueServer).EnqueueBatch(&queueEnqueueBatchServer{stream})
}

type Queue_EnqueueBatchServer interface {
	Send(*Ack) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type queueEnqueueBatchServer struct {
	grpc.ServerStream
}

func (x *queueEnqueueBatchServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *queueEnqueueBatchServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Queue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.grpc.queue",
	HandlerType: (*QueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _Queue_Enqueue_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Consume",
			Handler:       _Queue_Consume_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EnqueueBatch",
			Handler:       _Queue_EnqueueBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
