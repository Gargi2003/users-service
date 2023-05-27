package main

import (
	handler "users/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/signup", handler.Signup)
	router.POST("/users/login", handler.Login)
	router.GET("/users/validate", handler.RequireAuth, handler.Validate)

	router.GET("/users/userid", handler.GetUser)
	router.GET("/users", handler.ListUser)
	router.PUT("/users/userid", handler.UpdateUser)
	router.DELETE("/users/userid", handler.DeleteUser)

	router.Run(":8080")

}
