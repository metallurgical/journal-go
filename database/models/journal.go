package models

import "github.com/jinzhu/gorm"

type Journal struct {
	gorm.Model
	Id int `gorm:"AUTO_INCREMENT"`
	Status int32
}