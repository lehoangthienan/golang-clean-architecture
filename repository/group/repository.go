package group

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Repository is interface for user repository
type Repository interface {
	Create(context.Context, *transaction.Pool, *domain.Group) (*domain.Group, error)
	Update(context.Context, *transaction.Pool, *requestModel.UpdateGroup) (*responseModel.UpdateGroup, error)
	Delete(context.Context, *transaction.Pool, *requestModel.DeleteGroup) (*responseModel.DeleteGroup, error)
	AssignHeroesGroup(ctx context.Context, pool *transaction.Pool, req []*domain.GroupHero) ([]*domain.GroupHero, error)
}
