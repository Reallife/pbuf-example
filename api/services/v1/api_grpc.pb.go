// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package services

import (
	context "context"
	v1 "github.com/Reallife/pbuf-example/api/messages/v1"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessagesAPIClient is the client API for MessagesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagesAPIClient interface {
	GetDirectMessages(ctx context.Context, in *User, opts ...grpc.CallOption) (*DirectMessages, error)
	SendDirectMessages(ctx context.Context, in *v1.DirectMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type messagesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagesAPIClient(cc grpc.ClientConnInterface) MessagesAPIClient {
	return &messagesAPIClient{cc}
}

func (c *messagesAPIClient) GetDirectMessages(ctx context.Context, in *User, opts ...grpc.CallOption) (*DirectMessages, error) {
	out := new(DirectMessages)
	err := c.cc.Invoke(ctx, "/services.MessagesAPI/GetDirectMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messagesAPIClient) SendDirectMessages(ctx context.Context, in *v1.DirectMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/services.MessagesAPI/SendDirectMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessagesAPIServer is the server API for MessagesAPI service.
// All implementations must embed UnimplementedMessagesAPIServer
// for forward compatibility
type MessagesAPIServer interface {
	GetDirectMessages(context.Context, *User) (*DirectMessages, error)
	SendDirectMessages(context.Context, *v1.DirectMessage) (*empty.Empty, error)
	mustEmbedUnimplementedMessagesAPIServer()
}

// UnimplementedMessagesAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMessagesAPIServer struct {
}

func (UnimplementedMessagesAPIServer) GetDirectMessages(context.Context, *User) (*DirectMessages, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDirectMessages not implemented")
}
func (UnimplementedMessagesAPIServer) SendDirectMessages(context.Context, *v1.DirectMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDirectMessages not implemented")
}
func (UnimplementedMessagesAPIServer) mustEmbedUnimplementedMessagesAPIServer() {}

// UnsafeMessagesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagesAPIServer will
// result in compilation errors.
type UnsafeMessagesAPIServer interface {
	mustEmbedUnimplementedMessagesAPIServer()
}

func RegisterMessagesAPIServer(s grpc.ServiceRegistrar, srv MessagesAPIServer) {
	s.RegisterService(&MessagesAPI_ServiceDesc, srv)
}

func _MessagesAPI_GetDirectMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesAPIServer).GetDirectMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.MessagesAPI/GetDirectMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesAPIServer).GetDirectMessages(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessagesAPI_SendDirectMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.DirectMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessagesAPIServer).SendDirectMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.MessagesAPI/SendDirectMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessagesAPIServer).SendDirectMessages(ctx, req.(*v1.DirectMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// MessagesAPI_ServiceDesc is the grpc.ServiceDesc for MessagesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessagesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.MessagesAPI",
	HandlerType: (*MessagesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDirectMessages",
			Handler:    _MessagesAPI_GetDirectMessages_Handler,
		},
		{
			MethodName: "SendDirectMessages",
			Handler:    _MessagesAPI_SendDirectMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/services/v1/api.proto",
}
