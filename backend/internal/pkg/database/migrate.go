package database

import (
	"fmt"
	"github.com/zyy125/my-blog/backend/internal/model"
)

// AutoMigrate 自动迁移所有数据表
func AutoMigrate() error {
	models := []interface{}{
		&model.Category{},
		&model.Tag{},
		&model.Article{},
		&model.Comment{},  
	}
	
	if err := DB.AutoMigrate(models...); err != nil {
		return fmt.Errorf("数据表迁移失败: %w", err)
	}
	
	fmt.Println("✅ 数据表迁移成功")
	return nil
}