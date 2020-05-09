package auth

import (
	"github.com/go-kit/kit/endpoint"
)

// AuthMiddleware struct
type AuthMiddleware struct {
	AuthModerator endpoint.Middleware
	AuthAdmin     endpoint.Middleware
}
