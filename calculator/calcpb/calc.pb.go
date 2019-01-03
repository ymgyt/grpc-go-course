// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type CalcRequest struct {
	Arg1                 int64    `protobuf:"varint,1,opt,name=Arg1,proto3" json:"Arg1,omitempty"`
	Arg2                 int64    `protobuf:"varint,2,opt,name=Arg2,proto3" json:"Arg2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{0}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetArg1() int64 {
	if m != nil {
		return m.Arg1
	}
	return 0
}

func (m *CalcRequest) GetArg2() int64 {
	if m != nil {
		return m.Arg2
	}
	return 0
}

type CalcResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{1}
}

func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (m *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(m, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeNumberDecomposeRequest struct {
	PrimeNumber          int64    `protobuf:"varint,1,opt,name=PrimeNumber,proto3" json:"PrimeNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecomposeRequest) Reset()         { *m = PrimeNumberDecomposeRequest{} }
func (m *PrimeNumberDecomposeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposeRequest) ProtoMessage()    {}
func (*PrimeNumberDecomposeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{2}
}

func (m *PrimeNumberDecomposeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Unmarshal(m, b)
}
func (m *PrimeNumberDecomposeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecomposeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecomposeRequest.Merge(m, src)
}
func (m *PrimeNumberDecomposeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecomposeRequest.Size(m)
}
func (m *PrimeNumberDecomposeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecomposeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecomposeRequest proto.InternalMessageInfo

func (m *PrimeNumberDecomposeRequest) GetPrimeNumber() int64 {
	if m != nil {
		return m.PrimeNumber
	}
	return 0
}

type PrimeNumberDecomposeResponse struct {
	Factor               int64    `protobuf:"varint,1,opt,name=Factor,proto3" json:"Factor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeNumberDecomposeResponse) Reset()         { *m = PrimeNumberDecomposeResponse{} }
func (m *PrimeNumberDecomposeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeNumberDecomposeResponse) ProtoMessage()    {}
func (*PrimeNumberDecomposeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{3}
}

func (m *PrimeNumberDecomposeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Unmarshal(m, b)
}
func (m *PrimeNumberDecomposeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeNumberDecomposeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeNumberDecomposeResponse.Merge(m, src)
}
func (m *PrimeNumberDecomposeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeNumberDecomposeResponse.Size(m)
}
func (m *PrimeNumberDecomposeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeNumberDecomposeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeNumberDecomposeResponse proto.InternalMessageInfo

func (m *PrimeNumberDecomposeResponse) GetFactor() int64 {
	if m != nil {
		return m.Factor
	}
	return 0
}

type AverageRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageRequest) Reset()         { *m = AverageRequest{} }
func (m *AverageRequest) String() string { return proto.CompactTextString(m) }
func (*AverageRequest) ProtoMessage()    {}
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{4}
}

func (m *AverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageRequest.Unmarshal(m, b)
}
func (m *AverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageRequest.Marshal(b, m, deterministic)
}
func (m *AverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageRequest.Merge(m, src)
}
func (m *AverageRequest) XXX_Size() int {
	return xxx_messageInfo_AverageRequest.Size(m)
}
func (m *AverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AverageRequest proto.InternalMessageInfo

func (m *AverageRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type AverageResponse struct {
	Average              float64  `protobuf:"fixed64,1,opt,name=Average,proto3" json:"Average,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AverageResponse) Reset()         { *m = AverageResponse{} }
func (m *AverageResponse) String() string { return proto.CompactTextString(m) }
func (*AverageResponse) ProtoMessage()    {}
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{5}
}

func (m *AverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AverageResponse.Unmarshal(m, b)
}
func (m *AverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AverageResponse.Marshal(b, m, deterministic)
}
func (m *AverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AverageResponse.Merge(m, src)
}
func (m *AverageResponse) XXX_Size() int {
	return xxx_messageInfo_AverageResponse.Size(m)
}
func (m *AverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AverageResponse proto.InternalMessageInfo

func (m *AverageResponse) GetAverage() float64 {
	if m != nil {
		return m.Average
	}
	return 0
}

type FindMaximumRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumRequest) Reset()         { *m = FindMaximumRequest{} }
func (m *FindMaximumRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaximumRequest) ProtoMessage()    {}
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{6}
}

