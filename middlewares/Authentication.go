package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/services"
)

func AuthorizeBearer() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.GetHeader("Authorization")
		tokenFromHeader := strings.Split(bearToken, " ")
		fmt.Println("Token After Split : " + tokenFromHeader[1])
		if len(tokenFromHeader) == 2 {
			token, err := services.VerifyToken(tokenFromHeader[1])
			fmt.Println(token)
			if token != nil && token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
			} else {
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err.Error()})
				return
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Bearer wrong format"})
		}
	}
}
