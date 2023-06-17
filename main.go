package main

import (
	_ "users/docs"
	handler "users/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Users Service
// @description Users API in go using gin-framework
// @version 1.0
// @host localhost:8080
// @BasePath /api
func main() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/users/signup", handler.Signup)
	router.POST("/users/login", handler.Login)
	router.GET("/users/validate", handler.RequireAuth, handler.Validate)
	router.GET("/users/get", handler.GetUsernameById)
	router.GET("/users", handler.ListUser)
	router.PUT("/users/userid", handler.UpdateUser)
	router.DELETE("/users/userid", handler.DeleteUser)

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:4200"}
	// config.AllowMethods = []string{"POST"}
	// config.AllowHeaders = []string{"Content-Type"}
	// router.Use(cors.New(config))
	router.Run(":8080")

}
