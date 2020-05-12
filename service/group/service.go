package group

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
)

// Service interface
type Service interface {
	Create(context.Context, req.CreateGroup) (*res.CreateGroup, error)
	Update(context.Context, req.UpdateGroup) (*res.UpdateGroup, error)
	Delete(context.Context, req.DeleteGroup) (*res.DeleteGroup, error)
	AssignHeroesGroup(ctx context.Context, req *req.AssignHeroesGroup) ([]*domain.GroupHero, error)
}

// Middleware func
type Middleware func(Service) Service
