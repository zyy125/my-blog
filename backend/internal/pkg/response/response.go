package response

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`    // 状态码
	Msg  string      `json:"msg"`     // 提示信息
	Data interface{} `json:"data"`    // 数据
}

// PageData 分页响应结构
type PageData struct {
	List     interface{} `json:"list"`      // 数据列表
	Total    int64       `json:"total"`     // 总记录数
	Page     int         `json:"page"`      // 当前页码
	PageSize int         `json:"page_size"` // 每页数量
}

// Success 成功响应（带数据）
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:  http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}

// SuccessWithMsg 成功响应（自定义消息）
func SuccessWithMsg(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  msg,
		Data: data,
	})
}

// Error 错误响应（400 客户端错误）
func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: nil,
	})
}

// ServerError 服务器错误响应（500）
func ServerError(c *gin.Context, msg string) {
	c.JSON(http.StatusInternalServerError, Response{
		Code: http.StatusInternalServerError,
		Msg:   msg,
		Data: nil,
	})
}

// NotFound 资源不存在响应（404）
func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, Response{
		Code: http.StatusNotFound,
		Msg:  msg,
		Data: nil,
	})
}

// PageSuccess 分页成功响应
func PageSuccess(c *gin. Context, list interface{}, total int64, page, pageSize int) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:   "success",
		Data:  PageData{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		},
	})
}