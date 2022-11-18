package routes

import (
	"meal_management/auth"
	"meal_management/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MealRoute(route *gin.Engine, db *gorm.DB) {

	mealController := controllers.NewMealController(db)
	jwtMiddleware, _ := auth.InitJwt(db)
	authHelper := auth.InitHelper(db)

	privateRoute := route.Group("/api", authHelper.VerifyToken, jwtMiddleware.MiddlewareFunc())
	{
		privateRoute.POST("/meal", authHelper.CheckRoles([]string{"ADMIN"}), mealController.CreateMeal)
		privateRoute.GET("/meal", authHelper.CheckRoles([]string{"USER", "ADMIN"}), mealController.GetMeals)
		privateRoute.GET("/meal/:mealday", authHelper.CheckRoles([]string{"USER", "ADMIN"}), mealController.GetMealByMealday)
		privateRoute.DELETE("/meal", authHelper.CheckRoles([]string{"ADMIN"}), mealController.DeleteMealById)
	}

}
