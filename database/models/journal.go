package models

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	goslug "github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	intrand "math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Journal struct {
	ID                uint `gorm:"primary_key"`
	ReferenceNo       interface{}
	UserId            uint64
	JournalCategoryId uint64
	Title             string
	Slug              string `gorm:"unique"`
	Overview          string
	Status            int32
	PublishedAt       *time.Time
	CreatedAt         time.Time
	UpdatedAt         *time.Time
}

// Get existing journal from database.
func (j *Journal) GetJournal(db gorm.DB, id int64) (error) {
	if err := db.First(j, id).Error; gorm.IsRecordNotFoundError(err) {
		return err
	}
	return nil
}

// Create new journal resource.
func (j *Journal) CreateJournal(db gorm.DB, req *pb.JournalCreateRequest) (interface{}, error) {
	var originalSlug, slug string
	var count = 0

	originalSlug = goslug.Make(req.Title)
	db.Model(&Journal{}).Where("slug = ?", originalSlug).Count(&count)
	if count > 0 {
		strSlug := make(chan string)
		go randToken(10, strSlug)
		slug = <-strSlug
		originalSlug = originalSlug + strings.ToLower(slug)
	}

	ref := make(chan int)
	go randRef(db, ref)
	refNo := <-ref

	journal := &Journal{
		ReferenceNo:       os.Getenv("JOURNAL_PREFIX_NO") + strconv.Itoa(refNo),
		Title:             req.Title,
		Slug:              originalSlug,
		Overview:          req.Overview,
		UserId:            req.UserId,
		JournalCategoryId: req.JournalCategoryId,
		Status:            0,
	}
	if !db.NewRecord(&journal) {
		return nil, errors.New("Operation failed. Please try again.")
	}
	db.Create(&journal)

	return journal, nil
}

// Generate random string.
func randToken(length int, strSlug chan<- string) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	strSlug <- base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

// Generate random number.
func randRef(db gorm.DB, ref chan<- int) {
	intrand.Seed(time.Now().UnixNano())
	max := 10000000000
	min := 1111111111
	refNo := min + intrand.Intn(max-min)
	count := 0
	db.Model(&Journal{}).Where("reference_no = ?", refNo).Count(&count)

	if count > 0 {
		randRef(db, ref)
	}
	ref <- refNo
}
