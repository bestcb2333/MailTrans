package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/McaxDev/MailTrans/config"
	"github.com/McaxDev/MailTrans/handler"
	"github.com/abiosoft/ishell"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 程序的入口函数
func main() {

	// 读取程序当前所在的路径
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("读取程序所在路径失败", err.Error())
	}

	// 将程序的执行目录更改为当前目录
	if err := os.Chdir(filepath.Dir(exePath)); err != nil {
		log.Fatal("更改程序基准目录失败", err.Error())
	}

	// 加载配置文件
	if err := config.ConfigInit(); err != nil {
		log.Fatal("配置文件加载失败", err.Error())
	}

	// 启动后端
	go func() {
		r := gin.Default()

		// 允许CORS跨域
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"*"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		//设置路由
		r.GET("/all", handler.AllMail)
		r.GET("/mail", handler.SpecificMail)
		r.Run(":" + config.Config.Port)
	}()

	// 启动命令行工具
	shell := ishell.New()
	shell.Run()
}
