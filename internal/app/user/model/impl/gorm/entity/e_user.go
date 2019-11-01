package entity

import (
	"caas-micro/internal/app/user/pkg/gormplus"
	"caas-micro/proto/user"
	"context"
	"github.com/golang/protobuf/ptypes"
)

// GetUserDB 获取用户存储
func GetUserDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, User{})
}

// GetUserRoleDB 获取用户角色关联存储
func GetUserRoleDB(ctx context.Context, defDB *gormplus.DB) *gormplus.DB {
	return getDBWithModel(ctx, defDB, UserRole{})
}

// SchemaUser 用户对象
type SchemaUser user.UserSchema

// ToUser 转换为用户实体
func (a SchemaUser) ToUser() *User {
	item := &User{
		RecordID: a.RecordID,
		UserName: &a.UserName,
		RealName: &a.RealName,
		Password: &a.Password,
		Status:   &a.Status,
		Creator:  &a.Creator,
		Email:    &a.Email,
		Phone:    &a.Phone,
	}
	return item
}

// ToUserRoles 转换为用户角色关联列表
func (a SchemaUser) ToUserRoles() []*UserRole {
	list := make([]*UserRole, len(a.Roles))
	for i, item := range a.Roles {
		list[i] = &UserRole{
			UserID: a.RecordID,
			RoleID: item.RoleID,
		}
	}
	return list
}

// User 用户实体
type User struct {
	Model
	RecordID string  `gorm:"column:record_id;size:36;index;"` // 记录内码
	UserName *string `gorm:"column:user_name;size:64;index;"` // 用户名
	RealName *string `gorm:"column:real_name;size:64;index;"` // 真实姓名
	Password *string `gorm:"column:password;size:40;"`        // 密码(sha1(md5(明文))加密)
	Email    *string `gorm:"column:email;size:255;index;"`    // 邮箱
	Phone    *string `gorm:"column:phone;size:20;index;"`     // 手机号
	Status   *int64  `gorm:"column:status;index;"`            // 状态(1:启用 2:停用)
	Creator  *string `gorm:"column:creator;size:36;"`         // 创建者
}

func (a User) String() string {
	return toString(a)
}

// TableName 表名
func (a User) TableName() string {
	return a.Model.TableName("user")
}

// ToSchemaUser 转换为用户对象
func (a User) ToSchemaUser() *user.UserSchema {
	ts, _ := ptypes.TimestampProto(a.CreatedAt)
	item := &user.UserSchema{
		RecordID:  a.RecordID,
		UserName:  *a.UserName,
		RealName:  *a.RealName,
		Password:  *a.Password,
		Status:    *a.Status,
		Creator:   *a.Creator,
		Email:     *a.Email,
		Phone:     *a.Phone,
		CreatedAt: ts,
	}
	return item
}

// Users 用户实体列表
type Users []*User

// ToSchemaUsers 转换为用户对象列表
func (a Users) ToSchemaUsers() []*user.UserSchema {
	list := make([]*user.UserSchema, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUser()
	}
	return list
}

// UserRole 用户角色关联实体
type UserRole struct {
	Model
	UserID string `gorm:"column:user_id;size:36;index;"` // 用户内码
	RoleID string `gorm:"column:role_id;size:36;index;"` // 角色内码
}

// TableName 表名
func (a UserRole) TableName() string {
	return a.Model.TableName("user_role")
}

// ToSchemaUserRole 转换为用户角色对象
func (a UserRole) ToSchemaUserRole() *user.UserRole {
	return &user.UserRole{
		RoleID: a.RoleID,
	}
}

// UserRoles 用户角色关联列表
type UserRoles []*UserRole

// GetByUserID 根据用户ID获取用户角色对象列表
func (a UserRoles) GetByUserID(userID string) []*user.UserRole {
	var list []*user.UserRole
	for _, item := range a {
		if item.UserID == userID {
			list = append(list, item.ToSchemaUserRole())
		}
	}
	return list
}

// ToSchemaUserRoles 转换为用户角色对象列表
func (a UserRoles) ToSchemaUserRoles() []*user.UserRole {
	list := make([]*user.UserRole, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaUserRole()
	}
	return list
}
