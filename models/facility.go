package models

import (
	"time"

	"gorm.io/gorm"
)

type Facility struct {
	ID          int `gorm:"primarykey;AUTO_INCREMENT"`
	Homestay_ID int
	Feature_ID  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
