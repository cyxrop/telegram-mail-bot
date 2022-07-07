package grpc

import (
	service "gitlab.ozon.dev/cyxrop/homework-2/internal/app/service/user"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
)

type UserServiceServer struct {
	api.UnimplementedUserServiceServer

	service service.UserService
}

func NewUserServiceServer(s service.UserService) *UserServiceServer {
	return &UserServiceServer{
		service: s,
	}
}
