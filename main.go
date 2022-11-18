package main

import (
	"log"
	"meal_management/database"
	"meal_management/routes"
	"meal_management/swagger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello!! Welcome to meal management app",
		})
	})
	//.../docs route for swagger documentation
	swagger.SwaggerDocRoute(route)

	db := database.InitDb()

	routes.UserRoute(route, db)
	routes.UserRoleRoute(route, db)
	routes.MealRoute(route, db)

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
