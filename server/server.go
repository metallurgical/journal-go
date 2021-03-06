package server

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	journal "github.com/metallurgical/journal-go/database/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type JournalServer struct {
	DB gorm.DB
}

type RevisionServer struct {
	DB gorm.DB
}

type PublicationServer struct {
	DB gorm.DB
}

// Publish approved journal.
func (s *JournalServer) Publish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	if journal.Status == 2 {
		return nil, status.Errorf(codes.PermissionDenied, "No action taken as journal already published.")
	}
	if journal.Status != 1 {
		return nil, status.Errorf(codes.PermissionDenied, "Only approved journal that was able to publish")
	}
	dateNow := time.Now()
	journal.Status = req.Status
	journal.PublishedAt = &dateNow

	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Un-publish published journal.
func (s *JournalServer) UnPublish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	if journal.Status == 3 {
		return nil, status.Errorf(codes.PermissionDenied, "No action taken as journal already unpublished.")
	}
	if journal.Status != 2 {
		return nil, status.Errorf(codes.PermissionDenied, "Only published journal that was able to unpublish.")
	}
	journal.Status = req.Status;
	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Approve pending journal.
func (s *JournalServer) Approve(ctx context.Context, req *pb.JournalApproveRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	if journal.Status == 1 {
		return nil, status.Errorf(codes.PermissionDenied, "No action taken as journal already approved.")
	}
	journal.Status = req.Status;
	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Delete journal.
func (s *JournalServer) Destroy(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Resource not found. Please try again"}, nil
	}

	if err := s.DB.Delete(&journal).Error; err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Create Journal.
func (s *JournalServer) Create(ctx context.Context, req *pb.JournalCreateRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	createdJournal, err := journal.CreateJournal(s.DB, req);
	if err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: err.Error(),
		}, nil
	}
	resp, err := json.Marshal(&createdJournal)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Operation failed, please try again")
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
		Journal: string(resp),
	}, nil
}

// Edit Journal.
func (s *JournalServer) Edit(ctx context.Context, req *pb.JournalEditRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	editedJournal, err := journal.Edit(s.DB, req)
	if err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: err.Error(),
		}, nil
	}
	resp, err := json.Marshal(&editedJournal)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Operation failed, please try again")
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
		Journal: string(resp),
	}, nil
}

// Review journal.
func (s *JournalServer) Review(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	if journal.Status == 4 {
		return nil, status.Errorf(codes.PermissionDenied, "No action taken as journal already in-review.")
	}
	journal.Status = req.Status;
	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Reject journal.
func (s *JournalServer) Reject(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	journal := &journal.Journal{}
	if err := journal.GetJournal(s.DB, req.Id); err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	if journal.Status == 5 {
		return nil, status.Errorf(codes.PermissionDenied, "No action taken as journal already rejected.")
	}
	journal.Status = req.Status;
	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.JournalResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// Create revision
func (s *RevisionServer) Create(ctx context.Context, req *pb.RevisionRequest) (*pb.RevisionResponse, error) {
	rev := journal.Revision{}
	createdRev, err := rev.Create(s.DB, req)
	if err != nil {
		return &pb.RevisionResponse{
			Flag:    "error",
			Message: err.Error(),
		}, nil
	}
	resp, err := json.Marshal(&createdRev)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Operation failed, please try again")
	}

	return &pb.RevisionResponse{
		Flag:     "success",
		Message:  "success",
		Revision: string(resp),
	}, nil
}

// Create publication
func (s *PublicationServer) Create(ctx context.Context, req *pb.PublicationRequest) (*pb.PublicationResponse, error) {
	pub := journal.Publication{}
	createdRev, err := pub.Create(s.DB, req)
	if err != nil {
		return &pb.PublicationResponse{
			Flag:    "error",
			Message: err.Error(),
		}, nil
	}
	resp, err := json.Marshal(&createdRev)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Operation failed, please try again")
	}

	return &pb.PublicationResponse{
		Flag:        "success",
		Message:     "success",
		Publication: string(resp),
	}, nil
}

// Update publication
func (s *PublicationServer) Update(ctx context.Context, req *pb.PublicationRequest) (*pb.PublicationResponse, error) {
	pub := &journal.Publication{}
	editedPublication, err := pub.Update(s.DB, req)
	if err != nil {
		return &pb.PublicationResponse{
			Flag:    "error",
			Message: err.Error(),
		}, nil
	}
	resp, err := json.Marshal(&editedPublication)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Operation failed, please try again")
	}

	return &pb.PublicationResponse{
		Flag:    "success",
		Message: "success",
		Publication: string(resp),
	}, nil
}

// Delete publication
func (s *PublicationServer) Delete(ctx context.Context, req *pb.PublicationRequest) (*pb.PublicationResponse, error) {
	pub := &journal.Publication{}
	if err := pub.Get(s.DB, req.Id); err != nil {
		return &pb.PublicationResponse{
			Flag:    "error",
			Message: "Resource not found. Please try again"}, nil
	}

	if err := s.DB.Delete(&pub).Error; err != nil {
		return &pb.PublicationResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}

	return &pb.PublicationResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}

// SetStatus change publication's status
func (s *PublicationServer) SetStatus(ctx context.Context, req *pb.PublicationRequest) (*pb.PublicationResponse, error) {
	pub := &journal.Publication{}
	if err := pub.Get(s.DB, req.Id); err != nil {
		return &pb.PublicationResponse{Flag: "error", Message: "Resource not found. Please try again"}, nil
	}
	pub.Status = req.Status;
	if err := s.DB.Save(&pub).Error; err != nil {
		return &pb.PublicationResponse{
			Flag:    "error",
			Message: "Operation failed, please try again",
		}, nil
	}

	return &pb.PublicationResponse{
		Flag:    "success",
		Message: "success",
	}, nil
}
