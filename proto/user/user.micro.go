// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResult, error)
	QueryShow(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*UserShowQueryResult, error)
	Create(ctx context.Context, in *UserSchema, opts ...client.CallOption) (*UserSchema, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*NullResult, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*UserSchema, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResult, error) {
	req := c.c.NewRequest(c.name, "User.Query", in)
	out := new(QueryResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) QueryShow(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*UserShowQueryResult, error) {
	req := c.c.NewRequest(c.name, "User.QueryShow", in)
	out := new(UserShowQueryResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Create(ctx context.Context, in *UserSchema, opts ...client.CallOption) (*UserSchema, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(UserSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Delete(ctx context.Context, in *DeleteUserRequest, opts ...client.CallOption) (*NullResult, error) {
	req := c.c.NewRequest(c.name, "User.Delete", in)
	out := new(NullResult)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Get(ctx context.Context, in *GetUserRequest, opts ...client.CallOption) (*UserSchema, error) {
	req := c.c.NewRequest(c.name, "User.Get", in)
	out := new(UserSchema)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Query(context.Context, *QueryRequest, *QueryResult) error
	QueryShow(context.Context, *QueryRequest, *UserShowQueryResult) error
	Create(context.Context, *UserSchema, *UserSchema) error
	Delete(context.Context, *DeleteUserRequest, *NullResult) error
	Get(context.Context, *GetUserRequest, *UserSchema) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Query(ctx context.Context, in *QueryRequest, out *QueryResult) error
		QueryShow(ctx context.Context, in *QueryRequest, out *UserShowQueryResult) error
		Create(ctx context.Context, in *UserSchema, out *UserSchema) error
		Delete(ctx context.Context, in *DeleteUserRequest, out *NullResult) error
		Get(ctx context.Context, in *GetUserRequest, out *UserSchema) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Query(ctx context.Context, in *QueryRequest, out *QueryResult) error {
	return h.UserHandler.Query(ctx, in, out)
}

func (h *userHandler) QueryShow(ctx context.Context, in *QueryRequest, out *UserShowQueryResult) error {
	return h.UserHandler.QueryShow(ctx, in, out)
}

func (h *userHandler) Create(ctx context.Context, in *UserSchema, out *UserSchema) error {
	return h.UserHandler.Create(ctx, in, out)
}

func (h *userHandler) Delete(ctx context.Context, in *DeleteUserRequest, out *NullResult) error {
	return h.UserHandler.Delete(ctx, in, out)
}

func (h *userHandler) Get(ctx context.Context, in *GetUserRequest, out *UserSchema) error {
	return h.UserHandler.Get(ctx, in, out)
}
