package hero

import (
	"context"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/hero"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
	"github.com/lehoangthienan/marvel-heroes-backend/util/helper"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

type heroRepo struct {
	db *gorm.DB
}

// NewRepo func
func NewRepo(db *gorm.DB) Repository {
	return &heroRepo{db: db}
}

func (r *heroRepo) Create(ctx context.Context, pool *transaction.Pool, hero *domain.Hero) (*domain.Hero, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	heroExisted := &domain.Hero{Name: hero.Name}

	err := db.Find(heroExisted, heroExisted).Error
	if err == nil {
		return nil, errors.HeronameIsExistedError
	}

	ctxUserID, check := ctx.Value(contextkey.UserIDContextKey).(string)

	if !check {
		return nil, errors.NotLoggedInError
	}

	creatorID, _ := domain.UUIDFromString(ctxUserID)

	hero.CreatorID = &creatorID

	err = db.Create(&hero).Error

	if err != nil {
		return nil, err
	}

	return hero, err
}

func (r *heroRepo) Update(ctx context.Context, pool *transaction.Pool, req *requestModel.UpdateHero) (*responseModel.UpdateHero, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	heroID, _ := domain.UUIDFromString(req.ParamHeroID)

	hero := &domain.Hero{
		Model: domain.Model{
			ID: heroID,
		},
	}

	err := db.Find(hero, hero).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.HeroNotExistError
		}
		return nil, err
	}

	if req.Name != "" {
		hero.Name = req.Name
	}

	if req.Power != "" {
		hero.Power = req.Power
	}

	err = db.Save(hero).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.UpdateHero{
		Hero: hero,
	}, err
}

func (r *heroRepo) Delete(ctx context.Context, pool *transaction.Pool, req *requestModel.DeleteHero) (*responseModel.DeleteHero, error) {
	db, _ := helper.UseDBConn(r.db, pool)

	heroID, _ := domain.UUIDFromString(req.ParamHeroID)

	hero := &domain.Hero{
		Model: domain.Model{
			ID: heroID,
		},
	}

	err := db.Find(hero, hero).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = errors.HeroNotExistError
		}
		return nil, err
	}

	err = db.Delete(hero).Error

	if err != nil {
		return nil, err
	}

	return &responseModel.DeleteHero{
		Hero: hero,
	}, err
}

func (r *heroRepo) GetHeroes(ctx context.Context, req *requestModel.GetHeroes) (*responseModel.GetHeroes, error) {
	res := &responseModel.GetHeroes{}
	db, _ := helper.UseDBConn(r.db, nil)

	skip := req.Skip
	limit := req.Limit

	if req.Skip == "" {
		skip = "-1"
	}

	if req.Limit == "" {
		limit = "-1"
	}

	hereos := []domain.Hero{}
	var total int

	var err error

	findHeroWG := sync.WaitGroup{}

	findHeroWG.Add(1)
	go func() {
		defer findHeroWG.Done()
		errFunc := db.Limit(limit).Offset(skip).Order("created_at desc").Find(&hereos).Error
		if errFunc != nil {
			err = errFunc
		}
	}()

	findHeroWG.Add(1)
	go func() {
		defer findHeroWG.Done()
		errFunc := db.Model(&domain.Hero{}).Count(&total).Error
		if errFunc != nil {
			err = errFunc
		}
	}()

	findHeroWG.Wait()

	if err != nil {
		return res, err
	}

	heroesRes := make([]*domain.Hero, len(hereos))

	wg := sync.WaitGroup{}
	for idx, hero := range hereos {
		wg.Add(1)
		go func(product domain.Hero, idx int) {
			defer wg.Done()
			heroesRes[idx] = &hero
		}(hero, idx)
	}
	wg.Wait()

	res.Heroes = heroesRes
	res.Total = total

	return res, err
}
