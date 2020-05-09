package hero

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
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

func (mw validatingMiddleware) Create(ctx context.Context, req req.CreateHero) (*res.CreateHero, error) {
	if req.Name == "" {
		return nil, errors.MissingNameHeroError
	}

	if req.Power == "" {
		return nil, errors.MissingHeroPowerError
	}

	return mw.Service.Create(ctx, req)
}

func (mw validatingMiddleware) Update(ctx context.Context, req req.UpdateHero) (*res.UpdateHero, error) {
	if req.Name != "" && len(req.Name) < 6 {
		return nil, errors.LengthNameHeroError
	}

	if req.Power != "" && len(req.Power) < 6 {
		return nil, errors.LengthHeroPowerError
	}

	return mw.Service.Update(ctx, req)
}
