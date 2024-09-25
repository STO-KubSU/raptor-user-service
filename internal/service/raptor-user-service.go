package service

import (
	"context"
	"fmt"

	pb "github.com/STO-KubSU/raptor-proto/userpb"
)

func i32Ptr(i int32) *int32 {
	return &i
}

type UserService struct {
	pb.UnimplementedUserServiceServer
	dbUsers map[int32]*pb.User
}

func NewUserService() *UserService {
	return &UserService{dbUsers: make(map[int32]*pb.User)}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.dbUsers[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return &pb.GetUserResponse{User: user}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	userId := int32(len(s.dbUsers))
	user := &pb.User{Id: i32Ptr(userId), Name: req.User.Name, Email: req.User.Email}
	s.dbUsers[userId] = user
	return &pb.CreateUserResponse{Id: *user.Id}, nil
}
