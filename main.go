package main

import (
	"errors"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	_ "star/internal/packed"

	"star/internal/cmd"
)

func main() {
	var err error

	err = connDb()
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}

// connDb检查数据库连接是否正常
func connDb() error {
	err := g.DB().PingMaster()
	if err != nil {
		return errors.New("数据库连接失败")
	}
	return nil
}
