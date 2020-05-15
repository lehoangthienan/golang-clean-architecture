package service

import (
	"github.com/lehoangthienan/marvel-heroes-backend/service/auth"
	"github.com/lehoangthienan/marvel-heroes-backend/service/group"
	"github.com/lehoangthienan/marvel-heroes-backend/service/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/service/image"
	"github.com/lehoangthienan/marvel-heroes-backend/service/user"
)

// Service define list of all services in project
type Service struct {
	UserService  user.Service
	AuthService  auth.Service
	HeroService  hero.Service
	GroupService group.Service
	ImageService image.Service
}
