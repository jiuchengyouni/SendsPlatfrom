// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: yearBill.proto

package yearBill

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	YearBillService_InfoCheck_FullMethodName      = "/YearBillService/InfoCheck"
	YearBillService_GetCertificate_FullMethodName = "/YearBillService/GetCertificate"
	YearBillService_YearBillPing_FullMethodName   = "/YearBillService/YearBillPing"
	YearBillService_CheckState_FullMethodName     = "/YearBillService/CheckState"
	YearBillService_Appraise_FullMethodName       = "/YearBillService/Appraise"
	YearBillService_DataStorage_FullMethodName    = "/YearBillService/DataStorage"
	YearBillService_GetLearnData_FullMethodName   = "/YearBillService/GetLearnData"
	YearBillService_GetPayData_FullMethodName     = "/YearBillService/GetPayData"
	YearBillService_DataInit_FullMethodName       = "/YearBillService/DataInit"
	YearBillService_GetRank_FullMethodName        = "/YearBillService/GetRank"
)

// YearBillServiceClient is the client API for YearBillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YearBillServiceClient interface {
	InfoCheck(ctx context.Context, in *InfoCheckRequest, opts ...grpc.CallOption) (*InfoCheckResponse, error)
	GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error)
	YearBillPing(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*YearBillPingResponse, error)
	// rpc PayDataInit(PayDataInitRequest)returns(PayDataInitResponse);
	// rpc BookDataInit(BookDataInitRequest)returns(BookDataInitResponse);
	// rpc LearnDataInit(LearnDataInitRequest)returns(LearnDataInitResponse);
	CheckState(ctx context.Context, in *CheckStateRequest, opts ...grpc.CallOption) (*CheckStateResponse, error)
	Appraise(ctx context.Context, in *AppraiseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DataStorage(ctx context.Context, in *DataStorageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetLearnData(ctx context.Context, in *GetLearnDataRequest, opts ...grpc.CallOption) (*GetLearnDataResponse, error)
	GetPayData(ctx context.Context, in *GetPayDataRequest, opts ...grpc.CallOption) (*GetPayDataResponse, error)
	DataInit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRank(ctx context.Context, in *GetRankRequest, opts ...grpc.CallOption) (*GetRankResponse, error)
}

type yearBillServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYearBillServiceClient(cc grpc.ClientConnInterface) YearBillServiceClient {
	return &yearBillServiceClient{cc}
}

func (c *yearBillServiceClient) InfoCheck(ctx context.Context, in *InfoCheckRequest, opts ...grpc.CallOption) (*InfoCheckResponse, error) {
	out := new(InfoCheckResponse)
	err := c.cc.Invoke(ctx, YearBillService_InfoCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) GetCertificate(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*GetCertificateResponse, error) {
	out := new(GetCertificateResponse)
	err := c.cc.Invoke(ctx, YearBillService_GetCertificate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) YearBillPing(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*YearBillPingResponse, error) {
	out := new(YearBillPingResponse)
	err := c.cc.Invoke(ctx, YearBillService_YearBillPing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) CheckState(ctx context.Context, in *CheckStateRequest, opts ...grpc.CallOption) (*CheckStateResponse, error) {
	out := new(CheckStateResponse)
	err := c.cc.Invoke(ctx, YearBillService_CheckState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) Appraise(ctx context.Context, in *AppraiseRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, YearBillService_Appraise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) DataStorage(ctx context.Context, in *DataStorageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, YearBillService_DataStorage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) GetLearnData(ctx context.Context, in *GetLearnDataRequest, opts ...grpc.CallOption) (*GetLearnDataResponse, error) {
	out := new(GetLearnDataResponse)
	err := c.cc.Invoke(ctx, YearBillService_GetLearnData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) GetPayData(ctx context.Context, in *GetPayDataRequest, opts ...grpc.CallOption) (*GetPayDataResponse, error) {
	out := new(GetPayDataResponse)
	err := c.cc.Invoke(ctx, YearBillService_GetPayData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) DataInit(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, YearBillService_DataInit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearBillServiceClient) GetRank(ctx context.Context, in *GetRankRequest, opts ...grpc.CallOption) (*GetRankResponse, error) {
	out := new(GetRankResponse)
	err := c.cc.Invoke(ctx, YearBillService_GetRank_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YearBillServiceServer is the server API for YearBillService service.
// All implementations must embed UnimplementedYearBillServiceServer
// for forward compatibility
type YearBillServiceServer interface {
	InfoCheck(context.Context, *InfoCheckRequest) (*InfoCheckResponse, error)
	GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error)
	YearBillPing(context.Context, *emptypb.Empty) (*YearBillPingResponse, error)
	// rpc PayDataInit(PayDataInitRequest)returns(PayDataInitResponse);
	// rpc BookDataInit(BookDataInitRequest)returns(BookDataInitResponse);
	// rpc LearnDataInit(LearnDataInitRequest)returns(LearnDataInitResponse);
	CheckState(context.Context, *CheckStateRequest) (*CheckStateResponse, error)
	Appraise(context.Context, *AppraiseRequest) (*emptypb.Empty, error)
	DataStorage(context.Context, *DataStorageRequest) (*emptypb.Empty, error)
	GetLearnData(context.Context, *GetLearnDataRequest) (*GetLearnDataResponse, error)
	GetPayData(context.Context, *GetPayDataRequest) (*GetPayDataResponse, error)
	DataInit(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetRank(context.Context, *GetRankRequest) (*GetRankResponse, error)
	mustEmbedUnimplementedYearBillServiceServer()
}

// UnimplementedYearBillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYearBillServiceServer struct {
}

func (UnimplementedYearBillServiceServer) InfoCheck(context.Context, *InfoCheckRequest) (*InfoCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoCheck not implemented")
}
func (UnimplementedYearBillServiceServer) GetCertificate(context.Context, *GetCertificateRequest) (*GetCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCertificate not implemented")
}
func (UnimplementedYearBillServiceServer) YearBillPing(context.Context, *emptypb.Empty) (*YearBillPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method YearBillPing not implemented")
}
func (UnimplementedYearBillServiceServer) CheckState(context.Context, *CheckStateRequest) (*CheckStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckState not implemented")
}
func (UnimplementedYearBillServiceServer) Appraise(context.Context, *AppraiseRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Appraise not implemented")
}
func (UnimplementedYearBillServiceServer) DataStorage(context.Context, *DataStorageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataStorage not implemented")
}
func (UnimplementedYearBillServiceServer) GetLearnData(context.Context, *GetLearnDataRequest) (*GetLearnDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLearnData not implemented")
}
func (UnimplementedYearBillServiceServer) GetPayData(context.Context, *GetPayDataRequest) (*GetPayDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayData not implemented")
}
func (UnimplementedYearBillServiceServer) DataInit(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataInit not implemented")
}
func (UnimplementedYearBillServiceServer) GetRank(context.Context, *GetRankRequest) (*GetRankResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRank not implemented")
}
func (UnimplementedYearBillServiceServer) mustEmbedUnimplementedYearBillServiceServer() {}

// UnsafeYearBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YearBillServiceServer will
// result in compilation errors.
type UnsafeYearBillServiceServer interface {
	mustEmbedUnimplementedYearBillServiceServer()
}

func RegisterYearBillServiceServer(s grpc.ServiceRegistrar, srv YearBillServiceServer) {
	s.RegisterService(&YearBillService_ServiceDesc, srv)
}

func _YearBillService_InfoCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).InfoCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_InfoCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).InfoCheck(ctx, req.(*InfoCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_GetCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).GetCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_GetCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).GetCertificate(ctx, req.(*GetCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_YearBillPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).YearBillPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_YearBillPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).YearBillPing(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_CheckState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).CheckState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_CheckState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).CheckState(ctx, req.(*CheckStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_Appraise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppraiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).Appraise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_Appraise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).Appraise(ctx, req.(*AppraiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_DataStorage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataStorageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).DataStorage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_DataStorage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).DataStorage(ctx, req.(*DataStorageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_GetLearnData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLearnDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).GetLearnData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_GetLearnData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).GetLearnData(ctx, req.(*GetLearnDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_GetPayData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).GetPayData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_GetPayData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).GetPayData(ctx, req.(*GetPayDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_DataInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).DataInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_DataInit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).DataInit(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _YearBillService_GetRank_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRankRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YearBillServiceServer).GetRank(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: YearBillService_GetRank_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YearBillServiceServer).GetRank(ctx, req.(*GetRankRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YearBillService_ServiceDesc is the grpc.ServiceDesc for YearBillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YearBillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "YearBillService",
	HandlerType: (*YearBillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InfoCheck",
			Handler:    _YearBillService_InfoCheck_Handler,
		},
		{
			MethodName: "GetCertificate",
			Handler:    _YearBillService_GetCertificate_Handler,
		},
		{
			MethodName: "YearBillPing",
			Handler:    _YearBillService_YearBillPing_Handler,
		},
		{
			MethodName: "CheckState",
			Handler:    _YearBillService_CheckState_Handler,
		},
		{
			MethodName: "Appraise",
			Handler:    _YearBillService_Appraise_Handler,
		},
		{
			MethodName: "DataStorage",
			Handler:    _YearBillService_DataStorage_Handler,
		},
		{
			MethodName: "GetLearnData",
			Handler:    _YearBillService_GetLearnData_Handler,
		},
		{
			MethodName: "GetPayData",
			Handler:    _YearBillService_GetPayData_Handler,
		},
		{
			MethodName: "DataInit",
			Handler:    _YearBillService_DataInit_Handler,
		},
		{
			MethodName: "GetRank",
			Handler:    _YearBillService_GetRank_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yearBill.proto",
}
