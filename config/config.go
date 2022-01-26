package config

import (
	"fmt"
	"project1/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbconfig := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "",
		"DB_HOST":     "localhost",
		"DB_PORT":     "3306",
		"DB_NAME":     "project1"}
	// Sesuaikan dengan database kalian
	connect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbconfig["DB_USERNAME"],
		dbconfig["DB_PASSWORD"],
		dbconfig["DB_HOST"],
		dbconfig["DB_PORT"],
		dbconfig["DB_NAME"])

	// Sesuaikan dengan database kalian
	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitalMigration()
}

// Function Initial Migration (Tabel)
func InitalMigration() {
	DB.AutoMigrate(&models.Product{})
}
