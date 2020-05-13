package group

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/group"
	repo "github.com/lehoangthienan/marvel-heroes-backend/repository"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type groupService struct {
	repo repo.Repository
	tx   transaction.TXService
}

// NewService func
func NewService(repo repo.Repository, tx transaction.TXService) Service {
	return &groupService{repo: repo, tx: tx}
}

func (s *groupService) Create(ctx context.Context, req requestModel.CreateGroup) (*responseModel.CreateGroup, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	group, err := s.repo.GroupRepository.Create(ctx, pool, &domain.Group{
		Name: req.Name,
	})
	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.CreateGroupFailedError
	}

	return &responseModel.CreateGroup{Group: group}, s.tx.TXCommit(pool)
}

func (s *groupService) Update(ctx context.Context, req requestModel.UpdateGroup) (*responseModel.UpdateGroup, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	group, err := s.repo.GroupRepository.Update(ctx, pool, &req)

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.UpdateGroupFailedError
	}

	return &responseModel.UpdateGroup{Group: group.Group}, s.tx.TXCommit(pool)
}

func (s *groupService) Delete(ctx context.Context, req requestModel.DeleteGroup) (*responseModel.DeleteGroup, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	group, err := s.repo.GroupRepository.Delete(ctx, pool, &req)

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.UpdateGroupFailedError
	}

	return &responseModel.DeleteGroup{Group: group.Group}, s.tx.TXCommit(pool)
}

func (s *groupService) AssignHeroesGroup(ctx context.Context, req *requestModel.AssignHeroesGroup) ([]*domain.GroupHero, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	groupID, err := domain.UUIDFromString(req.GroupID)
	if err != nil {
		return nil, err
	}

	addList := []*domain.GroupHero{}

	for i := 0; i < len(req.Heroes); i++ {
		heroID, err := domain.UUIDFromString(req.Heroes[i])
		if err != nil {
			return nil, err
		}
		addList = append(addList, &domain.GroupHero{GroupID: &groupID, HeroID: &heroID})
	}

	if len(addList) > 0 {
		addList, err = s.repo.GroupRepository.AssignHeroesGroup(ctx, pool, addList)
		if err != nil {
			s.tx.TXRollBack(pool)
			return nil, err
		}
	}

	return addList, s.tx.TXCommit(pool)
}
