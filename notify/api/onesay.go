package api

import (
	"github.com/Fromsko/gouitls/knet"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
	"net/http"
)

// GetEveryDay 获取每日一句
func GetEveryDay() string {
	var equiangular string
	Spider := knet.SendRequest{
		FetchURL: "http://open.iciba.com/dsapi/?date",
	}
	Spider.Send(func(body []byte, c []*http.Cookie, err error) {
		if err != nil {
			color.Red("获取每日一句失败")
			equiangular = "千里之堤, 始于足下。"
			return
		}
		equiangular = gjson.Get(string(body), "note").String()
	})
	return equiangular
}
