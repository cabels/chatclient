// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chatpb

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

type SendMessageRequest struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type SendMessageResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(m, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo

func (m *SendMessageResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type ReceiveMessagesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReceiveMessagesRequest) Reset()         { *m = ReceiveMessagesRequest{} }
func (m *ReceiveMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessagesRequest) ProtoMessage()    {}
func (*ReceiveMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *ReceiveMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMessagesRequest.Unmarshal(m, b)
}
func (m *ReceiveMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMessagesRequest.Marshal(b, m, deterministic)
}
func (m *ReceiveMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMessagesRequest.Merge(m, src)
}
func (m *ReceiveMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_ReceiveMessagesRequest.Size(m)
}
func (m *ReceiveMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMessagesRequest proto.InternalMessageInfo

type ReceiveMessagesResponse struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReceiveMessagesResponse) Reset()         { *m = ReceiveMessagesResponse{} }
func (m *ReceiveMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*ReceiveMessagesResponse) ProtoMessage()    {}
func (*ReceiveMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *ReceiveMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiveMessagesResponse.Unmarshal(m, b)
}
func (m *ReceiveMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiveMessagesResponse.Marshal(b, m, deterministic)
}
func (m *ReceiveMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiveMessagesResponse.Merge(m, src)
}
func (m *ReceiveMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_ReceiveMessagesResponse.Size(m)
}
func (m *ReceiveMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiveMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiveMessagesResponse proto.InternalMessageInfo

func (m *ReceiveMessagesResponse) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Message struct {
	Sender               string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{4}
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

func (m *Message) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "chatpb.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "chatpb.SendMessageResponse")
	proto.RegisterType((*ReceiveMessagesRequest)(nil), "chatpb.ReceiveMessagesRequest")
	proto.RegisterType((*ReceiveMessagesResponse)(nil), "chatpb.ReceiveMessagesResponse")
	proto.RegisterType((*Message)(nil), "chatpb.Message")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0xb1, 0x0b, 0x92, 0x94, 0xec, 0xb9, 0x84,
	0x82, 0x53, 0xf3, 0x52, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x83, 0x52, 0x0b, 0x4b, 0x53,
	0x8b, 0x4b, 0x84, 0x34, 0xb9, 0xd8, 0x73, 0x21, 0x22, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46,
	0xfc, 0x7a, 0x10, 0xf5, 0x7a, 0x30, 0x85, 0x30, 0x79, 0x25, 0x27, 0x2e, 0x61, 0x14, 0x03, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xb4, 0xb9, 0x38, 0xa0, 0x2a, 0x8a, 0x25, 0x18, 0x15, 0x98,
	0xb1, 0x19, 0x01, 0x57, 0xa0, 0x24, 0xc1, 0x25, 0x16, 0x94, 0x9a, 0x9c, 0x9a, 0x59, 0x96, 0x0a,
	0x95, 0x2b, 0x86, 0x3a, 0x44, 0xc9, 0x8d, 0x4b, 0x1c, 0x43, 0x86, 0x1c, 0x1b, 0x4c, 0xb9, 0xd8,
	0xa1, 0x82, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x60, 0xaf, 0x71, 0x06,
	0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0x25, 0xa9, 0x15, 0x25, 0x12, 0x4c, 0x60, 0x51, 0x30, 0xdb,
	0x68, 0x05, 0x23, 0x17, 0x97, 0x73, 0x46, 0x62, 0x49, 0x70, 0x6a, 0x51, 0x59, 0x6a, 0x91, 0x90,
	0x1b, 0x17, 0x37, 0x92, 0x5f, 0x85, 0xa4, 0x60, 0xf6, 0x61, 0x86, 0xa0, 0x94, 0x34, 0x56, 0x39,
	0xa8, 0xd3, 0x83, 0xb8, 0xf8, 0xd1, 0x7c, 0x25, 0x24, 0x07, 0x53, 0x8f, 0x3d, 0x20, 0xa4, 0xe4,
	0x71, 0xca, 0x43, 0xcc, 0x4c, 0x62, 0x03, 0xc7, 0xab, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x39,
	0x7c, 0x6c, 0x6b, 0xe5, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServerClient is the client API for ChatServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServerClient interface {
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	ReceiveMessages(ctx context.Context, in *ReceiveMessagesRequest, opts ...grpc.CallOption) (*ReceiveMessagesResponse, error)
}

type chatServerClient struct {
	cc *grpc.ClientConn
}

func NewChatServerClient(cc *grpc.ClientConn) ChatServerClient {
	return &chatServerClient{cc}
}

func (c *chatServerClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/chatpb.ChatServer/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServerClient) ReceiveMessages(ctx context.Context, in *ReceiveMessagesRequest, opts ...grpc.CallOption) (*ReceiveMessagesResponse, error) {
	out := new(ReceiveMessagesResponse)
	err := c.cc.Invoke(ctx, "/chatpb.ChatServer/ReceiveMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServerServer is the server API for ChatServer service.
type ChatServerServer interface {
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	ReceiveMessages(context.Context, *ReceiveMessagesRequest) (*ReceiveMessagesResponse, error)
}

// UnimplementedChatServerServer can be embedded to have forward compatible implementations.
type UnimplementedChatServerServer struct {
}

func (*UnimplementedChatServerServer) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (*UnimplementedChatServerServer) ReceiveMessages(ctx context.Context, req *ReceiveMessagesRequest) (*ReceiveMessagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveMessages not implemented")
}

func RegisterChatServerServer(s *grpc.Server, srv ChatServerServer) {
	s.RegisterService(&_ChatServer_serviceDesc, srv)
}

func _ChatServer_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServerServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.ChatServer/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServerServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServer_ReceiveMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServerServer).ReceiveMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatpb.ChatServer/ReceiveMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServerServer).ReceiveMessages(ctx, req.(*ReceiveMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chatpb.ChatServer",
	HandlerType: (*ChatServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatServer_SendMessage_Handler,
		},
		{
			MethodName: "ReceiveMessages",
			Handler:    _ChatServer_ReceiveMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
