package hero

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	heroDecode "github.com/lehoangthienan/marvel-heroes-backend/transport/http/decode/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/middlewares"
)

// Router represents heroes-specific routings
func Router(
	endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.HeroEndpoint.GetHeros,
			heroDecode.GetHeroesRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.HeroEndpoint.CreateHero,
			),
			heroDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{heroID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthModerator(
				endpoints.HeroEndpoint.UpdateHero,
			),
			heroDecode.UpdateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{heroID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.HeroEndpoint.DeleteHero,
			),
			heroDecode.DeleteRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
