package model

import (
	"caas-micro/internal/app/user/model/impl/gorm"
	"caas-micro/pkg/errors"
	"caas-micro/proto/user"
	"context"
)

// NewUser 创建用户存储实例
func NewUser(db *gorm.DB) *User {
	return &User{db}
}

// User 用户存储
type User struct {
	db *gorm.DB
}

func (a *User) getQueryOption(opts ...user.UserQueryOptions) user.UserQueryOptions {
	var opt user.UserQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *User) Query(ctx context.Context, params user.QueryRequest, opts ...user.UserQueryOptions) (*user.QueryResult, error) {
	db := gorm.GetUserDB(ctx, a.db).DB
	if v := params.UserName; v != "" {
		db = db.Where("user_name=?", v)
	}
	if v := params.LikeUserName; v != "" {
		db = db.Where("user_name LIKE ?", "%"+v+"%")
	}
	if v := params.LikeRealName; v != "" {
		db = db.Where("real_name LIKE ?", "%"+v+"%")
	}
	if v := params.Email; v != "" {
		db = db.Where("email=?", v)
	}
	if v := params.Status; v > 0 {
		db = db.Where("status=?", v)
	}
	if v := params.RoleIDS; len(v) > 0 {
		subQuery := gorm.GetUserRoleDB(ctx, a.db).Select("user_id").Where("role_id IN(?)", v).SubQuery()
		db = db.Where("record_id IN(?)", subQuery)
	}
	db = db.Order("id DESC")

	opt := a.getQueryOption(opts...)
	var list gorm.Users
	pr, err := WrapPageQuery(db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	qr := &user.QueryResult{
		PageResult: pr,
		Data:       list.ToSchemaUsers(),
	}

	err = a.fillSchemaUsers(ctx, qr.Data, opts...)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (a *User) Create(ctx context.Context, item user.UserSchema) error {
	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		sitem := gorm.SchemaUser(item)
		result := gorm.GetUserDB(ctx, a.db).Create(sitem.ToUser())
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}

		for _, eitem := range sitem.ToUserRoles() {
			result := gorm.GetUserRoleDB(ctx, a.db).Create(eitem)
			if err := result.Error; err != nil {
				return errors.WithStack(err)
			}
		}
		return nil
	})
	//return ExecTrans(ctx, a.db, func(ctx context.Context) error {
	//	sitem := entity.SchemaUser(item)
	//	result := entity.GetUserDB(ctx, a.db).Create(sitem.ToUser())
	//	if err := result.Error; err != nil {
	//		return errors.WithStack(err)
	//	}
	//
	//	for _, eitem := range sitem.ToUserRoles() {
	//		result := entity.GetUserRoleDB(ctx, a.db).Create(eitem)
	//		if err := result.Error; err != nil {
	//			return errors.WithStack(err)
	//		}
	//	}
	//	return nil
	//})
}

func (a *User) Get(ctx context.Context, recordID string, opts ...user.UserQueryOptions) (*user.UserSchema, error) {
	var item gorm.User
	ok, err := a.db.FindOne(gorm.GetUserDB(ctx, a.db).Where("record_id=?", recordID), &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}
	sitem := item.ToSchemaUser()
	err = a.fillSchemaUsers(ctx, []*user.UserSchema{sitem}, opts...)
	if err != nil {
		return nil, err
	}
	return sitem, nil
}

func (a *User) Delete(ctx context.Context, recordID string) error {

	return ExecTrans(ctx, a.db, func(ctx context.Context) error {
		result := gorm.GetUserDB(ctx, a.db).Where("record_id=?", recordID).Delete(gorm.User{})
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}

		result = gorm.GetUserRoleDB(ctx, a.db).Where("user_id=?", recordID).Delete(gorm.UserRole{})
		if err := result.Error; err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
}

func (a *User) fillSchemaUsers(ctx context.Context, items []*user.UserSchema, opts ...user.UserQueryOptions) error {
	opt := a.getQueryOption(opts...)

	if opt.IncludeRoles {
		userIDs := make([]string, len(items))
		for i, item := range items {
			userIDs[i] = item.RecordID
		}

		var roleList gorm.UserRoles
		if opt.IncludeRoles {
			items, err := a.queryRoles(ctx, userIDs...)
			if err != nil {
				return err
			}
			roleList = items
		}

		for i, item := range items {
			if len(roleList) > 0 {
				items[i].Roles = roleList.GetByUserID(item.RecordID)
			}
		}
	}

	return nil
}

func (a *User) queryRoles(ctx context.Context, userIDs ...string) (gorm.UserRoles, error) {
	var list gorm.UserRoles
	result := gorm.GetUserRoleDB(ctx, a.db).Where("user_id IN(?)", userIDs).Find(&list)
	if err := result.Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return list, nil
}
