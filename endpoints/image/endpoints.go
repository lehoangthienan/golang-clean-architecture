package image

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// ImageEndpoint struct
type ImageEndpoint struct {
	CreateImage  endpoint.Endpoint
	GetImageFile endpoint.Endpoint
}

// NewEndpoint func
func NewEndpoint(s service.Service) ImageEndpoint {
	return ImageEndpoint{
		CreateImage:  MakeCreateImageEndpoint(s),
		GetImageFile: MakeGetImageFile(s),
	}
}
