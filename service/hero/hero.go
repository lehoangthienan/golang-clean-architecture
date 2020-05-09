package hero

import (
	"context"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
	repo "github.com/lehoangthienan/marvel-heroes-backend/repository"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type heroService struct {
	repo repo.Repository
	tx   transaction.TXService
}

// NewService func
func NewService(repo repo.Repository, tx transaction.TXService) Service {
	return &heroService{repo: repo, tx: tx}
}

func (s *heroService) Create(ctx context.Context, req requestModel.CreateHero) (*responseModel.CreateHero, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	hero, err := s.repo.HeroRepository.Create(ctx, pool, &domain.Hero{
		Name:  req.Name,
		Power: req.Power,
	})
	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.CreateHeroFailedError
	}

	return &responseModel.CreateHero{Hero: hero}, s.tx.TXCommit(pool)
}

func (s *heroService) Update(ctx context.Context, req requestModel.UpdateHero) (*responseModel.UpdateHero, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	hero, err := s.repo.HeroRepository.Update(ctx, pool, &req)

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.UpdateHeroFailedError
	}

	return &responseModel.UpdateHero{Hero: hero.Hero}, s.tx.TXCommit(pool)
}

func (s *heroService) Delete(ctx context.Context, req requestModel.DeleteHero) (*responseModel.DeleteHero, error) {
	pool, err := s.tx.TXBegin()
	if err != nil {
		return nil, err
	}

	hero, err := s.repo.HeroRepository.Delete(ctx, pool, &req)

	if err != nil {
		s.tx.TXRollBack(pool)
		return nil, errors.UpdateHeroFailedError
	}

	return &responseModel.DeleteHero{Hero: hero.Hero}, s.tx.TXCommit(pool)
}

func (s *heroService) GetHeroes(ctx context.Context, req requestModel.GetHeroes) (*responseModel.GetHeroes, error) {
	heroes, err := s.repo.HeroRepository.GetHeroes(ctx, &req)

	if err != nil {
		return nil, errors.UpdateHeroFailedError
	}

	return heroes, nil
}
