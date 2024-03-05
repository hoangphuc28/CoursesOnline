// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: File-Service/file.proto

package file

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error)
	UploadAsset(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadAssetClient, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) UploadAvatar(ctx context.Context, in *UploadAvatarRequest, opts ...grpc.CallOption) (*UploadAvatarResponse, error) {
	out := new(UploadAvatarResponse)
	err := c.cc.Invoke(ctx, "/file.FileService/UploadAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UploadAsset(ctx context.Context, opts ...grpc.CallOption) (FileService_UploadAssetClient, error) {
	stream, err := c.cc.NewStream(ctx, &FileService_ServiceDesc.Streams[0], "/file.FileService/UploadAsset", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileServiceUploadAssetClient{stream}
	return x, nil
}

type FileService_UploadAssetClient interface {
	Send(*UploadAssetRequest) error
	Recv() (*UploadAssetResponse, error)
	grpc.ClientStream
}

type fileServiceUploadAssetClient struct {
	grpc.ClientStream
}

func (x *fileServiceUploadAssetClient) Send(m *UploadAssetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileServiceUploadAssetClient) Recv() (*UploadAssetResponse, error) {
	m := new(UploadAssetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResponse, error)
	UploadAsset(FileService_UploadAssetServer) error
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) UploadAvatar(context.Context, *UploadAvatarRequest) (*UploadAvatarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedFileServiceServer) UploadAsset(FileService_UploadAssetServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadAsset not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadAvatarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FileService/UploadAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UploadAvatar(ctx, req.(*UploadAvatarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UploadAsset_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileServiceServer).UploadAsset(&fileServiceUploadAssetServer{stream})
}

type FileService_UploadAssetServer interface {
	Send(*UploadAssetResponse) error
	Recv() (*UploadAssetRequest, error)
	grpc.ServerStream
}

type fileServiceUploadAssetServer struct {
	grpc.ServerStream
}

func (x *fileServiceUploadAssetServer) Send(m *UploadAssetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileServiceUploadAssetServer) Recv() (*UploadAssetRequest, error) {
	m := new(UploadAssetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadAvatar",
			Handler:    _FileService_UploadAvatar_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadAsset",
			Handler:       _FileService_UploadAsset_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "File-Service/file.proto",
}