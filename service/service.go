package service

import (
	"github.com/lehoangthienan/marvel-heroes-backend/service/user"
)

// Service define list of all services in project
type Service struct {
	UserService user.Service
}
