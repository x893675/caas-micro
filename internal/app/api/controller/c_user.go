package controller

import (
	"caas-micro/internal/app/api/pkg/ginplus"
	"caas-micro/internal/app/api/schema"
	"caas-micro/pkg/errors"
	"caas-micro/proto/user"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println("in querypage 1")
	var params user.QueryRequest
	//params.LikeUserName = c.Query("user_name")
	//params.LikeRealName = c.Query("real_name")
	//params.Email = c.Query("email")
	//if v := util.S(c.Query("status")).DefaultInt64(0); v > 0 {
	//	params.Status = v
	//}
	//
	//if v := c.Query("role_ids"); v != "" {
	//	params.RoleIDS = strings.Split(v, ",")
	//}
	var opt user.UserQueryOptions
	opt.IncludeRoles = true
	opt.PageParam = &user.PaginationParam{
		PageIndex: 1,
		PageSize:  10,
	}
	//params.QueryOpt.IncludeRoles = true
	//params.QueryOpt.PageParam.PageIndex = 1
	//params.QueryOpt.PageParam.PageSize = 10
	params.QueryOpt = &opt
	//params.QueryOpt.PageParam = ginplus.GetPaginationParam(c)
	fmt.Println("in querypage 2")
	result, err := a.UserSvc.QueryShow(context.TODO(), &params)
	if err != nil {
		fmt.Println("err is ", err.Error())
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Create 创建数据
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.User true
// @Success 200 schema.User
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router POST /api/v1/users
func (a *UserController) Create(c *gin.Context) {
	var item schema.CreateUserParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	param := &user.UserSchema{
		UserName: item.UserName,
		RealName: item.RealName,
		Password: item.Password,
		Phone:    item.Phone,
		Email:    item.Email,
		Status:   int64(item.Status),
		Roles:    item.Roles,
	}
	fmt.Println("role id is ", param.Roles[0].RoleID)
	nitem, err := a.UserSvc.Create(context.TODO(), param)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
	//ginplus.ResSuccess(c, nitem.CleanSecure())
}

// Delete 删除数据
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.HTTPStatus "{status:OK}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router DELETE /api/v1/users/{id}
func (a *UserController) Delete(c *gin.Context) {
	_, err := a.UserSvc.Delete(context.TODO(), &user.DeleteUserRequest{
		Uid: c.Param("id"),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)

	//err := a.UserBll.Delete(ginplus.NewContext(c), c.Param("id"))
	//if err != nil {
	//	ginplus.ResError(c, err)
	//	return
	//}
	//ginplus.ResOK(c)
}

// Get 查询指定数据
// Get 查询指定数据
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 schema.User
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router GET /api/v1/users/{id}
func (a *UserController) Get(c *gin.Context) {
	item, err := a.UserSvc.Get(context.TODO(), &user.GetUserRequest{
		Uid: c.Param("id"),
		QueryOpt: &user.UserQueryOptions{
			IncludeRoles: true,
		},
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
	//item, err := a.UserBll.Get(ginplus.NewContext(c), c.Param("id"), schema.UserQueryOptions{
	//	IncludeRoles: true,
	//})
	//if err != nil {
	//	ginplus.ResError(c, err)
	//	return
	//}
	//ginplus.ResSuccess(c, item.CleanSecure())
}

// Update 更新数据
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.User true
// @Success 200 schema.User
// @Failure 400 schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router PUT /api/v1/users/{id}
func (a *UserController) Update(c *gin.Context) {
	var item schema.CreateUserParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	param := &user.UpdateUserRequest{
		Uid: c.Param("id"),
		User: &user.UserSchema{
			UserName: item.UserName,
			RealName: item.RealName,
			Password: item.Password,
			Phone:    item.Phone,
			Email:    item.Email,
			Status:   int64(item.Status),
			Roles:    item.Roles,
		},
	}
	nitem, err := a.UserSvc.Update(context.TODO(), param)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
	//var item schema.User
	//if err := ginplus.ParseJSON(c, &item); err != nil {
	//	ginplus.ResError(c, err)
	//	return
	//}
	//
	//nitem, err := a.UserBll.Update(ginplus.NewContext(c), c.Param("id"), item)
	//if err != nil {
	//	ginplus.ResError(c, err)
	//	return
	//}
	//ginplus.ResSuccess(c, nitem.CleanSecure())
}
