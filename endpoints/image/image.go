package image

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// MakeCreateImageEndpoint func
func MakeCreateImageEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		req := r.(request.Images)
		res, err := s.ImageService.Create(ctx, &req)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
}
