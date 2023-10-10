package api

import (
	"notify/utils"
)

// InitTemplateMessage 模板数据
func InitTemplateMessage() map[string]any {
	cnameInfo := GetCnameData()
	weather := SearchWeather("吉首")
	onesay := GetEveryDay()
	cnameInfo["Week"] = map[string]string{"value": utils.Weekly}
	cnameInfo["City"] = map[string]string{"value": weather.Local}
	cnameInfo["Weather"] = map[string]string{"value": weather.WeatherInfo.Text}
	cnameInfo["Temp"] = map[string]string{"value": weather.WeatherInfo.Temp}
	cnameInfo["Onesay"] = map[string]string{"value": onesay}
	return cnameInfo
}
