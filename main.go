package main

import(
	pb "github.com/metallurgical/journal-go/api/golang"
	svr "github.com/metallurgical/journal-go/server"
	mysqlDB "github.com/metallurgical/journal-go/database"
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
	dbConn, err := mysqlDB.Connect()

	pb.RegisterJournalServer(grpcServer, &svr.JournalServer{DB: dbConn})

	grpcServer.Serve(conn)
}