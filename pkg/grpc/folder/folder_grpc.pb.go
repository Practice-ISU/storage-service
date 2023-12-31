// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: folder.proto

package folder

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

// FolderServiceClient is the client API for FolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FolderServiceClient interface {
	AddFolder(ctx context.Context, in *FolderAddDTO, opts ...grpc.CallOption) (*FolderResponse, error)
	DeleteFolder(ctx context.Context, in *FolderDeleteDTO, opts ...grpc.CallOption) (*Details, error)
	RenameFolder(ctx context.Context, in *FolderRenameDTO, opts ...grpc.CallOption) (*FolderResponse, error)
	GetFolder(ctx context.Context, in *FolderGetDTO, opts ...grpc.CallOption) (*FolderResponse, error)
	GetAllFolders(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AllFoldersResponse, error)
}

type folderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFolderServiceClient(cc grpc.ClientConnInterface) FolderServiceClient {
	return &folderServiceClient{cc}
}

func (c *folderServiceClient) AddFolder(ctx context.Context, in *FolderAddDTO, opts ...grpc.CallOption) (*FolderResponse, error) {
	out := new(FolderResponse)
	err := c.cc.Invoke(ctx, "/folderservice.FolderService/AddFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) DeleteFolder(ctx context.Context, in *FolderDeleteDTO, opts ...grpc.CallOption) (*Details, error) {
	out := new(Details)
	err := c.cc.Invoke(ctx, "/folderservice.FolderService/DeleteFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) RenameFolder(ctx context.Context, in *FolderRenameDTO, opts ...grpc.CallOption) (*FolderResponse, error) {
	out := new(FolderResponse)
	err := c.cc.Invoke(ctx, "/folderservice.FolderService/RenameFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) GetFolder(ctx context.Context, in *FolderGetDTO, opts ...grpc.CallOption) (*FolderResponse, error) {
	out := new(FolderResponse)
	err := c.cc.Invoke(ctx, "/folderservice.FolderService/GetFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) GetAllFolders(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AllFoldersResponse, error) {
	out := new(AllFoldersResponse)
	err := c.cc.Invoke(ctx, "/folderservice.FolderService/GetAllFolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FolderServiceServer is the server API for FolderService service.
// All implementations must embed UnimplementedFolderServiceServer
// for forward compatibility
type FolderServiceServer interface {
	AddFolder(context.Context, *FolderAddDTO) (*FolderResponse, error)
	DeleteFolder(context.Context, *FolderDeleteDTO) (*Details, error)
	RenameFolder(context.Context, *FolderRenameDTO) (*FolderResponse, error)
	GetFolder(context.Context, *FolderGetDTO) (*FolderResponse, error)
	GetAllFolders(context.Context, *UserId) (*AllFoldersResponse, error)
	mustEmbedUnimplementedFolderServiceServer()
}

// UnimplementedFolderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFolderServiceServer struct {
}

func (UnimplementedFolderServiceServer) AddFolder(context.Context, *FolderAddDTO) (*FolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFolder not implemented")
}
func (UnimplementedFolderServiceServer) DeleteFolder(context.Context, *FolderDeleteDTO) (*Details, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFolder not implemented")
}
func (UnimplementedFolderServiceServer) RenameFolder(context.Context, *FolderRenameDTO) (*FolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFolder not implemented")
}
func (UnimplementedFolderServiceServer) GetFolder(context.Context, *FolderGetDTO) (*FolderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFolder not implemented")
}
func (UnimplementedFolderServiceServer) GetAllFolders(context.Context, *UserId) (*AllFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFolders not implemented")
}
func (UnimplementedFolderServiceServer) mustEmbedUnimplementedFolderServiceServer() {}

// UnsafeFolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FolderServiceServer will
// result in compilation errors.
type UnsafeFolderServiceServer interface {
	mustEmbedUnimplementedFolderServiceServer()
}

func RegisterFolderServiceServer(s grpc.ServiceRegistrar, srv FolderServiceServer) {
	s.RegisterService(&FolderService_ServiceDesc, srv)
}

func _FolderService_AddFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderAddDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).AddFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folderservice.FolderService/AddFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).AddFolder(ctx, req.(*FolderAddDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_DeleteFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderDeleteDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).DeleteFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folderservice.FolderService/DeleteFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).DeleteFolder(ctx, req.(*FolderDeleteDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_RenameFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderRenameDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).RenameFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folderservice.FolderService/RenameFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).RenameFolder(ctx, req.(*FolderRenameDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_GetFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FolderGetDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).GetFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folderservice.FolderService/GetFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).GetFolder(ctx, req.(*FolderGetDTO))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_GetAllFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).GetAllFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/folderservice.FolderService/GetAllFolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).GetAllFolders(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// FolderService_ServiceDesc is the grpc.ServiceDesc for FolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "folderservice.FolderService",
	HandlerType: (*FolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFolder",
			Handler:    _FolderService_AddFolder_Handler,
		},
		{
			MethodName: "DeleteFolder",
			Handler:    _FolderService_DeleteFolder_Handler,
		},
		{
			MethodName: "RenameFolder",
			Handler:    _FolderService_RenameFolder_Handler,
		},
		{
			MethodName: "GetFolder",
			Handler:    _FolderService_GetFolder_Handler,
		},
		{
			MethodName: "GetAllFolders",
			Handler:    _FolderService_GetAllFolders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "folder.proto",
}
