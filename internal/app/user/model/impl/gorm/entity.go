package gorm

import (
	icontext "caas-micro/internal/app/user/context"
	"caas-micro/pkg/util"
	"context"
	"fmt"
	"time"
)

// 表名前缀
var tablePrefix string

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

// GetTablePrefix 获取表名前缀
func GetTablePrefix() string {
	return tablePrefix
}

// Model base model
type Model struct {
	ID        uint       `gorm:"column:id;primary_key;auto_increment;"`
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", GetTablePrefix(), name)
}

func toString(v interface{}) string {
	return util.JSONMarshalToString(v)
}
func getDBWithModel(ctx context.Context, defDB *DB, m interface{}) *DB {
	return Wrap(getDB(ctx, defDB).Model(m))
}

func getDB(ctx context.Context, defDB *DB) *DB {
	trans, ok := icontext.FromTrans(ctx)
	if ok {
		db, ok := trans.(*DB)
		if ok {
			return db
		}
	}
	return defDB
}
