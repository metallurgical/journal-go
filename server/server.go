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
	if journal.Status != 0 {
		return nil, status.Errorf(codes.PermissionDenied, "Only pending journal that was able to approved.")
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
