package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Meal struct {
	gorm.Model
	Vegetable string `json:"vegetable"`
	Vorta     string `json:"vorta"`
	Meat      string `json:"meat"`
	Fish      string `json:"fish"`
	ExtraItem string `json:"extraitem"`
	Price     int    `json:"price"`
	MealDay   string `json:"mealday"`
}

func CreateMeal(db *gorm.DB, newMeal *Meal) (err error) {
	err = db.Create(newMeal).Error
	if err != nil {
		return err
	}
	return nil
}

func GetMeals(db *gorm.DB, meals *[]Meal) (err error) {
	err = db.Find(meals).Error
	if err != nil {
		return err
	}
	return nil
}
func GetMealByMeadday(db *gorm.DB, meals *[]Meal, meal_day string) (err error) {
	err = db.Where("meal_day = ?", meal_day).Find(meals).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteMealById(db *gorm.DB, meal *Meal, id int) (err error) {
	fmt.Println("id debug", id)
	tx := db.Begin()
	if tx.Where("id = ?", id).Delete(&Meal{}); tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return tx.Commit().Error
}
