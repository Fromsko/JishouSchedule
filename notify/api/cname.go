package api

import (
	"fmt"
	"net/http"
	"strings"

	"notify/enum"
	"notify/utils"

	"github.com/Fromsko/gouitls/knet"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

// GetCnameData 获取课表数据
func GetCnameData() (result map[string]any) {
	// 存储数据
	result = make(map[string]any)
	// 获取数据
	Spider := knet.SendRequest{
		FetchURL: enum.CnameData + utils.Week,
	}
	Spider.Send(func(resp []byte, Cookies []*http.Cookie, err error) {
		// 判断是否请求成功
		if statusCode := gjson.GetBytes(resp, "code").Int(); statusCode != 200 || err != nil {
			color.Red("课表数据获取失败!")
			return
		}

		// 获取周次
		weekInfo := gjson.GetBytes(resp, "data.周次").String()
		result["周次"] = map[string]string{"value": weekInfo}

		// 遍历本周数据
		for key, value := range gjson.GetBytes(resp, "data.课程信息.课程数据."+utils.Weekly).Map() {
			if value.String() != "没课哟" {
				course := fmt.Sprintf("%s %s %s",
					value.Get("课程名"),
					strings.Split(value.Get("老师").String(), "(")[0],
					value.Get("教室"),
				)
				result[key] = map[string]string{"value": course}
			} else {
				result[key] = map[string]string{"value": value.String()}
			}
		}
	})
	return result
}
