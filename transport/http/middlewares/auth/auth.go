package auth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// MakeAuthModeratorMiddleware func
func MakeAuthModeratorMiddleware(s service.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			err := s.AuthService.AuthenticateModerator(ctx)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}

// MakeAuthAdminMiddleware func
func MakeAuthAdminMiddleware(s service.Service) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			err := s.AuthService.AuthenticateAdmin(ctx)
			if err != nil {
				return nil, err
			}
			return next(ctx, req)
		}
	}
}
