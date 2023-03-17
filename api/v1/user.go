package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api/v1/ 这里是定义req 和 res 的地方
// tag没搞懂啥意思

type UserListReq struct {
	g.Meta `path:"/user/list" tags:"UserService" method:"get" summary:"获取用户列表的请求request"`
}

type UserListRes struct {
	g.Meta `mime:"text/html" example:"string"`
}


