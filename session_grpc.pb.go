// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: session.proto

package __

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

// SessionServiceClient is the client API for SessionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SessionServiceClient interface {
	// for user session
	SetUserSession(ctx context.Context, in *UserSessionPayload, opts ...grpc.CallOption) (*SetUserSessionResponse, error)
	DeleteUserSession(ctx context.Context, in *DeleteUserSessionPayload, opts ...grpc.CallOption) (*OkResponse, error)
	ClearUserAllSession(ctx context.Context, in *ClearUserAllSessionPayload, opts ...grpc.CallOption) (*OkResponse, error)
	GetUserAllSession(ctx context.Context, in *GetUserAllSessionPayload, opts ...grpc.CallOption) (*GetUserAllSessionResponse, error)
	GetUserSessionRefreshToken(ctx context.Context, in *GetUserSessionRefreshTokenPayload, opts ...grpc.CallOption) (*SetUserSessionResponse, error)
	// for roles
	SetRole(ctx context.Context, in *RoleObject, opts ...grpc.CallOption) (*OkResponse, error)
	GetRole(ctx context.Context, in *GetRoleParam, opts ...grpc.CallOption) (*GetRoleReponse, error)
	// path map
	SetServiceMap(ctx context.Context, opts ...grpc.CallOption) (SessionService_SetServiceMapClient, error)
	GetServiceMap(ctx context.Context, in *NoPayload, opts ...grpc.CallOption) (SessionService_GetServiceMapClient, error)
	// for access check
	HaveAccess(ctx context.Context, in *HaveAccessParam, opts ...grpc.CallOption) (*HaveAccessResponse, error)
	VerifyUserSession(ctx context.Context, in *VerifyUserSessionParams, opts ...grpc.CallOption) (*VerifyUserSessionResponse, error)
	// One Time Password and Temporary link
	SetOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error)
	GetOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error)
	GetAndExpireOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error)
	// Mailing
	SendMail(ctx context.Context, in *MailSendTwoFactorAuthOtpPayload, opts ...grpc.CallOption) (*OkResponse, error)
}

type sessionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSessionServiceClient(cc grpc.ClientConnInterface) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) SetUserSession(ctx context.Context, in *UserSessionPayload, opts ...grpc.CallOption) (*SetUserSessionResponse, error) {
	out := new(SetUserSessionResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/SetUserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) DeleteUserSession(ctx context.Context, in *DeleteUserSessionPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/DeleteUserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) ClearUserAllSession(ctx context.Context, in *ClearUserAllSessionPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/ClearUserAllSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetUserAllSession(ctx context.Context, in *GetUserAllSessionPayload, opts ...grpc.CallOption) (*GetUserAllSessionResponse, error) {
	out := new(GetUserAllSessionResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/GetUserAllSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetUserSessionRefreshToken(ctx context.Context, in *GetUserSessionRefreshTokenPayload, opts ...grpc.CallOption) (*SetUserSessionResponse, error) {
	out := new(SetUserSessionResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/GetUserSessionRefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) SetRole(ctx context.Context, in *RoleObject, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/SetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetRole(ctx context.Context, in *GetRoleParam, opts ...grpc.CallOption) (*GetRoleReponse, error) {
	out := new(GetRoleReponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) SetServiceMap(ctx context.Context, opts ...grpc.CallOption) (SessionService_SetServiceMapClient, error) {
	stream, err := c.cc.NewStream(ctx, &SessionService_ServiceDesc.Streams[0], "/user_session_service.SessionService/SetServiceMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionServiceSetServiceMapClient{stream}
	return x, nil
}

type SessionService_SetServiceMapClient interface {
	Send(*ServiceMapData) error
	CloseAndRecv() (*OkResponse, error)
	grpc.ClientStream
}

type sessionServiceSetServiceMapClient struct {
	grpc.ClientStream
}

func (x *sessionServiceSetServiceMapClient) Send(m *ServiceMapData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sessionServiceSetServiceMapClient) CloseAndRecv() (*OkResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(OkResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionServiceClient) GetServiceMap(ctx context.Context, in *NoPayload, opts ...grpc.CallOption) (SessionService_GetServiceMapClient, error) {
	stream, err := c.cc.NewStream(ctx, &SessionService_ServiceDesc.Streams[1], "/user_session_service.SessionService/GetServiceMap", opts...)
	if err != nil {
		return nil, err
	}
	x := &sessionServiceGetServiceMapClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SessionService_GetServiceMapClient interface {
	Recv() (*ServiceMapData, error)
	grpc.ClientStream
}

type sessionServiceGetServiceMapClient struct {
	grpc.ClientStream
}

func (x *sessionServiceGetServiceMapClient) Recv() (*ServiceMapData, error) {
	m := new(ServiceMapData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sessionServiceClient) HaveAccess(ctx context.Context, in *HaveAccessParam, opts ...grpc.CallOption) (*HaveAccessResponse, error) {
	out := new(HaveAccessResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/HaveAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) VerifyUserSession(ctx context.Context, in *VerifyUserSessionParams, opts ...grpc.CallOption) (*VerifyUserSessionResponse, error) {
	out := new(VerifyUserSessionResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/VerifyUserSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) SetOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/SetOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/GetOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) GetAndExpireOTP(ctx context.Context, in *OTPPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/GetAndExpireOTP", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) SendMail(ctx context.Context, in *MailSendTwoFactorAuthOtpPayload, opts ...grpc.CallOption) (*OkResponse, error) {
	out := new(OkResponse)
	err := c.cc.Invoke(ctx, "/user_session_service.SessionService/SendMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SessionServiceServer is the server API for SessionService service.
// All implementations must embed UnimplementedSessionServiceServer
// for forward compatibility
type SessionServiceServer interface {
	// for user session
	SetUserSession(context.Context, *UserSessionPayload) (*SetUserSessionResponse, error)
	DeleteUserSession(context.Context, *DeleteUserSessionPayload) (*OkResponse, error)
	ClearUserAllSession(context.Context, *ClearUserAllSessionPayload) (*OkResponse, error)
	GetUserAllSession(context.Context, *GetUserAllSessionPayload) (*GetUserAllSessionResponse, error)
	GetUserSessionRefreshToken(context.Context, *GetUserSessionRefreshTokenPayload) (*SetUserSessionResponse, error)
	// for roles
	SetRole(context.Context, *RoleObject) (*OkResponse, error)
	GetRole(context.Context, *GetRoleParam) (*GetRoleReponse, error)
	// path map
	SetServiceMap(SessionService_SetServiceMapServer) error
	GetServiceMap(*NoPayload, SessionService_GetServiceMapServer) error
	// for access check
	HaveAccess(context.Context, *HaveAccessParam) (*HaveAccessResponse, error)
	VerifyUserSession(context.Context, *VerifyUserSessionParams) (*VerifyUserSessionResponse, error)
	// One Time Password and Temporary link
	SetOTP(context.Context, *OTPPayload) (*OkResponse, error)
	GetOTP(context.Context, *OTPPayload) (*OkResponse, error)
	GetAndExpireOTP(context.Context, *OTPPayload) (*OkResponse, error)
	// Mailing
	SendMail(context.Context, *MailSendTwoFactorAuthOtpPayload) (*OkResponse, error)
	mustEmbedUnimplementedSessionServiceServer()
}

// UnimplementedSessionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSessionServiceServer struct {
}

func (UnimplementedSessionServiceServer) SetUserSession(context.Context, *UserSessionPayload) (*SetUserSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserSession not implemented")
}
func (UnimplementedSessionServiceServer) DeleteUserSession(context.Context, *DeleteUserSessionPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserSession not implemented")
}
func (UnimplementedSessionServiceServer) ClearUserAllSession(context.Context, *ClearUserAllSessionPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearUserAllSession not implemented")
}
func (UnimplementedSessionServiceServer) GetUserAllSession(context.Context, *GetUserAllSessionPayload) (*GetUserAllSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAllSession not implemented")
}
func (UnimplementedSessionServiceServer) GetUserSessionRefreshToken(context.Context, *GetUserSessionRefreshTokenPayload) (*SetUserSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSessionRefreshToken not implemented")
}
func (UnimplementedSessionServiceServer) SetRole(context.Context, *RoleObject) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRole not implemented")
}
func (UnimplementedSessionServiceServer) GetRole(context.Context, *GetRoleParam) (*GetRoleReponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedSessionServiceServer) SetServiceMap(SessionService_SetServiceMapServer) error {
	return status.Errorf(codes.Unimplemented, "method SetServiceMap not implemented")
}
func (UnimplementedSessionServiceServer) GetServiceMap(*NoPayload, SessionService_GetServiceMapServer) error {
	return status.Errorf(codes.Unimplemented, "method GetServiceMap not implemented")
}
func (UnimplementedSessionServiceServer) HaveAccess(context.Context, *HaveAccessParam) (*HaveAccessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HaveAccess not implemented")
}
func (UnimplementedSessionServiceServer) VerifyUserSession(context.Context, *VerifyUserSessionParams) (*VerifyUserSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyUserSession not implemented")
}
func (UnimplementedSessionServiceServer) SetOTP(context.Context, *OTPPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOTP not implemented")
}
func (UnimplementedSessionServiceServer) GetOTP(context.Context, *OTPPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOTP not implemented")
}
func (UnimplementedSessionServiceServer) GetAndExpireOTP(context.Context, *OTPPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAndExpireOTP not implemented")
}
func (UnimplementedSessionServiceServer) SendMail(context.Context, *MailSendTwoFactorAuthOtpPayload) (*OkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedSessionServiceServer) mustEmbedUnimplementedSessionServiceServer() {}

// UnsafeSessionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SessionServiceServer will
// result in compilation errors.
type UnsafeSessionServiceServer interface {
	mustEmbedUnimplementedSessionServiceServer()
}

func RegisterSessionServiceServer(s grpc.ServiceRegistrar, srv SessionServiceServer) {
	s.RegisterService(&SessionService_ServiceDesc, srv)
}

func _SessionService_SetUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSessionPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SetUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/SetUserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SetUserSession(ctx, req.(*UserSessionPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_DeleteUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserSessionPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).DeleteUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/DeleteUserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).DeleteUserSession(ctx, req.(*DeleteUserSessionPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_ClearUserAllSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearUserAllSessionPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).ClearUserAllSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/ClearUserAllSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).ClearUserAllSession(ctx, req.(*ClearUserAllSessionPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetUserAllSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAllSessionPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetUserAllSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/GetUserAllSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetUserAllSession(ctx, req.(*GetUserAllSessionPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetUserSessionRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSessionRefreshTokenPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetUserSessionRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/GetUserSessionRefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetUserSessionRefreshToken(ctx, req.(*GetUserSessionRefreshTokenPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_SetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/SetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SetRole(ctx, req.(*RoleObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetRole(ctx, req.(*GetRoleParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_SetServiceMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SessionServiceServer).SetServiceMap(&sessionServiceSetServiceMapServer{stream})
}

type SessionService_SetServiceMapServer interface {
	SendAndClose(*OkResponse) error
	Recv() (*ServiceMapData, error)
	grpc.ServerStream
}

type sessionServiceSetServiceMapServer struct {
	grpc.ServerStream
}

func (x *sessionServiceSetServiceMapServer) SendAndClose(m *OkResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sessionServiceSetServiceMapServer) Recv() (*ServiceMapData, error) {
	m := new(ServiceMapData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SessionService_GetServiceMap_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NoPayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SessionServiceServer).GetServiceMap(m, &sessionServiceGetServiceMapServer{stream})
}

type SessionService_GetServiceMapServer interface {
	Send(*ServiceMapData) error
	grpc.ServerStream
}

type sessionServiceGetServiceMapServer struct {
	grpc.ServerStream
}

func (x *sessionServiceGetServiceMapServer) Send(m *ServiceMapData) error {
	return x.ServerStream.SendMsg(m)
}

func _SessionService_HaveAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HaveAccessParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).HaveAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/HaveAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).HaveAccess(ctx, req.(*HaveAccessParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_VerifyUserSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyUserSessionParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).VerifyUserSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/VerifyUserSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).VerifyUserSession(ctx, req.(*VerifyUserSessionParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_SetOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SetOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/SetOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SetOTP(ctx, req.(*OTPPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/GetOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetOTP(ctx, req.(*OTPPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_GetAndExpireOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OTPPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).GetAndExpireOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/GetAndExpireOTP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).GetAndExpireOTP(ctx, req.(*OTPPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailSendTwoFactorAuthOtpPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).SendMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_session_service.SessionService/SendMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).SendMail(ctx, req.(*MailSendTwoFactorAuthOtpPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// SessionService_ServiceDesc is the grpc.ServiceDesc for SessionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SessionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_session_service.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUserSession",
			Handler:    _SessionService_SetUserSession_Handler,
		},
		{
			MethodName: "DeleteUserSession",
			Handler:    _SessionService_DeleteUserSession_Handler,
		},
		{
			MethodName: "ClearUserAllSession",
			Handler:    _SessionService_ClearUserAllSession_Handler,
		},
		{
			MethodName: "GetUserAllSession",
			Handler:    _SessionService_GetUserAllSession_Handler,
		},
		{
			MethodName: "GetUserSessionRefreshToken",
			Handler:    _SessionService_GetUserSessionRefreshToken_Handler,
		},
		{
			MethodName: "SetRole",
			Handler:    _SessionService_SetRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _SessionService_GetRole_Handler,
		},
		{
			MethodName: "HaveAccess",
			Handler:    _SessionService_HaveAccess_Handler,
		},
		{
			MethodName: "VerifyUserSession",
			Handler:    _SessionService_VerifyUserSession_Handler,
		},
		{
			MethodName: "SetOTP",
			Handler:    _SessionService_SetOTP_Handler,
		},
		{
			MethodName: "GetOTP",
			Handler:    _SessionService_GetOTP_Handler,
		},
		{
			MethodName: "GetAndExpireOTP",
			Handler:    _SessionService_GetAndExpireOTP_Handler,
		},
		{
			MethodName: "SendMail",
			Handler:    _SessionService_SendMail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetServiceMap",
			Handler:       _SessionService_SetServiceMap_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetServiceMap",
			Handler:       _SessionService_GetServiceMap_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "session.proto",
}
