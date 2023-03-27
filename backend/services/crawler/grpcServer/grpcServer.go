package grpcserver

import (
	"log"
	"net"

	pb "github.com/neil-go-phan/Football-news-aggregation/backend/grpcfile"
	"google.golang.org/grpc"
)

func GRPCServer() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}


}