func (m *FindMaximumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumRequest.Unmarshal(m, b)
}
func (m *FindMaximumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumRequest.Marshal(b, m, deterministic)
}
func (m *FindMaximumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumRequest.Merge(m, src)
}
func (m *FindMaximumRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaximumRequest.Size(m)
}
func (m *FindMaximumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumRequest proto.InternalMessageInfo

func (m *FindMaximumRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaximumResponse struct {
	Maximum              int64    `protobuf:"varint,1,opt,name=Maximum,proto3" json:"Maximum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaximumResponse) Reset()         { *m = FindMaximumResponse{} }
func (m *FindMaximumResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaximumResponse) ProtoMessage()    {}
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{7}
}

func (m *FindMaximumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaximumResponse.Unmarshal(m, b)
}
func (m *FindMaximumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaximumResponse.Marshal(b, m, deterministic)
}
func (m *FindMaximumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaximumResponse.Merge(m, src)
}
func (m *FindMaximumResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaximumResponse.Size(m)
}
func (m *FindMaximumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaximumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaximumResponse proto.InternalMessageInfo

func (m *FindMaximumResponse) GetMaximum() int64 {
	if m != nil {
		return m.Maximum
	}
	return 0
}

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{8}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquareRootResponse struct {
	NumberRoot           float64  `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootResponse) Reset()         { *m = SquareRootResponse{} }
func (m *SquareRootResponse) String() string { return proto.CompactTextString(m) }
func (*SquareRootResponse) ProtoMessage()    {}
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{9}
}

func (m *SquareRootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootResponse.Unmarshal(m, b)
}
func (m *SquareRootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootResponse.Marshal(b, m, deterministic)
}
func (m *SquareRootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootResponse.Merge(m, src)
}
func (m *SquareRootResponse) XXX_Size() int {
	return xxx_messageInfo_SquareRootResponse.Size(m)
}
func (m *SquareRootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootResponse proto.InternalMessageInfo

func (m *SquareRootResponse) GetNumberRoot() float64 {
	if m != nil {
		return m.NumberRoot
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "calc.CalcRequest")
	proto.RegisterType((*CalcResponse)(nil), "calc.CalcResponse")
	proto.RegisterType((*PrimeNumberDecomposeRequest)(nil), "calc.PrimeNumberDecomposeRequest")
	proto.RegisterType((*PrimeNumberDecomposeResponse)(nil), "calc.PrimeNumberDecomposeResponse")
	proto.RegisterType((*AverageRequest)(nil), "calc.AverageRequest")
	proto.RegisterType((*AverageResponse)(nil), "calc.AverageResponse")
	proto.RegisterType((*FindMaximumRequest)(nil), "calc.FindMaximumRequest")
	proto.RegisterType((*FindMaximumResponse)(nil), "calc.FindMaximumResponse")
	proto.RegisterType((*SquareRootRequest)(nil), "calc.SquareRootRequest")
	proto.RegisterType((*SquareRootResponse)(nil), "calc.SquareRootResponse")
}

func init() { proto.RegisterFile("calculator/calcpb/calc.proto", fileDescriptor_cfc74e58bc0fa04b) }

var fileDescriptor_cfc74e58bc0fa04b = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4f, 0x4f, 0xfa, 0x40,
	0x10, 0xfd, 0x95, 0x1f, 0x56, 0x33, 0x35, 0x1a, 0x46, 0xc4, 0x5a, 0x49, 0xc4, 0x3d, 0x18, 0x12,
	0x0c, 0x20, 0x06, 0x0f, 0x5e, 0x0c, 0x6a, 0x88, 0x17, 0x8d, 0x29, 0x37, 0x2f, 0x66, 0xa9, 0x1b,
	0xd2, 0xa4, 0x65, 0xcb, 0xb6, 0x25, 0x7e, 0x66, 0x3f, 0x85, 0x69, 0x77, 0xfb, 0x07, 0x21, 0x78,
	0xea, 0xbc, 0x37, 0x6f, 0xde, 0x34, 0xfb, 0x76, 0xa1, 0xe9, 0x50, 0xcf, 0x89, 0x3d, 0x1a, 0x71,
	0xd1, 0x4b, 0xca, 0x60, 0x9a, 0x7e, 0xba, 0x81, 0xe0, 0x11, 0xc7, 0x6a, 0x52, 0x93, 0x21, 0x18,
	0x8f, 0xd4, 0x73, 0x6c, 0xb6, 0x88, 0x59, 0x18, 0x21, 0x42, 0x75, 0x24, 0x66, 0xd7, 0xa6, 0xd6,
	0xd2, 0xda, 0xff, 0xed, 0xb4, 0x56, 0xdc, 0xc0, 0xac, 0xe4, 0xdc, 0x80, 0x5c, 0xc2, 0xbe, 0x1c,
	0x0b, 0x03, 0x3e, 0x0f, 0x19, 0x36, 0x40, 0xb7, 0x59, 0x18, 0x7b, 0x91, 0x9a, 0x54, 0x88, 0xdc,
	0xc3, 0xd9, 0x9b, 0x70, 0x7d, 0xf6, 0x1a, 0xfb, 0x53, 0x26, 0x9e, 0x98, 0xc3, 0xfd, 0x80, 0x87,
	0x2c, 0x5b, 0xd7, 0x02, 0xa3, 0xd4, 0x56, 0xb3, 0x65, 0x8a, 0xdc, 0x42, 0x73, 0xb3, 0x41, 0xb1,
	0x78, 0x4c, 0x9d, 0x88, 0x67, 0xc3, 0x0a, 0x91, 0x36, 0x1c, 0x8c, 0x96, 0x4c, 0xd0, 0x59, 0xbe,
	0xab, 0x01, 0xfa, 0xca, 0x1a, 0x85, 0x48, 0x07, 0x0e, 0x73, 0xa5, 0x32, 0x35, 0x61, 0x57, 0x51,
	0xa9, 0x56, 0xb3, 0x33, 0x48, 0xae, 0x00, 0xc7, 0xee, 0xfc, 0xf3, 0x85, 0x7e, 0xb9, 0x7e, 0xec,
	0xff, 0x65, 0xdd, 0x83, 0xa3, 0x15, 0x75, 0x61, 0xaf, 0x28, 0xa5, 0xcf, 0x20, 0xe9, 0x40, 0x6d,
	0xb2, 0x88, 0xa9, 0x60, 0x36, 0xe7, 0x51, 0xc9, 0x7d, 0x5e, 0xb8, 0xef, 0xd8, 0x0a, 0x91, 0x21,
	0x60, 0x59, 0xac, 0xcc, 0xcf, 0xc1, 0x90, 0xfd, 0x0f, 0xc1, 0x79, 0xa4, 0xfe, 0x1f, 0x24, 0x95,
	0x08, 0x07, 0xdf, 0x15, 0x19, 0xf9, 0x84, 0x89, 0xa5, 0xeb, 0x30, 0xec, 0x41, 0x35, 0x81, 0x58,
	0xeb, 0xa6, 0x97, 0xa3, 0x74, 0x1b, 0x2c, 0x2c, 0x53, 0xd2, 0x9f, 0xfc, 0x43, 0x0a, 0xf5, 0x4d,
	0x91, 0xe0, 0x85, 0x54, 0x6f, 0xc9, 0xdb, 0x22, 0xdb, 0x24, 0xd9, 0x82, 0xbe, 0x86, 0x77, 0x79,
	0x00, 0x58, 0x97, 0x23, 0xab, 0x61, 0x5a, 0xc7, 0xbf, 0xd8, 0x6c, 0xb6, 0xad, 0xe1, 0x33, 0x18,
	0xa5, 0x43, 0x47, 0x53, 0x2a, 0xd7, 0x53, 0xb3, 0x4e, 0x37, 0x74, 0x0a, 0x9f, 0xbe, 0x86, 0x23,
	0x80, 0xe2, 0x80, 0xf1, 0x44, 0xca, 0xd7, 0xf2, 0xb1, 0xcc, 0xf5, 0x46, 0x66, 0xf3, 0xb0, 0xf7,
	0xae, 0xcb, 0x97, 0x37, 0xd5, 0xd3, 0x57, 0x77, 0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xdf,
	0x02, 0xbc, 0x95, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposeRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecomposeClient, error)
	Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error)
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error)
	// error handling
	// this RPC will throw an exeception if the sent number is negative
	// The error being sent is of type INVALID_ARGUMENT
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calcServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalcServiceClient(cc *grpc.ClientConn) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeNumberDecompose(ctx context.Context, in *PrimeNumberDecomposeRequest, opts ...grpc.CallOption) (CalcService_PrimeNumberDecomposeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calc.CalcService/PrimeNumberDecompose", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeNumberDecomposeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeNumberDecomposeClient interface {
	Recv() (*PrimeNumberDecomposeResponse, error)
	grpc.ClientStream
}

type calcServicePrimeNumberDecomposeClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeNumberDecomposeClient) Recv() (*PrimeNumberDecomposeResponse, error) {
	m := new(PrimeNumberDecomposeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (CalcService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[1], "/calc.CalcService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceAverageClient{stream}
	return x, nil
}

type CalcService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calcServiceAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (CalcService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[2], "/calc.CalcService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceFindMaximumClient{stream}
	return x, nil
}

type CalcService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type calcServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *calcServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calc.CalcService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
	PrimeNumberDecompose(*PrimeNumberDecomposeRequest, CalcService_PrimeNumberDecomposeServer) error
	Average(CalcService_AverageServer) error
	FindMaximum(CalcService_FindMaximumServer) error
	// error handling
	// this RPC will throw an exeception if the sent number is negative
	// The error being sent is of type INVALID_ARGUMENT
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeNumberDecompose_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeNumberDecomposeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeNumberDecompose(m, &calcServicePrimeNumberDecomposeServer{stream})
}

type CalcService_PrimeNumberDecomposeServer interface {
	Send(*PrimeNumberDecomposeResponse) error
	grpc.ServerStream
}

type calcServicePrimeNumberDecomposeServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeNumberDecomposeServer) Send(m *PrimeNumberDecomposeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).Average(&calcServiceAverageServer{stream})
}

type CalcService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calcServiceAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).FindMaximum(&calcServiceFindMaximumServer{stream})
}

type CalcService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type calcServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *calcServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calc.CalcService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalcService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeNumberDecompose",
			Handler:       _CalcService_PrimeNumberDecompose_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _CalcService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _CalcService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}