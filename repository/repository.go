package repository

import (
	"github.com/lehoangthienan/marvel-heroes-backend/repository/user"
)

// Repository struct
type Repository struct {
	UserRepository user.Repository
}
