package main

import (
	_ "star/internal/logic"

	"errors"
	_ "star/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"

	"star/internal/cmd"
)

func main() {
	var err error

	// 全局设置i18n
	g.I18n().SetLanguage("zh-CN")
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
