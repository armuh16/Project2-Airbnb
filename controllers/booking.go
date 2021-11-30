package controllers

import (
	"alta/airbnb/lib/database"
	responses "alta/airbnb/lib/response"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// Controller untuk memasukkan barang baru ke Booking
func CreateBookingController(c echo.Context) error {
	input := models.Booking{}
	c.Bind(&input)
	format := "2006-01-02"

	checkIn, _ := time.Parse(format, input.CheckIn)
	checkOut, _ := time.Parse(format, input.CheckOut)
	fmt.Println("INPUT", input)
	logged := middlewares.ExtractTokenUserId(c)
	input.User_ID = int(logged)
	book, err := database.CreateBooking(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	database.AddCheckOut(checkIn, checkOut, book.ID)
	database.AddHargaToReservation(input.Homestay_ID, book.ID)
	return c.JSON(http.StatusOK, responses.StatusSuccess("success"))
}

func GetBookingControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	userId, _ := database.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.StatusUnauthorized())
	}
	booking, _ := database.GetReservation(id)
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success", booking))
}

func CancelBookingController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.FalseParamResponse())
	}
	userId, _ := database.GetReservationOwner(id)
	if err != nil || userId == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	logged := middlewares.ExtractTokenUserId(c)
	if int(logged) != userId {
		return c.JSON(http.StatusBadRequest, responses.StatusUnauthorized())
	}
	database.CancelReservation(id)
	return c.JSON(http.StatusOK, responses.StatusSuccess("success"))
}

// Fungsi untuk melakukan pengecekan availability suatu room
// func RoomReservationCheck(c echo.Context) error {
// 	input := models.BodyCheckIn{}
// 	c.Bind(&input)

// 	// Mendapatkan seluruh tanggal reservation room tertentu
// 	dateList, err := database.RoomReservationList(input.Homestay_ID)
// 	if err != nil || dateList == nil {
// 		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
// 	}
// 	// Pengecekan ketersediaan room untuk tanggal check_in dan check_out yang diinginkan
// 	for _, date := range dateList {
// 		input_checkin := input.CheckIn.Unix()
// 		fmt.Println("INI HASILNYA : ", input_checkin)
// 		input_checkout := input.CheckOut.Unix()
// 		date_checkin := date.CheckIn.Unix()
// 		fmt.Println("INI HASILNYA DATE : ", date_checkin)
// 		date_checkout := date.CheckOut.Unix()
// 		if (input_checkin >= date_checkin && input_checkin <= date_checkout) || (input_checkout >= date_checkin && input_checkout <= date_checkout) {
// 			return c.JSON(http.StatusBadRequest, responses.StatusFailed("Room Unavailable"))
// 		}
// 	}
// 	return c.JSON(http.StatusOK, responses.StatusSuccess("Room Available"))
// }

func CekReservationControllers(c echo.Context) error {
	cek := models.BodyCheckIn{}
	c.Bind(&cek)

	_, err := database.GetHomeStayDetail(int(cek.Homestay_ID))
	data, er := database.CekStatusReservation(cek.Homestay_ID, cek.CheckIn, cek.CheckOut)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}

	if er != nil || data == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("failed"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success", data))

}
