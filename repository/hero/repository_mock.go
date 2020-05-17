package hero

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	CreateFunc    func(ctx context.Context, pool *transaction.Pool, hero *domain.Hero) (*domain.Hero, error)
	UpdateFunc    func(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateHero) (*responseModel.UpdateHero, error)
	DeleteFunc    func(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteHero) (*responseModel.DeleteHero, error)
	GetHeroesFunc func(ctx context.Context, req *requestModel.GetHeroes) (*responseModel.GetHeroes, error)
}

// Create function
func (rm *RepositoryMock) Create(ctx context.Context, pool *transaction.Pool, hero *domain.Hero) (*domain.Hero, error) {
	if rm.CreateFunc == nil {
		panic("RepositoryMock not declare Create function")
	}
	return rm.CreateFunc(ctx, pool, hero)
}

// Update function
func (rm *RepositoryMock) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateHero) (*responseModel.UpdateHero, error) {
	if rm.UpdateFunc == nil {
		panic("RepositoryMock not declare Update function")
	}
	return rm.UpdateFunc(ctx, pool, req)
}

// Delete function
func (rm *RepositoryMock) Delete(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteHero) (*responseModel.DeleteHero, error) {
	if rm.DeleteFunc == nil {
		panic("RepositoryMock not declare Delete function")
	}
	return rm.DeleteFunc(ctx, pool, req)
}

// GetHeroes function
func (rm *RepositoryMock) GetHeroes(ctx context.Context, req *requestModel.GetHeroes) (*responseModel.GetHeroes, error) {
	if rm.GetHeroesFunc == nil {
		panic("RepositoryMock not declare GetHeroes function")
	}
	return rm.GetHeroesFunc(ctx, req)
}
