package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Validate godoc
// @Summary Validate User
// @Description Validates the user information
// @Tags Authentication
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} LoginResponse
// @Router /validate [get]
func Validate(c *gin.Context) {

	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

}
