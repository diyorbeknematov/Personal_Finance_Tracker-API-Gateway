// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: budgeting_service/reporting_and_notifications.proto

package budgeting

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ReportingNotificationService_GetSepending_FullMethodName         = "/reporting_notification.ReportingNotificationService/GetSepending"
	ReportingNotificationService_GetIncome_FullMethodName            = "/reporting_notification.ReportingNotificationService/GetIncome"
	ReportingNotificationService_GetBudgetPerformance_FullMethodName = "/reporting_notification.ReportingNotificationService/GetBudgetPerformance"
	ReportingNotificationService_GoalProgress_FullMethodName         = "/reporting_notification.ReportingNotificationService/GoalProgress"
	ReportingNotificationService_SendNotification_FullMethodName     = "/reporting_notification.ReportingNotificationService/SendNotification"
	ReportingNotificationService_GetNotificationList_FullMethodName  = "/reporting_notification.ReportingNotificationService/GetNotificationList"
	ReportingNotificationService_GetNotification_FullMethodName      = "/reporting_notification.ReportingNotificationService/GetNotification"
	ReportingNotificationService_UpdateNotification_FullMethodName   = "/reporting_notification.ReportingNotificationService/UpdateNotification"
	ReportingNotificationService_DeleteNotification_FullMethodName   = "/reporting_notification.ReportingNotificationService/DeleteNotification"
)

// ReportingNotificationServiceClient is the client API for ReportingNotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportingNotificationServiceClient interface {
	GetSepending(ctx context.Context, in *GetSependingReq, opts ...grpc.CallOption) (*GetSependingResp, error)
	GetIncome(ctx context.Context, in *GetIncomeReportReq, opts ...grpc.CallOption) (*GetIncomeReportResp, error)
	GetBudgetPerformance(ctx context.Context, in *GetBudgetPerformanceReq, opts ...grpc.CallOption) (*GetBudgetPerformanceResp, error)
	GoalProgress(ctx context.Context, in *GetGoalProgressReq, opts ...grpc.CallOption) (*GetGoalProgressResp, error)
	// Notification
	SendNotification(ctx context.Context, in *SendNotificationReq, opts ...grpc.CallOption) (*SendNotificationResp, error)
	GetNotificationList(ctx context.Context, in *GetNotificationsListReq, opts ...grpc.CallOption) (*GetNotificationsListResp, error)
	GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationResp, error)
	UpdateNotification(ctx context.Context, in *UpdateNotificationReq, opts ...grpc.CallOption) (*UpdateNotificationResp, error)
	DeleteNotification(ctx context.Context, in *DeleteNotificationReq, opts ...grpc.CallOption) (*DeleteNotificationResp, error)
}

type reportingNotificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportingNotificationServiceClient(cc grpc.ClientConnInterface) ReportingNotificationServiceClient {
	return &reportingNotificationServiceClient{cc}
}

func (c *reportingNotificationServiceClient) GetSepending(ctx context.Context, in *GetSependingReq, opts ...grpc.CallOption) (*GetSependingResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSependingResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GetSepending_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) GetIncome(ctx context.Context, in *GetIncomeReportReq, opts ...grpc.CallOption) (*GetIncomeReportResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetIncomeReportResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GetIncome_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) GetBudgetPerformance(ctx context.Context, in *GetBudgetPerformanceReq, opts ...grpc.CallOption) (*GetBudgetPerformanceResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBudgetPerformanceResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GetBudgetPerformance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) GoalProgress(ctx context.Context, in *GetGoalProgressReq, opts ...grpc.CallOption) (*GetGoalProgressResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGoalProgressResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GoalProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) SendNotification(ctx context.Context, in *SendNotificationReq, opts ...grpc.CallOption) (*SendNotificationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendNotificationResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_SendNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) GetNotificationList(ctx context.Context, in *GetNotificationsListReq, opts ...grpc.CallOption) (*GetNotificationsListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationsListResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GetNotificationList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) GetNotification(ctx context.Context, in *GetNotificationReq, opts ...grpc.CallOption) (*GetNotificationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotificationResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_GetNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) UpdateNotification(ctx context.Context, in *UpdateNotificationReq, opts ...grpc.CallOption) (*UpdateNotificationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNotificationResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_UpdateNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportingNotificationServiceClient) DeleteNotification(ctx context.Context, in *DeleteNotificationReq, opts ...grpc.CallOption) (*DeleteNotificationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteNotificationResp)
	err := c.cc.Invoke(ctx, ReportingNotificationService_DeleteNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportingNotificationServiceServer is the server API for ReportingNotificationService service.
// All implementations must embed UnimplementedReportingNotificationServiceServer
// for forward compatibility
type ReportingNotificationServiceServer interface {
	GetSepending(context.Context, *GetSependingReq) (*GetSependingResp, error)
	GetIncome(context.Context, *GetIncomeReportReq) (*GetIncomeReportResp, error)
	GetBudgetPerformance(context.Context, *GetBudgetPerformanceReq) (*GetBudgetPerformanceResp, error)
	GoalProgress(context.Context, *GetGoalProgressReq) (*GetGoalProgressResp, error)
	// Notification
	SendNotification(context.Context, *SendNotificationReq) (*SendNotificationResp, error)
	GetNotificationList(context.Context, *GetNotificationsListReq) (*GetNotificationsListResp, error)
	GetNotification(context.Context, *GetNotificationReq) (*GetNotificationResp, error)
	UpdateNotification(context.Context, *UpdateNotificationReq) (*UpdateNotificationResp, error)
	DeleteNotification(context.Context, *DeleteNotificationReq) (*DeleteNotificationResp, error)
	mustEmbedUnimplementedReportingNotificationServiceServer()
}

// UnimplementedReportingNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReportingNotificationServiceServer struct {
}

