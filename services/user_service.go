package services

import (
	"context"
	"log"

	pb "app/protos/user"
)

// UserService 用户服务
type UserService struct{}

// Index 用户列表
func (s *UserService) Index(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}

// Store 用户添加
func (s *UserService) Store(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}

// Show ...
func (s *UserService) Show(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}

// Update ...
func (s *UserService) Update(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}

// Destroy ...
func (s *UserService) Destroy(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}
