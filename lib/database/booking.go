package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
	"fmt"
	"log"
	"time"
)

// Fungsi untuk membuat data shopping carts
func CreateBooking(booking *models.Reserve) (interface{}, error) {
	if err := config.DB.Create(&booking.Bookings).Error; err != nil {
		return nil, err
	}
	req_reservation := models.Booking{
		Homestay_ID: booking.Bookings.Homestay_ID,
		CheckIn:     booking.Bookings.CheckIn,
		CheckOut:    booking.Bookings.CheckOut,
	}

	req_credit := models.PaymentRequest{
		Type:   booking.PaymentReq.Type,
		Name:   booking.PaymentReq.Name,
		Number: booking.PaymentReq.Number,
		Cvv:    booking.PaymentReq.Cvv,
		Month:  booking.PaymentReq.Month,
		Year:   booking.PaymentReq.Year,
	}

	Create_Res := models.Reserve{
		Bookings:   req_reservation,
		PaymentReq: req_credit,
	}
	booking.PaymentReq.Booking_ID = booking.Bookings.ID
	config.DB.Create(&booking.PaymentReq)
	return Create_Res, nil
}

func GetPriceIDuserHomestay(id, day int) (int, int, error) {
	homestay := models.Homestay{}
	err := config.DB.Find(&homestay, id)
	if err.Error != nil {
		return 0, 0, err.Error
	}
	log.Println("harga : ", homestay.Price)
	return homestay.Price * day, homestay.User_ID, nil
}

func SearchDay(in, out string) int {
	format := "2006-01-02"
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	diff := end.Sub(start)
	return int(diff.Hours() / 24)
}

func CheckStatusReservation(id_home int, check_in, check_out string) (interface{}, error) {
	var check []models.Booking
	var result string
	checkTime := CekTimeBefore(check_in, check_out)
	if checkTime == true {

		err := config.DB.Table("bookings").Select("*").Where("bookings.homestay_id = ?", id_home).Find(&check)
		if err.Error != nil || err.RowsAffected == 0 {
			return 0, err.Error
		}
		fmt.Println("check row = ", err.RowsAffected)

		for i, _ := range check {
			result = SearchAvailableDay(check[i].CheckIn, check[i].CheckOut, check_in, check_out)
			if result == "not available" {
				break
			}
		}
		return result, nil
	}
	return 0, nil
}

func CekTimeBefore(check_in, check_out string) bool {
	format := "2006-01-02"
	start, _ := time.Parse(format, check_in)
	end, _ := time.Parse(format, check_out)
	if start.Before(end) && time.Now().Before(start) {
		return true
	}
	return false
}

func SearchAvailableDay(in, out, cek_in, cek_out string) string {
	format := "2006-01-02"

	cek_start, _ := time.Parse(format, cek_in)
	cek_end, _ := time.Parse(format, cek_out)
	start, _ := time.Parse(format, in)
	end, _ := time.Parse(format, out)

	hasil := "available"
	if (start.Before(cek_start) && end.After(cek_start)) || (start.Before(cek_end) && end.After(cek_end)) {
		hasil = "not available"
		return hasil
	} else if start.Equal(cek_start) || end.Equal(cek_start) || start.Equal(cek_end) || end.Equal(cek_end) {
		hasil = "not available"
		return hasil
	}

	fmt.Println("hasil", hasil)
	return hasil
}

// Fungsi untuk mendapatkan data seluruh booking dari user
// func GetBookingId(idUser int) (interface{}, error) {
// 	var bookings models.Booking
// 	query := config.DB.Find(&bookings).Where("user_id = ?", idUser)
// 	if query.Error != nil {
// 		return 0, query.Error
// 	}
// 	if query.RowsAffected == 0 {
// 		return 0, nil
// 	}
// 	return bookings, nil
// }

func GetReservation(id int) (interface{}, error) {
	var get_reservation []models.GetReserve

	query := config.DB.Table("bookings").Select("*").Where("bookings.user_id = ?", id).Find(&get_reservation)
	if query.Error != nil || query.RowsAffected == 0 {
		return nil, query.Error
	}
	query_homestay := config.DB.Table("homestays").Select("bookings.user_id,bookings.homestay_id,homestays.name,bookings.check_in,bookings.check_out,homestays.price,bookings.total_price").Joins("join bookings on homestays.id = bookings.homestay_id").Find(&get_reservation)
	if query_homestay.Error != nil {
		return nil, query_homestay.Error
	}
	log.Println("gethomestay :", get_reservation[0].CheckIn)
	return get_reservation, nil
}
