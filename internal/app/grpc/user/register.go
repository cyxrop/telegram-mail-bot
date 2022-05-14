package grpc

import (
	"context"
	"errors"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s UserServiceServer) Register(ctx context.Context, r *api.RegisterUserRequest) (*api.ID, error) {
	ID, err := s.service.Register(ctx, r.TelegramUsername, r.TelegramChatId)
	if errors.Is(err, repository.ErrUniqueViolation) {
		return nil, status.Error(codes.AlreadyExists, "user with the specified telegram username already exists")
	}

	return &api.ID{
		ID: ID,
	}, err
}
