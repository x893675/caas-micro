package user

import (
	"caas-micro/internal/app/user/model"
	"caas-micro/proto/auth"
	"caas-micro/proto/user"
	"context"
	"fmt"
	"github.com/google/wire"
)

//var (
//	JWT_SIGNING_METHOD = util.GetEnvironment("JWT_SIGNING_METHOD", "HS512")
//)

type UserServer struct {
	authSvc   auth.AuthService
	userModel model.IUser
}

func NewUserServer(a auth.AuthService, user model.IUser) (*UserServer, error) {
	return &UserServer{
		authSvc:   a,
		userModel: user,
	}, nil
}

func (u *UserServer) MigrateDB() error {

}

func (u *UserServer) Query(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResult) error {
	fmt.Println("in user srv query: ", req.UserName)

	if req.UserName != "hanamichi" {
		return nil
	}

	//item := user.UserEntity{
	//	RecordID:  "1234567",
	//	UserName:  "hanamichi",
	//	RealName:  "hanamichi",
	//	Password:  "hanamichi",
	//	Phone:     "hanamichi",
	//	Email:     "dfjkldjfkld",
	//	Status:    0,
	//	CreatedAt: nil,
	//	Roles:     nil,
	//}
	//rsp.Data = append(rsp.Data, &item)
	return nil
}

var ProviderSet = wire.NewSet(NewUserServer)
