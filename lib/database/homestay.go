package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
	"alta/airbnb/util"
	"strconv"
	"unicode"
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
	if err := config.DB.Save(&homestay).Error; err != nil {
		return nil, err
	}

	return &homestay, nil
}

func InsertFasilities(feature_id []int, homestay_id int) (*models.Facility, error) {
	facility := make([]models.Facility, len(feature_id))
	for i := 0; i < len(feature_id); i++ {
		facility[i].Feature_ID = feature_id[i]
		facility[i].Homestay_ID = homestay_id
	}
	if err := config.DB.Create(&facility).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func GetHomeStayDetail(homestay_id int) (*models.HomeStayResponDetail, error) {
	homestay := models.Homestay{}
	tx := config.DB.Find(&homestay, homestay_id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected < 1 {
		return nil, nil
	}
	fasilitas := []models.Feature{}
	config.DB.Table("facilities").Select(
		"features.feature_name").Joins("join features on features.id = facilities.feature_id").
		Where("facilities.deleted_at IS NULL and homestay_id=?", homestay_id).Find(&fasilitas)

	features := make([]string, len(fasilitas))
	for i := 0; i < len(fasilitas); i++ {
		features[i] = fasilitas[i].Feature_name
	}
	homedetail := models.HomeStayResponDetail{
		ID:          homestay.ID,
		Name:        homestay.Name,
		Type:        homestay.Type,
		Description: homestay.Description,
		Price:       homestay.Price,
		Address:     homestay.Address,
		Latitude:    homestay.Latitude,
		Longitude:   homestay.Longitude,
	}
	homedetail.Features = features
	return &homedetail, nil
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

func GetFeatureIdByFacility(facilities string) (*models.Feature, error) {
	isdigit := unicode.IsDigit(rune(facilities[0]))
	var digit int
	if isdigit {
		digit, _ = strconv.Atoi(facilities)
	}
	feature := models.Feature{}
	if err := config.DB.Where("feature_name=? or id=?", facilities, digit).Find(&feature).Error; err != nil {
		return nil, err
	}
	return &feature, nil
}

func GetHomeStayByFacility(facilities string) ([]models.HomeStayRespon, error) {
	homestay := []models.HomeStayRespon{}
	home, e := GetFeatureIdByFacility(facilities)
	if e != nil {
		return nil, e
	}

	tx := config.DB.Table("homestays").Select(
		"homestays.id, homestays.name, homestays.type, homestays.description, homestays.price, homestays.address, homestays.latitude, homestays.longitude").Joins(
		"join facilities on facilities.homestay_id = homestays.id").Where("facilities.deleted_at IS NULL and feature_id=?", home.ID).Find(&homestay)
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
