package config

import (
	"alta/airbnb/models"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConfig() (config map[string]string) {
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conf
}

var DB *gorm.DB

// Initial Database
func InitDB() {
	dbconfig := GetConfig()
	// Sesuaikan dengan database kalian
	connect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbconfig["DB_USERNAME"],
		dbconfig["DB_PASSWORD"],
		dbconfig["DB_HOST"],
		dbconfig["DB_PORT"],
		dbconfig["DB_NAME"])

	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitalMigration()
}

// Function Initial Migration (Tabel)
func InitalMigration() {
	DB.AutoMigrate(&models.Users{})
}

func InitDBTest() {
	// Sesuaikan dengan database kalian
	connect := "root:yourpasswd@tcp(localhost:3306)/usertesting?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitalMigrationtest()
}

func InitalMigrationtest() {
	DB.Migrator().DropTable(&models.Users{})
	DB.AutoMigrate(&models.Users{})
}
