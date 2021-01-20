package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {

	// userRepo := repositories.RepositoriesConstruct(database.DB)
	// userService := services.NewService(userRepo)
	// userController := controller.UserControllerInit(userService)

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("Have request coming")
	})

	user := r.Group("/v1/users")
	{
		user.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusNotFound, "ID : Not Found")
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
