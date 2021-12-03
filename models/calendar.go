package models

import (
	"time"

	"gorm.io/gorm"
)

type Calendar struct {
	ID          int `gorm:"primarykey"`
	Booking_ID  int
	Homestay_ID int
	DateIn      time.Time `gorm:"type:datetime;not null" json:"datein" form:"datein"`
	DateOut     time.Time `gorm:"type:datetime;not null" json:"dateout" form:"dateout"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
