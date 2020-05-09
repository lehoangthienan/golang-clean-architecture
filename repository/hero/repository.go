package hero

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Repository is interface for user repository
type Repository interface {
	Create(context.Context, *transaction.Pool, *domain.Hero) (*domain.Hero, error)
	Update(context.Context, *transaction.Pool, *requestModel.UpdateHero) (*responseModel.UpdateHero, error)
	Delete(context.Context, *transaction.Pool, *requestModel.DeleteHero) (*responseModel.DeleteHero, error)
	GetHeroes(context.Context, *requestModel.GetHeroes) (*responseModel.GetHeroes, error)
}
