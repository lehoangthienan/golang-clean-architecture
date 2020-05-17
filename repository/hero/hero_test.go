package hero

import (
	"context"
	"reflect"
	"testing"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
	userRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/user"
	pgConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	tx "github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

func Test_heroRepo_Create(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_create_hero",
		PassWord: "123456",
		Role:     "admin",
	}

	hero := &domain.Hero{
		Name:  "Captian",
		Power: "shield",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx  context.Context
		user *domain.User
		hero *domain.Hero
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Hero
		wantErr error
	}{
		{
			name: "Create hero success",
			args: args{
				ctx:  context.TODO(),
				user: user,
				hero: hero,
			},
			want: hero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &heroRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			got, err := repo.Create(ctx, nil, tt.args.hero)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heroRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heroRepo_Update(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_update_hero",
		PassWord: "123456",
		Role:     "admin",
	}

	hero := &domain.Hero{
		Name:  "Captian",
		Power: "shield",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx  context.Context
		user *domain.User
		hero *domain.Hero
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Hero
		wantErr error
	}{
		{
			name: "Update hero success",
			args: args{
				ctx:  context.TODO(),
				user: user,
				hero: hero,
			},
			want: hero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &heroRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			heroCreate, err := repo.Create(ctx, nil, tt.args.hero)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			heroUpdate := &requestModel.UpdateHero{
				ParamHeroID: heroCreate.ID.String(),
				Name:        "Ironman",
				Power:       "SuperPower",
			}

			got, err := repo.Update(ctx, nil, heroUpdate)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Hero.Name, heroUpdate.Name) {
				t.Errorf("heroRepo.Update() = %v, want %v", got, tt.want)
			}

			if !reflect.DeepEqual(got.Hero.Power, heroUpdate.Power) {
				t.Errorf("heroRepo.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heroRepo_Delete(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_update_hero",
		PassWord: "123456",
		Role:     "admin",
	}

	hero := &domain.Hero{
		Name:  "Captian FS",
		Power: "shield",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx  context.Context
		user *domain.User
		hero *domain.Hero
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Hero
		wantErr error
	}{
		{
			name: "Delete hero success",
			args: args{
				ctx:  context.TODO(),
				user: user,
				hero: hero,
			},
			want: hero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &heroRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			heroCreate, err := repo.Create(ctx, nil, tt.args.hero)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			heroDelete := &requestModel.DeleteHero{
				ParamHeroID: heroCreate.ID.String(),
			}

			_, err = repo.Delete(ctx, nil, heroDelete)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_heroRepo_Get(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_get_hero",
		PassWord: "123456",
		Role:     "admin",
	}

	hero := &domain.Hero{
		Name:  "Captian FS A",
		Power: "shield",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx  context.Context
		user *domain.User
		hero *domain.Hero
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Hero
		wantErr error
	}{
		{
			name: "Get hero success",
			args: args{
				ctx:  context.TODO(),
				user: user,
				hero: hero,
			},
			want: hero,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &heroRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			_, err = repo.Create(ctx, nil, tt.args.hero)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			getPararms := &requestModel.GetHeroes{
				Skip:  "0",
				Limit: "10",
			}

			heroes, err := repo.GetHeroes(ctx, getPararms)

			if err != nil && err != tt.wantErr {
				t.Errorf("heroRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(heroes.Heroes) > 10 {
				t.Errorf("heroRepo.Get() len(heroes.Heroes) = %v, want %v", len(heroes.Heroes), 10)
			}
		})
	}
}
