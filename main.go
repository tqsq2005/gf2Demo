package main

import (
	_ "gf2-demo/internal/logic"
	_ "gf2-demo/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gf2-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
