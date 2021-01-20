package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/controller"
	database "github.com/twistycs/pokemon-go-backend/databases"
	"github.com/twistycs/pokemon-go-backend/middlewares"
	"github.com/twistycs/pokemon-go-backend/repositories"
	"github.com/twistycs/pokemon-go-backend/services"
)

func SetUpRoutes() *gin.Engine {

	userRepo := repositories.UserRepositories(database.DB)
	userService := services.NewUserService(userRepo)
	// userController := controller.UserControllerInit(userService)

	logInController := controller.LogInConstructor(userService)

	r := gin.Default()

	login := r.Group("/v1/login")
	{
		login.POST("/", logInController.LogInController)
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
