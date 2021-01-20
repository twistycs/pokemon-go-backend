package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/twistycs/pokemon-go-backend/imp"

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
