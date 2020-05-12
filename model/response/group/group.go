package group

import "github.com/lehoangthienan/marvel-heroes-backend/model/domain"

// Group struct
type Group struct {
	*domain.Group
}

// CreateGroup struct
type CreateGroup struct {
	Group *domain.Group `json:"group,omitempty"`
}

// UpdateGroup struct
type UpdateGroup struct {
	Group *domain.Group `json:"group,omitempty"`
}

// DeleteGroup struct
type DeleteGroup struct {
	Group *domain.Group `json:"group,omitempty"`
}

// GetHeroes struct
type GetHeroes struct {
	Heroes []*domain.Hero `json:"heroes"`
	Total  int            `json:"total,omitempty"`
}

// AssignHeroesGroup Struct
type AssignHeroesGroup struct {
	Group  *domain.Group  `json:"group,omitempty"`
	Heroes []*domain.Hero `json:"heroes,omitempty"`
}
