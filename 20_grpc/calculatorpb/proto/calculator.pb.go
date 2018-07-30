// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/calculator.proto

package calculatorpb

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

type CalculatorRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_0c0b5626957dde3f, []int{0}
}
func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (dst *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(dst, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CalculatorResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_0c0b5626957dde3f, []int{1}
}
func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (dst *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(dst, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatorAverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorAverageResponse) Reset()         { *m = CalculatorAverageResponse{} }
func (m *CalculatorAverageResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorAverageResponse) ProtoMessage()    {}
func (*CalculatorAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_0c0b5626957dde3f, []int{2}
}
func (m *CalculatorAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorAverageResponse.Unmarshal(m, b)
}
func (m *CalculatorAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorAverageResponse.Marshal(b, m, deterministic)
}
func (dst *CalculatorAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorAverageResponse.Merge(dst, src)
}
func (m *CalculatorAverageResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorAverageResponse.Size(m)
}
func (m *CalculatorAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorAverageResponse proto.InternalMessageInfo

func (m *CalculatorAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type CalculatorOddEvenResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorOddEvenResponse) Reset()         { *m = CalculatorOddEvenResponse{} }
func (m *CalculatorOddEvenResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorOddEvenResponse) ProtoMessage()    {}
func (*CalculatorOddEvenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_calculator_0c0b5626957dde3f, []int{3}
}
func (m *CalculatorOddEvenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorOddEvenResponse.Unmarshal(m, b)
}
func (m *CalculatorOddEvenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorOddEvenResponse.Marshal(b, m, deterministic)
}
func (dst *CalculatorOddEvenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorOddEvenResponse.Merge(dst, src)
}
func (m *CalculatorOddEvenResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorOddEvenResponse.Size(m)
}
func (m *CalculatorOddEvenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorOddEvenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorOddEvenResponse proto.InternalMessageInfo

func (m *CalculatorOddEvenResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *CalculatorOddEvenResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
	proto.RegisterType((*CalculatorAverageResponse)(nil), "calculator.CalculatorAverageResponse")
	proto.RegisterType((*CalculatorOddEvenResponse)(nil), "calculator.CalculatorOddEvenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	// Unary
	Square(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	// Server Streaming
	PrimeFactors(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorsClient, error)
	// Client Streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error)
	// BiDirectional Streaming
	OddEven(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_OddEvenClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Square(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Square", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) PrimeFactors(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (CalculatorService_PrimeFactorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/PrimeFactors", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeFactorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeFactorsClient interface {
	Recv() (*CalculatorResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeFactorsClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeFactorsClient) Recv() (*CalculatorResponse, error) {
	m := new(CalculatorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceAverageClient{stream}
	return x, nil
}

type CalculatorService_AverageClient interface {
	Send(*CalculatorRequest) error
	CloseAndRecv() (*CalculatorAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceAverageClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceAverageClient) CloseAndRecv() (*CalculatorAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CalculatorAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) OddEven(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_OddEvenClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/OddEven", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceOddEvenClient{stream}
	return x, nil
}

type CalculatorService_OddEvenClient interface {
	Send(*CalculatorRequest) error
	Recv() (*CalculatorOddEvenResponse, error)
	grpc.ClientStream
}

type calculatorServiceOddEvenClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceOddEvenClient) Send(m *CalculatorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceOddEvenClient) Recv() (*CalculatorOddEvenResponse, error) {
	m := new(CalculatorOddEvenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	// Unary
	Square(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	// Server Streaming
	PrimeFactors(*CalculatorRequest, CalculatorService_PrimeFactorsServer) error
	// Client Streaming
	Average(CalculatorService_AverageServer) error
	// BiDirectional Streaming
	OddEven(CalculatorService_OddEvenServer) error
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Square_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Square(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Square",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Square(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_PrimeFactors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalculatorRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).PrimeFactors(m, &calculatorServicePrimeFactorsServer{stream})
}

type CalculatorService_PrimeFactorsServer interface {
	Send(*CalculatorResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeFactorsServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeFactorsServer) Send(m *CalculatorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).Average(&calculatorServiceAverageServer{stream})
}

type CalculatorService_AverageServer interface {
	SendAndClose(*CalculatorAverageResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceAverageServer) SendAndClose(m *CalculatorAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceAverageServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_OddEven_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).OddEven(&calculatorServiceOddEvenServer{stream})
}

type CalculatorService_OddEvenServer interface {
	Send(*CalculatorOddEvenResponse) error
	Recv() (*CalculatorRequest, error)
	grpc.ServerStream
}

type calculatorServiceOddEvenServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceOddEvenServer) Send(m *CalculatorOddEvenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceOddEvenServer) Recv() (*CalculatorRequest, error) {
	m := new(CalculatorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Square",
			Handler:    _CalculatorService_Square_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeFactors",
			Handler:       _CalculatorService_PrimeFactors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalculatorService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "OddEven",
			Handler:       _CalculatorService_OddEven_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/calculator.proto",
}

func init() { proto.RegisterFile("proto/calculator.proto", fileDescriptor_calculator_0c0b5626957dde3f) }

var fileDescriptor_calculator_0c0b5626957dde3f = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x03, 0x0b, 0x08,
	0x71, 0x21, 0x44, 0x94, 0xb4, 0xb9, 0x04, 0x9d, 0xe1, 0xbc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x31, 0x2e, 0xb6, 0xbc, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xd6, 0x20, 0x28, 0x4f, 0x49, 0x87, 0x4b, 0x08, 0x59, 0x71, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a,
	0x48, 0x75, 0x51, 0x6a, 0x71, 0x69, 0x4e, 0x09, 0x4c, 0x35, 0x84, 0xa7, 0x64, 0xcc, 0x25, 0x89,
	0x50, 0xed, 0x58, 0x96, 0x5a, 0x94, 0x98, 0x9e, 0x8a, 0x43, 0x13, 0x13, 0x5c, 0x93, 0x3b, 0xb2,
	0x26, 0xff, 0x94, 0x14, 0xd7, 0xb2, 0xd4, 0x3c, 0x42, 0x36, 0x09, 0x09, 0x71, 0xb1, 0x94, 0x54,
	0x16, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x46, 0x77, 0x99, 0x90, 0x7d,
	0x16, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xc9, 0xc5, 0x16, 0x5c, 0x58, 0x9a, 0x58,
	0x94, 0x2a, 0x24, 0xab, 0x87, 0x14, 0x2e, 0x18, 0x41, 0x20, 0x25, 0x87, 0x4b, 0x1a, 0xe2, 0x14,
	0x25, 0x06, 0xa1, 0x40, 0x2e, 0x9e, 0x80, 0xa2, 0xcc, 0xdc, 0x54, 0xb7, 0xc4, 0xe4, 0x92, 0xfc,
	0xa2, 0x62, 0x8a, 0x0d, 0x34, 0x60, 0x14, 0x0a, 0xe6, 0x62, 0x87, 0x86, 0x13, 0x21, 0xd3, 0x54,
	0xb1, 0x4b, 0xa3, 0x85, 0xb2, 0x12, 0x83, 0x06, 0xa3, 0x50, 0x28, 0x17, 0x3b, 0x34, 0x1c, 0xc9,
	0x34, 0x14, 0x2d, 0x16, 0x40, 0x86, 0x1a, 0x30, 0x3a, 0xf1, 0x45, 0xf1, 0x20, 0x54, 0x17, 0x24,
	0x25, 0xb1, 0x81, 0xd3, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x13, 0x65, 0x47, 0x62, 0x75,
	0x02, 0x00, 0x00,
}
