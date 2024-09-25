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
	users := make(map[int32]*pb.User)
	users[0] = &pb.User{Id: i32Ptr(0)}
	return &UserService{dbUsers: users}
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
	*s.dbUsers[userId].Id = userId
	user := &pb.User{Id: i32Ptr(userId), Name: req.User.Name, Email: req.User.Email}
	s.dbUsers[userId] = user
	return &pb.CreateUserResponse{Id: *user.Id}, nil
}
