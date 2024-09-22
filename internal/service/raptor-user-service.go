package service

import (
	"context"

	pb "github.com/STO-KubSU/raptor-proto/userpb"
)

func i32Ptr(i int32) *int32 {
	return &i
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user := &pb.User{Id: i32Ptr(req.Id), Name: "Alice"}
	return &pb.GetUserResponse{User: user}, nil
}
