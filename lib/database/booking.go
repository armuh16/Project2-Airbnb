package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
)

// Fungsi untuk membuat data shopping carts
func CreateBooking(booking models.Booking) (models.Booking, error) {
	query := config.DB.Save(&booking)
	if query.Error != nil {
		return booking, query.Error
	} else {
		return booking, nil
	}
}

// Fungsi untuk mendapatkan data seluruh booking dari user
func GetBookingId(idUser int) (int, error) {
	var bookings models.Booking
	query := config.DB.Find(&bookings, idUser)
	if query.Error != nil {
		return 0, query.Error
	}
	if query.RowsAffected == 0 {
		return 0, nil
	}
	return int(bookings.ID), nil
}
