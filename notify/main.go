package main

import (
	"notify/api"
	"notify/core"
	"notify/utils"

	"github.com/robfig/cron/v3"
)

// 程序入口
func main() {
	Schedule, Serve := core.NewRegister(
		"tCwXBXHh1f9m5SOiLkPvtx6-rdPQk1mSnZPbYSBu2Kw",
		core.RegisterServer{
			NickName: "小茹",
			UserID:   "oyDlz6OhrlXhuk0NOvlOOiyNeW9c",
		},
		core.RegisterServer{
			NickName: "自己",
			UserID:   "oyDlz6NDeZZ0yGE6KRH_Nj_XwNnQ",
		},
	)

	AutoTask("0 0 7 * * ?", func() {
		// 获取 Token
		Serve.GetToken()
		// 获取数据
		Serve.TempInfo = api.InitTemplateMessage()
		// 推送任务
		Schedule.PushSchedule(*Serve)
	})
}

// 自动任务
func AutoTask(Timer string, Task func()) {
	c := cron.New(cron.WithSeconds())

	// 每天早晨7:00
	if _, err := c.AddFunc(Timer, Task); err != nil {
		utils.Log.Debugf("添加任务时出错：%v", err)
		return
	} else {
		c.Start()
		utils.Log.Info("程序启动成功🚀")
		utils.Log.Info("项目地址: https://github.com/Fromsko/Jishouschedule")
	}
	select {}
}
