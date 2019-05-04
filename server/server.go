package server

import (
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	journal "github.com/metallurgical/journal-go/database/models"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	_ "google.golang.org/grpc/status"
	"context"
)

type JournalServer struct {
	DB gorm.DB
}

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
	journal.Status = req.Status;
	if err := s.DB.Save(&journal).Error; err != nil {
		return &pb.JournalResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

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
		return &pb.JournalResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

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
		return &pb.JournalResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}