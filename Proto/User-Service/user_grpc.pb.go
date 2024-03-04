// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: User-Service/user.proto

package User_Service

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	GetTokenResetPass(ctx context.Context, in *GetTokenResetPassRequest, opts ...grpc.CallOption) (*GetTokenResetPassResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...grpc.CallOption) (*UpdateAvatarResponse, error)
	NewInstructor(ctx context.Context, in *NewInstructorRequest, opts ...grpc.CallOption) (*Instructor, error)
	GetInstructor(ctx context.Context, in *GetInstructorInformationRequest, opts ...grpc.CallOption) (*GetInstructorInformationResponse, error)
	GetProfile(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetProfileResponse, error)
	GetProfileInstructor(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetProfileInstructorResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetTokenResetPass(ctx context.Context, in *GetTokenResetPassRequest, opts ...grpc.CallOption) (*GetTokenResetPassResponse, error) {
	out := new(GetTokenResetPassResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetTokenResetPass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateAvatar(ctx context.Context, in *UpdateAvatarRequest, opts ...grpc.CallOption) (*UpdateAvatarResponse, error) {
	out := new(UpdateAvatarResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) NewInstructor(ctx context.Context, in *NewInstructorRequest, opts ...grpc.CallOption) (*Instructor, error) {
	out := new(Instructor)
	err := c.cc.Invoke(ctx, "/user.UserService/NewInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetInstructor(ctx context.Context, in *GetInstructorInformationRequest, opts ...grpc.CallOption) (*GetInstructorInformationResponse, error) {
	out := new(GetInstructorInformationResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfile(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetProfileResponse, error) {
	out := new(GetProfileResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetProfileInstructor(ctx context.Context, in *GetUserInformationRequest, opts ...grpc.CallOption) (*GetProfileInstructorResponse, error) {
	out := new(GetProfileInstructorResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/GetProfileInstructor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	GetTokenResetPass(context.Context, *GetTokenResetPassRequest) (*GetTokenResetPassResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	UpdateAvatar(context.Context, *UpdateAvatarRequest) (*UpdateAvatarResponse, error)
	NewInstructor(context.Context, *NewInstructorRequest) (*Instructor, error)
	GetInstructor(context.Context, *GetInstructorInformationRequest) (*GetInstructorInformationResponse, error)
	GetProfile(context.Context, *GetUserInformationRequest) (*GetProfileResponse, error)
	GetProfileInstructor(context.Context, *GetUserInformationRequest) (*GetProfileInstructorResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) GetTokenResetPass(context.Context, *GetTokenResetPassRequest) (*GetTokenResetPassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenResetPass not implemented")
}
func (UnimplementedUserServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserServiceServer) UpdateAvatar(context.Context, *UpdateAvatarRequest) (*UpdateAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAvatar not implemented")
}
func (UnimplementedUserServiceServer) NewInstructor(context.Context, *NewInstructorRequest) (*Instructor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewInstructor not implemented")
}
func (UnimplementedUserServiceServer) GetInstructor(context.Context, *GetInstructorInformationRequest) (*GetInstructorInformationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstructor not implemented")
}
func (UnimplementedUserServiceServer) GetProfile(context.Context, *GetUserInformationRequest) (*GetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedUserServiceServer) GetProfileInstructor(context.Context, *GetUserInformationRequest) (*GetProfileInstructorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileInstructor not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetTokenResetPass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenResetPassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetTokenResetPass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetTokenResetPass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetTokenResetPass(ctx, req.(*GetTokenResetPassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateAvatar(ctx, req.(*UpdateAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_NewInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewInstructorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).NewInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/NewInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).NewInstructor(ctx, req.(*NewInstructorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstructorInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInstructor(ctx, req.(*GetInstructorInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfile(ctx, req.(*GetUserInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetProfileInstructor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetProfileInstructor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetProfileInstructor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetProfileInstructor(ctx, req.(*GetUserInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "GetTokenResetPass",
			Handler:    _UserService_GetTokenResetPass_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "UpdateAvatar",
			Handler:    _UserService_UpdateAvatar_Handler,
		},
		{
			MethodName: "NewInstructor",
			Handler:    _UserService_NewInstructor_Handler,
		},
		{
			MethodName: "GetInstructor",
			Handler:    _UserService_GetInstructor_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _UserService_GetProfile_Handler,
		},
		{
			MethodName: "GetProfileInstructor",
			Handler:    _UserService_GetProfileInstructor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "User-Service/user.proto",
}