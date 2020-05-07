package grpc

import (
	"github.com/go-kit/kit/log"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	"google.golang.org/grpc"
)

func NewGRPCHandler(
	endpoints endpoints.Endpoints,
	logger log.Logger,
	grpcServer *grpc.Server,
) {
}
