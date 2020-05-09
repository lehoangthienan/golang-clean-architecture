package hero

import "github.com/lehoangthienan/marvel-heroes-backend/model/domain"

// Hero struct
type Hero struct {
	*domain.Hero
}

// CreateHero struct
type CreateHero struct {
	Hero *domain.Hero `json:"hero,omitempty"`
}

// UpdateHero struct
type UpdateHero struct {
	Hero *domain.Hero `json:"hero,omitempty"`
}

// DeleteHero struct
type DeleteHero struct {
	Hero *domain.Hero `json:"hero,omitempty"`
}

// GetHeroes struct
type GetHeroes struct {
	Heroes []*domain.Hero `json:"heroes"`
	Total  int            `json:"total,omitempty"`
}
