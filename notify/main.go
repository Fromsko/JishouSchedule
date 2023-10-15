package main

import (
	"notify/api"
	"notify/core"
	"notify/plugin"
	_ "notify/statik"
)

// 程序入口
func main() {
	Schedule, Serve := core.NewRegister(&core.Server{
		TemplateID: "tCwXBXHh1f9m5SOiLkPvtx6-rdPQk1mSnZPbYSBu2Kw",
		RegisterServerList: []core.RegisterServer{
			{
				NickName: "小茹",
				UserID:   "oyDlz6OhrlXhuk0NOvlOOiyNeW9c",
			},
			{
				NickName: "自己",
				UserID:   "oyDlz6NDeZZ0yGE6KRH_Nj_XwNnQ",
			},
		},
	})

	go plugin.AutoTask("0 0 7 * * ?", func() {
		// 获取 Token
		Serve.GetToken()
		// 获取数据
		Serve.TempInfo = api.InitTemplateMessage()
		// 推送任务
		Schedule.PushSchedule(*Serve)
	})

	go plugin.HtmlServer(
		":80",
		// ":443",
	)
	select {}
}