func (UnimplementedReportingNotificationServiceServer) GetSepending(context.Context, *GetSependingReq) (*GetSependingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSepending not implemented")
}
func (UnimplementedReportingNotificationServiceServer) GetIncome(context.Context, *GetIncomeReportReq) (*GetIncomeReportResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIncome not implemented")
}
func (UnimplementedReportingNotificationServiceServer) GetBudgetPerformance(context.Context, *GetBudgetPerformanceReq) (*GetBudgetPerformanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBudgetPerformance not implemented")
}
func (UnimplementedReportingNotificationServiceServer) GoalProgress(context.Context, *GetGoalProgressReq) (*GetGoalProgressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoalProgress not implemented")
}
func (UnimplementedReportingNotificationServiceServer) SendNotification(context.Context, *SendNotificationReq) (*SendNotificationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendNotification not implemented")
}
func (UnimplementedReportingNotificationServiceServer) GetNotificationList(context.Context, *GetNotificationsListReq) (*GetNotificationsListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotificationList not implemented")
}
func (UnimplementedReportingNotificationServiceServer) GetNotification(context.Context, *GetNotificationReq) (*GetNotificationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotification not implemented")
}
func (UnimplementedReportingNotificationServiceServer) UpdateNotification(context.Context, *UpdateNotificationReq) (*UpdateNotificationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotification not implemented")
}
func (UnimplementedReportingNotificationServiceServer) DeleteNotification(context.Context, *DeleteNotificationReq) (*DeleteNotificationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotification not implemented")
}
func (UnimplementedReportingNotificationServiceServer) mustEmbedUnimplementedReportingNotificationServiceServer() {
}

// UnsafeReportingNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportingNotificationServiceServer will
// result in compilation errors.
type UnsafeReportingNotificationServiceServer interface {
	mustEmbedUnimplementedReportingNotificationServiceServer()
}

func RegisterReportingNotificationServiceServer(s grpc.ServiceRegistrar, srv ReportingNotificationServiceServer) {
	s.RegisterService(&ReportingNotificationService_ServiceDesc, srv)
}

func _ReportingNotificationService_GetSepending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSependingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GetSepending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GetSepending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GetSepending(ctx, req.(*GetSependingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_GetIncome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIncomeReportReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GetIncome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GetIncome_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GetIncome(ctx, req.(*GetIncomeReportReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_GetBudgetPerformance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBudgetPerformanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GetBudgetPerformance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GetBudgetPerformance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GetBudgetPerformance(ctx, req.(*GetBudgetPerformanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_GoalProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoalProgressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GoalProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GoalProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GoalProgress(ctx, req.(*GetGoalProgressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_SendNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).SendNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_SendNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).SendNotification(ctx, req.(*SendNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_GetNotificationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationsListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GetNotificationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GetNotificationList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GetNotificationList(ctx, req.(*GetNotificationsListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_GetNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).GetNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_GetNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).GetNotification(ctx, req.(*GetNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_UpdateNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).UpdateNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_UpdateNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).UpdateNotification(ctx, req.(*UpdateNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportingNotificationService_DeleteNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNotificationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportingNotificationServiceServer).DeleteNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportingNotificationService_DeleteNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportingNotificationServiceServer).DeleteNotification(ctx, req.(*DeleteNotificationReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportingNotificationService_ServiceDesc is the grpc.ServiceDesc for ReportingNotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportingNotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reporting_notification.ReportingNotificationService",
	HandlerType: (*ReportingNotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSepending",
			Handler:    _ReportingNotificationService_GetSepending_Handler,
		},
		{
			MethodName: "GetIncome",
			Handler:    _ReportingNotificationService_GetIncome_Handler,
		},
		{
			MethodName: "GetBudgetPerformance",
			Handler:    _ReportingNotificationService_GetBudgetPerformance_Handler,
		},
		{
			MethodName: "GoalProgress",
			Handler:    _ReportingNotificationService_GoalProgress_Handler,
		},
		{
			MethodName: "SendNotification",
			Handler:    _ReportingNotificationService_SendNotification_Handler,
		},
		{
			MethodName: "GetNotificationList",
			Handler:    _ReportingNotificationService_GetNotificationList_Handler,
		},
		{
			MethodName: "GetNotification",
			Handler:    _ReportingNotificationService_GetNotification_Handler,
		},
		{
			MethodName: "UpdateNotification",
			Handler:    _ReportingNotificationService_UpdateNotification_Handler,
		},
		{
			MethodName: "DeleteNotification",
			Handler:    _ReportingNotificationService_DeleteNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "budgeting_service/reporting_and_notifications.proto",
}
