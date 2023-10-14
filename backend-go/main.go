package main

import (
	"log"
	"net"

	pb "github.com/syahrizalfauzi/go-grpc-tryout/proto"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello World")
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	log.Println("Server is running on port 9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
