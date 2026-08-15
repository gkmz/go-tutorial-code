package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserResponse 是用户接口的响应结构。
type UserResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, UserResponse{ID: 1, Name: c.Param("id")})
	})
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
