package handlers

import (
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Users []struct {
		Username string `json:"username"`
	} `json:"users"`
}

func CreateUser(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusBadRequest, "Error binding req object")
		return
	}
	for _, user := range req.Users {
		if _, err := db.Query("INSERT INTO users (username) VALUES (?)", user.Username); err != nil {
			c.JSON(http.StatusBadRequest, "Error executing query")
			return
		}
	}

	c.JSON(http.StatusAccepted, "User Added Successfully")
}
