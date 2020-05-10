package group

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/helper"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type groupRepo struct {
	db *gorm.DB
}

// NewRepo func
func NewRepo(db *gorm.DB) Repository {
	return &groupRepo{db: db}
}

func (r *groupRepo) Create(ctx context.Context, pool *transaction.Pool, group *domain.Group) (*domain.Group, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	groupExisted := &domain.Group{Name: group.Name}

	err := db.Find(groupExisted, groupExisted).Error
	if err == nil {
		return nil, errors.GroupnameIsExistedError
	}

	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)

	if !check {
		return nil, errors.NotLoggedInError
	}

	creatorID, _ := domain.UUIDFromString(ctxUserID)

	group.CreatorID = &creatorID

	err = db.Create(&group).Error

	if err != nil {
		return nil, err
	}

	return group, err
}

func (r *groupRepo) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateGroup) (*responseModel.UpdateGroup, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	groupID, _ := domain.UUIDFromString(req.ParamGroupID)

	group := &domain.Group{
		Model: domain.Model{
			ID: groupID,
		},
	}

	err := db.Find(group, group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.GroupNotExistError
		}
		return nil, err
	}

	if req.Name != "" {
		group.Name = req.Name
	}

	err = db.Save(group).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.UpdateGroup{
		Group: group,
	}, err
}

func (r *groupRepo) Delete(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteGroup) (*responseModel.DeleteGroup, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	groupID, _ := domain.UUIDFromString(req.ParamGroupID)

	group := &domain.Group{
		Model: domain.Model{
			ID: groupID,
		},
	}

	err := db.Find(group, group).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.GroupNotExistError
		}
		return nil, err
	}

	err = db.Delete(group).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.DeleteGroup{
		Group: group,
	}, err
}
