package hero

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// MakeCreateHeroEndpoint func
func MakeCreateHeroEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.CreateHero)
		res, err := s.HeroService.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeUpdateEndpoint func
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.UpdateHero)
		res, err := s.HeroService.Update(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeDeleteEndpoint func
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.DeleteHero)
		res, err := s.HeroService.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}

// MakeGetsEndpoint func
func MakeGetsEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.GetHeroes)
		res, err := s.HeroService.GetHeroes(ctx, req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
