package mealfooditems

import (
	"github.com/yourusername/yourproject/internal/fooditems"
	"gorm.io/gorm"
)

type MealFoodItem struct {
	gorm.Model
	MealID   int64              `gorm:"not null" json:"mealId"`
	FoodItem fooditems.FoodItem `gorm:"foreignKey:FoodItemID" json:"foodItem"`
	Quantity int                `gorm:"not null" json:"quantity"`
}