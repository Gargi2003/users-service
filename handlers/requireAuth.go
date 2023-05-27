package handlers

import (
	"fmt"
	"net/http"
	"time"
	utils "users/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	//get the cookie from the req
	tokenString, err := c.Cookie("Authorization")
	utils.Logger.Info().Msg(tokenString)
	if err != nil {
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
		db.QueryRow("select * from users where id=?", claims["sub"]).Scan(&user.ID, &user.UserName, &user.Password)
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
