// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/operations.proto

package investapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OperationsServiceClient is the client API for OperationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsServiceClient interface {
	// Метод получения списка операций по счёту.При работе с данным методом необходимо учитывать
	// [особенности взаимодействия](/investAPI/operations_problems) с данным методом.
	GetOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error)
	// Метод получения портфеля по счёту.
	GetPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error)
	// Метод получения списка позиций по счёту.
	GetPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error)
	// Метод получения доступного остатка для вывода средств.
	GetWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error)
	// Метод получения брокерского отчёта.
	GetBrokerReport(ctx context.Context, in *BrokerReportRequest, opts ...grpc.CallOption) (*BrokerReportResponse, error)
	// Метод получения отчёта "Справка о доходах за пределами РФ".
	GetDividendsForeignIssuer(ctx context.Context, in *GetDividendsForeignIssuerRequest, opts ...grpc.CallOption) (*GetDividendsForeignIssuerResponse, error)
	// Метод получения списка операций по счёту с пагинацией. При работе с данным методом необходимо учитывать
	// [особенности взаимодействия](/investAPI/operations_problems) с данным методом.
	GetOperationsByCursor(ctx context.Context, in *GetOperationsByCursorRequest, opts ...grpc.CallOption) (*GetOperationsByCursorResponse, error)
}

type operationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsServiceClient(cc grpc.ClientConnInterface) OperationsServiceClient {
	return &operationsServiceClient{cc}
}

