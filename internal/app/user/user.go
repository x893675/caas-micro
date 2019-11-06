package user

import (
	"caas-micro/internal/app/user/model"
	"caas-micro/pkg/errors"
	"caas-micro/pkg/util"
	"caas-micro/proto/auth"
	"caas-micro/proto/user"
	"context"
	"fmt"
	"github.com/google/wire"
)

type UserServer struct {
	authSvc   auth.AuthService
	userModel model.IUser
	roleModel model.IRole
	//menuModel  model.IMenu
	//transModel model.ITrans
}

func NewUserServer(a auth.AuthService, user model.IUser, role model.IRole) (*UserServer, error) {
	return &UserServer{
		authSvc:   a,
		userModel: user,
		roleModel: role,
		//menuModel:  menu,
		//transModel: trans,
	}, nil
}

//// MenuTree 菜单树
//type MenuTree struct {
//	RecordID   string               `json:"record_id" swaggo:"false,记录ID"`
//	Name       string               `json:"name" binding:"required" swaggo:"true,菜单名称"`
//	Sequence   int                  `json:"sequence" swaggo:"false,排序值"`
//	Icon       string               `json:"icon" swaggo:"false,菜单图标"`
//	Router     string               `json:"router" swaggo:"false,访问路由"`
//	Hidden     int                  `json:"hidden" swaggo:"false,隐藏菜单(0:不隐藏 1:隐藏)"`
//	ParentID   string               `json:"parent_id" swaggo:"false,父级ID"`
//	ParentPath string               `json:"parent_path" swaggo:"false,父级路径"`
//	Resources  []*user.MenuResource `json:"resources" swaggo:"false,资源列表"`
//	Actions    []*user.MenuAction   `json:"actions" swaggo:"false,动作列表"`
//	Children   *[]*MenuTree         `json:"children,omitempty" swaggo:"false,子级树"`
//}
//
//// MenuTrees 菜单树列表
//type MenuTrees []*MenuTree
//
//func (u *UserServer) InitData(ctx context.Context) error {
//
//}
//
//func initMenuData(ctx context.Context, u *UserServer) error {
//	// 检查是否存在菜单数据，如果不存在则初始化
//	menuResult, err := u.menuModel.Query(ctx, user.MenuQueryParam{}, user.MenuQueryOptions{
//		PageParam: &user.PaginationParam{PageIndex: -1},
//	})
//	if err != nil {
//		return err
//	} else if menuResult.PageResult.Total > 0 {
//		return nil
//	}
//
//	var data MenuTrees
//	err = util.JSONUnmarshal([]byte(menuData), &data)
//	if err != nil {
//		return err
//	}
//
//	return createMenus(ctx, trans, menu, "", data)
//}
//
//func createMenus(ctx context.Context, u *UserServer, parentID string, list MenuTrees) error {
//	return u.transModel.Exec(ctx, func(ctx context.Context) error {
//		for _, item := range list {
//			sitem := user.MenuSchema{
//				Name:      item.Name,
//				Sequence:  int64(item.Sequence),
//				Icon:      item.Icon,
//				RecordID:  item.RecordID,
//				Router:    item.Router,
//				Hidden:    int64(item.Hidden),
//				ParentID:  parentID,
//				Actions:   item.Actions,
//				Resources: item.Resources,
//			}
//			nsitem, err := u.menuModel.Create(ctx, sitem)
//			if err != nil {
//				return err
//			}
//
//			if item.Children != nil && len(*item.Children) > 0 {
//				err := createMenus(ctx, trans, menu, nsitem.RecordID, *item.Children)
//				if err != nil {
//					return err
//				}
//			}
//		}
//
//		return nil
//	})
//}

