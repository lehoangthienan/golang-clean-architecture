package user

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	config "github.com/lehoangthienan/marvel-heroes-backend/util/config/env"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/helper"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, pool *transaction.Pool, user *domain.User) (*domain.User, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	userExisted := &domain.User{UserName: user.UserName}

	err := db.Find(userExisted, userExisted).Error
	if err == nil {
		return nil, errors.UsernameIsExistedError
	}

	err = db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *userRepo) LogIn(ctx context.Context, pool *transaction.Pool, req *requestModel.SignInUser) (*responseModel.SignInUser, error) {
	username := req.UserName
	password := req.PassWord

	db, _ := helper.UseDBConn(r.db, pool)

	user := &domain.User{UserName: username}
	err := db.Find(user, user).Error
	if err != nil && gorm.IsRecordNotFoundError(err) {
		return nil, errors.UserNotFoundError
	}
	checkPassword := user.ComparePassword(password)

	if checkPassword == false {
		return nil, errors.WrongPasswordError
	}
	clasms := helper.TokenClaims{
		UserID:   user.Model.ID.String(),
		Username: user.UserName,
		Role:     user.Role,
	}

	jwt, err := helper.GenerateToken(config.GetJWTSerectKeyEnv(), clasms)

	if err != nil {
		return nil, err
	}

	return &responseModel.SignInUser{
		User:  user,
		Token: jwt,
	}, err
}

func (r *userRepo) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateUser) (*responseModel.UpdateUser, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)
	if !check {
		return nil, errors.NotLoggedInError
	}

	userID, _ := domain.UUIDFromString(ctxUserID)

	user := &domain.User{
		Model: domain.Model{
			ID: userID,
		},
	}

	err := db.Find(user, user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.UserNotExistError
		}
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.UserName != nil {
		user.UserName = *req.UserName
	}

	if req.Role != nil {
		user.Role = *req.Role
	}

	if req.PassWord != nil {
		user.PassWord = *req.PassWord
	}

	err = db.Save(user).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.UpdateUser{
		User: user,
	}, err
}
