package user

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	repo "github.com/lehoangthienan/marvel-heroes-backend/repository"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type userService struct {
	repo repo.Repository
	tx   transaction.TXService
}

// NewService func
func NewService(repo repo.Repository, tx transaction.TXService) Service {
	return &userService{repo: repo, tx: tx}
}

func (s *userService) Create(ctx context.Context, req requestModel.CreateUser) (*responseModel.CreateUser, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.UserRepository.Create(ctx, pool, &domain.User{
		Name:     req.Name,
		UserName: req.UserName,
		PassWord: req.PassWord,
		Role:     req.Role,
	})

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.CreateUserFailedError
	}

	return &responseModel.CreateUser{User: user}, s.tx.TXCommit(pool)
}

func (s *userService) LogIn(ctx context.Context, req requestModel.SignInUser) (*responseModel.SignInUser, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.UserRepository.LogIn(ctx, pool, &requestModel.SignInUser{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.CreateUserFailedError
	}

	return &responseModel.SignInUser{User: user.User, Token: user.Token}, s.tx.TXCommit(pool)
}

func (s *userService) Update(ctx context.Context, req requestModel.UpdateUser) (*responseModel.UpdateUser, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.UserRepository.Update(ctx, pool, &req)

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.UpdateUserFailedError
	}

	return &responseModel.UpdateUser{User: user.User}, s.tx.TXCommit(pool)
}
