package schema

import (
	"caas-micro/proto/user"
	"time"
)

// LoginParam 登录参数
type LoginParam struct {
	UserName string `json:"user_name" binding:"required" swaggo:"true,用户名"`
	Password string `json:"password" binding:"required" swaggo:"true,密码(md5加密)"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

//// UserRole 用户角色
//type UserRoleParam struct {
//	RoleID string `json:"role_id" swaggo:"true,角色ID"`
//}

// User 用户对象
type CreateUserParam struct {
	RecordID  string           `json:"record_id" swaggo:"false,记录ID"`
	UserName  string           `json:"user_name" binding:"required" swaggo:"true,用户名"`
	RealName  string           `json:"real_name" binding:"required" swaggo:"true,真实姓名"`
	Password  string           `json:"password" swaggo:"false,密码"`
	Phone     string           `json:"phone" swaggo:"false,手机号"`
	Email     string           `json:"email" binding:"required" swaggo:"true,邮箱"`
	Status    int              `json:"status" binding:"required,max=2,min=1" swaggo:"true,用户状态(1:启用 2:停用)"`
	Creator   string           `json:"creator" swaggo:"false,创建者"`
	CreatedAt time.Time        `json:"created_at" swaggo:"false,创建时间"`
	Roles     []*user.UserRole `json:"roles" binding:"required,gt=0" swaggo:"true,角色授权"`
}
