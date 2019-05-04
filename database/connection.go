package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

func Connect() (gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"))
	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	return *db, err
}
