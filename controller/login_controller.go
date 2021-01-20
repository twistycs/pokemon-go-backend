package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"github.com/twistycs/pokemon-go-backend/services"
)

type LogInController struct {
	userService imp.UserService
}

func LogInConstructor(userService imp.UserService) *LogInController {
	return &LogInController{
		userService: userService,
	}
}

func (logInController *LogInController) LogInController(c *gin.Context) {
	var jsonInputUser models.User
	if err := c.ShouldBindJSON(&jsonInputUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if jsonInputUser.UserName == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Username required"})
		return
	}
	resp, err := logInController.userService.GetUserByUserName(jsonInputUser.UserName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	} else if (resp != models.User{}) {
		claims := services.TokenClaim{
			jsonInputUser.UserName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
				Issuer:    jsonInputUser.UserName,
				IssuedAt:  time.Now().Unix(),
			},
		}
		token, err := services.GenerateToken(claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, token)
		}
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, "User not found")
		return
	}
}
