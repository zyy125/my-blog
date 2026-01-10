package upload

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

// UploadConfig 上传配置
type UploadConfig struct {
	SavePath    string   // 保存路径
	MaxSize     int64    // 最大文件大小（字节）
	AllowedExts []string // 允许的扩展名
}

// DefaultConfig 默认配置
var DefaultConfig = &UploadConfig{
	SavePath:    "./uploads",
	MaxSize:     10 * 1024 * 1024, // 10MB
	AllowedExts: []string{".jpg", ".jpeg", ".png", ".gif", ".webp"},
}

// SaveImage 保存图片
func SaveImage(file *multipart.FileHeader) (string, error) {
	return SaveFile(file, DefaultConfig)
}

// SaveFile 保存文件
func SaveFile(file *multipart.FileHeader, config *UploadConfig) (string, error) {
	// 1. 验证文件大小
	if file.Size > config.MaxSize {
		return "", fmt.Errorf("文件大小超过限制（最大%dMB）", config.MaxSize/1024/1024)
	}
	
	// 2. 验证文件扩展名
	ext := strings.ToLower(path.Ext(file.Filename))
	allowed := false
	for _, allowedExt := range config.AllowedExts {
		if ext == allowedExt {
			allowed = true
			break
		}
	}
	if !allowed {
		return "", fmt.Errorf("不支持的文件类型：%s", ext)
	}
	
	// 3. 生成文件名（按日期分目录）
	now := time.Now()
	datePath := now.Format("2006/01/02")
	filename := fmt.Sprintf("%d%s", now.UnixNano(), ext)
	
	// 4. 创建目录
	savePath := path.Join(config. SavePath, datePath)
	if err := os.MkdirAll(savePath, os.ModePerm); err != nil {
		return "", fmt.Errorf("创建目录失败: %v", err)
	}
	
	// 5. 保存文件
	dst := path.Join(savePath, filename)
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()
	
	out, err := os.Create(dst)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %v", err)
	}
	defer out.Close()
	
	if _, err := io.Copy(out, src); err != nil {
		return "", fmt.Errorf("保存文件失败: %v", err)
	}
	
	// 6. 返回相对路径（用于前端访问）
	return "/" + strings.ReplaceAll(dst, "\\", "/"), nil
}