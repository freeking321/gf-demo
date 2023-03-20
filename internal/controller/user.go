package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/model"
	"demo/internal/service"
)

var User = cUser{}

type cUser struct {}

func (c *cUser) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	// 调用service
	getListRes, err := service.User().GetUserList(ctx, model.UserGetListInput{
		Page : req.Page,
		Size : req.Size,
		UserId: req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &v1.UserListRes{
		List : getListRes.List,
		Page : getListRes.Page,
		Size : getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
