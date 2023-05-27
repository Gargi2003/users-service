package handlers

import (
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

type UpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UpdateUser(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	var req UpdateRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusBadRequest, "Error binding req object")
		return
	}
	id := c.Query("id")

	if _, err := db.Query("update users set username=? where id=?", req.Username, id); err != nil {
		c.JSON(http.StatusBadRequest, "Error executing query")
		return
	}

	c.JSON(http.StatusAccepted, "User Updated Successfully")

}
