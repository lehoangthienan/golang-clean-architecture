package auth

import (
	"context"
)

// Service interface
type Service interface {
	AuthenticateModerator(ctx context.Context) error
	AuthenticateAdmin(ctx context.Context) error
}
