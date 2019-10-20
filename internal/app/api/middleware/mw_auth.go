package middleware

import (
	"caas-micro/internal/app/api/pkg/ginplus"
	"caas-micro/pkg/errors"
	"caas-micro/proto/auth"
	"context"

	"github.com/gin-gonic/gin"
)

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.AuthService, skipper ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userID string
		if t := ginplus.GetToken(c); t != "" {
			// id, err := a.ParseUserID(t)
			// if err != nil {
			// 	if err == auth.ErrInvalidToken {
			// 		ginplus.ResError(c, errors.ErrNoPerm)
			// 		return
			// 	}
			// 	ginplus.ResError(c, errors.WithStack(err))
			// 	return
			// }
			resp, err := a.VertifyToken(context.TODO(), &auth.Request{
				Username: t,
				Password: t,
			})
			if err != nil {
				ginplus.ResError(c, err)
			}
			userID = resp.GetMsg()
		}

		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}

		if userID == "" {
			// if config.GetGlobalConfig().RunMode == "debug" {
			// 	c.Set(ginplus.UserIDKey, config.GetGlobalConfig().Root.UserName)
			// 	c.Next()
			// 	return
			// }
			ginplus.ResError(c, errors.ErrNoPerm)
		}
	}
}
