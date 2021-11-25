package models

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	ID        int    `gorm:"primarykey"`
	Type      string `gorm:"type:varchar(100);not null" json:"type" form:"type"`
	Name      string `gorm:"type:varchar(255)" json:"name" form:"name"`
	Cvv       int    `gorm:"type:int(10)" json:"cvv" form:"cvv"`
	Month     string `gorm:"type:varchar(100)" json:"month" form:"month"`
	Year      string `gorm:"type:varchar(100)" json:"year" form:"year"`
	Number    string `gorm:"type:varchar(100)" json:"number" form:"number"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
