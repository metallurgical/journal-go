package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Journal struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Status int32
}

func (j *Journal) GetJournal(db gorm.DB, id int64) (error) {
	if err := db.First(j, id).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}