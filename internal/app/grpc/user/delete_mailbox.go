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

func (s UserServiceServer) DeleteMailbox(ctx context.Context, r *api.DeleteMailboxRequest) (*empty.Empty, error) {
	err := s.service.DeleteUserMailbox(ctx, r.TelegramUsername, r.Mail)

	if errors.Is(err, repository.ErrNotFound) {
		return nil, status.Error(codes.NotFound, "mailbox with the specified mail not found")
	}

	return &emptypb.Empty{}, err
}
