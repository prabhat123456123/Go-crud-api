package routes

import (
	"go-crud-api/controllers"
	"go-crud-api/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/", middlewares.AuthMiddleware(),controllers.CreateUser)
		user.GET("/", controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id",
	middlewares.AuthMiddleware(),
	middlewares.AllowRoles("Admin"),
	controllers.DeleteUser,
)

	}
}
