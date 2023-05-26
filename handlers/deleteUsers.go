package handlers

import (
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	id := c.Query("id")
	rows, err := db.Exec("delete from users where id=?", id)
	if err != nil {
		utils.Logger.Err(err).Msg("Error executing query")
		c.JSON(http.StatusBadRequest, "Error executing query")
		return
	}
	resultSet, err := rows.RowsAffected()
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while getting the number of affected rows")
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if resultSet == 0 {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}
	utils.Logger.Info().Msg("User deleted successfully!")
	c.JSON(http.StatusOK, "User Deleted !!!")
}
