package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID          int       `gorm:"primarykey;AUTO_INCREMENT"`
	User_ID     int       `json:"userid" form:"userid"`
	Homestay_ID int       `json:"homestayid" form:"homestayid"`
	CheckIn     time.Time `gorm:"type:datetime;not null" json:"checkin" form:"checkin"`
	CheckOut    time.Time `gorm:"type:datetime;not null" json:"checkout" form:"checkout"`
	LongStay    int       `json:"longstay" form:"longstay"`
	Total_Price int       `json:"totalprice" form:"totalprice"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type BodyCheckIn struct {
	Homestay_ID int    `json:"homestayid" form:"homestayid"`
	CheckIn     string `json:"checkin" form:"checkin"`
	CheckOut    string `json:"checkout" form:"checkout"`
}

type ReservationDate struct {
	CheckIn  time.Time
	CheckOut time.Time
}
