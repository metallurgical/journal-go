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

const DateFormat = "2006-01-02 15:04:05"

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
	publishedAt, err := time.Parse(DateFormat, req.PublishedAt)
	if err != nil {
		publishedAt = time.Now() // default to current UTC date
	}
	rev := &Publication{
		Volume:      req.Volume,
		No:          req.No,
		CreatedBy:   req.CreatedBy,
		Status:      req.Status,
		PublishedAt: publishedAt,
	}
	if !db.NewRecord(&rev) {
		return nil, errors.New("Operation failed. Please try again.")
	}
	db.Create(&rev)

	return rev, nil
}

// Edit existing resource.
func (j *Publication) Update(db gorm.DB, req *pb.PublicationRequest) (interface{}, error) {
	publishedAt, err := time.Parse(DateFormat, req.PublishedAt)
	if err != nil {
		publishedAt = time.Now() // default to current UTC date
	}
	db.First(&j, req.Id)
	db.Model(&j).Updates(Publication{
		Volume:      req.Volume,
		No:          req.No,
		Status:      req.Status,
		PublishedAt: publishedAt,
	})
	return j, nil
}
