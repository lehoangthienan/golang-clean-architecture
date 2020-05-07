package user

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Repository is interface for user repository
type Repository interface {
	Create(context.Context, *transaction.Pool, *domain.User) (*domain.User, error)
	LogIn(context.Context, *transaction.Pool, *requestModel.SignInUser) (*responseModel.SignInUser, error)
	Update(context.Context, *transaction.Pool, *requestModel.UpdateUser) (*responseModel.UpdateUser, error)
}
