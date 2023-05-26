package main

import (
	handler "users/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/users/userid", handler.GetUser)
	router.GET("/users", handler.ListUser)
	router.PUT("/users/userid", handler.UpdateUser)
	router.DELETE("/users/userid", handler.DeleteUser)
	router.POST("/users", handler.CreateUser)

	router.Run(":8080")

}
