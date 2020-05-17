package user

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

// Make sure RepositoryMock implement Repository interface
var _ Repository = &RepositoryMock{}

// RepositoryMock struct
type RepositoryMock struct {
	CreateFunc func(ctx context.Context, pool *transaction.Pool, user *domain.User) (*domain.User, error)
	LogInFunc  func(ctx context.Context, pool *transaction.Pool, req *requestModel.SignInUser) (*responseModel.SignInUser, error)
	UpdateFunc func(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateUser) (*responseModel.UpdateUser, error)
}

// Create function
func (rm *RepositoryMock) Create(ctx context.Context, pool *transaction.Pool, user *domain.User) (*domain.User, error) {
	if rm.CreateFunc == nil {
		panic("RepositoryMock not declare Create function")
	}
	return rm.CreateFunc(ctx, pool, user)
}

// LogIn function
func (rm *RepositoryMock) LogIn(ctx context.Context, pool *transaction.Pool, req *requestModel.SignInUser) (*responseModel.SignInUser, error) {
	if rm.LogInFunc == nil {
		panic("RepositoryMock not declare LogIn function")
	}
	return rm.LogInFunc(ctx, pool, req)
}

// Update function
func (rm *RepositoryMock) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateUser) (*responseModel.UpdateUser, error) {
	if rm.UpdateFunc == nil {
		panic("RepositoryMock not declare Update function")
	}
	return rm.UpdateFunc(ctx, pool, req)
}
