package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// api/v1/ 这里是定义req 和 res 的地方
// tag没搞懂啥意思

type UserListReq struct {
	g.Meta `path:"/user/list" tags:"UserService" method:"get" summary:"获取用户列表的请求request"` // 定义接口路由和请求方式
	Page int `v:"required|length:1,16"`  // 定义接收参数和验证器
	Size int `v:"required|length:1,16"`  // 定义接收参数和验证器
	UserId int `json:"user_id" v:"required"`
}

type UserListRes struct { //定义返回结构
	List interface{} `json:"list"`
	Page  int `json:"page"`  // 分页码
	Size  int `json:"size"`  // 分页数量
	Total int `json:"total"` // 数据总数
}


