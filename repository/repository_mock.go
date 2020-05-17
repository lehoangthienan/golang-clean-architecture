package repository

import (
	userRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/user"
)

// RepositoryMock struct
type RepositoryMock struct {
	User *userRepo.RepositoryMock
}

// NewRepositoryMock func will return repository mock
func NewRepositoryMock(repoMock RepositoryMock) Repository {
	return Repository{
		UserRepository: repoMock.User,
	}
}
