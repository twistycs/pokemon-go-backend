package routes

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/middlewares"
	"github.com/twistycs/pokemon-go-backend/services"
)

func SetUpRoutes() *gin.Engine {

	// userRepo := repositories.RepositoriesConstruct(database.DB)
	// userService := services.NewService(userRepo)
	// userController := controller.UserControllerInit(userService)

	r := gin.Default()

	login := r.Group("/v1/login")
	{
		login.POST("/", func(c *gin.Context) {
			claims := services.TokenClaim{
				"bar",
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
					Issuer:    "Wit",
					IssuedAt:  time.Now().Unix(),
				},
			}
			token, err := services.GenerateToken(claims)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			} else {
				c.JSON(http.StatusOK, token)
			}
		})
	}

	user := r.Group("/v1/users")
	user.Use(middlewares.AuthorizeBearer())
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, "TEST")
		})
	}

	// {
	// 	user.GET("/:id", userController.GetUserByIdController)
	// }

	// {
	// 	user.POST("/", userController.InsertUserController)
	// }

	// {
	// 	user.PUT("/:id", userController.UpdateUserController)
	// }

	// {
	// 	user.DELETE("/:id", userController.DeleteUserByIdController)
	// }

	return r
}
