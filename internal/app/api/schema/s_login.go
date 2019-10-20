package schema

// LoginParam 登录参数
type LoginParam struct {
	UserName string `json:"user_name" binding:"required" swaggo:"true,用户名"`
	Password string `json:"password" binding:"required" swaggo:"true,密码(md5加密)"`
}
