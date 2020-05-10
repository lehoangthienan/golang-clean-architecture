package hero

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lehoangthienan/marvel-heroes-backend/service"
)

// HeroEndpoint struct
type HeroEndpoint struct {
	CreateHero endpoint.Endpoint
	UpdateHero endpoint.Endpoint
	DeleteHero endpoint.Endpoint
	GetHeros   endpoint.Endpoint
}

// NewEndpoint func
func NewEndpoint(s service.Service) HeroEndpoint {
	return HeroEndpoint{
		CreateHero: MakeCreateHeroEndpoint(s),
		UpdateHero: MakeUpdateEndpoint(s),
		DeleteHero: MakeDeleteEndpoint(s),
		GetHeros:   MakeGetsEndpoint(s),
	}
}
