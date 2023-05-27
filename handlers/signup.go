package handlers

import (
	"net/http"
	utils "users/common"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(c *gin.Context) {

	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	//get email/password from req body
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusBadRequest, "Error binding req object")
		return
	}

	//hash the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		utils.Logger.Err(err).Msg("Error hashing password")
		c.JSON(http.StatusInternalServerError, "Error hashing password")
		return
	}
	//create user
	if _, err := db.Query("INSERT INTO users (username,password) VALUES (?,?)", req.Username, string(hashedPass)); err != nil {
		utils.Logger.Err(err).Msg("Error executing query")
		c.JSON(http.StatusBadRequest, "Error executing query:  "+err.Error())
		return
	}

	//respond
	c.JSON(http.StatusAccepted, "User Added Successfully")
}
