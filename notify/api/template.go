package api

import (
	"notify/core"
	"notify/utils"
)

// InitTemplateMessage 模板数据
func InitTemplateMessage() map[string]any {
	cnameInfo := GetCnameData()
	weather := SearchWeather("吉首")
	onesay := GetEveryDay()
	cnameInfo["Week"] = map[string]string{"value": utils.GetWeekly()}
	cnameInfo["City"] = map[string]string{"value": weather.Local}
	cnameInfo["Weather"] = map[string]string{"value": weather.WeatherInfo.Text}
	cnameInfo["Temp"] = map[string]string{"value": weather.WeatherInfo.Temp}
	cnameInfo["Onesay"] = map[string]string{"value": onesay}
	return cnameInfo
}

func Register() (*core.ScheduleService, *core.Service) {
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
	return Schedule, Serve
}
