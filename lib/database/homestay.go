package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
	"alta/airbnb/util"
)

func InsertHomestay(homestay models.Homestay, user_id int) (*models.Homestay, error) {
	homestay.User_ID = user_id
	tx := config.DB.Where("name=? AND user_id=?", homestay.Name, user_id).Find(&models.Homestay{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return nil, nil
	}
	lat, lng, e := util.GetGeocodeLocations(homestay.Address)
	if e != nil {
		return nil, e
	}
	homestay.Latitude = lat
	homestay.Longitude = lng
	if err := config.DB.Save(&homestay).Error; err != nil {
		return nil, err
	}
	return &models.Homestay{}, nil
}

func GetHomeStayDetail(homestay_id int) (*models.HomeStayRespon, error) {
	homestay := models.HomeStayRespon{}
	tx := config.DB.Table("homestays").Select(
		"homestays.id, homestays.name, homestays.type, homestays.description, homestays.price, homestays.address, homestays.latitude, homestays.longitude").
		Where("homestays.deleted_at IS NULL and id=?", homestay_id).Find(&homestay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	return &homestay, nil
}

func GetHomeStayByType(tipe string) ([]models.HomeStayRespon, error) {
	homestay := []models.HomeStayRespon{}
	tx := config.DB.Table("homestays").Select(
		"homestays.id, homestays.name, homestays.type, homestays.description, homestays.price, homestays.address, homestays.latitude, homestays.longitude").
		Where("homestays.deleted_at IS NULL and type=?", tipe).Find(&homestay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return homestay, nil
}

func GetMyHometay(user_id int) ([]models.HomeStayRespon, error) {
	homestay := []models.HomeStayRespon{}
	tx := config.DB.Table("homestays").Select(
		"homestays.id, homestays.name, homestays.type, homestays.description, homestays.price, homestays.address, homestays.latitude, homestays.longitude").
		Where("homestays.deleted_at IS NULL and user_id=?", user_id).Find(&homestay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return homestay, nil
}

func GetAllHomeStay() ([]models.HomeStayRespon, error) {
	homestays := []models.HomeStayRespon{}
	tx := config.DB.Table("homestays").Select(
		"homestays.id, homestays.name, homestays.type, homestays.description, homestays.price, homestays.address, homestays.latitude, homestays.longitude").
		Where("homestays.deleted_at IS NULL").Find(&homestays)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return homestays, nil
}

func EditHomestay(homerequest *models.HomeStayRespon, id int, user_id int) (*models.Homestay, error) {
	homestay := models.Homestay{}
	tx := config.DB.Where("user_id=?", user_id).Find(&homestay, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	lat, lng, e := util.GetGeocodeLocations(homestay.Address)
	if e != nil {
		return nil, e
	}
	homestay.Name = homerequest.Name
	homestay.Type = homerequest.Type
	homestay.Description = homerequest.Description
	homestay.Price = homerequest.Price
	homestay.Latitude = lat
	homestay.Longitude = lng
	if tx.RowsAffected > 0 {
		if err := config.DB.Save(&homestay).Error; err != nil {
			return nil, err
		} else {
			return &homestay, nil
		}
	}
	return nil, nil
}

func DeleteHomestay(id int, user_id int) (*models.Homestay, error) {
	tx := config.DB.Where("id=? and user_id=?", id, user_id).Delete(&models.Homestay{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &models.Homestay{}, nil
	}
	return nil, nil
}
