package model

import (
	imodel "caas-micro/internal/app/user/model/impl/gorm/model"
	"caas-micro/proto/user"
	"context"
	"github.com/google/wire"
)

// IUser 用户对象存储接口
type IUser interface {
	// 查询数据
	Query(ctx context.Context, params user.QueryRequest, opts ...user.UserQueryOptions) (*user.QueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...user.UserQueryOptions) (*user.UserSchema, error)
	// 创建数据
	Create(ctx context.Context, item user.UserSchema) error
	//// 更新数据
	//Update(ctx context.Context, recordID string, item schema.User) error
	// 删除数据
	Delete(ctx context.Context, recordID string) error
	//// 更新状态
	//UpdateStatus(ctx context.Context, recordID string, status int) error
	//// 更新密码
	//UpdatePassword(ctx context.Context, recordID, password string) error
	////通过email查询
	//GetByEmail(ctx context.Context, email string) (*schema.User, error)
}

var (
	ProviderProductionSet = wire.NewSet(imodel.ProviderSet, wire.Bind(new(IUser), new(*imodel.User)),
		wire.Bind(new(IRole), new(*imodel.Role)), wire.Bind(new(IMenu), new(*imodel.Menu)),
		wire.Bind(new(ITrans), new(*imodel.Trans)))
)
