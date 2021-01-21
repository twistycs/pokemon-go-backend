package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/twistycs/pokemon-go-backend/controller"
	database "github.com/twistycs/pokemon-go-backend/databases"
	"github.com/twistycs/pokemon-go-backend/middlewares"
	"github.com/twistycs/pokemon-go-backend/repositories"
	"github.com/twistycs/pokemon-go-backend/services"
)

func SetUpRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	userRepo := repositories.UserRepositories(database.DB)
	userService := services.NewUserService(userRepo)
	userController := controller.UserControllerInit(userService)

	logInController := controller.LogInConstructor(userService)

	login := r.Group("/v1/login")
	{
		login.POST("/", logInController.LogInController)
	}

	user := r.Group("/v1/users")
	{
		user.POST("/", userController.InsertUserController)
	}

	user.Use(middlewares.AuthorizeBearer()) //authen token
	{
		user.GET("/", userController.GetAllUserController)
	}

	// {
	// 	user.GET("/:id", userController.GetUserByIdController)
	// }

	// {
	// 	user.PUT("/:id", userController.UpdateUserController)
	// }

	// {
	// 	user.DELETE("/:id", userController.DeleteUserByIdController)
	// }
	return r
}
