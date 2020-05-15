package image

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	userDecode "github.com/lehoangthienan/marvel-heroes-backend/transport/http/decode/image"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
)

// Router represents user-specific routings
func Router(endpoints endpoints.Endpoints, options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.ImageEndpoint.CreateImage,
			userDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
