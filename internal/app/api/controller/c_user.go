package controller

import (
	"caas-micro/internal/app/api/pkg/ginplus"
	"caas-micro/pkg/errors"
	"caas-micro/pkg/util"
	"caas-micro/proto/user"
	"context"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserController struct {
	UserSvc user.UserService
}

func NewUserController(user user.UserService) *UserController {
	return &UserController{
		UserSvc: user,
	}
}

// Query 查询数据
func (a *UserController) Query(c *gin.Context) {
	switch c.Query("q") {
	case "page":
		a.QueryPage(c)
	default:
		ginplus.ResError(c, errors.ErrUnknownQuery)
	}
}

// QueryPage 查询分页数据
// @Summary 查询分页数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" 1
// @Param pageSize query int true "分页大小" 10
// @Param user_name query string false "用户名(模糊查询)"
// @Param real_name query string false "真实姓名(模糊查询)"
// @Param email query string false "邮箱"
// @Param role_ids query string false "角色ID(多个以英文逗号分隔)"
// @Param status query int false "状态(1:启用 2:停用)"
// @Success 200 []schema.UserShow "分页查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 400 schema.HTTPError "{error:{code:0,message:未知的查询类型}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/users?q=page
func (a *UserController) QueryPage(c *gin.Context) {
	var params user.QueryRequest
	params.LikeUserName = c.Query("user_name")
	params.LikeRealName = c.Query("real_name")
	params.Email = c.Query("email")
	if v := util.S(c.Query("status")).DefaultInt64(0); v > 0 {
		params.Status = v
	}

	if v := c.Query("role_ids"); v != "" {
		params.RoleIDS = strings.Split(v, ",")
	}

	params.QueryOpt.IncludeRoles = true
	params.QueryOpt.PageParam = ginplus.GetPaginationParam(c)
	result, err := a.UserSvc.QueryShow(context.TODO(), &params)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResPage(c, result.Data, result.PageResult)
}
