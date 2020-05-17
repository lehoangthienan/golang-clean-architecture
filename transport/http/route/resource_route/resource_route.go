package user

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	imageDecode "github.com/lehoangthienan/marvel-heroes-backend/transport/http/decode/image"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
)

// Router represents user-specific routings
func Router(endpoints endpoints.Endpoints, options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/{filePath}", httptransport.NewServer(
			endpoints.ImageEndpoint.GetImageFile,
			imageDecode.GetImageFileRequest,
			encode.EncodeFileResponse,
			options...,
		).ServeHTTP)
	}
}
