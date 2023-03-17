package main

import (
	"context"
	"demo/internal/controller"
	_ "demo/internal/packed"
	"fmt"
	_ "github.com/gin-gonic/gin"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/hello" method:"get"`
	Name string `v:"required" dc:"your name"`
}

type HelloRes struct {
	Reply string `dc:"reply content"`
}

type Hello struct {}

// 这个function 我还不太理解 应该是接口？

func (Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	res = &HelloRes{
		Reply: fmt.Sprintf(`hi %s`, req),
	}

	return
}

func main() {
	//cmd.Main.Run(gctx.New())

	s := g.Server()

	s.Use(ghttp.MiddlewareHandlerResponse) // 如果 response 没有定义返回格式 中间件给定一个默认的格式
	s.BindHandler("/ping", func(r *ghttp.Request) {
		r.Response.Write("hello world !")
	})

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(new(Hello))
		group.Bind(controller.Hello)  // 这里绑定注册路由？


	})

	s.Run()
}
