package group

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	CreateFunc            func(ctx context.Context, pool *transaction.Pool, group *domain.Group) (*domain.Group, error)
	UpdateFunc            func(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateGroup) (*responseModel.UpdateGroup, error)
	DeleteFunc            func(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteGroup) (*responseModel.DeleteGroup, error)
	AssignHeroesGroupFunc func(ctx context.Context, pool *transaction.Pool, groupHeros []*domain.GroupHero) ([]*domain.GroupHero, error)
}

// Create function
func (rm *RepositoryMock) Create(ctx context.Context, pool *transaction.Pool, group *domain.Group) (*domain.Group, error) {
	if rm.CreateFunc == nil {
		panic("RepositoryMock not declare Create function")
	}
	return rm.CreateFunc(ctx, pool, group)
}

// Update function
func (rm *RepositoryMock) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateGroup) (*responseModel.UpdateGroup, error) {
	if rm.UpdateFunc == nil {
		panic("RepositoryMock not declare Update function")
	}
	return rm.UpdateFunc(ctx, pool, req)
}

// Delete function
func (rm *RepositoryMock) Delete(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteGroup) (*responseModel.DeleteGroup, error) {
	if rm.DeleteFunc == nil {
		panic("RepositoryMock not declare Delete function")
	}
	return rm.DeleteFunc(ctx, pool, req)
}

// AssignHeroesGroup function
func (rm *RepositoryMock) AssignHeroesGroup(ctx context.Context, pool *transaction.Pool, groupHeros []*domain.GroupHero) ([]*domain.GroupHero, error) {
	if rm.AssignHeroesGroupFunc == nil {
		panic("RepositoryMock not declare AssignHeroes function")
	}
	return rm.AssignHeroesGroupFunc(ctx, pool, groupHeros)
}
