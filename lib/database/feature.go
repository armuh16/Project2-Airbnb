package database

import (
	"alta/airbnb/config"
	"alta/airbnb/models"
)

func InsertFeature(feature models.Feature) (models.Feature, error) {
	tx := config.DB.Save(&feature)
	if tx.Error != nil {
		return feature, tx.Error
	}
	return feature, nil
}

func GetFeatureByName(Feature_name string) (int64, error) {
	tx := config.DB.Where("feature_name = ?", Feature_name).Find(&models.Feature{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return tx.RowsAffected, nil
	}
	return 0, nil
}
