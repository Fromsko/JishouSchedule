package core

import (
	"Toch/define"
	"Toch/utils"
)

var log = utils.Log

// Login 登录
func (web *WebObject) Login() bool {
	loginPage := web.Page

	log.Info("进入登录页面")
	// 导航到目标页面
	loginPage.MustNavigate(define.LoginPageURL)

	// 等待页面加载完毕
	loginPage.MustWaitLoad()
	loginPage.MustWaitStable().MustScreenshot(define.ImgHomePage)

	// 填写登录信息(账号|密码)
	loginPage.MustElement(define.InputLogin).MustInput(utils.Config.Username)
	loginPage.MustElement(define.InputLoginTwo).MustInput(utils.Config.Password)

	// 点击立即登录按钮
	loginPage.MustElement(define.LoginButton).MustClick()

	// 检验是否登录成功
	status, _ := loginPage.Element(define.LoginStatus)
	if text, _ := status.Text(); text != "登录成功" {
		log.Error("登录失败: ", text)
		loginPage.MustScreenshot(define.ImgFailed)
		return false
	}

	// 获取基本数据
	log.Info("登录成功")
	loginPage.MustElement(define.WeatherPage).MustWaitStable().MustScreenshot(define.ImgWeather)
	loginPage.MustElement(define.PeopleInfo).MustWaitStable().MustScreenshot(define.ImgPeopleInfo)
	loginPage.MustScreenshot(define.ImgSuccess)
	return true
}

// Logout 退出
func (web *WebObject) Logout() {
	log.Info("正在退出登录...")

	// 找到下拉框元素
	search, _ := web.Page.Search("设置")
	search.First.MustClick()
	//page.MustElementR(ExitMenu, "设置").MustClick()

	// 退出
	search, _ = web.Page.Search("退出")
	parent, _ := search.First.Parent()
	parentText, _ := parent.Text()
	parent.MustClick()

	log.Infof("成功%s登录!", parentText)

	defer web.Browser.MustClose()
}
