package main

import (
	"fmt"
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/config"
)

func main() {
	// ========== 1. 加载配置 ==========
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log. Fatalf("配置加载失败: %v", err)
	}
	
	// 打印配置信息（开发阶段用于调试）
	fmt.Printf("服务器端口: %s\n", config.App.Server.Port)
	fmt.Printf("运行模式: %s\n", config.App.Server.Mode)
	fmt.Printf("数据库地址: %s:%d\n", config.App. Database.Host, config.App. Database.Port)
	
	// ========== 2. 设置Gin模式 ==========
	gin.SetMode(config.App.Server.Mode)
	
	// ========== 3. 创建路由 ==========
	r := gin.Default()
	
	// 测试路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	
	// 新增：显示配置信息的路由（开发测试用）
	r.GET("/config", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"server_mode": config.App.Server.Mode,
			"database":  gin.H{
				"host":    config.App.Database.Host,
				"dbname": config.App.Database.DBName,
			},
		})
	})
	
	// ========== 4. 启动服务器（使用配置的端口）==========
	fmt.Printf("服务器启动在 http://localhost%s\n", config.App.Server.Port)
	if err := r.Run(config.App.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}