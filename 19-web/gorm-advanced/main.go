package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Author 是文章作者模型。
type Author struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Posts []Post
}

// Post 是文章模型，展示一对多关联和软删除。
type Post struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	AuthorID  uint
	DeletedAt gorm.DeletedAt
}

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&Author{}, &Post{}); err != nil {
		log.Fatal(err)
	}
	author := Author{Name: "老墨", Posts: []Post{{Title: "Go 事务"}, {Title: "GORM 查询"}}}
	if err := db.Transaction(func(tx *gorm.DB) error { return tx.Create(&author).Error }); err != nil {
		log.Fatal(err)
	}
	var loaded Author
	if err := db.Preload("Posts").First(&loaded, author.ID).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(loaded.Name, len(loaded.Posts))
}
