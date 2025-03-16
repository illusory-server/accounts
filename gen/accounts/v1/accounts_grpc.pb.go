// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: accounts/v1/accounts.proto

package v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AccountsService_Create_FullMethodName               = "/accounts.v1.AccountsService/Create"
	AccountsService_UpdateInfoById_FullMethodName       = "/accounts.v1.AccountsService/UpdateInfoById"
	AccountsService_UpdateNicknameById_FullMethodName   = "/accounts.v1.AccountsService/UpdateNicknameById"
	AccountsService_UpdateRoleById_FullMethodName       = "/accounts.v1.AccountsService/UpdateRoleById"
	AccountsService_UpdateEmailById_FullMethodName      = "/accounts.v1.AccountsService/UpdateEmailById"
	AccountsService_UpdatePasswordById_FullMethodName   = "/accounts.v1.AccountsService/UpdatePasswordById"
	AccountsService_AddAvatarLink_FullMethodName        = "/accounts.v1.AccountsService/AddAvatarLink"
	AccountsService_DeleteById_FullMethodName           = "/accounts.v1.AccountsService/DeleteById"
	AccountsService_DeleteManyById_FullMethodName       = "/accounts.v1.AccountsService/DeleteManyById"
	AccountsService_GetAccountById_FullMethodName       = "/accounts.v1.AccountsService/GetAccountById"
	AccountsService_GetAccountByEmail_FullMethodName    = "/accounts.v1.AccountsService/GetAccountByEmail"
	AccountsService_GetAccountByNickname_FullMethodName = "/accounts.v1.AccountsService/GetAccountByNickname"
	AccountsService_GetAccountsByIds_FullMethodName     = "/accounts.v1.AccountsService/GetAccountsByIds"
	AccountsService_GetAccountsByQuery_FullMethodName   = "/accounts.v1.AccountsService/GetAccountsByQuery"
)

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsServiceClient interface {
	Create(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	UpdateInfoById(ctx context.Context, in *UpdateInfoByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateNicknameById(ctx context.Context, in *UpdateNickByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateRoleById(ctx context.Context, in *UpdateRoleByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateEmailById(ctx context.Context, in *UpdateEmailByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdatePasswordById(ctx context.Context, in *UpdatePasswordByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddAvatarLink(ctx context.Context, in *AddAvatarLinkRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteManyById(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAccountById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Account, error)
	GetAccountByEmail(ctx context.Context, in *String, opts ...grpc.CallOption) (*Account, error)
	GetAccountByNickname(ctx context.Context, in *String, opts ...grpc.CallOption) (*Account, error)
	GetAccountsByIds(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Accounts, error)
	GetAccountsByQuery(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error)
}

type accountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsServiceClient(cc grpc.ClientConnInterface) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) Create(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, AccountsService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) UpdateInfoById(ctx context.Context, in *UpdateInfoByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_UpdateInfoById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) UpdateNicknameById(ctx context.Context, in *UpdateNickByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_UpdateNicknameById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) UpdateRoleById(ctx context.Context, in *UpdateRoleByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_UpdateRoleById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) UpdateEmailById(ctx context.Context, in *UpdateEmailByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_UpdateEmailById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) UpdatePasswordById(ctx context.Context, in *UpdatePasswordByIdRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_UpdatePasswordById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) AddAvatarLink(ctx context.Context, in *AddAvatarLinkRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_AddAvatarLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) DeleteById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_DeleteById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) DeleteManyById(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*empty.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AccountsService_DeleteManyById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccountById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, AccountsService_GetAccountById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccountByEmail(ctx context.Context, in *String, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, AccountsService_GetAccountByEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccountByNickname(ctx context.Context, in *String, opts ...grpc.CallOption) (*Account, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Account)
	err := c.cc.Invoke(ctx, AccountsService_GetAccountByNickname_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccountsByIds(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Accounts, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Accounts)
	err := c.cc.Invoke(ctx, AccountsService_GetAccountsByIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) GetAccountsByQuery(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryAccountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryAccountsResponse)
	err := c.cc.Invoke(ctx, AccountsService_GetAccountsByQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
// All implementations must embed UnimplementedAccountsServiceServer
// for forward compatibility.
type AccountsServiceServer interface {
	Create(context.Context, *CreateAccountRequest) (*Account, error)
	UpdateInfoById(context.Context, *UpdateInfoByIdRequest) (*empty.Empty, error)
	UpdateNicknameById(context.Context, *UpdateNickByIdRequest) (*empty.Empty, error)
	UpdateRoleById(context.Context, *UpdateRoleByIdRequest) (*empty.Empty, error)
	UpdateEmailById(context.Context, *UpdateEmailByIdRequest) (*empty.Empty, error)
	UpdatePasswordById(context.Context, *UpdatePasswordByIdRequest) (*empty.Empty, error)
	AddAvatarLink(context.Context, *AddAvatarLinkRequest) (*empty.Empty, error)
	DeleteById(context.Context, *Id) (*empty.Empty, error)
	DeleteManyById(context.Context, *Ids) (*empty.Empty, error)
	GetAccountById(context.Context, *Id) (*Account, error)
	GetAccountByEmail(context.Context, *String) (*Account, error)
	GetAccountByNickname(context.Context, *String) (*Account, error)
	GetAccountsByIds(context.Context, *Ids) (*Accounts, error)
	GetAccountsByQuery(context.Context, *QueryRequest) (*QueryAccountsResponse, error)
	mustEmbedUnimplementedAccountsServiceServer()
}

// UnimplementedAccountsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccountsServiceServer struct{}

func (UnimplementedAccountsServiceServer) Create(context.Context, *CreateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServiceServer) UpdateInfoById(context.Context, *UpdateInfoByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInfoById not implemented")
}
func (UnimplementedAccountsServiceServer) UpdateNicknameById(context.Context, *UpdateNickByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNicknameById not implemented")
}
func (UnimplementedAccountsServiceServer) UpdateRoleById(context.Context, *UpdateRoleByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoleById not implemented")
}
func (UnimplementedAccountsServiceServer) UpdateEmailById(context.Context, *UpdateEmailByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEmailById not implemented")
}
func (UnimplementedAccountsServiceServer) UpdatePasswordById(context.Context, *UpdatePasswordByIdRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePasswordById not implemented")
}
func (UnimplementedAccountsServiceServer) AddAvatarLink(context.Context, *AddAvatarLinkRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAvatarLink not implemented")
}
func (UnimplementedAccountsServiceServer) DeleteById(context.Context, *Id) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedAccountsServiceServer) DeleteManyById(context.Context, *Ids) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteManyById not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccountById(context.Context, *Id) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountById not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccountByEmail(context.Context, *String) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByEmail not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccountByNickname(context.Context, *String) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountByNickname not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccountsByIds(context.Context, *Ids) (*Accounts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsByIds not implemented")
}
func (UnimplementedAccountsServiceServer) GetAccountsByQuery(context.Context, *QueryRequest) (*QueryAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountsByQuery not implemented")
}
func (UnimplementedAccountsServiceServer) mustEmbedUnimplementedAccountsServiceServer() {}
func (UnimplementedAccountsServiceServer) testEmbeddedByValue()                         {}

// UnsafeAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServiceServer will
// result in compilation errors.
type UnsafeAccountsServiceServer interface {
	mustEmbedUnimplementedAccountsServiceServer()
}

func RegisterAccountsServiceServer(s grpc.ServiceRegistrar, srv AccountsServiceServer) {
	// If the following call pancis, it indicates UnimplementedAccountsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccountsService_ServiceDesc, srv)
}

func _AccountsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Create(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_UpdateInfoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInfoByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).UpdateInfoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_UpdateInfoById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).UpdateInfoById(ctx, req.(*UpdateInfoByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_UpdateNicknameById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNickByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).UpdateNicknameById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_UpdateNicknameById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).UpdateNicknameById(ctx, req.(*UpdateNickByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_UpdateRoleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).UpdateRoleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_UpdateRoleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).UpdateRoleById(ctx, req.(*UpdateRoleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_UpdateEmailById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateEmailByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).UpdateEmailById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_UpdateEmailById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).UpdateEmailById(ctx, req.(*UpdateEmailByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_UpdatePasswordById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).UpdatePasswordById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_UpdatePasswordById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).UpdatePasswordById(ctx, req.(*UpdatePasswordByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_AddAvatarLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAvatarLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).AddAvatarLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_AddAvatarLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).AddAvatarLink(ctx, req.(*AddAvatarLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_DeleteById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).DeleteById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_DeleteManyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).DeleteManyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_DeleteManyById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).DeleteManyById(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccountById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccountById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_GetAccountById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccountById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccountByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccountByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_GetAccountByEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccountByEmail(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccountByNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccountByNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_GetAccountByNickname_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccountByNickname(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccountsByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ids)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccountsByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_GetAccountsByIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccountsByIds(ctx, req.(*Ids))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_GetAccountsByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).GetAccountsByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccountsService_GetAccountsByQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).GetAccountsByQuery(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountsService_ServiceDesc is the grpc.ServiceDesc for AccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accounts.v1.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AccountsService_Create_Handler,
		},
		{
			MethodName: "UpdateInfoById",
			Handler:    _AccountsService_UpdateInfoById_Handler,
		},
		{
			MethodName: "UpdateNicknameById",
			Handler:    _AccountsService_UpdateNicknameById_Handler,
		},
		{
			MethodName: "UpdateRoleById",
			Handler:    _AccountsService_UpdateRoleById_Handler,
		},
		{
			MethodName: "UpdateEmailById",
			Handler:    _AccountsService_UpdateEmailById_Handler,
		},
		{
			MethodName: "UpdatePasswordById",
			Handler:    _AccountsService_UpdatePasswordById_Handler,
		},
		{
			MethodName: "AddAvatarLink",
			Handler:    _AccountsService_AddAvatarLink_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _AccountsService_DeleteById_Handler,
		},
		{
			MethodName: "DeleteManyById",
			Handler:    _AccountsService_DeleteManyById_Handler,
		},
		{
			MethodName: "GetAccountById",
			Handler:    _AccountsService_GetAccountById_Handler,
		},
		{
			MethodName: "GetAccountByEmail",
			Handler:    _AccountsService_GetAccountByEmail_Handler,
		},
		{
			MethodName: "GetAccountByNickname",
			Handler:    _AccountsService_GetAccountByNickname_Handler,
		},
		{
			MethodName: "GetAccountsByIds",
			Handler:    _AccountsService_GetAccountsByIds_Handler,
		},
		{
			MethodName: "GetAccountsByQuery",
			Handler:    _AccountsService_GetAccountsByQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accounts/v1/accounts.proto",
}
