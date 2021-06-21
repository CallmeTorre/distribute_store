// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rpc_definition

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

// DistributeStoreClient is the client API for DistributeStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DistributeStoreClient interface {
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (DistributeStore_GetClient, error)
	Put(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
	Append(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error)
}

type distributeStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewDistributeStoreClient(cc grpc.ClientConnInterface) DistributeStoreClient {
	return &distributeStoreClient{cc}
}

func (c *distributeStoreClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (DistributeStore_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &DistributeStore_ServiceDesc.Streams[0], "/rpc_definition.DistributeStore/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &distributeStoreGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributeStore_GetClient interface {
	Recv() (*Value, error)
	grpc.ClientStream
}

type distributeStoreGetClient struct {
	grpc.ClientStream
}

func (x *distributeStoreGetClient) Recv() (*Value, error) {
	m := new(Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributeStoreClient) Put(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc_definition.DistributeStore/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributeStoreClient) Append(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc_definition.DistributeStore/Append", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DistributeStoreServer is the server API for DistributeStore service.
// All implementations must embed UnimplementedDistributeStoreServer
// for forward compatibility
type DistributeStoreServer interface {
	Get(*Key, DistributeStore_GetServer) error
	Put(context.Context, *Message) (*Empty, error)
	Append(context.Context, *Message) (*Empty, error)
	mustEmbedUnimplementedDistributeStoreServer()
}

// UnimplementedDistributeStoreServer must be embedded to have forward compatible implementations.
type UnimplementedDistributeStoreServer struct {
}

func (UnimplementedDistributeStoreServer) Get(*Key, DistributeStore_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDistributeStoreServer) Put(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedDistributeStoreServer) Append(context.Context, *Message) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Append not implemented")
}
func (UnimplementedDistributeStoreServer) mustEmbedUnimplementedDistributeStoreServer() {}

// UnsafeDistributeStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DistributeStoreServer will
// result in compilation errors.
type UnsafeDistributeStoreServer interface {
	mustEmbedUnimplementedDistributeStoreServer()
}

func RegisterDistributeStoreServer(s grpc.ServiceRegistrar, srv DistributeStoreServer) {
	s.RegisterService(&DistributeStore_ServiceDesc, srv)
}

func _DistributeStore_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Key)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributeStoreServer).Get(m, &distributeStoreGetServer{stream})
}

type DistributeStore_GetServer interface {
	Send(*Value) error
	grpc.ServerStream
}

type distributeStoreGetServer struct {
	grpc.ServerStream
}

func (x *distributeStoreGetServer) Send(m *Value) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributeStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributeStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_definition.DistributeStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributeStoreServer).Put(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributeStore_Append_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributeStoreServer).Append(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc_definition.DistributeStore/Append",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributeStoreServer).Append(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// DistributeStore_ServiceDesc is the grpc.ServiceDesc for DistributeStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DistributeStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc_definition.DistributeStore",
	HandlerType: (*DistributeStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _DistributeStore_Put_Handler,
		},
		{
			MethodName: "Append",
			Handler:    _DistributeStore_Append_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _DistributeStore_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpc_definition/distribute_store.proto",
}