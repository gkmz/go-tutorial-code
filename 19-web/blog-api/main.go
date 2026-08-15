package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Blog 是博客服务的应用容器。
type Blog struct{ DB *gorm.DB }

// Post 是博客文章模型。
type Post struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
}

// NewBlog 创建使用内存数据库的演示博客服务。
func NewBlog() (*Blog, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&Post{}); err != nil {
		return nil, err
	}
	return &Blog{DB: db}, nil
}

// Router 创建博客 API 路由。
func (b *Blog) Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/healthz", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/posts", func(c *gin.Context) {
		var posts []Post
		if err := b.DB.Find(&posts).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	})
	r.POST("/posts", func(c *gin.Context) {
		var post Post
		if err := c.ShouldBindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := b.DB.Create(&post).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, post)
	})
	return r
}

func main() {
	blog, err := NewBlog()
	if err != nil {
		panic(err)
	}
	_ = blog.Router().Run(":8086")
}
