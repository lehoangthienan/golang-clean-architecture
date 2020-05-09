package user

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	userDecode "github.com/lehoangthienan/marvel-heroes-backend/transport/http/decode/user"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
)

// Router represents user-specific routings
func Router(endpoints endpoints.Endpoints, options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.UserEndpoint.CreateUser,
			userDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/sign-in", httptransport.NewServer(
			endpoints.UserEndpoint.SignIn,
			userDecode.SignInRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/", httptransport.NewServer(
			endpoints.UserEndpoint.UpdateUser,
			userDecode.UpdateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
