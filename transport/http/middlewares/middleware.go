package middlewares

import (
	"github.com/lehoangthienan/marvel-heroes-backend/service"
	"github.com/lehoangthienan/marvel-heroes-backend/transport/http/middlewares/auth"
)

// Middlewares struct
type Middlewares struct {
	AuthMiddleware auth.AuthMiddleware
}

// MakeHTTPpMiddleware func
func MakeHTTPpMiddleware(s service.Service) Middlewares {
	return Middlewares{
		AuthMiddleware: auth.AuthMiddleware{
			AuthModerator: auth.MakeAuthModeratorMiddleware(s),
			AuthAdmin:     auth.MakeAuthAdminMiddleware(s),
		},
	}
}
