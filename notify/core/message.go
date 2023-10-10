package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"notify/enum"
	"notify/utils"

	"github.com/Fromsko/gouitls/knet"
	"github.com/fatih/color"
	"github.com/tidwall/gjson"
)

type Service struct {
	AccessToken string           // Token
	Template    *TemplateMessage // 模板
	TempInfo    map[string]any   // 模板信息
}

type TemplateMessage struct {
	ToUser     string         `json:"touser"`
	TemplateID string         `json:"template_id"`
	Url        string         `json:"url"`
	Data       map[string]any `json:"data"`
}

func (s *Service) SendMsg(msg map[string]any, callBack func(resp string)) {
	s.Template.Data = msg

	content, err := json.Marshal(s.Template)
	if err != nil {
		color.Red(err.Error())
	}

	Spider := knet.SendRequest{
		Method:   "POST",
		FetchURL: enum.Template + s.AccessToken,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1 Edg/117.0.0.0",
		},
		Data: strings.NewReader(string(content)),
	}

	Spider.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		if err != nil {
			color.Red(err.Error())
			return
		}
		callBack(string(resp))
	})
}

func (s *Service) GetToken() {
	Spider := knet.SendRequest{
		FetchURL: fmt.Sprintf(
			enum.TokenURL,
			enum.AppID,
			enum.AppSecret,
		),
	}
	Spider.Send(func(resp []byte, cookies []*http.Cookie, err error) {
		statusCode := gjson.GetBytes(resp, "errcode").Int()
		if statusCode != 0 || err != nil {
			color.Red("获取Access Token失败，错误代码: %d\n", statusCode)
			return
		}
		s.AccessToken = gjson.GetBytes(resp, "access_token").String()
	})
}

func InitServe(templateID string) *Service {
	s := &Service{
		Template: &TemplateMessage{
			TemplateID: templateID,
			Data:       make(map[string]any),
			Url:        enum.CnameImage + utils.Week,
		},
	}
	return s
}
