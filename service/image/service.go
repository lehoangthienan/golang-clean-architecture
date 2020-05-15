package image

import (
	"context"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/image"
)

// Service interface
type Service interface {
	Create(context.Context, *req.Images) (*res.Images, error)
}

// Middleware func
type Middleware func(Service) Service
