package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"websocket-chat/initializers"
	"websocket-chat/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const AUTH_PREFIX string = "Bearer "

func errMsg(msg string) gin.H {
	return gin.H{"error": msg}
}

func RequireAuth(c *gin.Context) {
	defaultErrMsg := errMsg("Unauthorized")
	authHeader := c.GetHeader("Authorization")

    if ! strings.HasPrefix(authHeader, AUTH_PREFIX) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, defaultErrMsg)
		return
    }

	tokenString := strings.TrimPrefix(authHeader, AUTH_PREFIX)
	token, err := jwt.Parse(tokenString, keyFunc)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, errMsg(err.Error()))
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, defaultErrMsg)
			return
		}
		var user models.User
		initializers.DB.First(&user,claims["sub"])

		if user.ID == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, defaultErrMsg)
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, defaultErrMsg)
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
		return []byte(os.Getenv("SECRET")), nil
	}

	return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
}
