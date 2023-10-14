package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	pb "github.com/syahrizalfauzi/go-grpc-tryout/proto"
)

type server struct {
	pb.UnimplementedUserServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {

	log.Printf("Received login request")
	return &pb.LoginResponse{
		Token: "tokenasdasdasd",
		Profile: &pb.UserProfile{
			Username: in.Username,
			Email:    in.Username + "@domain.com",
			Phone:    fmt.Sprint(rand.Intn(9999999999)),
			Address:  "Jl. Jalan",
		},
	}, nil
}
