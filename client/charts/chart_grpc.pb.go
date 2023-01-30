// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/chart.proto

package charts

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

// ChartDataServiceClient is the client API for ChartDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChartDataServiceClient interface {
	GetLineChartData(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (ChartDataService_GetLineChartDataClient, error)
	GetBarChartData(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (ChartDataService_GetBarChartDataClient, error)
}

type chartDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChartDataServiceClient(cc grpc.ClientConnInterface) ChartDataServiceClient {
	return &chartDataServiceClient{cc}
}

func (c *chartDataServiceClient) GetLineChartData(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (ChartDataService_GetLineChartDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChartDataService_ServiceDesc.Streams[0], "/charts.ChartDataService/GetLineChartData", opts...)
	if err != nil {
		return nil, err
	}
	x := &chartDataServiceGetLineChartDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChartDataService_GetLineChartDataClient interface {
	Recv() (*ChartResponse, error)
	grpc.ClientStream
}

type chartDataServiceGetLineChartDataClient struct {
	grpc.ClientStream
}

func (x *chartDataServiceGetLineChartDataClient) Recv() (*ChartResponse, error) {
	m := new(ChartResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chartDataServiceClient) GetBarChartData(ctx context.Context, in *ChartRequest, opts ...grpc.CallOption) (ChartDataService_GetBarChartDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChartDataService_ServiceDesc.Streams[1], "/charts.ChartDataService/GetBarChartData", opts...)
	if err != nil {
		return nil, err
	}
	x := &chartDataServiceGetBarChartDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChartDataService_GetBarChartDataClient interface {
	Recv() (*ChartResponse, error)
	grpc.ClientStream
}

type chartDataServiceGetBarChartDataClient struct {
	grpc.ClientStream
}

func (x *chartDataServiceGetBarChartDataClient) Recv() (*ChartResponse, error) {
	m := new(ChartResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChartDataServiceServer is the server API for ChartDataService service.
// All implementations must embed UnimplementedChartDataServiceServer
// for forward compatibility
type ChartDataServiceServer interface {
	GetLineChartData(*ChartRequest, ChartDataService_GetLineChartDataServer) error
	GetBarChartData(*ChartRequest, ChartDataService_GetBarChartDataServer) error
	mustEmbedUnimplementedChartDataServiceServer()
}

// UnimplementedChartDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChartDataServiceServer struct {
}

func (UnimplementedChartDataServiceServer) GetLineChartData(*ChartRequest, ChartDataService_GetLineChartDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLineChartData not implemented")
}
func (UnimplementedChartDataServiceServer) GetBarChartData(*ChartRequest, ChartDataService_GetBarChartDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBarChartData not implemented")
}
func (UnimplementedChartDataServiceServer) mustEmbedUnimplementedChartDataServiceServer() {}

// UnsafeChartDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChartDataServiceServer will
// result in compilation errors.
type UnsafeChartDataServiceServer interface {
	mustEmbedUnimplementedChartDataServiceServer()
}

func RegisterChartDataServiceServer(s grpc.ServiceRegistrar, srv ChartDataServiceServer) {
	s.RegisterService(&ChartDataService_ServiceDesc, srv)
}

func _ChartDataService_GetLineChartData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChartDataServiceServer).GetLineChartData(m, &chartDataServiceGetLineChartDataServer{stream})
}

type ChartDataService_GetLineChartDataServer interface {
	Send(*ChartResponse) error
	grpc.ServerStream
}

type chartDataServiceGetLineChartDataServer struct {
	grpc.ServerStream
}

func (x *chartDataServiceGetLineChartDataServer) Send(m *ChartResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChartDataService_GetBarChartData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChartRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChartDataServiceServer).GetBarChartData(m, &chartDataServiceGetBarChartDataServer{stream})
}

type ChartDataService_GetBarChartDataServer interface {
	Send(*ChartResponse) error
	grpc.ServerStream
}

type chartDataServiceGetBarChartDataServer struct {
	grpc.ServerStream
}

func (x *chartDataServiceGetBarChartDataServer) Send(m *ChartResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ChartDataService_ServiceDesc is the grpc.ServiceDesc for ChartDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChartDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "charts.ChartDataService",
	HandlerType: (*ChartDataServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLineChartData",
			Handler:       _ChartDataService_GetLineChartData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetBarChartData",
			Handler:       _ChartDataService_GetBarChartData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/chart.proto",
}