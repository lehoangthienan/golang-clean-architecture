package hero

// CreateHero Struct
type CreateHero struct {
	Name  string `json:"name,omitempty"`
	Power string `json:"power,omitempty"`
}

// UpdateHero Struct
type UpdateHero struct {
	ParamHeroID string `json:"heroID,omitempty"`
	Name        string `json:"name,omitempty"`
	Power       string `json:"power,omitempty"`
}

// DeleteHero Struct
type DeleteHero struct {
	ParamHeroID string `json:"heroID,omitempty"`
}

// GetHeroes Struct
type GetHeroes struct {
	Skip  string `json:"skip,omitempty"`
	Limit string `json:"limit,omitempty"`
}
