package handler

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zyy125/my-blog/backend/config"
	"github.com/zyy125/my-blog/backend/internal/pkg/response"
)

// TokenStore 存储有效的 token
type TokenStore struct {
	tokens map[string]time.Time // token -> 过期时间
	mu     sync.RWMutex
}

var tokenStore = &TokenStore{
	tokens: make(map[string]time.Time),
}

// AuthHandler 认证处理器
type AuthHandler struct{}

// NewAuthHandler 创建认证处理器
func NewAuthHandler() *AuthHandler {
	// 启动清理过期 token 的协程
	go tokenStore.cleanExpiredTokens()
	return &AuthHandler{}
}

// VerifyRequest 验证请求结构
type VerifyRequest struct {
	SecretKey string `json:"secret_key" binding:"required"`
}

// VerifyResponse 验证响应结构
type VerifyResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"` // Unix 时间戳
}

// VerifyKey 验证密钥并返回临时 token
// POST /api/admin/auth/verify
func (h *AuthHandler) VerifyKey(c *gin.Context) {
	var req VerifyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, "请求参数错误")
		return
	}

	// 验证密钥
	if req.SecretKey != config.App.Admin.SecretKey {
		response.Error(c, "密钥错误")
		return
	}

	// 生成临时 token
	token, err := generateToken()
	if err != nil {
		response.Error(c, "生成 Token 失败")
		return
	}

	// 设置 token 过期时间（默认 2 小时）
	expiresAt := time.Now().Add(time.Duration(config.App.Admin.TokenExpireHours) * time.Hour)
	tokenStore.addToken(token, expiresAt)

	response.Success(c, VerifyResponse{
		Token:     token,
		ExpiresAt: expiresAt.Unix(),
	})
}

// CheckToken 检查 token 是否有效（内部使用）
func CheckToken(token string) bool {
	return tokenStore.isValid(token)
}

// generateToken 生成随机 token
func generateToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// addToken 添加 token
func (ts *TokenStore) addToken(token string, expiresAt time.Time) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.tokens[token] = expiresAt
}

// isValid 检查 token 是否有效
func (ts *TokenStore) isValid(token string) bool {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	expiresAt, exists := ts.tokens[token]
	if !exists {
		return false
	}

	// 检查是否过期
	if time.Now().After(expiresAt) {
		return false
	}

	return true
}

// cleanExpiredTokens 定期清理过期 token
func (ts *TokenStore) cleanExpiredTokens() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		ts.mu.Lock()
		now := time.Now()
		for token, expiresAt := range ts.tokens {
			if now.After(expiresAt) {
				delete(ts.tokens, token)
			}
		}
		ts.mu.Unlock()
	}
}
