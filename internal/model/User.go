package model

type UserGetListOutput struct {
	List []UserItem `json:"list"`
	Page  int                      `json:"page"`  // 分页码
	Size  int                      `json:"size"`  // 分页数量
	Total int                      `json:"total"` // 数据总数
	Stats map[string]int // 发布内容数量
}

type UserGetListInput struct {
	Page int
	Size int
	Sort string
	UserId int
}

type UserItem struct {
	Id int `json:"id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
}