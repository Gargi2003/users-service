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
	Email    string `json:"email"`
}

// Signup creates a new user account
// @Summary Create User
// @Description Create a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body CreateRequest true "User object"
// @Success 202 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /signup [post]
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
	if _, err := db.Query("INSERT INTO users (username,password,email) VALUES (?,?,?)", req.Username, string(hashedPass), req.Email); err != nil {
		utils.Logger.Err(err).Msg("Error executing query")
		c.JSON(http.StatusBadRequest, "Error executing query:  "+err.Error())
		return
	}

	//respond
	c.JSON(http.StatusAccepted, "User Added Successfully")
}
