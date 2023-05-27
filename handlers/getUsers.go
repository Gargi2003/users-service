package handlers

import (
	"database/sql"
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()
	id := c.Query("id")
	query := "select * from users where id=?"
	rows, err := db.Query(query, id)
	if err != nil {
		utils.Logger.Err(err).Msgf("Error executing query: ", err)
		return
	}
	users := parseRows(rows)
	c.JSON(http.StatusOK, &users)
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
