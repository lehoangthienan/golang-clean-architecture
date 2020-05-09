package options

import (
	"context"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc/metadata"
)

func LogRequestInfo(logger log.Logger) func(ctx context.Context, m metadata.MD) context.Context {
	return func(ctx context.Context, m metadata.MD) context.Context {
		auth := m.Get("request")
		if len(auth) > 0 {
			logger.Log("Request", auth[0])
		}
		return ctx
	}
}
