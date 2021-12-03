package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
)

func InsertPayment(booking_id int, paymentReq models.Payment) (models.Payment, error) {
	payment := models.Payment{
		Type:   paymentReq.Type,
		Name:   paymentReq.Name,
		Cvv:    paymentReq.Cvv,
		Month:  paymentReq.Month,
		Year:   paymentReq.Year,
		Number: paymentReq.Number,
	}
	if err := config.DB.Create(&payment).Error; err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}
