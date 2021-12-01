package models

import (
	"time"

	"gorm.io/gorm"
)

type Feature struct {
	ID           int        `gorm:"primarykey"`
	Feature_name string     `gorm:"type:varchar(100);not null" json:"featurename" form:"featurename"`
	Facilities   []Facility `gorm:"foreignKey:Feature_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Homestays    []*Homestay `gorm:"many2many:feature_homestays;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
