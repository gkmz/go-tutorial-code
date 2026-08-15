package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 是 GORM 映射的用户模型。
type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		log.Fatal(err)
	}
	if err := db.Create(&User{Name: "老墨"}).Error; err != nil {
		log.Fatal(err)
	}
	var user User
	if err := db.First(&user, "name = ?", "老墨").Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d %s\n", user.ID, user.Name)
}
