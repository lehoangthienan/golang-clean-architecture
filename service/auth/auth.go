package auth

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	"github.com/lehoangthienan/marvel-heroes-backend/util/constants"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
)

type authService struct {
	db *gorm.DB
}

// NewAuthService func
func NewAuthService(db *gorm.DB) Service {
	return &authService{
		db: db,
	}
}

func (s *authService) AuthenticateModerator(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return errors.NotLoggedInError
	}

	userID, err := domain.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}

	user := &domain.User{Model: domain.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return errors.AccountNotFoundError
	}

	if user.Role != constants.ADMIN && user.Role != constants.MODERATOR {
		return errors.AccessDeniedError
	}

	return nil
}

func (s *authService) AuthenticateAdmin(ctx context.Context) error {
	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return errors.NotLoggedInError
	}

	userID, err := domain.UUIDFromString(ctxUserID)
	if err != nil {
		return err
	}

	user := &domain.User{Model: domain.Model{
		ID: userID,
	}}

	err = s.db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return errors.AccountNotFoundError
	}

	if user.Role != constants.ADMIN {
		return errors.AccessDeniedError
	}

	return nil
}
