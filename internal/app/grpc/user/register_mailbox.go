package grpc

import (
	"context"
	"errors"

	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/models"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s UserServiceServer) RegisterMailbox(ctx context.Context, r *api.RegisterMailboxRequest) (*api.ID, error) {
	ID, err := s.service.CreateUserMailbox(ctx, r.TelegramUsername, models.Mailbox{
		Mail:     r.Mail,
		Password: r.Password,
	})

	if errors.Is(err, repository.ErrUniqueViolation) {
		return nil, status.Error(codes.AlreadyExists, "mailbox with the specified mail already exists")
	}

	if errors.Is(err, repository.ErrNotNullViolation) {
		return nil, status.Error(codes.NotFound, "user with the specified telegram username not found")
	}

	return &api.ID{
		ID: ID,
	}, err
}
