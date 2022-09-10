package Controllers

import (
	"net/http"
	"strconv"

	"github.com/ahmadeyamin/gocrud/Core"
	"github.com/ahmadeyamin/gocrud/Models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, c *gin.Context) {

	if !Core.IsFilled(c, "name") || !Core.IsFilled(c, "email") {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"success": false,
		})
		return
	}
	var count int64

	db.Model(&Models.User{}).Where("email = ?", Core.Input(c, "email")).Count(&count)

	if count > 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"success": false,
			"message": "User already exists",
		})
		return
	}
	user := Models.User{
		Name:  Core.Input(c, "name"),
		Email: Core.Input(c, "email"),
	}

	db.Create(&user)

	c.JSON(http.StatusOK, user)
}

func AllUsers(db *gorm.DB, c *gin.Context) {

	users := []Models.User{}
	db.Limit(20).Find(&users)

	c.JSON(http.StatusOK, users)
}

func UpdateUser(db *gorm.DB, c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	var count int64

	user := Models.User{}
	user.ID = uint(id)

	db.Model(&Models.User{}).Find(&user).Count(&count)

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	user.Name = Core.Input(c, "name")
	db.Save(&user)

	c.JSON(http.StatusOK, user)
}

func DeleteUser(db *gorm.DB, c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var count int64

	user := Models.User{}
	user.ID = uint(id)

	db.Model(&Models.User{}).Find(&user).Count(&count)

	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "User not found",
		})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})

}
