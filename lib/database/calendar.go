package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
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
	if n <= 0 {
		n = 1
	}
	dateCalendar := make([]models.Calendar, n)
	now := time.Now()
	intervalOut := int(bookInfo.CheckOut.Sub(now).Hours())
	// intervalDay := bookInfo.CheckIn.Day() - now.Day()
	intervalInt := int(bookInfo.CheckIn.Sub(now).Hours()) % 24
	Day := int(bookInfo.CheckIn.Sub(now).Hours()) / 24
	if intervalInt >= 14 {
		Day++
	}
	// Ketika checkin dan checkout di hari yang sama, berlaku untuk yang reservasi subuh
	if intervalOut <= 12 {
		dateCalendar[0].Homestay_ID = homestay_id
		dateCalendar[0].DateIn = time.Now().AddDate(0, 0, -1)
		dateCalendar[0].DateOut = bookInfo.CheckOut
	} else {
		for i := 0; i < n; i++ {
			dateCalendar[i].Homestay_ID = homestay_id
			dateCalendar[i].DateIn = time.Now().AddDate(0, 0, Day+i)
			dateCalendar[i].DateOut = time.Now().AddDate(0, 0, Day+i+1)
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
	now := time.Now()
	zona, _ := now.Zone()
	format := "2006-01-02 15:04:05 MST"
	timeIn := " 14:00:00 " + zona
	timeOut := " 12:00:00 " + zona
	checkIn, _ := time.Parse(format, request.CheckIn+timeIn)
	if request.CheckIn == request.CheckOut {
		checkIn = time.Now()
	}

	checkOut, _ := time.Parse(format, request.CheckOut+timeOut)
	longstayInt := int(checkOut.Sub(checkIn).Hours() / 22)
	// intervalDay := checkIn.Day() - now.Day()
	intervalInt := int(checkIn.Sub(now).Hours()) % 24
	Day := int(checkIn.Sub(now).Hours()) / 24
	if intervalInt >= 14 {
		Day++
	}
	for i := 0; i < longstayInt; i++ {
		date := time.Now().AddDate(0, 0, Day+i)
		datef := date.Format("2006-01-02")
		if row := CheckDate(request.Homestay_ID, datef); row > 0 {
			return row
		}
	}
	return 0
}
