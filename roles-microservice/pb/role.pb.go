// Code generated by protoc-gen-go. DO NOT EDIT.
// source: role.proto

/*
Package role is a generated protocol buffer package.

It is generated from these files:
	role.proto

It has these top-level messages:
	EmptyRequest
	GetUserRoleRequest
	RolesReply
	Role
	UserRoleReply
*/
package role

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Requests
type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetUserRoleRequest struct {
	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserRoleRequest) Reset()                    { *m = GetUserRoleRequest{} }
func (m *GetUserRoleRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRoleRequest) ProtoMessage()               {}
func (*GetUserRoleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetUserRoleRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type RolesReply struct {
	Roles []*Role `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
}

func (m *RolesReply) Reset()                    { *m = RolesReply{} }
func (m *RolesReply) String() string            { return proto.CompactTextString(m) }
func (*RolesReply) ProtoMessage()               {}
func (*RolesReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RolesReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	Id   int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Role) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserRoleReply struct {
	UserId int32   `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Roles  []*Role `protobuf:"bytes,2,rep,name=roles" json:"roles,omitempty"`
}

func (m *UserRoleReply) Reset()                    { *m = UserRoleReply{} }
func (m *UserRoleReply) String() string            { return proto.CompactTextString(m) }
func (*UserRoleReply) ProtoMessage()               {}
func (*UserRoleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserRoleReply) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserRoleReply) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "role.EmptyRequest")
	proto.RegisterType((*GetUserRoleRequest)(nil), "role.GetUserRoleRequest")
	proto.RegisterType((*RolesReply)(nil), "role.RolesReply")
	proto.RegisterType((*Role)(nil), "role.Role")
	proto.RegisterType((*UserRoleReply)(nil), "role.UserRoleReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Roles service

type RolesClient interface {
	GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error)
	GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error)
}

type rolesClient struct {
	cc *grpc.ClientConn
}

func NewRolesClient(cc *grpc.ClientConn) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) GetRoles(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*RolesReply, error) {
	out := new(RolesReply)
	err := grpc.Invoke(ctx, "/role.Roles/GetRoles", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetUserRole(ctx context.Context, in *GetUserRoleRequest, opts ...grpc.CallOption) (*UserRoleReply, error) {
	out := new(UserRoleReply)
	err := grpc.Invoke(ctx, "/role.Roles/GetUserRole", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Roles service

type RolesServer interface {
	GetRoles(context.Context, *EmptyRequest) (*RolesReply, error)
	GetUserRole(context.Context, *GetUserRoleRequest) (*UserRoleReply, error)
}

func RegisterRolesServer(s *grpc.Server, srv RolesServer) {
	s.RegisterService(&_Roles_serviceDesc, srv)
}

func _Roles_GetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Roles/GetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetRoles(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Roles/GetUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetUserRole(ctx, req.(*GetUserRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Roles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "role.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoles",
			Handler:    _Roles_GetRoles_Handler,
		},
		{
			MethodName: "GetUserRole",
			Handler:    _Roles_GetUserRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}

func init() { proto.RegisterFile("role.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x50, 0x3d, 0x4f, 0x84, 0x40,
	0x14, 0x04, 0x04, 0xd4, 0x41, 0x89, 0x79, 0x16, 0x12, 0x2a, 0xb2, 0x15, 0x31, 0x91, 0x02, 0x7b,
	0x3b, 0x63, 0xb4, 0xdc, 0xc4, 0xda, 0x68, 0xd8, 0xc2, 0x04, 0x04, 0xd9, 0xa5, 0x20, 0xf7, 0xe7,
	0x6f, 0x3f, 0x2e, 0x07, 0x97, 0x0b, 0xdd, 0xec, 0x7b, 0xb3, 0xf3, 0x66, 0x06, 0x18, 0xfb, 0x56,
	0x54, 0xc3, 0xd8, 0xab, 0x9e, 0x42, 0x83, 0x59, 0x8a, 0x9b, 0xd7, 0x6e, 0x50, 0x33, 0x17, 0xff,
	0x93, 0x90, 0x8a, 0x3d, 0x81, 0xde, 0x84, 0xfa, 0x94, 0x62, 0xe4, 0x7a, 0x7d, 0x98, 0xd2, 0x03,
	0x2e, 0x27, 0x3d, 0xfa, 0xfa, 0x6d, 0x32, 0xbf, 0xf0, 0xcb, 0x88, 0xc7, 0xe6, 0xf9, 0xde, 0xb0,
	0x0a, 0x30, 0x3c, 0xc9, 0xc5, 0xd0, 0xce, 0x54, 0x20, 0x32, 0xa2, 0x52, 0x93, 0x2e, 0xca, 0xa4,
	0x46, 0x65, 0xcf, 0x59, 0x21, 0xb7, 0x60, 0x8f, 0x08, 0xcd, 0x93, 0x52, 0x04, 0x47, 0x2d, 0x8d,
	0x88, 0x10, 0xfe, 0x7d, 0x77, 0x22, 0x0b, 0xf4, 0xe4, 0x9a, 0x5b, 0xcc, 0x3e, 0x70, 0xbb, 0xf8,
	0x30, 0xf2, 0x5b, 0x2e, 0x96, 0xbb, 0xc1, 0xc6, 0xdd, 0x7a, 0x87, 0xc8, 0xfa, 0xa4, 0x1a, 0x57,
	0x3a, 0x9f, 0xc3, 0xe4, 0x78, 0xeb, 0xfc, 0xf9, 0xdd, 0xf2, 0xd7, 0x85, 0x62, 0x1e, 0xbd, 0x20,
	0x59, 0x75, 0x42, 0x99, 0xa3, 0x9c, 0xd7, 0x94, 0xdf, 0xbb, 0xcd, 0x89, 0x6b, 0xe6, 0xfd, 0xc4,
	0xb6, 0xf0, 0xe7, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x07, 0xb6, 0x37, 0x35, 0x7e, 0x01, 0x00,
	0x00,
}
