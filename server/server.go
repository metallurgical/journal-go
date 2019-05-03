package server

import (
	pb "github.com/metallurgical/journal-go/api/golang"
	"context"
	"fmt"
)

type Journal struct {}

func (server *Journal) Publish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

func (server *Journal) UnPublish(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}

func (server *Journal) Approve(ctx context.Context, req *pb.JournalRequest) (*pb.JournalResponse, error) {
	fmt.Printf("Received request %s", req.Id)

	return &pb.JournalResponse{Flag: "success", Message: "success"}, nil
}