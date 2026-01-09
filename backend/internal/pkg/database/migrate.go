package database

import (
	"fmt"
	"github.com/zyy125/my-blog/backend/internal/model"
)

// AutoMigrate 自动迁移所有数据表
func AutoMigrate() error {
	models := []interface{}{
		&model.Category{},  // ✅ 新增：分类表
		&model.Tag{},       // ✅ 新增：标签表
		&model.Article{},   // 文章表（会自动创建 article_tags 中间表）
	}
	
	if err := DB.AutoMigrate(models... ); err != nil {
		return fmt.Errorf("数据表迁移失败: %w", err)
	}
	
	fmt.Println("✅ 数据表迁移成功")
	return nil
}