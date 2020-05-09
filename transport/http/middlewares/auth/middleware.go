package auth

import (
	"github.com/go-kit/kit/endpoint"
)

type AuthMiddleware struct {
	AuthUser  endpoint.Middleware
	AuthAdmin endpoint.Middleware
}
