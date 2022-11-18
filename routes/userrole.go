package routes

import (
	"meal_management/auth"
	"meal_management/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoleRoute(route *gin.Engine, db *gorm.DB) {

	userRoleController := controllers.NewUserRoleController(db)
	jwtMiddleware, _ := auth.InitJwt(db)
	authHelper := auth.InitHelper(db)

	privateRoute := route.Group("/api", authHelper.VerifyToken, jwtMiddleware.MiddlewareFunc())
	{
		// roles
		privateRoute.POST("/addrole", userRoleController.AddRoleToUser)
		privateRoute.POST("/removerole", userRoleController.DeleteUserRole)
	}
}
