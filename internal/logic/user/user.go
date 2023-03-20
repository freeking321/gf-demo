package user

import (
	"context"
	"demo/internal/model/entity"
	"demo/internal/service"

	"demo/internal/dao"
	"demo/internal/model"
)

type sUser struct {
}

func init() {
	// 注册service层的 register
	user := New()
	service.RegisterUser(user)
}

func New() *sUser {
	return &sUser{}
}

// GetUserList 获取用户列表  logic 里面编写业务逻辑代码
func (s *sUser) GetUserList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error) {
	out = &model.UserGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	m := dao.User.Ctx(ctx) // 实例化模型

	m = m.Where(dao.User.Columns().Id, in.UserId) // 筛选条件

	listModel := m.Page(in.Page, in.Size) //分页

	switch in.Sort { // 排序
	case "id":
		listModel = listModel.OrderDesc(dao.User.Columns().Id)
	case "account":
		listModel = listModel.OrderDesc(dao.User.Columns().Account)
	default:
		listModel = listModel.OrderDesc(dao.User.Columns().Id)
	}

	// 执行sql -- 为了拿count 感觉有些许复杂
	var list []*entity.User
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	if len(list) == 0 {
		return out, nil
	}
	out.Total = len(list)

	// 获取content
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}
