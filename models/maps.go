package models

type Address struct {
	ID          int `gorm:"primarykey"`
	Homestay_ID int
	Street      string `gorm:"type:varchar(255);not null" json:"street" form:"street"`
	City        string `gorm:"type:varchar(255);not null" json:"city" form:"city"`
	County      string `gorm:"type:varchar(255);not null" json:"county" form:"county"`
	State       string `gorm:"type:varchar(255);not null" json:"state" form:"state"`
	Country     string `gorm:"type:varchar(255);not null" json:"country" form:"country"`
	PostalCode  string `gorm:"type:varchar(255);not null" json:"postalcode" form:"postalcode"`
}
