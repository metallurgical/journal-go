package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	pb "github.com/metallurgical/journal-go/api/golang"
	mysqlDB "github.com/metallurgical/journal-go/database"
	svr "github.com/metallurgical/journal-go/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"flag"
	"go/build"
)

func main() {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", 7000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	env  := flag.String("env", "", "Replace default .env path")
	flag.Parse()
	envPath := ""
	if *env == "" {
		envPath = fmt.Sprintf("%s/src/github.com/metallurgical/journal-go/.env", build.Default.GOPATH)
	} else {
		envPath = *env
	}
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file")
	}
	grpcServer := grpc.NewServer()
	dbConn, err := mysqlDB.Connect()
	if err != nil {
		log.Fatalf("Error connecting to database. Please check database crendetials.")
	}
	pb.RegisterJournalServer(grpcServer, &svr.JournalServer{DB: dbConn})

	grpcServer.Serve(conn)
}
