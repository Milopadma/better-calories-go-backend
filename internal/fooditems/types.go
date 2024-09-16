package fooditems

import "gorm.io/gorm"

type FoodItem struct {
	gorm.Model
	Name         string `gorm:"not null" json:"name"`
	Calories     int    `gorm:"not null" json:"calories"`
	Protein      int    `gorm:"not null" json:"protein"`
	Carbohydrate int    `gorm:"not null" json:"carbohydrate"`
	Fat          int    `gorm:"not null" json:"fat"`
	Sugar        int    `gorm:"not null" json:"sugar"`
}