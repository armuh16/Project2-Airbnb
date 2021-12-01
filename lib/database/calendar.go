package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
	"fmt"
	"time"
)

func GetDataBooking(booking_id int) (models.Booking, error) {
	book := models.Booking{}
	if err := config.DB.Find(&book, booking_id).Error; err != nil {
		return models.Booking{}, err
	}
	return book, nil
}

func InsertDateToCalendar(homestay_id int, booking_id int) ([]models.Calendar, error) {
	bookInfo, _ := GetDataBooking(booking_id)
	n := bookInfo.LongStay
	dateCalendar := make([]models.Calendar, n)
	now := time.Now()
	interval := bookInfo.CheckIn.Sub(now)
	intervalInt := int(interval.Hours() / 24)
	if now.Hour() >= 12 && interval.Hours() >= 12 {
		intervalInt = intervalInt + 1
	}
	if bookInfo.LongStay == 0 {
		dateCalendar[0].Homestay_ID = homestay_id
		dateCalendar[0].DateIn = time.Now().AddDate(0, 0, -1)
		dateCalendar[0].DateOut = time.Now()
	} else {
		for i := 0; i < n; i++ {
			dateCalendar[i].Homestay_ID = homestay_id
			dateCalendar[i].DateIn = time.Now().AddDate(0, 0, intervalInt+i)
			dateCalendar[i].DateOut = time.Now().AddDate(0, 0, intervalInt+i+1)
		}
	}
	if err := config.DB.Create(&dateCalendar).Error; err != nil {
		return []models.Calendar{}, err
	}
	return dateCalendar, nil
}

func CheckDate(homestay_id int, date string) int64 {
	calendar := models.Calendar{}
	tx := config.DB.Where("date(date_in)=? AND homestay_id=?", date, homestay_id).Find(&calendar)
	if tx.Error != nil {
		return -1
	}
	if tx.RowsAffected > 0 {
		return tx.RowsAffected
	}
	return 0
}

func CheckAvailability(request models.BodyCheckIn) int64 {
	format := "2006-01-02"
	now := time.Now()
	checkIn, _ := time.Parse(format, request.CheckIn)
	checkOut, _ := time.Parse(format, request.CheckOut)
	longstay := checkOut.Sub(checkIn)
	interval := checkIn.Sub(now)
	longstayInt := int(longstay.Hours() / 24)
	intervalInt := int(interval.Hours() / 24)
	if now.Hour() >= 12 && interval.Hours() >= 12 {
		intervalInt = intervalInt + 1
	}
	for i := 0; i < longstayInt; i++ {
		date := time.Now().AddDate(0, 0, intervalInt+i)
		datef := date.Format(format)
		fmt.Println("DATEEEEEEEEEEEEEE >>>>>>>>>>>>", date)
		if row := CheckDate(request.Homestay_ID, datef); row > 0 {
			return row
		}
	}
	return 0
}
