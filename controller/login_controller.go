package controller

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"github.com/twistycs/pokemon-go-backend/services"
	"golang.org/x/crypto/bcrypt"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	resp, err := logInController.userService.GetUserByUserName(jsonInputUser.UserName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(jsonInputUser.Password))

	fmt.Println(jsonInputUser.UserName)
	fmt.Println(resp)
	if (err != nil || jsonInputUser.UserName == "" || resp == models.User{}) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errorMessage": "Username or password incorrect."})
		return
	}

	if (resp != models.User{}) {
		claims := services.TokenClaim{
			jsonInputUser.UserName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
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
	}
}
