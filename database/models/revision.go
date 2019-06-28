package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	"time"
)

type Revision struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	Remarks   string
	JournalId int64
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (Revision) TableName() string {
	return "document_revisions"
}

// Get existing resource from database.
func (j *Revision) Get(db gorm.DB, id int64) (error) {
	if err := db.First(j, id).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}

// Create new resource.
func (j *Revision) Create(db gorm.DB, req *pb.RevisionRequest) (interface{}, error) {
	rev := &Revision{
		Name:      req.Name,
		Remarks:   req.Remarks,
		JournalId: req.JournalId,
	}
	if !db.NewRecord(&rev) {
		return nil, errors.New("Operation failed. Please try again.")
	}
	db.Create(&rev)

	return rev, nil
}
