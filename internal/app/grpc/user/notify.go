package grpc

import (
	"context"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"gitlab.ozon.dev/cyxrop/homework-2/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s UserServiceServer) Notify(ctx context.Context, r *api.ByTelegramUsername) (*empty.Empty, error) {
	user, err := s.service.Get(ctx, r.TelegramUsername)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	errs := s.service.Notify(ctx, user)
	if len(errs) == 0 {
		return &empty.Empty{}, nil
	}

	sb := strings.Builder{}
	for _, err := range errs {
		sb.WriteString(err.Error())
		sb.WriteRune(';')
	}

	return &empty.Empty{}, status.Error(codes.Internal, sb.String())
}
