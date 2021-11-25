package models

import (
	"time"

	"gorm.io/gorm"
)

type Homestay struct {
	ID          int     `gorm:"primarykey"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name" form:"name"`
	Type        string  `gorm:"type:varchar(100);not null" json:"type" form:"type"`
	Description string  `gorm:"type:varchar(255);not null" json:"description" form:"description"`
	Status      string  `gorm:"type:varchar(255);not null" json:"status" form:"status"`
	Price       int     `gorm:"type:int;not null" json:"price" form:"price"`
	Latitude    float64 `gorm:"type:decimal(5,2);not null" json:"latitude" form:"latitude"`
	Longitude   float64 `gorm:"type:decimal(5,2);not null" json:"longitude" form:"longitude"`
	User_ID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PostHomestay struct {
	Name        string  `json:"name" form:"name"`
	Type        string  `json:"type" form:"type"`
	Description string  `json:"description" form:"description"`
	Price       int     `json:"price" form:"price"`
	Latitude    float64 `json:"latitude" form:"latitude"`
	Longitude   float64 `json:"longitude" form:"longitude"`
}
