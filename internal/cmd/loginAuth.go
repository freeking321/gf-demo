package cmd

import (
	"demo/internal/consts"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func StartGToken() (gfToken *gtoken.GfToken, err error) {
	gfToken = &gtoken.GfToken{
		CacheMode:       consts.CacheModeRedis,    //token 缓存模式
		ServerName:      consts.BackendServerName, //服务名称
		LoginPath:       "/login",                 // 这里应该会路由到 LoginBeforeFunc
		LoginBeforeFunc: loginFunc,                //认证
		//LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/user/logout",
		AuthPaths:        g.SliceStr{"/v1/hello"},
		AuthExcludePaths: g.SliceStr{"/v1/user/list"}, // 不拦截路径 /user/info,/system/user/info,/system/user,
		//AuthAfterFunc:    authAfterFunc,
		//MultiLogin:       consts.MultiLogin,
	}
	err = gfToken.Start()
	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	// loginPath 路由过来的
	name := r.Get("name").String()
	password := r.Get("password").String()

	//ctx := context.TODO() // 根上下文

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFaulMsg))
		r.ExitAll() // 结束
	}

	// 这里是验证账号密码的逻辑
	// todo

	// 这里返回啥重要么？
	return name, ""
}
