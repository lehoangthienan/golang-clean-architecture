package image

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/image"
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

func (mw validatingMiddleware) Create(ctx context.Context, req *req.Images) (*res.Images, error) {
	if req.Images == nil {
		return nil, errors.MissingImagesError
	}

	if len(req.Images) == 0 {
		return nil, errors.LengthImagesError
	}

	return mw.Service.Create(ctx, req)
}
