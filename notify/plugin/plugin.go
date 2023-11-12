package plugin

import (
	"html/template"
	"net/http"
	"notify/enum"
	_ "notify/statik"
	"notify/utils"

	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"github.com/robfig/cron/v3"
)

var StatikFS http.FileSystem

func init() {
	StatikFS, _ = fs.New()
	gin.SetMode(gin.ReleaseMode)
}

// AutoTask 自动任务
func AutoTask(Timer string, Task func()) {
	c := cron.New(cron.WithSeconds())

	// 每天早晨7:00
	if _, err := c.AddFunc(Timer, Task); err != nil {
		utils.Log.Debugf("添加任务时出错：%v", err)
		return
	}

	c.Start()
}

// HtmlServer 页面服务
func HtmlServer(port ...string) {
	Engine := gin.Default()
	Engine.StaticFS("/", StatikFS)
	Engine.NoRoute(renderReturnHome)

	if err := Engine.Run(port...); err == nil {
		utils.Log.Info("程序启动成功🚀")
		utils.Log.Info("当前版本: " + enum.VERSION)
		utils.Log.Info("项目地址: https://github.com/Fromsko/Jishouschedule")
	}
}

func renderReturnHome(ctx *gin.Context) {
	const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>stray</title>
    <style>
        body {
            font-family: 'Arial', sans-serif;
            background-color: #f4f4f4;
            text-align: center;
            padding: 50px;
        }

        .message {
            font-size: 24px;
            color: #333;
            margin-bottom: 20px;
        }

        .button {
            padding: 10px 20px;
            font-size: 16px;
            background-color: #4caf50;
            color: #fff;
            text-decoration: none;
            border-radius: 5px;
        }
    </style>
</head>
<body>
    <div class="message">是不是迷路了, 点击按钮找到回家的路~</div>
    <a href="/" class="button">回家</a>
</body>
</html>
`
	// 解析
	tmpl, err := template.New("return-home").Parse(htmlTemplate)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	// 渲染页面
	data := map[string]interface{}{}
	err = tmpl.Execute(ctx.Writer, data)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
