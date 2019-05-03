package main

import(
	pb "github.com/metallurgical/journal-go/api/golang"
	svr "github.com/metallurgical/journal-go/server"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"google.golang.org/grpc"
	"fmt"
	"net"
	"log"
)

func main() {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", 7000))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterJournalServer(grpcServer, &svr.Journal{})

	grpcServer.Serve(conn)
}