package controller

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/twistycs/pokemon-go-backend/imp"
	"github.com/twistycs/pokemon-go-backend/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService imp.UserService
}

func UserControllerInit(userService imp.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
func (u *UserController) GetAllUserController(c *gin.Context) {
	log.Info("Get All USER !!!")
	resp, err := u.userService.GetAllUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *UserController) GetUserByNameController(c *gin.Context) {
	userName := c.Params.ByName("userName")
	resp, err := u.userService.GetUserByUserName(userName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func (u *UserController) InsertUserController(c *gin.Context) {
	var jsonInputUser models.User
	if err := c.ShouldBindJSON(&jsonInputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bytePassword := []byte(jsonInputUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	jsonInputUser.Password = string(hashedPassword)
	jsonInputUser.UserName = strings.ToLower(jsonInputUser.UserName)
	user, err := u.userService.GetUserByUserName(jsonInputUser.UserName)
	if (user != models.User{}) {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Username already existing"})
		return
	}

	err = u.userService.InsertUser(&jsonInputUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.AbortWithStatus(http.StatusCreated)
}
