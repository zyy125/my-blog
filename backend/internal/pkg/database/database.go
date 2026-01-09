package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库实例
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dsn string) error {
	var err error
	
	// 配置 GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印 SQL 日志
	}
	
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}
	
	// 获取底层 SQL DB 对象，配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %w", err)
	}
	
	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期
	
	log.Println("数据库连接成功!")
	return nil
}