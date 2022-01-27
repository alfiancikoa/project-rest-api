package config

import (
	"fmt"
	"log"
	"project1/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConfig() map[string]string {
	conf, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return conf
}

func InitDB() {
	dbconfig := GetConfig()
	connect := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		dbconfig["MYSQL_USER"],
		dbconfig["MYSQL_PASSWORD"],
		dbconfig["MYSQL_HOST"],
		dbconfig["MYSQL_PORT"],
		dbconfig["MYSQL_DBNAME"])

	// Sesuaikan dengan database kalian
	var err error
	DB, err = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("DATABASE is CONNECTED")
	}
	InitalMigration()
}

// Function Initial Migration (Tabel)
func InitalMigration() {
	DB.AutoMigrate(&models.Product{})
}

func InitDBTest() {
	dbconfig := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "password",
		"DB_HOST":     "172.17.0.2",
		"DB_PORT":     "3306",
		"DB_NAME":     "dbproject_test"}
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
	InitalMigrationTest()
}

// Function Initial Migration (Tabel)
func InitalMigrationTest() {
	DB.Migrator().DropTable(&models.Product{})
	DB.AutoMigrate(&models.Product{})
}
