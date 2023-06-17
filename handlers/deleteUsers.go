package handlers

import (
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

// DeleteUser deletes a user by ID
// @Summary Delete a user
// @Description Delete a user by providing the user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {string} string "User Deleted!"
// @Failure 400 {string} string "Error executing query"
// @Failure 404 {string} string "User not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/delete [delete]
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
