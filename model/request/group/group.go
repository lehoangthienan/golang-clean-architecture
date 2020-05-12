package group

// CreateGroup Struct
type CreateGroup struct {
	Name string `json:"name,omitempty"`
}

// UpdateGroup Struct
type UpdateGroup struct {
	ParamGroupID string `json:"groupID,omitempty"`
	Name         string `json:"name,omitempty"`
}

// DeleteGroup Struct
type DeleteGroup struct {
	ParamGroupID string `json:"groupID,omitempty"`
}

// GetGroupHeroes Struct
type GetGroupHeroes struct {
	ParamGroupID string `json:"groupID,omitempty"`
	Skip         string `json:"skip,omitempty"`
	Limit        string `json:"limit,omitempty"`
}

// AssignHeroesGroup Struct
type AssignHeroesGroup struct {
	GroupID string   `json:"groupID,omitempty"`
	Heroes  []string `json:"heroes,omitempty"`
}
