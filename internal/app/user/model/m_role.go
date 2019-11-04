package model

import (
	"caas-micro/proto/user"
	"context"
)

// IRole 角色管理
type IRole interface {
	// 查询数据
	Query(ctx context.Context, params user.RoleQueryParam, opts ...user.RoleQueryOptions) (*user.RoleQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...user.RoleQueryOptions) (*user.RoleSchema, error)
	// 创建数据
	Create(ctx context.Context, item user.RoleSchema) error
	// 更新数据
	Update(ctx context.Context, recordID string, item user.RoleSchema) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
