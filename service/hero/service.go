package hero

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
)

// Service interface
type Service interface {
	Create(context.Context, req.CreateHero) (*res.CreateHero, error)
	Update(context.Context, req.UpdateHero) (*res.UpdateHero, error)
	Delete(context.Context, req.DeleteHero) (*res.DeleteHero, error)
	GetHeroes(context.Context, req.GetHeroes) (*res.GetHeroes, error)
}

// Middleware func
type Middleware func(Service) Service
