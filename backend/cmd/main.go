package main

import (
	"fmt"
	"log"

	"github.com/zyy125/my-blog/backend/config"
	"github.com/zyy125/my-blog/backend/internal/pkg/database"
	"github.com/zyy125/my-blog/backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// ========== 1. 加载配置 ==========
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log. Fatalf("配置加载失败: %v", err)
	}
	fmt.Println("✅ 配置加载成功")

	// ========== 2. 初始化数据库 ==========
	dsn := config.App.Database.DSN()
	if err := database.InitDB(dsn); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	fmt.Println("✅ 数据库连接成功")

	// ========== 3. 自动迁移 ==========
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("数据表迁移失败: %v", err)
	}
	fmt.Println("✅ 数据表迁移成功")

	// ========== 4. 设置路由 ==========
	gin.SetMode(config.App. Server.Mode)
	r := router.SetupRouter()

	// ========== 5. 启动服务器 ==========
	fmt.Printf("🚀 服务器启动在 http://localhost%s\n", config.App.Server.Port)
	if err := r.Run(config.App.Server.Port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}