package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type InfoReq struct {
	g.Meta `path:"account/info" method:"get" sm:"获取信息" tags:"用户"`
}

type InfoRes struct {
	Username string      `json:"username" dc:"用户名"`
	Password string      `json:"password" dc:"邮箱"`
	CreateAt *gtime.Time `json:"create_at" dc:"创建时间"`
	UpdateAt *gtime.Time `json:"update_at" dc:"更新时间"`
}
