package user

import (
	"caas-micro/internal/app/user/model"
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

var ProviderSet = wire.NewSet(NewUserServer)
