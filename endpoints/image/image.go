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

// MakeGetImageFile func
func MakeGetImageFile(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var path = request.(string)
		file, err := s.ImageService.GetImageFile(ctx, path)

		if err != nil {
			return nil, err
		}

		return file, nil
	}
}
