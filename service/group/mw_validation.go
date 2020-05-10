package group

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
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

func (mw validatingMiddleware) Create(ctx context.Context, req req.CreateGroup) (*res.CreateGroup, error) {
	if req.Name == "" {
		return nil, errors.MissingNameGroupError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validatingMiddleware) Update(ctx context.Context, req req.UpdateGroup) (*res.UpdateGroup, error) {
	if req.Name != "" && len(req.Name) < 6 {
		return nil, errors.LengthNameGroupError
	}

	return mw.Service.Update(ctx, req)
}
