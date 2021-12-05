package database

import (
	"alta/airbnb/config"
	"alta/airbnb/middlewares"
	"alta/airbnb/models"

	"golang.org/x/crypto/bcrypt"
)

var user models.Users

func GetUser(userID int) (*models.Users, error) {
	var userid models.Users
	tx := config.DB.Find(&userid, userID)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &userid, nil
	}
	return nil, nil
}

func GetUserByEmail(email string) (int64, error) {
	tx := config.DB.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return tx.RowsAffected, nil
	}
	return 0, nil
}

func RegisterUser(user models.Users) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int) (interface{}, error) {
	var userid models.Users
	if err := config.DB.Where("id = ?", id).Delete(&userid).Error; err != nil {
		return nil, err
	}
	return userid, nil
}

func UpdateUser(id int, User models.Users) (models.Users, error) {
	var user models.Users

	if err := config.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	expass, _ := GeneratehashPassword(User.Password)
	user.Name = User.Name
	user.Email = User.Email
	user.Password = expass
	user.PhoneNumber = User.PhoneNumber
	user.Gender = User.Gender

	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func LoginUsers(user *models.UserLogin) (*models.Users, error) {
	var err error
	userpassword := models.Users{}
	if err = config.DB.Where("email = ?", user.Email).First(&userpassword).Error; err != nil {
		return nil, err
	}
	check := CheckPasswordHash(user.Password, userpassword.Password)
	if !check {
		return nil, nil
	}

	userpassword.Token, err = middlewares.CreateToken(int(userpassword.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(userpassword).Error; err != nil {
		return nil, err
	}
	return &userpassword, nil
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
