// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: filehandler.proto

package filehandler

import (
	context "context"
	common "github.com/thomas-osgood/ofs/protobufs/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileserviceClient is the client API for Fileservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileserviceClient interface {
	// rpc designed to delete a file located within the
	// uploads directory of the server.
	DeleteFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*common.StatusMessage, error)
	// rpc designed to upload a file from the machine the
	// agent is running on to the control server.
	DownloadFile(ctx context.Context, opts ...grpc.CallOption) (Fileservice_DownloadFileClient, error)
	// rpc designed to gather and return a list of files
	// that can be downloaded by the client.
	ListFiles(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (Fileservice_ListFilesClient, error)
	// rpc designed to create a subdirectory within the
	// uploads directory on the server.
	MakeDirectory(ctx context.Context, in *MakeDirectoryRequest, opts ...grpc.CallOption) (*common.StatusMessage, error)
	// rpc designed to rename a file as requested by
	// the client.
	RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*common.StatusMessage, error)
	// rpc designed to download a file to the machine the
	// agent is running on from the control server.
	UploadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (Fileservice_UploadFileClient, error)
}

type fileserviceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileserviceClient(cc grpc.ClientConnInterface) FileserviceClient {
	return &fileserviceClient{cc}
}

func (c *fileserviceClient) DeleteFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*common.StatusMessage, error) {
	out := new(common.StatusMessage)
	err := c.cc.Invoke(ctx, "/filehandler.Fileservice/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileserviceClient) DownloadFile(ctx context.Context, opts ...grpc.CallOption) (Fileservice_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fileservice_ServiceDesc.Streams[0], "/filehandler.Fileservice/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileserviceDownloadFileClient{stream}
	return x, nil
}

type Fileservice_DownloadFileClient interface {
	Send(*FileChunk) error
	CloseAndRecv() (*common.StatusMessage, error)
	grpc.ClientStream
}

type fileserviceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *fileserviceDownloadFileClient) Send(m *FileChunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileserviceDownloadFileClient) CloseAndRecv() (*common.StatusMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(common.StatusMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileserviceClient) ListFiles(ctx context.Context, in *common.Empty, opts ...grpc.CallOption) (Fileservice_ListFilesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fileservice_ServiceDesc.Streams[1], "/filehandler.Fileservice/ListFiles", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileserviceListFilesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fileservice_ListFilesClient interface {
	Recv() (*FileInfo, error)
	grpc.ClientStream
}

type fileserviceListFilesClient struct {
	grpc.ClientStream
}

func (x *fileserviceListFilesClient) Recv() (*FileInfo, error) {
	m := new(FileInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileserviceClient) MakeDirectory(ctx context.Context, in *MakeDirectoryRequest, opts ...grpc.CallOption) (*common.StatusMessage, error) {
	out := new(common.StatusMessage)
	err := c.cc.Invoke(ctx, "/filehandler.Fileservice/MakeDirectory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileserviceClient) RenameFile(ctx context.Context, in *RenameFileRequest, opts ...grpc.CallOption) (*common.StatusMessage, error) {
	out := new(common.StatusMessage)
	err := c.cc.Invoke(ctx, "/filehandler.Fileservice/RenameFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileserviceClient) UploadFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (Fileservice_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &Fileservice_ServiceDesc.Streams[2], "/filehandler.Fileservice/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileserviceUploadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Fileservice_UploadFileClient interface {
	Recv() (*FileChunk, error)
	grpc.ClientStream
}

type fileserviceUploadFileClient struct {
	grpc.ClientStream
}

func (x *fileserviceUploadFileClient) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileserviceServer is the server API for Fileservice service.
// All implementations must embed UnimplementedFileserviceServer
// for forward compatibility
type FileserviceServer interface {
	// rpc designed to delete a file located within the
	// uploads directory of the server.
	DeleteFile(context.Context, *FileRequest) (*common.StatusMessage, error)
	// rpc designed to upload a file from the machine the
	// agent is running on to the control server.
	DownloadFile(Fileservice_DownloadFileServer) error
	// rpc designed to gather and return a list of files
	// that can be downloaded by the client.
	ListFiles(*common.Empty, Fileservice_ListFilesServer) error
	// rpc designed to create a subdirectory within the
	// uploads directory on the server.
	MakeDirectory(context.Context, *MakeDirectoryRequest) (*common.StatusMessage, error)
	// rpc designed to rename a file as requested by
	// the client.
	RenameFile(context.Context, *RenameFileRequest) (*common.StatusMessage, error)
	// rpc designed to download a file to the machine the
	// agent is running on from the control server.
	UploadFile(*FileRequest, Fileservice_UploadFileServer) error
	mustEmbedUnimplementedFileserviceServer()
}

// UnimplementedFileserviceServer must be embedded to have forward compatible implementations.
type UnimplementedFileserviceServer struct {
}

func (UnimplementedFileserviceServer) DeleteFile(context.Context, *FileRequest) (*common.StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileserviceServer) DownloadFile(Fileservice_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedFileserviceServer) ListFiles(*common.Empty, Fileservice_ListFilesServer) error {
	return status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFileserviceServer) MakeDirectory(context.Context, *MakeDirectoryRequest) (*common.StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeDirectory not implemented")
}
func (UnimplementedFileserviceServer) RenameFile(context.Context, *RenameFileRequest) (*common.StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenameFile not implemented")
}
func (UnimplementedFileserviceServer) UploadFile(*FileRequest, Fileservice_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedFileserviceServer) mustEmbedUnimplementedFileserviceServer() {}

// UnsafeFileserviceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileserviceServer will
// result in compilation errors.
type UnsafeFileserviceServer interface {
	mustEmbedUnimplementedFileserviceServer()
}

func RegisterFileserviceServer(s grpc.ServiceRegistrar, srv FileserviceServer) {
	s.RegisterService(&Fileservice_ServiceDesc, srv)
}

func _Fileservice_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileserviceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filehandler.Fileservice/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileserviceServer).DeleteFile(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fileservice_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileserviceServer).DownloadFile(&fileserviceDownloadFileServer{stream})
}

type Fileservice_DownloadFileServer interface {
	SendAndClose(*common.StatusMessage) error
	Recv() (*FileChunk, error)
	grpc.ServerStream
}

type fileserviceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *fileserviceDownloadFileServer) SendAndClose(m *common.StatusMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileserviceDownloadFileServer) Recv() (*FileChunk, error) {
	m := new(FileChunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Fileservice_ListFiles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(common.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileserviceServer).ListFiles(m, &fileserviceListFilesServer{stream})
}

type Fileservice_ListFilesServer interface {
	Send(*FileInfo) error
	grpc.ServerStream
}

type fileserviceListFilesServer struct {
	grpc.ServerStream
}

func (x *fileserviceListFilesServer) Send(m *FileInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _Fileservice_MakeDirectory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeDirectoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileserviceServer).MakeDirectory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filehandler.Fileservice/MakeDirectory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileserviceServer).MakeDirectory(ctx, req.(*MakeDirectoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fileservice_RenameFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenameFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileserviceServer).RenameFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filehandler.Fileservice/RenameFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileserviceServer).RenameFile(ctx, req.(*RenameFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fileservice_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileserviceServer).UploadFile(m, &fileserviceUploadFileServer{stream})
}

type Fileservice_UploadFileServer interface {
	Send(*FileChunk) error
	grpc.ServerStream
}

type fileserviceUploadFileServer struct {
	grpc.ServerStream
}

func (x *fileserviceUploadFileServer) Send(m *FileChunk) error {
	return x.ServerStream.SendMsg(m)
}

// Fileservice_ServiceDesc is the grpc.ServiceDesc for Fileservice service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fileservice_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "filehandler.Fileservice",
	HandlerType: (*FileserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFile",
			Handler:    _Fileservice_DeleteFile_Handler,
		},
		{
			MethodName: "MakeDirectory",
			Handler:    _Fileservice_MakeDirectory_Handler,
		},
		{
			MethodName: "RenameFile",
			Handler:    _Fileservice_RenameFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadFile",
			Handler:       _Fileservice_DownloadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListFiles",
			Handler:       _Fileservice_ListFiles_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _Fileservice_UploadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "filehandler.proto",
}
