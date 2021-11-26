package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Controller untuk memasukkan barang baru ke Booking
func CreateBookingController(c echo.Context) error {
	// Mendapatkan data booking baru dari client
	input := models.Booking{}
	idToken := middlewares.ExtractTokenUserId(c)
	c.Bind(&input)
	if input.User_ID != idToken {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("Wrong Users ID"))
	}

	// Menyimpan data barang baru menggunakan fungsi CreateBooking
	booking, e := database.CreateBooking(input)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed to book"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success to book", booking))
}

// Controller untuk mendapatkan data Booking
func GetBookingController(c echo.Context) error {
	// Mendapatkan id user dari token
	idUser := middlewares.ExtractTokenUserId(c)
	// Mendapatkan data seluruh booking pada user tertentu menggunakan fungsi GetBookingId
	book, e := database.GetBookingId(idUser)
	if e != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed to fetch bookings"))
	}
	if book == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("booking not found"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success view booking user id", book))
}
