package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS 只允许配置中的前端来源，生产环境不要使用任意来源。
func CORS(allowedOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Origin") == allowedOrigin {
			c.Header("Access-Control-Allow-Origin", allowedOrigin)
			c.Header("Vary", "Origin")
		}
		c.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.New()
	r.Use(CORS("http://localhost:3000"))
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	_ = r.Run(":8084")
}
