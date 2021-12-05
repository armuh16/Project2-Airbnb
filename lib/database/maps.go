package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"

	"github.com/kelvins/geocoder"
)

// Insert Address to Maps Table
func InsertAddress(homestay_id int, address []geocoder.Address) (*models.Address, error) {
	maps := models.Address{
		Street:      address[2].Street,
		City:        address[2].City,
		County:      address[2].County,
		State:       address[2].State,
		Country:     address[2].Country,
		PostalCode:  address[2].PostalCode,
		Homestay_ID: homestay_id,
	}
	if err := config.DB.Create(&maps).Error; err != nil {
		return nil, err
	}
	return &maps, nil
}

// Edit Address in Maps Table
func EditAddress(homestay_id int, addressRequest []geocoder.Address) (*models.Address, error) {
	maps := models.Address{
		Street:     addressRequest[2].Street,
		City:       addressRequest[2].City,
		County:     addressRequest[2].County,
		State:      addressRequest[2].State,
		Country:    addressRequest[2].Country,
		PostalCode: addressRequest[2].PostalCode,
	}
	if err := config.DB.Model(&models.Address{}).Where("homestay_id=?", homestay_id).Updates(maps).Error; err != nil {
		return nil, err
	}
	return &maps, nil
}
