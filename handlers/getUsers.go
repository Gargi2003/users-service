package handlers

import (
	"database/sql"
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

// GetUsernameById retrieves the username by ID
// @Summary Get username by ID
// @Description Retrieve the username by providing the user ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {string} string "Username"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/username [get]
func GetUsernameById(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	userId := c.Query("id")
	rows, err := db.Query("SELECT * from users where id=?", userId)
	if err != nil {
		utils.Logger.Err(err).Msg("Error occurred while executing query")
		c.JSON(http.StatusInternalServerError, "Error occurred while executing query")
		return
	}
	user := utils.Users{}
	for rows.Next() {

		err := rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}
	}
	c.JSON(http.StatusOK, &user.UserName)
}

// ListUser godoc
// @Summary List Users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Success 200 {array} utils.Users
// @Router /users [get]
func ListUser(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	utils.Logger.Debug().Msg("connected to db")

	query := "select * from users"
	rows, err1 := db.Query(query)
	if err1 != nil {
		utils.Logger.Err(err1).Msg("Error executing query")
		c.JSON(http.StatusInternalServerError, "Errror executing query")
		return
	}
	users := parseRows(rows)
	c.JSON(http.StatusOK, &users)
}

func parseRows(rows *sql.Rows) *[]utils.Users {
	var users []utils.Users

	for rows.Next() {
		var user utils.Users
		rows.Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
		users = append(users, user)
	}
	return &users
}
