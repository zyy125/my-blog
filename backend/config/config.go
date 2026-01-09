package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// Config 全局配置结构体
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Admin    AdminConfig    `mapstructure:"admin"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`  // 服务端口
	Mode string `mapstructure:"mode"`  // 运行模式:  debug/release
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

// AdminConfig 管理员配置
type AdminConfig struct {
	Token      string   `mapstructure:"token"`       // 管理员Token
	AllowedIPs []string `mapstructure:"allowed_ips"` // IP白名单（可选）
}

// App 全局配置实例
var App *Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	// 1. 创建 Viper 实例
	v := viper.New()
	
	// 2. 设置配置文件路径和名称
	v.SetConfigFile(configPath)
	
	// 3. 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败:  %w", err)
	}
	
	// 4. 将配置解析到结构体
	if err := v.Unmarshal(&App); err != nil {
		return fmt.Errorf("解析配置文件失败: %w", err)
	}
	
	return nil
}

// DSN 生成 MySQL 连接字符串
func (db *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db.Username,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
	)
}