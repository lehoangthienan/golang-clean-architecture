package group

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// GroupEndpoint struct
type GroupEndpoint struct {
	CreateGroup         endpoint.Endpoint
	UpdateGroup         endpoint.Endpoint
	DeleteGroup         endpoint.Endpoint
	AssignHeroesToGroup endpoint.Endpoint
}

// NewEndpoint func
func NewEndpoint(s service.Service) GroupEndpoint {
	return GroupEndpoint{
		CreateGroup:         MakeCreateEndpoint(s),
		UpdateGroup:         MakeUpdateEndpoint(s),
		DeleteGroup:         MakeDeleteEndpoint(s),
		AssignHeroesToGroup: MakeAssignHeroesToGroupEndpoint(s),
	}
}
