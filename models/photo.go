package models

import (
	"time"

	"gorm.io/gorm"
)

type Photo struct {
	ID          int    `gorm:"primarykey"`
	Homestay_ID int    `gorm:"primarykey" json:"homestay_id" form:"homestay_id"`
	Photo_Name  string `gorm:"type:varchar(50);not null" json:"photo_name" form:"photo_name"`
	Url         string `gorm:"type:longtext" json:"url" form:"url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Get_Photo struct {
	Homestay_ID int
	Nama_Photo  string
}

type EditPhoto struct {
	Photo_Name string
	Url        string
}
