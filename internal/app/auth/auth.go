package auth

import (
	"caas-micro/internal/app/auth/auther"
	"caas-micro/internal/app/auth/auther/jwtauth"
	"caas-micro/internal/app/auth/auther/jwtauth/store/buntdb"
	"caas-micro/pkg/errors"
	"caas-micro/pkg/util"
	"caas-micro/proto/auth"
	"caas-micro/proto/user"
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
)

var (
	JWT_SIGNING_METHOD = util.GetEnvironment("JWT_SIGNING_METHOD", "HS512")

	JWT_SIGNING_KEY = util.GetEnvironment("JWT_SIGNING_KEY", "caasmicro")

	JWT_EXPIRED = util.GetEnvironment("JWT_EXPIRED", "7200")

	JWT_STORE = util.GetEnvironment("JWT_STORE", "file")

	JWT_STORE_PATH = util.GetEnvironment("JWT_STORE_PATH", "/jwt_auth.db")
)

type AuthServer struct {
	userSvc user.UserService
	auther  auther.Auther
}

func NewAuthServer(a auther.Auther, u user.UserService) (*AuthServer, error) {
	return &AuthServer{
		userSvc: u,
		auther:  a,
	}, nil
}

func NewAuther() (auther.Auther, error) {
	exp, _ := util.S(JWT_EXPIRED).Int()
	var opts []jwtauth.Option
	opts = append(opts, jwtauth.SetExpired(exp))
	opts = append(opts, jwtauth.SetSigningKey([]byte(JWT_SIGNING_KEY)))
	opts = append(opts, jwtauth.SetKeyfunc(func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, auther.ErrInvalidToken
		}
		return []byte(JWT_SIGNING_KEY), nil
	}))

	switch JWT_SIGNING_METHOD {
	case "HS256":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS256))
	case "HS384":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS384))
	case "HS512":
		opts = append(opts, jwtauth.SetSigningMethod(jwt.SigningMethodHS512))
	}

	var store jwtauth.Storer
	switch JWT_STORE {
	case "file":
		s, err := buntdb.NewStore(JWT_STORE_PATH)
		if err != nil {
			return nil, err
		}
		store = s
	}

	return jwtauth.New(store, opts...), nil
}

// func (a *AuthServer) GenerateToken(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
// 	log.Println("in GenerateToken")
// 	log.Println(req.Username)
// 	log.Println(req.Password)
// 	rsp.Msg = "Hello " + req.Username
// 	return nil
// }

func (a *AuthServer) generateToken(ctx context.Context, userID string, token *auth.Token) error {
	tokenInfo, err := a.auther.GenerateToken(userID)
	if err != nil {
		return errors.WithStack(err)
	}

	token.AccessToken = tokenInfo.GetAccessToken()
	token.TokenType = tokenInfo.GetTokenType()
	token.ExpiresAt = tokenInfo.GetExpiresAt()
	return nil
}

func (a *AuthServer) DestroyToken(ctx context.Context, req *auth.TokenString, rsp *auth.NullResult) error {
	log.Println("in DestroyToken")
	err := a.auther.DestroyToken(req.Token)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a *AuthServer) Verify(ctx context.Context, req *auth.LoginRequest, rsp *auth.Token) error {
	log.Println("in Verify")
	response, err := a.userSvc.Query(ctx, &user.QueryRequest{
		UserName: req.Username,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if len(response.Data) == 0 {
		return errors.ErrInvalidUserName
	}

	item := response.Data[0]
	if item.Password != util.SHA1HashString(req.Password) {
		return errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return errors.ErrUserDisable
	}

	err = a.generateToken(ctx, item.RecordID, rsp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (a *AuthServer) VertifyToken(ctx context.Context, req *auth.TokenString, rsp *auth.VerifyResponse) error {
	log.Println("in VertifyToken")
	log.Println(req.Token)

	uid, err := a.auther.ParseUserID(req.Token)

	if err != nil {
		if err == auther.ErrInvalidToken {
			return errors.ErrInvalidToken
		}
		return err
	}

	rsp.Uid = uid
	return nil
}

func (a *AuthServer) OpensfiftVerify(ctx context.Context, req *auth.LoginRequest, rsp *auth.OpenshiftAuthResult) error {
	log.Println("in OpensfiftVerify")
	response, err := a.userSvc.Query(ctx, &user.QueryRequest{
		UserName: req.Username,
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else if len(response.Data) == 0 {
		return errors.ErrInvalidUserName
	}

	item := response.Data[0]
	if item.Password != util.SHA1HashString(req.Password) {
		return errors.ErrInvalidPassword
	} else if item.Status != 1 {
		return errors.ErrUserDisable
	}

	rsp.RecordID = item.RecordID
	rsp.UserName = item.UserName
	rsp.RealName = item.RealName
	rsp.Email = item.Email
	return nil
}
