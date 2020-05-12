package group

import (
	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	groupDecode "github.com/lehoangthienan/marvel-heroes-backend/transport/http/decode/group"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/middlewares"
)

// Router represents groups-specific routings
func Router(
	endpoints endpoints.Endpoints,
	middlewares middlewares.Middlewares,
	options []httptransport.ServerOption) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.GroupEndpoint.CreateGroup,
			),
			groupDecode.CreateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/assign", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthModerator(
				endpoints.GroupEndpoint.AssignHeroesToGroup,
			),
			groupDecode.AssignHeroesToGroupRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{groupID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthModerator(
				endpoints.GroupEndpoint.UpdateGroup,
			),
			groupDecode.UpdateRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{groupID}", httptransport.NewServer(
			middlewares.AuthMiddleware.AuthAdmin(
				endpoints.GroupEndpoint.DeleteGroup,
			),
			groupDecode.DeleteRequest,
			encode.EncodeResponse,
			options...,
		).ServeHTTP)
	}
}
