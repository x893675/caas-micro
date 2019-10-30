package user

import (
	"caas-micro/pkg/util"
	"caas-micro/proto/auth"
	"caas-micro/proto/user"
	"context"
	"fmt"
)

var (
	JWT_SIGNING_METHOD = util.GetEnvironment("JWT_SIGNING_METHOD", "HS512")
)

type UserServer struct {
	authSvc auth.AuthService
}

func NewUserServer(a auth.AuthService) (*UserServer, error) {
	return &UserServer{
		authSvc: a,
	}, nil
}

func (u *UserServer) Query(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResult) error {
	fmt.Println("in user srv query: ", req.UserName)

	if req.UserName != "hanamichi" {
		return nil
	}

	item := user.UserEntity{
		RecordID:  "1234567",
		UserName:  "hanamichi",
		RealName:  "hanamichi",
		Password:  "hanamichi",
		Phone:     "hanamichi",
		Email:     "dfjkldjfkld",
		Status:    0,
		CreatedAt: nil,
		Roles:     nil,
	}
	rsp.Data = append(rsp.Data, &item)
	return nil
}
