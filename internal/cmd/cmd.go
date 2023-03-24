package cmd

import (
	"context"
	"demo/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller"
)

// 这个文件是注册路由的地方

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动gtoken
			gfToken, err := StartGToken()
			//开放平台api --这个分组可以写成一个分组树
			s.Group("/v1", func(group *ghttp.RouterGroup) {
				// 这里可以为路由组 使用多个中间件
				group.Middleware(ghttp.MiddlewareHandlerResponse) // 给定默认的返回
				group.Middleware(service.Middleware().CORS)       // 自定义的一个跨域中间件

				gfToken.Middleware(ctx, group) // 分组使用gftoken, 需要认证分组下的所有路由
				group.Bind(
					controller.Hello,
					controller.User,
				)
			})

			// 单独定义一个路由组
			s.Group("/user", func(group *ghttp.RouterGroup) {

			})

			s.Run()
			return nil
		},
	}
)
