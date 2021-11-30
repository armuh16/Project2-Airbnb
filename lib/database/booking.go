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
func AddCheckOut(checkIn, checkOut time.Time, idBooking int) {
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

func CekStatusReservation(id_home int, check_in, check_out string) (interface{}, error) {
	var check []models.Booking
	var hasil string

	if CekTimeBefore(check_in, check_out) == true {
		err := config.DB.Table("bookings").Select("*").Where("bookings.homestay_id = ?", id_home).Find(&check)
		if err.Error != nil {
			return 0, err.Error
		} else if err.RowsAffected == 0 {
			return 1, nil
		}

		for i, _ := range check {
			hasil = SearchAvailableDay(check[i].CheckIn, check[i].CheckOut, check_in, check_out)
			if hasil == "Not Available" {
				break
			}
		}
		return hasil, nil
	}
	return 0, nil
}

func SearchAvailableDay(in, out, check_in, check_out string) string {
	format := "2006-01-02"

	check_start, _ := time.Parse(format, check_in)
	check_end, _ := time.Parse(format, check_out)
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	hasil := "Available"
	if (start.Before(check_start) && end.After(check_start)) || (start.Before(check_end) && end.After(check_end)) {
		hasil = "Not Available"
		return hasil
	} else if start.Equal(check_start) || end.Equal(check_start) || start.Equal(check_end) || end.Equal(check_end) {
		hasil = "Not Available"
		return hasil
	}
	return hasil
}

func CekTimeBefore(check_start, check_end string) bool {
	format := "2006-01-02"
	start, _ := time.Parse(format, check_start)
	end, _ := time.Parse(format, check_end)
	if start.Before(end) && time.Now().Before(start) {
		return true
	}
	return false
}
