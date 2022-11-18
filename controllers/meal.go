package controllers

import (
	"fmt"
	"meal_management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MealRepo struct {
	Db *gorm.DB
}

func NewMealController(db *gorm.DB) *MealRepo {
	db.AutoMigrate(&models.Meal{})
	return &MealRepo{Db: db}
}

func (repository *MealRepo) CreateMeal(c *gin.Context) {
	var meal models.Meal

	if c.BindJSON(&meal) == nil {
		err := models.CreateMeal(repository.Db, &meal)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, meal)
	} else {
		fmt.Println(meal)
		c.JSON(http.StatusBadRequest, meal)
	}
}
func (repository *MealRepo) GetMeals(c *gin.Context) {
	var meals []models.Meal
	err := models.GetMeals(repository.Db, &meals)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, meals)
}
func (repository *MealRepo) GetMealByMealday(c *gin.Context) {
	var meals []models.Meal

	meal_day, _ := c.Params.Get("mealday")
	err := models.GetMealByMeadday(repository.Db, &meals, meal_day)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, meals)
}

func (repository *MealRepo) DeleteMealById(c *gin.Context) {
	var meal models.Meal

	id, _ := strconv.Atoi(c.Query("id"))
	err := models.DeleteMealById(repository.Db, &meal, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, meal)
}
