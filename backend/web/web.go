package web

import (
	"Toch/define"
	"Toch/utils"
	"Toch/web/middleware"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var (
	router *gin.Engine
	server *http.Server
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	utils.Log.Info("程序启动成功🚀")
	utils.Log.Info("当前版本: " + define.VERSION)
	utils.Log.Info("项目地址: https://github.com/Fromsko/Jishouschedule")
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

// StartServer 启动Web服务
func StartServer(port string) {
	// 创建Gin路由
	router = gin.Default()
	router.NoRoute(returnHome)
	router.Use(middleware.Cors())

	// 创建路由分组
	api := router.Group("/api/v1")

	// 处理查询参数/api/v1/get_cname_data
	api.GET("/get_cname_data", getCnameData)

	// 处理查询参数/api/v1/get_cname_table
	api.GET("/get_cname_table", getCnameTable)

	// 创建HTTP服务器
	server = &http.Server{
		Addr:    port,
		Handler: router,
	}

	// 启动Web服务
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			utils.Log.Errorf("Failed to start server: %s", err)
		}
	}()
}

// RestartServer 重启服务
func RestartServer(port string) {
	// 关闭当前服务器
	if server != nil {
		if err := server.Shutdown(nil); err != nil {
			fmt.Printf("Failed to shutdown server: %s", err)
		}
	}

	// 启动新的服务器
	StartServer(port)
}

func returnHome(ctx *gin.Context) {
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

// 处理查询参数/api/v1/get_cname_data
func getCnameData(c *gin.Context) {
	weekStr := c.Query("week")
	week, err := strconv.Atoi(weekStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
		return
	}

	// 读取指定目录下的JSON文件
	data, err := readJSONData(week)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read JSON data"})
		return
	}

	c.JSON(http.StatusOK, data)
}

// 处理查询参数/api/v1/get_cname_table
func getCnameTable(c *gin.Context) {
	weekStr := c.Query("week")
	week, err := strconv.Atoi(weekStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
		return
	}

	// 生成图片并返回
	imagePath, err := generateImage(week)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate image"})
		return
	}

	c.File(imagePath)
}

func readJSONData(week int) (result map[string]any, err error) {
	search := fmt.Sprintf("第%d周", week)
	err = utils.ReadFilesWithCallback(
		utils.GenPath("data", ""),
		search,
		func(filePath string) error {
			content, _ := os.ReadFile(filePath)
			_ = json.Unmarshal(content, &result)

			if result == nil {
				result = map[string]any{
					"search": search,
					"error":  "No data found for this week.",
				}
			}

			return nil
		},
	)

	return result, err
}

func generateImage(week int) (imagePath string, err error) {
	search := fmt.Sprintf("第%d周", week)
	err = utils.ReadFilesWithCallback(
		utils.GenPath("output", ""),
		search,
		func(filePath string) (err error) {
			fmt.Println(filePath)
			if filePath == "" {
				return errors.New("没找到")
			}
			imagePath = filePath
			return nil
		},
	)
	return imagePath, err
}
