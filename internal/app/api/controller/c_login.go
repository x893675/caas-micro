package controller

import (
	"caas-micro/internal/app/api/pkg/ginplus"
	"caas-micro/internal/app/api/schema"
	"caas-micro/proto/auth"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	AuthSvc auth.AuthService
}

func NewLoginController(auth auth.AuthService) *LoginController {
	return &LoginController{
		AuthSvc: auth,
	}
}

//func (s *LoginController) Anything(c *gin.Context) {
//	log.Print("Received Say.Anything API request")
//	c.JSON(200, map[string]string{
//		"message": "Hi, this is the Greeter API",
//	})
//}
//
//func (s *LoginController) Hello(c *gin.Context) {
//	log.Print("Received Say.Hello API request")
//
//	name := c.Param("name")
//
//	response, err := s.AuthSvc.DestroyToken(context.TODO(), &auth.Request{
//		Username: name,
//		Password: name,
//	})
//
//	if err != nil {
//		c.JSON(500, err)
//	}
//
//	c.JSON(200, response)
//}

func (s *LoginController) Login(c *gin.Context) {
	log.Print("Received api.Login request")

	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	response, err := s.AuthSvc.Verify(context.TODO(), &auth.LoginRequest{
		Username: item.UserName,
		Password: item.Password,
	})

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, response)

}

// Logout 用户登出
// @Summary 用户登出
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Router POST /api/v1/login/exit
func (a *LoginController) Logout(c *gin.Context) {
	// 检查用户是否处于登录状态，如果是则执行销毁
	userID := ginplus.GetUserID(c)
	if userID != "" {
		_, err := a.AuthSvc.DestroyToken(context.TODO(), &auth.TokenString{
			Token: ginplus.GetToken(c),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	ginplus.ResOK(c)
	//userID := ginplus.GetUserID(c)
	//if userID != "" {
	//	ctx := ginplus.NewContext(c)
	//	err := a.LoginBll.DestroyToken(ctx, ginplus.GetToken(c))
	//	if err != nil {
	//		logger.Errorf(ctx, err.Error())
	//	}
	//	logger.StartSpan(ginplus.NewContext(c), logger.SetSpanTitle("用户登出"), logger.SetSpanFuncName("Logout")).Infof("登出系统")
	//}
	//ginplus.ResOK(c)
}

// OpenshiftLogin openshift登录验证
// @Summary openshift登录验证
// @Param Authorization header string true "Basic 用户令牌"
// @Success 200 schema.OpenshiftUserShow
// @Failure 401 schema.HTTPError "{error:{code:400,message:无效的密码}}"
// @Router GET /api/v1/auth
func (a *LoginController) OpenshiftLogin(c *gin.Context) {

	username, password, err := ginplus.GetBasicToken(c)
	if err != nil {
		ginplus.ResError(c, err)
	}
	userInfo, err := a.AuthSvc.OpensfiftVerify(context.TODO(), &auth.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ginplus.ResOpenshiftLoginError(c, err)
		return
	}
	ginplus.ResSuccess(c, userInfo)
	//username, password, err := ginplus.GetBasicToken(c)
	//
	//if err != nil {
	//	ginplus.ResError(c, err)
	//}
	//
	//user, err := a.LoginBll.Verify(ginplus.NewContext(c), username, password)
	//if err != nil {
	//	ginplus.ResOpenshiftLoginError(c, err)
	//	return
	//}
	//
	//ginplus.ResSuccess(c, user.ToOpenshiftShows())
}
