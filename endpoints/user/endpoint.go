package user

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// UserEndpoint struct
type UserEndpoint struct {
	CreateUser endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	SignIn     endpoint.Endpoint
}

// NewEndpoint func
func NewEndpoint(s service.Service) UserEndpoint {
	return UserEndpoint{
		CreateUser: MakeCreateUserEndpoint(s),
		SignIn:     MakeSignInEndpoint(s),
		UpdateUser: MakeUpdateEndpoint(s),
	}
}
