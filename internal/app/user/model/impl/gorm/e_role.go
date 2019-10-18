package gorm

import (
	"caas-micro/proto/user"
	"context"
	"github.com/golang/protobuf/ptypes"
	"strings"
)

// GetRoleDB 获取角色存储
func GetRoleDB(ctx context.Context, defDB *DB) *DB {
	return getDBWithModel(ctx, defDB, Role{})
}

// GetRoleMenuDB 获取角色菜单关联存储
func GetRoleMenuDB(ctx context.Context, defDB *DB) *DB {
	return getDBWithModel(ctx, defDB, RoleMenu{})
}

// SchemaRole 角色对象
type SchemaRole user.RoleSchema

// ToRole 转换为角色实体
func (a SchemaRole) ToRole() *Role {
	item := &Role{
		RecordID: a.RecordID,
		Name:     &a.Name,
		Sequence: &a.Sequence,
		Memo:     &a.Memo,
		Creator:  &a.Creator,
	}
	return item
}

// ToRoleMenus 转换为角色菜单实体列表
func (a SchemaRole) ToRoleMenus() []*RoleMenu {
	list := make([]*RoleMenu, len(a.Menus))
	for i, item := range a.Menus {
		list[i] = SchemaRoleMenu(*item).ToRoleMenu(a.RecordID)
	}
	return list
}

// Role 角色实体
type Role struct {
	Model
	RecordID string  `gorm:"column:record_id;size:36;index;"` // 记录内码
	Name     *string `gorm:"column:name;size:100;index;"`     // 角色名称
	Sequence *int64  `gorm:"column:sequence;index;"`          // 排序值
	Memo     *string `gorm:"column:memo;size:200;"`           // 备注
	Creator  *string `gorm:"column:creator;size:36;"`         // 创建者
}

func (a Role) String() string {
	return toString(a)
}

// TableName 表名
func (a Role) TableName() string {
	return a.Model.TableName("role")
}

// ToSchemaRole 转换为角色对象
func (a Role) ToSchemaRole() *user.RoleSchema {
	ts, _ := ptypes.TimestampProto(a.CreatedAt)
	item := &user.RoleSchema{
		RecordID:  a.RecordID,
		Name:      *a.Name,
		Sequence:  *a.Sequence,
		Memo:      *a.Memo,
		Creator:   *a.Creator,
		CreatedAt: ts,
	}
	return item
}

// Roles 角色实体列表
type Roles []*Role

// ToSchemaRoles 转换为角色对象列表
func (a Roles) ToSchemaRoles() []*user.RoleSchema {
	list := make([]*user.RoleSchema, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRole()
	}
	return list
}

// SchemaRoleMenu 角色菜单对象
type SchemaRoleMenu user.RoleMenu

// ToRoleMenu 转换为角色菜单实体
func (a SchemaRoleMenu) ToRoleMenu(roleID string) *RoleMenu {
	item := &RoleMenu{
		RoleID: roleID,
		MenuID: a.MenuID,
	}

	var action string
	if v := a.Actions; len(v) > 0 {
		action = strings.Join(v, ",")
	}
	item.Action = &action

	var resource string
	if v := a.Resources; len(v) > 0 {
		resource = strings.Join(v, ",")
	}
	item.Resource = &resource

	return item
}

// RoleMenu 角色菜单关联实体
type RoleMenu struct {
	Model
	RoleID   string  `gorm:"column:role_id;size:36;index;"` // 角色内码
	MenuID   string  `gorm:"column:menu_id;size:36;index;"` // 菜单内码
	Action   *string `gorm:"column:action;size:2048;"`      // 动作权限(多个以英文逗号分隔)
	Resource *string `gorm:"column:resource;size:2048;"`    // 资源权限(多个以英文逗号分隔)
}

// TableName 表名
func (a RoleMenu) TableName() string {
	return a.Model.TableName("role_menu")
}

// ToSchemaRoleMenu 转换为角色菜单对象
func (a RoleMenu) ToSchemaRoleMenu() *user.RoleMenu {
	item := &user.RoleMenu{
		MenuID: a.MenuID,
	}

	if v := a.Action; v != nil && *v != "" {
		item.Actions = strings.Split(*v, ",")
	}
	if v := a.Resource; v != nil && *v != "" {
		item.Resources = strings.Split(*v, ",")
	}

	return item
}

// RoleMenus 角色菜单关联实体列表
type RoleMenus []*RoleMenu

// GetByRoleID 根据角色ID获取角色菜单对象列表
func (a RoleMenus) GetByRoleID(roleID string) []*user.RoleMenu {
	var list []*user.RoleMenu
	for _, item := range a {
		if item.RoleID == roleID {
			list = append(list, item.ToSchemaRoleMenu())
		}
	}
	return list
}

// ToSchemaRoleMenus 转换为角色菜单对象列表
func (a RoleMenus) ToSchemaRoleMenus() []*user.RoleMenu {
	list := make([]*user.RoleMenu, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaRoleMenu()
	}
	return list
}

// ToMap 转换为键值映射
func (a RoleMenus) ToMap() map[string]*RoleMenu {
	m := make(map[string]*RoleMenu)
	for _, item := range a {
		m[item.MenuID] = item
	}
	return m
}
