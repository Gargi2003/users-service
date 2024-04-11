package handlers

import (
	"fmt"
	"net/http"
	utils "users/common"

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

	user, ok := c.Get("user")
	fmt.Println(user)
	if !ok {
		utils.Logger.Err(fmt.Errorf("Error getting user"))
	}

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})

}
