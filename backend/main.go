package main

import (
	"Toch/core"
	"Toch/define"
	_ "Toch/utils"
	"Toch/web"

	"github.com/go-rod/rod"
)

var ServerPort = ":2000"

func main() {

	go web.AutoTask("0 0 */12 * * ?", func() {
		Browser := core.InitWeb(define.LoginPageURL)
		Img := core.InitImg()
		// 登录
		if loginStatus := Browser.Login(); loginStatus {
			Browser.NextPage()
			Browser.Extract(func(downloadPage *rod.Page, selectName string) {
				// 初始化课表
				cname, doc := core.InitCname(
					core.ReadHTML(
						Browser.Html(downloadPage),
					),
				)
				// 解析数据
				cname.Resolve(doc)
				// 写入文件
				result := cname.WriteFile(
					selectName,
					"法学3班",
				)
				// 存储图片
				core.SaveImg(Img.Create(result))
			})
			// 退出
			Browser.Logout()
			web.RestartServer(ServerPort)
		}
	})

	web.StartServer(ServerPort)
	select {}
}