func (c *operationsServiceClient) GetOperations(ctx context.Context, in *OperationsRequest, opts ...grpc.CallOption) (*OperationsResponse, error) {
	out := new(OperationsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetPortfolio(ctx context.Context, in *PortfolioRequest, opts ...grpc.CallOption) (*PortfolioResponse, error) {
	out := new(PortfolioResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetPositions(ctx context.Context, in *PositionsRequest, opts ...grpc.CallOption) (*PositionsResponse, error) {
	out := new(PositionsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetWithdrawLimits(ctx context.Context, in *WithdrawLimitsRequest, opts ...grpc.CallOption) (*WithdrawLimitsResponse, error) {
	out := new(WithdrawLimitsResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetBrokerReport(ctx context.Context, in *BrokerReportRequest, opts ...grpc.CallOption) (*BrokerReportResponse, error) {
	out := new(BrokerReportResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetBrokerReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetDividendsForeignIssuer(ctx context.Context, in *GetDividendsForeignIssuerRequest, opts ...grpc.CallOption) (*GetDividendsForeignIssuerResponse, error) {
	out := new(GetDividendsForeignIssuerResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetDividendsForeignIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsServiceClient) GetOperationsByCursor(ctx context.Context, in *GetOperationsByCursorRequest, opts ...grpc.CallOption) (*GetOperationsByCursorResponse, error) {
	out := new(GetOperationsByCursorResponse)
	err := c.cc.Invoke(ctx, "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperationsByCursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationsServiceServer is the server API for OperationsService service.
// All implementations must embed UnimplementedOperationsServiceServer
// for forward compatibility
type OperationsServiceServer interface {
	// Метод получения списка операций по счёту.При работе с данным методом необходимо учитывать
	// [особенности взаимодействия](/investAPI/operations_problems) с данным методом.
	GetOperations(context.Context, *OperationsRequest) (*OperationsResponse, error)
	// Метод получения портфеля по счёту.
	GetPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error)
	// Метод получения списка позиций по счёту.
	GetPositions(context.Context, *PositionsRequest) (*PositionsResponse, error)
	// Метод получения доступного остатка для вывода средств.
	GetWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error)
	// Метод получения брокерского отчёта.
	GetBrokerReport(context.Context, *BrokerReportRequest) (*BrokerReportResponse, error)
	// Метод получения отчёта "Справка о доходах за пределами РФ".
	GetDividendsForeignIssuer(context.Context, *GetDividendsForeignIssuerRequest) (*GetDividendsForeignIssuerResponse, error)
	// Метод получения списка операций по счёту с пагинацией. При работе с данным методом необходимо учитывать
	// [особенности взаимодействия](/investAPI/operations_problems) с данным методом.
	GetOperationsByCursor(context.Context, *GetOperationsByCursorRequest) (*GetOperationsByCursorResponse, error)
	mustEmbedUnimplementedOperationsServiceServer()
}

// UnimplementedOperationsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsServiceServer struct {
}

func (UnimplementedOperationsServiceServer) GetOperations(context.Context, *OperationsRequest) (*OperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedOperationsServiceServer) GetPortfolio(context.Context, *PortfolioRequest) (*PortfolioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortfolio not implemented")
}
func (UnimplementedOperationsServiceServer) GetPositions(context.Context, *PositionsRequest) (*PositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPositions not implemented")
}
func (UnimplementedOperationsServiceServer) GetWithdrawLimits(context.Context, *WithdrawLimitsRequest) (*WithdrawLimitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWithdrawLimits not implemented")
}
func (UnimplementedOperationsServiceServer) GetBrokerReport(context.Context, *BrokerReportRequest) (*BrokerReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBrokerReport not implemented")
}
func (UnimplementedOperationsServiceServer) GetDividendsForeignIssuer(context.Context, *GetDividendsForeignIssuerRequest) (*GetDividendsForeignIssuerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDividendsForeignIssuer not implemented")
}
func (UnimplementedOperationsServiceServer) GetOperationsByCursor(context.Context, *GetOperationsByCursorRequest) (*GetOperationsByCursorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperationsByCursor not implemented")
}
func (UnimplementedOperationsServiceServer) mustEmbedUnimplementedOperationsServiceServer() {}

// UnsafeOperationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsServiceServer will
// result in compilation errors.
type UnsafeOperationsServiceServer interface {
	mustEmbedUnimplementedOperationsServiceServer()
}

func RegisterOperationsServiceServer(s grpc.ServiceRegistrar, srv OperationsServiceServer) {
	s.RegisterService(&OperationsService_ServiceDesc, srv)
}

func _OperationsService_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetOperations(ctx, req.(*OperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetPortfolio_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortfolioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetPortfolio(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetPortfolio(ctx, req.(*PortfolioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetPositions(ctx, req.(*PositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetWithdrawLimits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawLimitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetWithdrawLimits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetWithdrawLimits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetWithdrawLimits(ctx, req.(*WithdrawLimitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetBrokerReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrokerReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetBrokerReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetBrokerReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetBrokerReport(ctx, req.(*BrokerReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetDividendsForeignIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDividendsForeignIssuerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetDividendsForeignIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetDividendsForeignIssuer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetDividendsForeignIssuer(ctx, req.(*GetDividendsForeignIssuerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationsService_GetOperationsByCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsByCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServiceServer).GetOperationsByCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tinkoff.public.invest.api.contract.v1.OperationsService/GetOperationsByCursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServiceServer).GetOperationsByCursor(ctx, req.(*GetOperationsByCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationsService_ServiceDesc is the grpc.ServiceDesc for OperationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.OperationsService",
	HandlerType: (*OperationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOperations",
			Handler:    _OperationsService_GetOperations_Handler,
		},
		{
			MethodName: "GetPortfolio",
			Handler:    _OperationsService_GetPortfolio_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _OperationsService_GetPositions_Handler,
		},
		{
			MethodName: "GetWithdrawLimits",
			Handler:    _OperationsService_GetWithdrawLimits_Handler,
		},
		{
			MethodName: "GetBrokerReport",
			Handler:    _OperationsService_GetBrokerReport_Handler,
		},
		{
			MethodName: "GetDividendsForeignIssuer",
			Handler:    _OperationsService_GetDividendsForeignIssuer_Handler,
		},
		{
			MethodName: "GetOperationsByCursor",
			Handler:    _OperationsService_GetOperationsByCursor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/operations.proto",
}

// OperationsStreamServiceClient is the client API for OperationsStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationsStreamServiceClient interface {
	// Server-side stream обновлений портфеля
	PortfolioStream(ctx context.Context, in *PortfolioStreamRequest, opts ...grpc.CallOption) (OperationsStreamService_PortfolioStreamClient, error)
	// Server-side stream обновлений информации по изменению позиций портфеля
	PositionsStream(ctx context.Context, in *PositionsStreamRequest, opts ...grpc.CallOption) (OperationsStreamService_PositionsStreamClient, error)
}

type operationsStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationsStreamServiceClient(cc grpc.ClientConnInterface) OperationsStreamServiceClient {
	return &operationsStreamServiceClient{cc}
}

func (c *operationsStreamServiceClient) PortfolioStream(ctx context.Context, in *PortfolioStreamRequest, opts ...grpc.CallOption) (OperationsStreamService_PortfolioStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OperationsStreamService_ServiceDesc.Streams[0], "/tinkoff.public.invest.api.contract.v1.OperationsStreamService/PortfolioStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationsStreamServicePortfolioStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationsStreamService_PortfolioStreamClient interface {
	Recv() (*PortfolioStreamResponse, error)
	grpc.ClientStream
}

type operationsStreamServicePortfolioStreamClient struct {
	grpc.ClientStream
}

func (x *operationsStreamServicePortfolioStreamClient) Recv() (*PortfolioStreamResponse, error) {
	m := new(PortfolioStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *operationsStreamServiceClient) PositionsStream(ctx context.Context, in *PositionsStreamRequest, opts ...grpc.CallOption) (OperationsStreamService_PositionsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &OperationsStreamService_ServiceDesc.Streams[1], "/tinkoff.public.invest.api.contract.v1.OperationsStreamService/PositionsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &operationsStreamServicePositionsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OperationsStreamService_PositionsStreamClient interface {
	Recv() (*PositionsStreamResponse, error)
	grpc.ClientStream
}

type operationsStreamServicePositionsStreamClient struct {
	grpc.ClientStream
}

func (x *operationsStreamServicePositionsStreamClient) Recv() (*PositionsStreamResponse, error) {
	m := new(PositionsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OperationsStreamServiceServer is the server API for OperationsStreamService service.
// All implementations must embed UnimplementedOperationsStreamServiceServer
// for forward compatibility
type OperationsStreamServiceServer interface {
	// Server-side stream обновлений портфеля
	PortfolioStream(*PortfolioStreamRequest, OperationsStreamService_PortfolioStreamServer) error
	// Server-side stream обновлений информации по изменению позиций портфеля
	PositionsStream(*PositionsStreamRequest, OperationsStreamService_PositionsStreamServer) error
	mustEmbedUnimplementedOperationsStreamServiceServer()
}

// UnimplementedOperationsStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperationsStreamServiceServer struct {
}

func (UnimplementedOperationsStreamServiceServer) PortfolioStream(*PortfolioStreamRequest, OperationsStreamService_PortfolioStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PortfolioStream not implemented")
}
func (UnimplementedOperationsStreamServiceServer) PositionsStream(*PositionsStreamRequest, OperationsStreamService_PositionsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method PositionsStream not implemented")
}
func (UnimplementedOperationsStreamServiceServer) mustEmbedUnimplementedOperationsStreamServiceServer() {
}

// UnsafeOperationsStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationsStreamServiceServer will
// result in compilation errors.
type UnsafeOperationsStreamServiceServer interface {
	mustEmbedUnimplementedOperationsStreamServiceServer()
}

func RegisterOperationsStreamServiceServer(s grpc.ServiceRegistrar, srv OperationsStreamServiceServer) {
	s.RegisterService(&OperationsStreamService_ServiceDesc, srv)
}

func _OperationsStreamService_PortfolioStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PortfolioStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationsStreamServiceServer).PortfolioStream(m, &operationsStreamServicePortfolioStreamServer{stream})
}

type OperationsStreamService_PortfolioStreamServer interface {
	Send(*PortfolioStreamResponse) error
	grpc.ServerStream
}

type operationsStreamServicePortfolioStreamServer struct {
	grpc.ServerStream
}

func (x *operationsStreamServicePortfolioStreamServer) Send(m *PortfolioStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _OperationsStreamService_PositionsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PositionsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OperationsStreamServiceServer).PositionsStream(m, &operationsStreamServicePositionsStreamServer{stream})
}

type OperationsStreamService_PositionsStreamServer interface {
	Send(*PositionsStreamResponse) error
	grpc.ServerStream
}

type operationsStreamServicePositionsStreamServer struct {
	grpc.ServerStream
}

func (x *operationsStreamServicePositionsStreamServer) Send(m *PositionsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// OperationsStreamService_ServiceDesc is the grpc.ServiceDesc for OperationsStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationsStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinkoff.public.invest.api.contract.v1.OperationsStreamService",
	HandlerType: (*OperationsStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PortfolioStream",
			Handler:       _OperationsStreamService_PortfolioStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PositionsStream",
			Handler:       _OperationsStreamService_PositionsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/operations.proto",
}
