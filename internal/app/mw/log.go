package mw

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)
	if err != nil {
		log.Println(info.FullMethod, req, err)
	} else {
		log.Println(info.FullMethod, req)
	}

	return resp, err
}
