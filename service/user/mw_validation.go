package user

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	"github.com/lehoangthienan/marvel-heroes-backend/util/constants"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
)

type validatingMiddleware struct {
	Service
}

// ValidatingMiddleware func
func ValidatingMiddleware() Middleware {
	return func(next Service) Service {
		return &validatingMiddleware{Service: next}
	}
}

func (mw validatingMiddleware) Create(ctx context.Context, req req.CreateUser) (*res.CreateUser, error) {
	if req.Name == "" {
		return nil, errors.MissingNameUserError
	}

	if req.UserName == "" {
		return nil, errors.MissingUserNameError
	}

	if req.PassWord == "" {
		return nil, errors.MissingPasswordError
	}

	if req.Role == "" {
		return nil, errors.MissingRoleError
	}

	if req.Role != constants.ADMIN && req.Role != constants.MODERATOR {
		return nil, errors.WrongRoleError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validatingMiddleware) Update(ctx context.Context, req req.UpdateUser) (*res.UpdateUser, error) {
	if req.Name != nil && len(*req.Name) < 6 {
		return nil, errors.LengthNameError
	}

	if req.UserName != nil && len(*req.UserName) < 6 {
		return nil, errors.LengthUsernameError
	}

	if req.PassWord != nil && len(*req.PassWord) < 6 {
		return nil, errors.LengthPasswordError
	}

	if req.Role != nil && *req.Role != constants.ADMIN && *req.Role != constants.MODERATOR {
		return nil, errors.WrongRoleError
	}

	return mw.Service.Update(ctx, req)
}
