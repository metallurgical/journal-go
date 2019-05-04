package server

import (
	"github.com/jinzhu/gorm"
	pb "github.com/metallurgical/journal-go/api/golang"
	journal "github.com/metallurgical/journal-go/database/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	"fmt"
)

type JournalServer struct {
	DB gorm.DB
}

func (s *JournalServer) Publish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)
	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

func (s *JournalServer) UnPublish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

func (s *JournalServer) Approve(ctx context.Context, req *pb.JournalApproveRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)

	tx := s.DB.Begin()
	journal := &journal.Journal{}
	if err:= s.DB.First(&journal, req.Id).Error; gorm.IsRecordNotFoundError(err) {
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
		tx.Rollback();
		return &pb.JournalResponse{Flag: "error", Message: "Operation failed, please try again"}, nil
	}
	tx.Commit()
	fmt.Print(journal)
	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}