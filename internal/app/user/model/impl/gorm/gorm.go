package gorm

import (
	igorm "github.com/jinzhu/gorm"
	"time"

	// gorm存储注入
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Config 配置参数
type Config struct {
	Debug        bool
	DBType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

func NewDB() (*DB, error) {
	c := &Config{
		Debug:        false,
		DBType:       "mysql",
		DSN:          "root:123456@tcp(mysql:3306)/user?charset=utf8mb4&parseTime=True&loc=Local&allowNativePasswords=true",
		MaxLifetime:  7200,
		MaxOpenConns: 150,
		MaxIdleConns: 50,
	}
	db, err := igorm.Open(c.DBType, c.DSN)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	SetTablePrefix("g_")
	autoMigrate(db)

	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	db.DB().SetMaxOpenConns(c.MaxOpenConns)
	db.DB().SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

// Wrap 包装gorm
func Wrap(db *igorm.DB) *DB {
	return &DB{db}
}

// DB gorm扩展DB
type DB struct {
	*igorm.DB
}

// FindPage 查询分页数据
func (d *DB) FindPage(db *igorm.DB, pageIndex, pageSize int, out interface{}) (int, error) {
	var count int
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return 0, err
	} else if count == 0 {
		return 0, nil
	}

	// 如果分页大小小于0，则不查询数据
	if pageSize < 0 || pageIndex < 0 {
		return count, nil
	}

	if pageIndex > 0 && pageSize > 0 {
		db = db.Offset((pageIndex - 1) * pageSize)
	}
	if pageSize > 0 {
		db = db.Limit(pageSize)
	}
	result = db.Find(out)
	if err := result.Error; err != nil {
		return 0, err
	}

	return count, nil
}

// FindOne 查询单条数据
func (d *DB) FindOne(db *igorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == igorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Check 检查数据是否存在
func (d *DB) Check(db *igorm.DB) (bool, error) {
	var count int
	result := db.Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// AutoMigrate 自动映射数据表
func autoMigrate(db *igorm.DB) error {
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
