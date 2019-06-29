package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	"time"
)

type Publication struct {
	ID          uint `gorm:"primary_key"`
	Volume      string
	No          int64
	CreatedBy   int64
	Status      int32
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (Publication) TableName() string {
	return "publications"
}

// Get existing resource from database.
func (j *Publication) Get(db gorm.DB, id int64) (error) {
	if err := db.First(j, id).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}

// Create new resource.
func (j *Publication) Create(db gorm.DB, req *pb.PublicationRequest) (interface{}, error) {
	rev := &Publication{
		Volume:    req.Volume,
		No:        req.No,
		CreatedBy: req.CreatedBy,
	}
	if !db.NewRecord(&rev) {
		return nil, errors.New("Operation failed. Please try again.")
	}
	db.Create(&rev)

	return rev, nil
}
