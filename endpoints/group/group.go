package group

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// MakeCreateEndpoint func
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreateGroup)
		res, err := s.GroupService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeUpdateEndpoint func
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.UpdateGroup)
		res, err := s.GroupService.Update(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeDeleteEndpoint func
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.DeleteGroup)
		res, err := s.GroupService.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeAssignHeroesToGroupEndpoint func
func MakeAssignHeroesToGroupEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.AssignHeroesGroup)
		res, err := s.GroupService.AssignHeroesGroup(ctx, &req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
