package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	ID             int       `gorm:"primarykey;AUTO_INCREMENT"`
	User_ID        int       `json:"userid" form:"userid"`
	Homestay_ID    int       `json:"homestayid" form:"homestayid"`
	CheckIn        time.Time `gorm:"type:datetime;not null" json:"checkin" form:"checkin"`
	CheckOut       time.Time `gorm:"type:datetime;not null" json:"checkout" form:"checkout"`
	LongStay       int       `json:"longstay" form:"longstay"`
	Total_Price    int       `json:"totalprice" form:"totalprice"`
	Status_Payment string    `gorm:"default:'not paided'" json:"statuspayment" form:"statuspayment"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type PostReservation struct {
	Homestay_ID int     `json:"homestayid" form:"homestayid"`
	CheckIn     string  `json:"checkin" form:"checkin"`
	CheckOut    string  `json:"checkout" form:"checkout"`
	Payment     Payment `json:"payment" form:"payment"`
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

type TempRespon struct {
	ID          int
	User_ID     int
	Homestay_ID int
	CheckIn     time.Time
	CheckOut    time.Time
	Total_Price int
}

type BookingRespon struct {
	ID          int
	Name        string
	Check_In    string
	Check_Out   string
	Long_Stay   int
	Price       int
	Total_Price int
}
type BookingDetailRespon struct {
	ID             int
	User_id        int
	Homestay_Id    int
	Name           string
	Check_In       string
	Check_Out      string
	Long_stay      int
	Total_Price    int
	Status_Payment string
}
