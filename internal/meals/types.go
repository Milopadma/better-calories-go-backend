package meals

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	DayID int64  `gorm:"not null" json:"dayId"`
	Name  string `gorm:"not null" json:"name"`
}