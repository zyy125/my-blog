package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
	"github.com/zyy125/my-blog/backend/internal/pkg/upload"
)

// UploadHandler 上传控制器
type UploadHandler struct{}

// NewUploadHandler 创建上传控制器实例
func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// UploadImage 上传图片
// POST /api/admin/upload/image
func (h *UploadHandler) UploadImage(c *gin.Context) {
	// 1. 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, "请选择要上传的文件")
		return
	}

	// 2. 保存文件
	filePath, err := upload.SaveImage(file)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	// 3. 返回文件URL
	response.Success(c, gin.H{
		"url": filePath,
	})
}
