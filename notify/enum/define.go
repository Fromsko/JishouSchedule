package enum

import "notify/utils"

// 课表校验
var (
	OneDay     = "http://open.iciba.com/dsapi/?date"
	CnameData  = utils.Conifg.GetString("CnameData")
	CnameImage = utils.Conifg.GetString("CnameImage")
	Template   = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="
	FlowerList = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid="
	TokenURL   = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

// 天气校验
var (
	// WeatherKey 天气Key
	WeatherKey = utils.Conifg.GetString("WeatherKey")
	// WeatherUrl 天气信息
	WeatherUrl = "https://devapi.qweather.com/v7/weather/now?location=%s&key=%s"
	// WeatherCityList 城市信息
	WeatherCityList = "https://geoapi.qweather.com/v2/city/lookup?location=%s&key=%s"
)

// 微信公众平台主体账号
var (
	AppID     = utils.Conifg.GetString("AppID")
	AppSecret = utils.Conifg.GetString("AppSecret")
)

/*

appid: 'wx13870eadc97c1366'
appsecret: '7ffee262abf55f1af9345cf91f980401'
cnamedata: 'http://1.117.154.114:20000/api/v1/get_cname_data?week='
cnameimage: 'http://1.117.154.114:20000/api/v1/get_cname_table?week='
weatherkey: 'f86106fc2f764c34884ab1f26ce427bc'

*/
