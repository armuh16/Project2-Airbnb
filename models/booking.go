package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID          int    `gorm:"primarykey;AUTO_INCREMENT"`
	User_ID     int    `json:"userid" form:"userid"`
	Homestay_ID int    `json:"homestayid" form:"homestayid"`
	CheckIn     string `gorm:"type:date;not null" json:"checkin" form:"checkin"`
	CheckOut    string `gorm:"type:date;not null" json:"checkout" form:"checkout"`
	Status      string `gorm:"type:varchar(100);not null" json:"status" form:"status"`
	Total_Price int
	Payments    PaymentRequest
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type PaymentRequest struct {
	Booking_ID int
	Type       string `json:"type" form:"type"`
	Name       string `json:"name" form:"name"`
	Cvv        int    `json:"cvv" form:"cvv"`
	Month      string `json:"month" form:"month"`
	Year       string `json:"year" form:"year"`
	Number     string `json:"number" form:"number"`
}

type Reserve struct {
	Bookings   Booking        `json:"bookings"`
	PaymentReq PaymentRequest `json:"paymentreq"`
}

type GetReserve struct {
	User_ID     int
	Homestay_ID int
	Name        string `json:"name" form:"name"`
	CheckIn     string
	CheckOut    string
	Price       int `json:"price" form:"price"`
	Total_Price int
}

type CheckHomestay struct {
	Homestay_ID int    `json:"homestayid" form:"homestayid"`
	CheckIn     string `json:"checkin" form:"checkin"`
	CheckOut    string `json:"checkout" form:"checkout"`
}
