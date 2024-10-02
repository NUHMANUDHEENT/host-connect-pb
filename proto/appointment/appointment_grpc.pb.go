// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package appointment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AppointmentServiceClient is the client API for AppointmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppointmentServiceClient interface {
	// Check available doctors and their next available time slot
	CheckAvailability(ctx context.Context, in *GetAvailabilityRequest, opts ...grpc.CallOption) (*GetAvailabilityResponse, error)
	// Confirm an appointment for a specific doctor and time
	ConfirmAppointment(ctx context.Context, in *ConfirmAppointmentRequest, opts ...grpc.CallOption) (*ConfirmAppointmentResponse, error)
	// Complete payment for an appointment
	CompletePayment(ctx context.Context, in *CompletePaymentRequest, opts ...grpc.CallOption) (*CompletePaymentResponse, error)
}

type appointmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppointmentServiceClient(cc grpc.ClientConnInterface) AppointmentServiceClient {
	return &appointmentServiceClient{cc}
}

func (c *appointmentServiceClient) CheckAvailability(ctx context.Context, in *GetAvailabilityRequest, opts ...grpc.CallOption) (*GetAvailabilityResponse, error) {
	out := new(GetAvailabilityResponse)
	err := c.cc.Invoke(ctx, "/appointment.AppointmentService/CheckAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) ConfirmAppointment(ctx context.Context, in *ConfirmAppointmentRequest, opts ...grpc.CallOption) (*ConfirmAppointmentResponse, error) {
	out := new(ConfirmAppointmentResponse)
	err := c.cc.Invoke(ctx, "/appointment.AppointmentService/ConfirmAppointment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appointmentServiceClient) CompletePayment(ctx context.Context, in *CompletePaymentRequest, opts ...grpc.CallOption) (*CompletePaymentResponse, error) {
	out := new(CompletePaymentResponse)
	err := c.cc.Invoke(ctx, "/appointment.AppointmentService/CompletePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppointmentServiceServer is the server API for AppointmentService service.
// All implementations must embed UnimplementedAppointmentServiceServer
// for forward compatibility
type AppointmentServiceServer interface {
	// Check available doctors and their next available time slot
	CheckAvailability(context.Context, *GetAvailabilityRequest) (*GetAvailabilityResponse, error)
	// Confirm an appointment for a specific doctor and time
	ConfirmAppointment(context.Context, *ConfirmAppointmentRequest) (*ConfirmAppointmentResponse, error)
	// Complete payment for an appointment
	CompletePayment(context.Context, *CompletePaymentRequest) (*CompletePaymentResponse, error)
	mustEmbedUnimplementedAppointmentServiceServer()
}

// UnimplementedAppointmentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppointmentServiceServer struct {
}

func (UnimplementedAppointmentServiceServer) CheckAvailability(context.Context, *GetAvailabilityRequest) (*GetAvailabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckAvailability not implemented")
}
func (UnimplementedAppointmentServiceServer) ConfirmAppointment(context.Context, *ConfirmAppointmentRequest) (*ConfirmAppointmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmAppointment not implemented")
}
func (UnimplementedAppointmentServiceServer) CompletePayment(context.Context, *CompletePaymentRequest) (*CompletePaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompletePayment not implemented")
}
func (UnimplementedAppointmentServiceServer) mustEmbedUnimplementedAppointmentServiceServer() {}

// UnsafeAppointmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppointmentServiceServer will
// result in compilation errors.
type UnsafeAppointmentServiceServer interface {
	mustEmbedUnimplementedAppointmentServiceServer()
}

func RegisterAppointmentServiceServer(s *grpc.Server, srv AppointmentServiceServer) {
	s.RegisterService(&_AppointmentService_serviceDesc, srv)
}

func _AppointmentService_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.AppointmentService/CheckAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CheckAvailability(ctx, req.(*GetAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_ConfirmAppointment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmAppointmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).ConfirmAppointment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.AppointmentService/ConfirmAppointment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).ConfirmAppointment(ctx, req.(*ConfirmAppointmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppointmentService_CompletePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppointmentServiceServer).CompletePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appointment.AppointmentService/CompletePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppointmentServiceServer).CompletePayment(ctx, req.(*CompletePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppointmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appointment.AppointmentService",
	HandlerType: (*AppointmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckAvailability",
			Handler:    _AppointmentService_CheckAvailability_Handler,
		},
		{
			MethodName: "ConfirmAppointment",
			Handler:    _AppointmentService_ConfirmAppointment_Handler,
		},
		{
			MethodName: "CompletePayment",
			Handler:    _AppointmentService_CompletePayment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appointment/appointment.proto",
}
