package handlers

import (
	"fmt"
	"net/http"
	"time"
	utils "users/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// RequireAuth middleware ensures that the request is authenticated with a valid JWT token
// @Summary Require Authentication
// @Description Middleware to check if the request is authenticated
// @Tags Authentication
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT token"
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /auth [get]
func RequireAuth(c *gin.Context) {
	//get the cookie from the req
	tokenString, err := c.Cookie("Authorization")
	utils.Logger.Info().Msg("Token string" + tokenString)
	if err != nil {
		utils.Logger.Info().Msg("Authorization not found")
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	//decode or validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(utils.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expiry
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.Logger.Info().Msg("token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub
		//connect to db
		db, err := utils.DBConn(utils.Username, utils.Password, utils.Dbname, utils.Port)
		if err != nil {
			utils.Logger.Err(err).Msg("Couldn't establish db connection")
		}
		defer db.Close()
		var user utils.Users
		db.QueryRow("select * from users where id=?", claims["sub"]).Scan(&user.ID, &user.UserName, &user.Password, &user.Email)
		if err != nil {
			utils.Logger.Err(err).Msg("Error executing query")
			c.JSON(http.StatusBadRequest, "Error executing query:  "+err.Error())
			return
		}
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to the req
		c.Set("user", &user)
		//continue
		c.Next()

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
