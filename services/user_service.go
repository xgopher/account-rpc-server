package services

import (
	"context"
	"log"

	pb "app/protos/user"
)

// UserService 用户服务
type UserService struct{}

// Store implements helloworld.GreeterServer
func (s *UserService) Store(ctx context.Context, in *pb.UserRequest) (*pb.UserReply, error) {
	log.Printf("Received: %v", in.Username)
	return &pb.UserReply{Message: "1XX " + in.Username}, nil
}