func (u *UserServer) Query(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResult) error {
	fmt.Println("in user srv query: ", req.UserName)

	//if req.UserName != "hanamichi" {
	//	return nil
	//}
	result, err := u.userModel.Query(ctx, *req)
	if err != nil {
		return err
	}
	rsp.Data = result.Data
	rsp.PageResult = result.PageResult

	//item := user.UserEntity{
	//	RecordID:  "1234567",
	//	UserName:  "hanamichi",
	//	RealName:  "hanamichi",
	//	Password:  "hanamichi",
	//	Phone:     "hanamichi",
	//	Email:     "dfjkldjfkld",
	//	Status:    0,
	//	CreatedAt: nil,
	//	Roles:     nil,
	//}
	//rsp.Data = append(rsp.Data, &item)
	return nil
}

func (u *UserServer) QueryShow(ctx context.Context, req *user.QueryRequest, rsp *user.UserShowQueryResult) error {
	//opts := user.UserQueryOptions{
	//	IncludeRoles: req.QueryOpt.IncludeRoles,
	//	PageParam: req.QueryOpt.PageParam,
	//}
	userResult, err := u.userModel.Query(ctx, *req, user.UserQueryOptions{
		IncludeRoles: req.QueryOpt.IncludeRoles,
		PageParam:    req.QueryOpt.PageParam,
	})
	if err != nil {
		return err
	} else if userResult == nil {
		return nil
	}

	//result := &user.UserShowQueryResult{
	//	PageResult: userResult.PageResult,
	//}
	rsp.PageResult = userResult.PageResult
	if len(userResult.Data) == 0 {
		return nil
	}
	roleResult, err := u.roleModel.Query(ctx, user.RoleQueryParam{
		RecordIDs: ToRoleIDs(userResult.Data),
	})
	if err != nil {
		return err
	}
	rsp.Data = userToUserShow(userResult.Data, rolesToMap(roleResult.Data))
	return nil
	//result.Data = userResult.Data.ToUserShows(roleResult.Data.ToMap())
	//return result, nil
}

func (u *UserServer) Create(ctx context.Context, req *user.UserSchema, rsp *user.UserSchema) error {
	if req.Password == "" {
		return errors.ErrUserNotEmptyPwd
	}

	err := u.checkUserName(ctx, req.UserName)
	if err != nil {
		return err
	}

	err = u.checkUserEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	fmt.Println("role id is ", req.Roles[0].RoleID)

	req.Password = util.SHA1HashString(req.Password)
	req.RecordID = util.MustUUID()
	err = u.userModel.Create(ctx, *req)
	if err != nil {
		return err
	}
	rsp.RecordID = req.RecordID
	rsp.UserName = req.UserName
	//rsp.Password = req.Password
	rsp.Email = req.Email
	rsp.Roles = req.Roles
	rsp.Status = req.Status

	return nil
	//item.Password = util.SHA1HashString(item.Password)
	//item.RecordID = util.MustUUID()
	//err = a.UserModel.Create(ctx, item)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return a.getUpdate(ctx, item.RecordID)
}

func (u *UserServer) Delete(ctx context.Context, req *user.DeleteUserRequest, rsp *user.NullResult) error {

	oldItem, err := u.userModel.Get(ctx, req.Uid)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	err = u.userModel.Delete(ctx, req.Uid)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServer) Get(ctx context.Context, req *user.GetUserRequest, rsp *user.UserSchema) error {

	item, err := u.userModel.Get(ctx, req.Uid, user.UserQueryOptions{
		IncludeRoles: req.QueryOpt.IncludeRoles,
	})
	if err != nil {
		return err
	} else if item == nil {
		return errors.ErrNotFound
	}

	rsp.RecordID = item.RecordID
	rsp.UserName = item.UserName
	//rsp.Password = req.Password
	rsp.Email = item.Email
	rsp.Roles = item.Roles
	rsp.Status = item.Status
	//item, err := a.UserModel.Get(ctx, recordID, opts...)
	//if err != nil {
	//	return nil, err
	//} else if item == nil {
	//	return nil, errors.ErrNotFound
	//}
	//return item, nil
	return nil
}

