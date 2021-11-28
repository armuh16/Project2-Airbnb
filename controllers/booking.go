package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller untuk memasukkan barang baru ke Booking
func CreateBookingController(c echo.Context) error {
	// Mendapatkan data booking baru dari client
	// input := models.PostBook{}
	// idToken := middlewares.ExtractTokenUserId(c)
	// c.Bind(&input)
	// if input.User_ID != idToken {
	// 	return c.JSON(http.StatusBadRequest, responses.StatusFailed("Wrong Users ID"))
	// }

	// // Menyimpan data barang baru menggunakan fungsi CreateBooking
	// booking, e := database.CreateBooking(input)
	// if e != nil {
	// 	return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed to book"))
	// }
	// return c.JSON(http.StatusOK, responses.StatusSuccessData("success to book", booking))

	Reservation := models.Reserve{}

	c.Bind(&Reservation)
	id := middlewares.ExtractTokenUserId(c)
	Reservation.Bookings.User_ID = int(id)

	homestay_id := Reservation.Bookings.Homestay_ID
	checkIn := Reservation.Bookings.CheckIn
	CheckOut := Reservation.Bookings.CheckOut

	// cek status reservasi
	data, er := database.CheckStatusReservation(homestay_id, checkIn, CheckOut)
	if er != nil || data == "not available" || data == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	} else {

		// mencari berapa hari reservasi
		day := database.SearchDay(checkIn, CheckOut)
		price, id_user_homestay, _ := database.GetPriceIDuserHomestay(Reservation.Bookings.Homestay_ID, day)

		log.Println("How Many Days : ", day)
		log.Println("id user homestay : ", id_user_homestay)

		//cek iduser di homestay
		if id == int(id_user_homestay) {
			return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
		}

		Reservation.Bookings.Status = "not available"
		Reservation.Bookings.Total_Price = price
		_, err := database.CreateBooking(&Reservation)
		if err != nil {
			return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
		}

		return c.JSON(http.StatusBadRequest, responses.StatusSuccess("success"))
	}
}

// Controller untuk mendapatkan data Booking
// func GetBookingController(c echo.Context) error {
// 	// Mendapatkan id user dari token
// 	idUser := middlewares.ExtractTokenUserId(c)
// 	// Mendapatkan data seluruh booking pada user tertentu menggunakan fungsi GetBookingId
// 	book, e := database.GetBookingId(idUser)
// 	if e != nil {
// 		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed to fetch bookings"))
// 	}
// 	if book == 0 {
// 		return c.JSON(http.StatusBadRequest, responses.StatusFailed("booking not found"))
// 	}
// 	return c.JSON(http.StatusOK, responses.StatusSuccessData("success view booking user id", book))
// }

func CheckReserveController(c echo.Context) error {
	cek := models.CheckHomestay{}
	c.Bind(&cek)

	data, er := database.CheckStatusReservation(cek.Homestay_ID, cek.CheckIn, cek.CheckOut)
	if er != nil || data == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success", data))
}

func GetBookingController(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	log.Println("id  :", id)
	data, e := database.GetReservation(id)
	if e != nil || data == nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success", data))
}
