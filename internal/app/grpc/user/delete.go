package grpc

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.ozon.dev/cyxrop/homework-2/internal/app/repository"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s UserServiceServer) DeleteByTelegramUsername(ctx context.Context, r *api.ByTelegramUsername) (*empty.Empty, error) {
	err := s.service.DeleteByTelegramUsername(ctx, r.TelegramUsername)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "user with the specified telegram username not found")
	}

	return &emptypb.Empty{}, err
}
