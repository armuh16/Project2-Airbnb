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
	var booking models.Booking
	tx := config.DB.Where("id = ?", id).Find(&booking)
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

// Fungsi untuk menghapus reservasi by reservasi id
func CancelReservation(id int) (interface{}, error) {
	var booking models.Booking
	if err := config.DB.Where("id = ?", id).Delete(&booking).Error; err != nil {
		return nil, err
	}
	return "deleted", nil
}

// Fungsi untuk mendapatkan tanggal check_in dan check_out suatu reservasi
func RoomReservationList(id int) ([]models.ReservationDate, error) {
	var dates []models.ReservationDate
	tx := config.DB.Table("bookings").Select("bookings.check_in, bookings.check_out").Where("bookings.homestay_id = ?", int(id)).Find(&dates)
	if tx.Error != nil || tx.RowsAffected < 1 {
		return nil, tx.Error
	}
	return dates, nil
}
