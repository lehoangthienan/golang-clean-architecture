package hero

import "github.com/lehoangthienan/marvel-heroes-backend/model/domain"

// CreateHero struct
type CreateHero struct {
	Hero *domain.Hero `json:"user,omitempty"`
}

// UpdateHero struct
type UpdateHero struct {
	Hero *domain.Hero `json:"user,omitempty"`
}

// DeleteHero struct
type DeleteHero struct {
	Hero *domain.Hero `json:"user,omitempty"`
}
