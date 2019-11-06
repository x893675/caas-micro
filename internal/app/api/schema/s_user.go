package schema

// UpdatePasswordParam 更新密码请求参数
type UpdatePasswordParam struct {
	OldPassword string `json:"old_password" binding:"required" swaggo:"true,旧密码(md5加密)"`
	NewPassword string `json:"new_password" binding:"required" swaggo:"true,新密码(md5加密)"`
}
