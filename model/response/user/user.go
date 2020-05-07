package user

import "github.com/lehoangthienan/marvel-heroes-backend/model/domain"

// CreateUser struct
type CreateUser struct {
	User *domain.User `json:"user,omitempty"`
}

// SignInUser struct
type SignInUser struct {
	User  *domain.User `json:"user,omitempty"`
	Token string       `json:"token,omitempty"`
}

// UpdateUser struct
type UpdateUser struct {
	User *domain.User `json:"user,omitempty"`
}
