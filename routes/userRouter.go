package routes

import (
	controller "RestaurantManagementAppGO/controller"
    middleware "RestaurantManagementAppGO/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.Login())


	incomingRoutes.GET("/users", middleware.Authentication(),controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",middleware.Authentication(),controller.GetUser())
	incomingRoutes.DELETE("/users/delete/:user_id",middleware.Authentication(),controller.DeleteUser())
}