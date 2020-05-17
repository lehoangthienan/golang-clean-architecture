package image

import (
	"context"
	"os"

	req "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
	res "github.com/lehoangthienan/marvel-heroes-backend/model/response/image"
)

// Service interface
type Service interface {
	Create(context.Context, *req.Images) (*res.Images, error)
	GetImageFile(context.Context, string) (*os.File, error)
}

// Middleware func
type Middleware func(Service) Service
