package api

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"notify/enum"

	"github.com/fatih/color"
	"github.com/tidwall/gjson"

	"github.com/Fromsko/gouitls/knet"
)

// WeatherObject 天气对象
type WeatherObject struct {
	Local       string
	WeatherID   string
	WeatherInfo struct {
		Text string
		Temp string
	}
	WeatherStatus int64
	WeatherDate   time.Time
}

// GetWeatherID 获取 天气ID
func (w *WeatherObject) GetWeatherID() {
	weather := knet.SendRequest{
		FetchURL: fmt.Sprintf(
			enum.WeatherCityList,
			url.QueryEscape(w.Local),
			enum.WeatherKey,
		),
	}
	weather.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		statusCode := gjson.GetBytes(resp, "code").Int()
		if statusCode != 200 || err != nil {
			color.Red("天气请求失败!")
		} else {
			location := gjson.GetBytes(resp, "location").Array()[0]
			ID := location.Get("id").String()
			w.WeatherID = ID
		}
		w.WeatherStatus = statusCode
	})
}

// GetWeatherInfo 获取天气信息
func (w *WeatherObject) GetWeatherInfo() {
	weather := knet.SendRequest{
		FetchURL: fmt.Sprintf(
			enum.WeatherUrl,
			w.WeatherID,
			enum.WeatherKey,
		),
	}
	weather.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		statusCode := gjson.GetBytes(resp, "code").Int()
		if statusCode != 200 || err != nil {
			color.Red("天气请求失败!")
			w.WeatherInfo.Text = "未知"
			w.WeatherInfo.Temp = "未知"
		} else {
			w.WeatherInfo.Text = gjson.GetBytes(resp, "now.text").String()
			w.WeatherInfo.Temp = gjson.GetBytes(resp, "now.temp").String()
		}
	})
}

// SearchWeather 获取指定地方的天气
func SearchWeather(local string) *WeatherObject {
	weather := &WeatherObject{
		Local: local,
	}
	weather.GetWeatherID()
	weather.GetWeatherInfo()
	return weather
}
