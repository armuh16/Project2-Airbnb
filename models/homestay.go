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
	Status      string  `gorm:"type:varchar(255);default:'available';not null" json:"status" form:"status"`
	Price       int     `gorm:"type:int;not null" json:"price" form:"price"`
	Latitude    float64 `gorm:"type:decimal(5,2);not null" json:"latitude" form:"latitude"`
	Longitude   float64 `gorm:"type:decimal(5,2);not null" json:"longitude" form:"longitude"`
	User_ID     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Bookings    []Booking      `gorm:"foreignKey:Homestay_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type HomeStayRespon struct {
	ID          int
	Name        string
	Type        string
	Description string
	Price       int
	Latitude    float64
	Longitude   float64
}
