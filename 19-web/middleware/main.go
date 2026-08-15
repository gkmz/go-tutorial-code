package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestID 为请求设置可追踪的请求 ID。
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Request-ID", time.Now().UTC().Format("20060102150405.000000000"))
		c.Next()
	}
}

// ErrorResponse 是统一错误响应结构。
type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), RequestID())
	r.GET("/error", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{Error: "invalid request"})
	})
	_ = r.Run(":8083")
}
