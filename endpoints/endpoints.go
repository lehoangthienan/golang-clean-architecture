package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints/index"
	"github.com/lehoangthienan/marvel-heroes-backend/endpoints/user"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// Endpoints struct
type Endpoints struct {
	IndexEndpoint endpoint.Endpoint
	UserEndpoint  user.UserEndpoint
}

// MakeServerEndpoints func
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		IndexEndpoint: index.MakeIndexEndpoints(),
		UserEndpoint:  user.NewEndpoint(s),
	}
}
