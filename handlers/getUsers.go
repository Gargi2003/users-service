package handlers

import (
	"database/sql"
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

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

		err := rows.Scan(&user.ID, &user.UserName, &user.Password)
		if err != nil {
			utils.Logger.Err(err).Msg("Error unmarshalling into struct from db")
		}
	}
	c.JSON(http.StatusOK, &user.UserName)
}

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
		utils.Logger.Err(err).Msgf("Error executing query: ", err)
		return
	}
	users := parseRows(rows)
	c.JSON(http.StatusOK, &users)
}

func parseRows(rows *sql.Rows) *[]utils.Users {
	var users []utils.Users

	for rows.Next() {
		var user utils.Users
		rows.Scan(&user.ID, &user.UserName, &user.Password)
		users = append(users, user)
	}
	return &users
}
