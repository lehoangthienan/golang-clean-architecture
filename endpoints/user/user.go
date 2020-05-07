package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// MakeCreateUserEndpoint func
func MakeCreateUserEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreateUser)
		res, err := s.UserService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeSignInEndpoint func
func MakeSignInEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.SignInUser)
		res, err := s.UserService.LogIn(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeUpdateEndpoint func
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.UpdateUser)
		res, err := s.UserService.Update(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
