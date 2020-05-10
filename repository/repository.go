package repository

import (
	"github.com/lehoangthienan/marvel-heroes-backend/repository/group"
	"github.com/lehoangthienan/marvel-heroes-backend/repository/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/repository/user"
)

// Repository struct
type Repository struct {
	UserRepository  user.Repository
	HeroRepository  hero.Repository
	GroupRepository group.Repository
}
