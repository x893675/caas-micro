package schema

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
