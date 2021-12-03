package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
	"time"
)

// Fungsi untuk membuat data booking
func CreateBooking(booking *models.Booking) (*models.Booking, error) {
	tx := config.DB.Create(&booking)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return booking, nil
}

// Fungsi untuk menambahkan tanggal checkout pada reservasi yang dibuat
func AddLongstay(checkIn, checkOut time.Time, idBooking int) {
	config.DB.Exec("UPDATE bookings SET long_stay = DATEDIFF(?, ?) WHERE id = ?", checkOut, checkIn, idBooking)
}

// Fungsi untuk menambahkan harga pada booking
func AddHargaToReservation(idHomestay, idBooking int) {
	config.DB.Exec("UPDATE bookings SET total_price = (SELECT price FROM homestays WHERE id = ?)*long_stay WHERE id = ?", idHomestay, idBooking)
}

// Fungsi untuk mendapatkan reservasi by reservasi id
func GetReservation(id int) (interface{}, error) {
	var booking models.BookingDetailRespon
	tx := config.DB.Table("bookings").Select("*, homestays.name").
		Joins("join homestays on homestays.id = bookings.homestay_id").
		Where("bookings.id = ?", id).Find(&booking)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return booking, nil
}

// Fungsi untuk mendapatkan reservasi owner
func GetReservationOwner(id int) (int, error) {
	var booking models.Booking
	tx := config.DB.Where("id = ?", id).Find(&booking)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return 0, tx.Error
	}
	return booking.User_ID, nil
}

// Fungsi untuk mendapatkan semua reservasi owner
func GetAllReservationOwner(user_id int) ([]models.BookingRespon, error) {
	var book []models.BookingRespon
	tx := config.DB.Table("bookings").
		Select(
			"bookings.id, bookings.check_in, bookings.check_out, bookings.total_price, bookings.long_stay, homestays.name, homestays.price").
		Joins("join homestays on homestays.id = bookings.homestay_id").
		Where("bookings.user_id=? and bookings.deleted_at IS NULL", user_id).Find(&book)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return book, nil
}

// Fungsi untuk menghapus reservasi by reservasi id
func CancelReservation(id int) (interface{}, error) {
	var booking models.Booking
	var calendar models.Calendar
	config.DB.Where("booking_id = ?", id).Delete(&calendar)
	config.DB.Model(&models.Booking{}).Where("id=?", id).Update("status_payment", "cancelled")
	if err := config.DB.Where("id = ?", id).Delete(&booking).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}
