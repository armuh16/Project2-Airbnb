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
	// Features     []*Feature `gorm:"many2many:feature_homestays;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
