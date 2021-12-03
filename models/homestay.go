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
	Address     string  `gorm:"type:varchar(255);not null" json:"address" form:"address"`
	Latitude    float64 `gorm:"not null" json:"latitude" form:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude" form:"longitude"`
	User_ID     int
	Facilities  []Facility `gorm:"foreignKey:Homestay_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Bookings    []Booking  `gorm:"foreignKey:Homestay_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Calendars   []Calendar `gorm:"foreignKey:Homestay_ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Maps        Address    `gorm:"foreignKey:Homestay_ID"`
	// Features     []*Feature `gorm:"many2many:feature_homestays;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PostHomestayRequest struct {
	Name        string `json:"name" form:"name"`
	Type        string `json:"type" form:"type"`
	Description string `json:"description" form:"description"`
	Facility    []int  `json:"facility" form:"facility"`
	Price       int    `json:"price" form:"price"`
	Address     string `json:"address" form:"address"`
}

type HomeStayResponDetail struct {
	ID          int
	Name        string
	Type        string
	Description string
	Price       int
	Address     string
	Latitude    float64
	Longitude   float64
	Features    []string
}
type HomeStayRespon struct {
	ID          int
	Name        string
	Type        string
	Description string
	Price       int
	Address     string
	Latitude    float64
	Longitude   float64
}
