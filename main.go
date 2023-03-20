package main

import (
	"demo/internal/cmd"
	_ "demo/internal/logic"
	_ "demo/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 总入口 只有这一个就行了
	cmd.Main.Run(gctx.New())
}
