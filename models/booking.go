package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID             int    `gorm:"primarykey"`
	User_ID        int    `json:"userid" form:"userid"`
	Homestay_ID    int    `json:"homestayid" form:"homestayid"`
	Payment_ID     int    `json:"paymentid" form:"paymentid"`
	CheckIn        string `gorm:"type:date;not null" json:"checkin" form:"checkin"`
	CheckOut       string `gorm:"type:date;not null" json:"checkout" form:"checkout"`
	Total_Price    int    `gorm:"type:int;not null" json:"totalprice" form:"totalprice"`
	Status_Payment string `gorm:"type:varchar(100);default:'Not Paided';not null" json:"statuspayment" form:"statuspayment"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Payment        Payment        `gorm:"foreignKey:Payment_ID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
