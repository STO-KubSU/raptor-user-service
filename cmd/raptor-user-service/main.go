package main

import (
	"log"
	"net"

	pb "github.com/STO-KubSU/raptor-proto/userpb"
	"github.com/STO-KubSU/raptor-user-service/internal/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{})

	log.Println("User service is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
