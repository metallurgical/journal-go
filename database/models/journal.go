package models

import (
	"errors"
	goslug "github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	"github.com/metallurgical/journal-go/server"
	"time"
)

type Journal struct {
	ID          uint `gorm:"primary_key"`
	Title       string
	Slug        string `gorm:"unique"`
	Overview    string
	Status      int32
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Get existing journal from database.
func (j *Journal) GetJournal(db gorm.DB, id int64) (error) {
	if err := db.First(j, id).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}

// Create new journal resource.
func (j *Journal) CreateJournal(s *server.JournalServer, req *pb.JournalCreateRequest) error {
	var slug string

	slug = goslug.Make(req.Title)
	journal := &Journal{
		Title:    req.Title,
		Slug:     slug,
		Overview: req.Overview,
		Status:   0,
	}
	if !s.DB.NewRecord(journal) {
		return errors.New("Operation failed. Please try again.")
	}
	return nil
}
