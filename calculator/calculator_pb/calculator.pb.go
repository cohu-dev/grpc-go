// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: calculator/calculator_pb/calculator.proto

package calculator_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberOne int32 `protobuf:"varint,1,opt,name=number_one,json=numberOne,proto3" json:"number_one,omitempty"`
	NumberTwo int32 `protobuf:"varint,2,opt,name=number_two,json=numberTwo,proto3" json:"number_two,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_pb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_pb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_pb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculatorRequest) GetNumberOne() int32 {
	if x != nil {
		return x.NumberOne
	}
	return 0
}

func (x *CalculatorRequest) GetNumberTwo() int32 {
	if x != nil {
		return x.NumberTwo
	}
	return 0
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_pb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_pb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_pb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calculator_pb_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculator_pb_calculator_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x51, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x77, 0x6f, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5b, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a,
	0x03, 0x53, 0x75, 0x6d, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculator_calculator_pb_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculator_pb_calculator_proto_rawDescData = file_calculator_calculator_pb_calculator_proto_rawDesc
)

func file_calculator_calculator_pb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculator_pb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculator_pb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculator_pb_calculator_proto_rawDescData)
	})
	return file_calculator_calculator_pb_calculator_proto_rawDescData
}

var file_calculator_calculator_pb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_calculator_pb_calculator_proto_goTypes = []interface{}{
	(*CalculatorRequest)(nil),  // 0: calculator.CalculatorRequest
	(*CalculatorResponse)(nil), // 1: calculator.CalculatorResponse
}
var file_calculator_calculator_pb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Sum:input_type -> calculator.CalculatorRequest
	1, // 1: calculator.CalculatorService.Sum:output_type -> calculator.CalculatorResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calculator_pb_calculator_proto_init() }
func file_calculator_calculator_pb_calculator_proto_init() {
	if File_calculator_calculator_pb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculator_pb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_pb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_calculator_pb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculator_pb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculator_pb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculator_pb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculator_pb_calculator_proto = out.File
	file_calculator_calculator_pb_calculator_proto_rawDesc = nil
	file_calculator_calculator_pb_calculator_proto_goTypes = nil
	file_calculator_calculator_pb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calculator/calculator_pb/calculator.proto",
}
