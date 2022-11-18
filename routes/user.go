package routes

import (
	"meal_management/auth"
	"meal_management/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoute(route *gin.Engine, db *gorm.DB) {

	userController := controllers.NewUserController(db)
	jwtMiddleware, _ := auth.InitJwt(db)

	publicRoute := route.Group("/api")
	{
		publicRoute.POST("/register", userController.Register)
		publicRoute.POST("/login", jwtMiddleware.LoginHandler)
	}
}
