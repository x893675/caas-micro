package entity

import (
	"caas-micro/internal/app/user/model/impl/gorm"
)

// AutoMigrate 自动映射数据表
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		new(User),
		new(UserRole),
		new(Role),
		new(RoleMenu),
		new(Menu),
		new(MenuAction),
		new(MenuResource),
	).Error
}

func NewEntity(db *gorm.DB) error {
	SetTablePrefix("g_")
	return autoMigrate(db)
}
