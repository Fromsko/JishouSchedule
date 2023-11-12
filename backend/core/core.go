package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Option func() (doc *goquery.Document, err error)

// WebObject 浏览器对象
type WebObject struct {
	BaseUrl  string
	Page     *rod.Page
	DownPage *rod.Page
	Browser  *rod.Browser
}

// CnameObject 课程总数据
type CnameObject struct {
	CnameResult  map[string]any `json:"课程信息"`
	CnameSpecial string         `json:"备注"`
	Cname        string         `json:"班级"`
	Weekly       string         `json:"周次"`
}

// InitWeb 初始化浏览器
func InitWeb(url string) (Web *WebObject) {
	//启动浏览器
	u := launcher.New().MustLaunch()
	// 创建浏览器实例
	browser := rod.New().ControlURL(u).MustConnect()
	// 新建一个页面
	page := browser.MustPage()

	return &WebObject{
		BaseUrl: url,
		Page:    page,
		Browser: browser,
	}
}

// InitCname 初始化课表对象
func InitCname(option Option) (cname *CnameObject, doc *goquery.Document) {
	doc, err := option()
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}

	return &CnameObject{
		CnameSpecial: "",
		CnameResult:  map[string]any{},
	}, doc
}

// ReadHTML 读取 Html 文件
func ReadHTML(content string) Option {
	reader := strings.NewReader(content)

	return func() (doc *goquery.Document, err error) {
		doc, err = goquery.NewDocumentFromReader(reader)
		if err != nil {
			return nil, fmt.Errorf("解析 HTML 文档时出错: %s", err)
		}
		return doc, nil
	}
}

// ReadFile 读取文件
func ReadFile(fileName string) Option {
	html, err := os.ReadFile(fileName)
	if err != nil {
		log.Errorf("无法读取HTML文件: %s", err)
	}

	return func() (doc *goquery.Document, err error) {
		return ReadHTML(string(html))()
	}
}
