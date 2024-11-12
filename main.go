package main

import (
	"github.com/gin-gonic/gin"
	"go-mall/common/logger"
	"go-mall/config"
	"net/http"
)

func main() {
	// 创建一个 Gin 引擎实例
	r := gin.New()
	// 中间件：日志记录和错误恢复
	r.Use(gin.Logger(), gin.Recovery())

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
		// 测试Zap 初始化的临时代码, 下节课会删掉
		logger.ZapLoggerTest(c)
		// 返回数据库配置信息
		c.JSON(http.StatusOK, gin.H{
			"type":     database.Type,
			"max_life": database.MaxLifeTime,
		})
	})

	// 启动服务，监听8080端口
	r.Run(":8080")
}
