// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: group.proto

package contactV1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	Create(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Get(ctx context.Context, in *SearchGroupRequest, opts ...grpc.CallOption) (*GroupsResponse, error)
	Member(ctx context.Context, in *SearchGroupMemberRequest, opts ...grpc.CallOption) (*GroupMemberResponse, error)
	Members(ctx context.Context, in *SearchGroupMemberRequest, opts ...grpc.CallOption) (*GroupMembersResponse, error)
	Join(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Quit(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	KickOut(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Mine(ctx context.Context, in *SearchGroupRequest, opts ...grpc.CallOption) (*GroupsResponse, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) Create(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Update(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Delete(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Get(ctx context.Context, in *SearchGroupRequest, opts ...grpc.CallOption) (*GroupsResponse, error) {
	out := new(GroupsResponse)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Member(ctx context.Context, in *SearchGroupMemberRequest, opts ...grpc.CallOption) (*GroupMemberResponse, error) {
	out := new(GroupMemberResponse)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Member", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Members(ctx context.Context, in *SearchGroupMemberRequest, opts ...grpc.CallOption) (*GroupMembersResponse, error) {
	out := new(GroupMembersResponse)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Members", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Join(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Quit(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Quit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) KickOut(ctx context.Context, in *UpdateGroupMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/KickOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) Mine(ctx context.Context, in *SearchGroupRequest, opts ...grpc.CallOption) (*GroupsResponse, error) {
	out := new(GroupsResponse)
	err := c.cc.Invoke(ctx, "/contactPb.v1.Group/Mine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	Create(context.Context, *UpdateGroupRequest) (*GroupResponse, error)
	Update(context.Context, *UpdateGroupRequest) (*emptypb.Empty, error)
	Delete(context.Context, *UpdateGroupRequest) (*emptypb.Empty, error)
	Get(context.Context, *SearchGroupRequest) (*GroupsResponse, error)
	Member(context.Context, *SearchGroupMemberRequest) (*GroupMemberResponse, error)
	Members(context.Context, *SearchGroupMemberRequest) (*GroupMembersResponse, error)
	Join(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error)
	Quit(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error)
	KickOut(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error)
	Mine(context.Context, *SearchGroupRequest) (*GroupsResponse, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) Create(context.Context, *UpdateGroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupServer) Update(context.Context, *UpdateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupServer) Delete(context.Context, *UpdateGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupServer) Get(context.Context, *SearchGroupRequest) (*GroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupServer) Member(context.Context, *SearchGroupMemberRequest) (*GroupMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Member not implemented")
}
func (UnimplementedGroupServer) Members(context.Context, *SearchGroupMemberRequest) (*GroupMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Members not implemented")
}
func (UnimplementedGroupServer) Join(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedGroupServer) Quit(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Quit not implemented")
}
func (UnimplementedGroupServer) KickOut(context.Context, *UpdateGroupMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KickOut not implemented")
}
func (UnimplementedGroupServer) Mine(context.Context, *SearchGroupRequest) (*GroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mine not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Create(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Update(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Delete(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Get(ctx, req.(*SearchGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Member_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Member(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Member",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Member(ctx, req.(*SearchGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Members_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Members(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Members",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Members(ctx, req.(*SearchGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Join(ctx, req.(*UpdateGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Quit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Quit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Quit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Quit(ctx, req.(*UpdateGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_KickOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).KickOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/KickOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).KickOut(ctx, req.(*UpdateGroupMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_Mine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Mine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/contactPb.v1.Group/Mine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Mine(ctx, req.(*SearchGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contactPb.v1.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Group_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Group_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Group_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Group_Get_Handler,
		},
		{
			MethodName: "Member",
			Handler:    _Group_Member_Handler,
		},
		{
			MethodName: "Members",
			Handler:    _Group_Members_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _Group_Join_Handler,
		},
		{
			MethodName: "Quit",
			Handler:    _Group_Quit_Handler,
		},
		{
			MethodName: "KickOut",
			Handler:    _Group_KickOut_Handler,
		},
		{
			MethodName: "Mine",
			Handler:    _Group_Mine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
