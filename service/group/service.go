package group

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
)

// Service interface
type Service interface {
	Create(context.Context, req.CreateGroup) (*res.CreateGroup, error)
	Update(context.Context, req.UpdateGroup) (*res.UpdateGroup, error)
	Delete(context.Context, req.DeleteGroup) (*res.DeleteGroup, error)
	// GetHeroes(context.Context, req.GetGroupHeroes) (*res.GetHeroes, error)
}

// Middleware func
type Middleware func(Service) Service
