package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get" summary:"You first hello api" tags:"Hello"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

// api/v1/ 这里是注册路由的地方
