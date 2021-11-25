package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID          int        `gorm:"primarykey"`
	Name        string     `gorm:"type:varchar(255)" json:"name" form:"name"`
	Email       string     `gorm:"type:varchar(100);unique;not null" json:"email" form:"email"`
	Password    string     `gorm:"type:varchar(255);not null" json:"password" form:"password"`
	PhoneNumber string     `gorm:"type:varchar(20);not null" json:"phonenumber" form:"phonenumber"`
	Gender      string     `gorm:"type:enum('male','female');" json:"gender" form:"gender"`
	Token       string     `gorm:"type:longtext;" json:"token" form:"token"`
	Role        string     `gorm:"type:varchar(100);" json:"" form:""`
	Homestays   []Homestay `gorm:"foreignKey:User_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GetUserResponse struct {
	Name        string
	Email       string
	PhoneNumber string
	Gender      string
}
