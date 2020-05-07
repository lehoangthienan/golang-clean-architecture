package user

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
)

// Service interface
type Service interface {
	Create(context.Context, req.CreateUser) (*res.CreateUser, error)
	LogIn(context.Context, req.SignInUser) (*res.SignInUser, error)
	Update(context.Context, req.UpdateUser) (*res.UpdateUser, error)
}

// Middleware func
type Middleware func(Service) Service
