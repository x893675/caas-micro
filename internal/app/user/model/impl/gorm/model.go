package gorm

import (
	icontext "caas-micro/internal/app/user/context"
	"caas-micro/internal/app/user/pkg/gormplus"
	"caas-micro/proto/user"
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// ExecTrans 执行事务
func ExecTrans(ctx context.Context, db *gormplus.DB, fn func(context.Context) error) error {
	if _, ok := icontext.FromTrans(ctx); ok {
		return fn(ctx)
	}

	transModel := NewTrans(db)
	trans, err := transModel.Begin(ctx)
	if err != nil {
		return err
	}

	err = fn(icontext.NewTrans(ctx, trans))
	if err != nil {
		_ = transModel.Rollback(ctx, trans)
		return err
	}
	return transModel.Commit(ctx, trans)
}

// WrapPageQuery 包装带有分页的查询
func WrapPageQuery(db *gorm.DB, pp *user.PaginationParam, out interface{}) (*user.PaginationResult, error) {
	if pp != nil {
		total, err := gormplus.Wrap(db).FindPage(db, int(pp.PageIndex), int(pp.PageSize), out)
		if err != nil {
			return nil, err
		}
		return &user.PaginationResult{
			Total: int64(total),
		}, nil
	}

	result := db.Find(out)
	return nil, result.Error
}

var ProviderSet = wire.NewSet(NewUser, NewTrans)
