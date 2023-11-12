package main

import (
	"notify/api"
	"notify/plugin"
	_ "notify/statik"
)

// 程序入口
func main() {

	go plugin.AutoTask("0 0 7 * * ?", func() {
		// 注册通知
		Schedule, Serve := api.Register()
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
