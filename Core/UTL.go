package Core

import (
	"github.com/ahmadeyamin/gocrud/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/hls?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Models.User{})

	return db
	// return db
}

func Input(c *gin.Context, key string) string {
	return c.PostForm(key)
}

func IsFilled(c *gin.Context, key string) bool {
	return c.PostForm(key) != ""
}
