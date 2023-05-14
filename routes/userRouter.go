package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/thaparoyal/golang-jwt/controllers"
	middleware "github.com/thaparoyal/golang-jwt/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticated())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
