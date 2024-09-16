package days

import (
	"time"

	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	UserID int64     `gorm:"not null" json:"userId"`
	Date   time.Time `gorm:"not null" json:"date"`
}