func (u *UserServer) Update(ctx context.Context, req *user.UpdateUserRequest, rsp *user.UserSchema) error {

	oldItem, err := u.userModel.Get(ctx, req.Uid)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.UserName != req.User.UserName {
		err := u.checkUserName(ctx, req.User.UserName)
		if err != nil {
			return err
		}
	} else if oldItem.Email != req.User.Email {
		err := u.checkUserEmail(ctx, req.User.Email)
		if err != nil {
			return err
		}
	}

	if req.User.Password != "" {
		req.User.Password = util.SHA1HashString(req.User.Password)
	}

	err = u.userModel.Update(ctx, req.Uid, *req.User)
	if err != nil {
		return err
	}

	nitem, err := u.userModel.Get(ctx, req.Uid, user.UserQueryOptions{
		IncludeRoles: true,
	})
	if err != nil {
		return err
	}

	rsp.RecordID = nitem.RecordID
	rsp.UserName = nitem.UserName
	//rsp.Password = req.Password
	rsp.Email = nitem.Email
	rsp.Roles = nitem.Roles
	rsp.Status = nitem.Status

	return nil
	//err = a.UserModel.Update(ctx, recordID, item)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return a.getUpdate(ctx, recordID)
}

func (u *UserServer) UpdataStatus(ctx context.Context, req *user.UpdateUserStatusRequest, rsp *user.NullResult) error {

	oldItem, err := u.userModel.Get(ctx, req.Uid, user.UserQueryOptions{
		IncludeRoles: true,
	})
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}
	err = u.userModel.UpdateStatus(ctx, req.Uid, int(req.Status))
	if err != nil {
		return err
	}
	//if status == 2 {
	//	a.Enforcer.DeleteUser(recordID)
	//} else {
	//	err = a.LoadPolicy(ctx, *oldItem)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//return nil
	return nil
}

//func (u *UserServer) getUpdate(ctx context.Context, recordID string) (*schema.User, error) {
//	nitem, err := a.Get(ctx, recordID, schema.UserQueryOptions{
//		IncludeRoles: true,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	err = a.LoadPolicy(ctx, *nitem)
//	if err != nil {
//		return nil, err
//	}
//	return nitem, nil
//}

func (u *UserServer) checkUserName(ctx context.Context, userName string) error {
	result, err := u.userModel.Query(ctx, user.QueryRequest{
		UserName: userName,
	}, user.UserQueryOptions{
		PageParam: &user.PaginationParam{
			PageSize: -1,
		},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.ErrUserNameExists
	}
	return nil
}

func (u *UserServer) checkUserEmail(ctx context.Context, userEmail string) error {
	result, err := u.userModel.Query(ctx, user.QueryRequest{
		Email: userEmail,
	}, user.UserQueryOptions{
		PageParam: &user.PaginationParam{
			PageSize: -1,
		},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.ErrEmailExists
	}
	return nil
}

func ToRoleIDs(a []*user.UserSchema) []string {
	var roleIDs []string
	for _, item := range a {
		roleIDs = append(roleIDs, roleToRoleIDs(item.Roles)...)
	}
	return roleIDs
}

// ToRoleIDs 转换为角色ID列表
func roleToRoleIDs(a []*user.UserRole) []string {
	list := make([]string, len(a))
	for i, item := range a {
		list[i] = item.RoleID
	}
	return list
}

func rolesToMap(a []*user.RoleSchema) map[string]*user.RoleSchema {
	m := make(map[string]*user.RoleSchema)
	for _, item := range a {
		m[item.RecordID] = item
	}
	return m
}

func userToUserShow(a []*user.UserSchema, mroles map[string]*user.RoleSchema) []*user.UserShow {
	list := make([]*user.UserShow, len(a))

	for i, item := range a {
		showItem := &user.UserShow{
			RecordID:  item.RecordID,
			RealName:  item.RealName,
			UserName:  item.UserName,
			Email:     item.Email,
			Phone:     item.Phone,
			Status:    item.Status,
			CreatedAt: item.CreatedAt,
		}

		var roles []*user.RoleSchema
		for _, roleID := range roleToRoleIDs(item.Roles) {
			if v, ok := mroles[roleID]; ok {
				roles = append(roles, v)
			}
		}
		showItem.Roles = roles
		list[i] = showItem
	}

	return list
}

var ProviderSet = wire.NewSet(NewUserServer)
