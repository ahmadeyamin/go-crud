package main

import (
	"net/http"

	"github.com/ahmadeyamin/gocrud/Controllers"
	"github.com/ahmadeyamin/gocrud/Core"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := Core.DBInit()

	router.GET("/users", func(c *gin.Context) {
		Controllers.AllUsers(db, c)
	})

	router.POST("/users/create", func(c *gin.Context) {
		Controllers.CreateUser(db, c)
	})

	router.POST("/users/update/:id", func(c *gin.Context) {
		Controllers.UpdateUser(db, c)
	})

	router.DELETE("/users/delete/:id", func(c *gin.Context) {
		Controllers.DeleteUser(db, c)
	})

	http.ListenAndServe(":8080", router)

}
