package http

import (
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/encode"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/options"
	userRoute "github.com/lehoangthienan/marvel-heroes-backend/transport/http/route/user"
	"github.com/lehoangthienan/marvel-heroes-backend/util/helper"
)

// NewHTTPHandler func
func NewHTTPHandler(
	endpoints endpoints.Endpoints,
	logger log.Logger,
) http.Handler {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})
	r.Use(cors.Handler)
	options := []httpTransport.ServerOption{
		httpTransport.ServerBefore(
			options.LogRequestInfo(logger),
			helper.VerifyToken,
		),
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(encode.EncodeError),
	}

	r.Get("/", httpTransport.NewServer(
		endpoints.IndexEndpoint,
		httpTransport.NopRequestDecoder,
		httpTransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/users", userRoute.Router(endpoints, options))

	return r
}
