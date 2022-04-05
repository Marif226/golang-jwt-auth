package routes

import (
	controller "golang-jwt-auth/controllers"
	middleware "golang-jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	//user gets token after login, so we have to check it and authenticate user
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.Login())

}