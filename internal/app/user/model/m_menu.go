package model

import (
	"caas-micro/proto/user"
	"context"
)

// IMenu 菜单管理存储接口
type IMenu interface {
	// 查询数据
	Query(ctx context.Context, params user.MenuQueryParam, opts ...user.MenuQueryOptions) (*user.MenuQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...user.MenuQueryOptions) (*user.MenuSchema, error)
	// 创建数据
	Create(ctx context.Context, item user.MenuSchema) error
	// 更新数据
	Update(ctx context.Context, recordID string, item user.MenuSchema) error
	// 更新父级路径
	UpdateParentPath(ctx context.Context, recordID, parentPath string) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
