package handlers

import (
	"net/http"
	"time"
	utils "users/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Login handles the user login request
// @Summary User Login
// @Description Logs in a user and returns a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body CreateRequest true "User login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {string} string
// @Router /login [post]
func Login(c *gin.Context) {
	//connect to db
	db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
	if err != nil {
		utils.Logger.Err(err).Msg("Couldn't establish db connection")
	}
	defer db.Close()

	// get email/password from req body
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		utils.Logger.Err(err).Msg("Error binding req object")
		c.JSON(http.StatusBadRequest, "Error binding req object")
		return
	}
	//look up requested user
	var userId string
	err1 := db.QueryRow("select id from users where username=?", req.Username).Scan(&userId)
	if err1 != nil {
		utils.Logger.Err(err1).Msg("Error executing query")
		c.JSON(http.StatusBadRequest, "Error executing query:  "+err1.Error())
		return
	}
	// compare sent in pass with saved user pass in db
	var actualPassword []byte
	db.QueryRow("select password from users where id=?", userId).Scan(&actualPassword)
	if err := bcrypt.CompareHashAndPassword(actualPassword, []byte(req.Password)); err != nil {
		utils.Logger.Err(err).Msg("Invalid username or password")
		c.JSON(http.StatusBadRequest, "Invalid username or password:  "+err.Error())
		return
	}
	expirationTime := time.Now().Add(time.Minute * 30).Unix()

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": expirationTime,
	})
	// Convert the token expiration time to a string
	expirationString := time.Unix(expirationTime, 0).String()

	// Log the token expiration
	utils.Logger.Info().Msgf("Token expiration: %s", expirationString)

	//sign and get the entire token using a secret key
	tokenString, err := token.SignedString([]byte(utils.SecretKey))
	if err != nil {
		utils.Logger.Err(err).Msg("Unable to create token")
		c.JSON(http.StatusBadRequest, "Unable to create token")
		return
	}

	//set cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 1800, "", "", false, true)

	//send it back
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"token":  tokenString,
	})

}
