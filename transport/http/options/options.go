package options

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/lehoangthienan/marvel-heroes-backend/util/constants"
)

// LogRequestInfo func
func LogRequestInfo(logger log.Logger) func(ctx context.Context, req *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		logger.Log("Method", req.Method, "Route", req.RequestURI)
		return ctx
	}
}

// SetAccessTokenToContext func
func SetAccessTokenToContext(ctx context.Context, req *http.Request) context.Context {
	accessToken := req.Header.Get("Authorization")
	ctx = context.WithValue(ctx, constants.TokenCTXKey, accessToken)
	return ctx
}
