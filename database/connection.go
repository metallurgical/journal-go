package database

import (
	"github.com/jinzhu/gorm"
	"log"
)

func Connect() (gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@/journal?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatalf("Could not connect to database :%v", err)
	}

	return *db, err
}