package main

import (
	"github.com/gin-gonic/gin"
	"go-mall/common/logger"
	"go-mall/config"
	"go-mall/middleware"
	"net/http"
)

func main() {
	// 创建一个 Gin 引擎实例
	r := gin.New()
	// 中间件：日志记录和错误恢复
	r.Use(gin.Logger(), middleware.StartTrace(), gin.Recovery())

	// 定义路由：
	// GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GET /config-read
	r.GET("/config-read", func(c *gin.Context) {
		// 从配置文件中读取数据库配置
		database := config.Database
		// 返回数据库配置信息
		c.JSON(http.StatusOK, gin.H{
			"type":     database.Type,
			"max_life": database.MaxLifeTime,
		})
	})
	r.GET("/logger-test", func(c *gin.Context) {
		logger.New(c).Info("logger test", "key", "KEY")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// 启动服务，监听8080端口
	r.Run(":8080")
}
