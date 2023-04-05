package handlers

import (
	"log"
	"net"

	pb "crawler/proto"

	"google.golang.org/grpc"
)

type gRPCServer struct {
	pb.UnimplementedCrawlerServiceServer
}

func GRPCServer(port string) {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	pb.RegisterCrawlerServiceServer(s, &gRPCServer{})
	log.Println("start server")
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
