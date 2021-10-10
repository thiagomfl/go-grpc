package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/thiagomfl/go-grpc/pb"
	"github.com/thiagomfl/go-grpc/services"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
