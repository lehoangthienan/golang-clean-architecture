package group

import (
	"context"
	"reflect"
	"testing"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
	userRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/user"
	pgConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	tx "github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

func Test_groupRepo_Create(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_create_group",
		PassWord: "123456",
		Role:     "admin",
	}

	group := &domain.Group{
		Name: "Avenger",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx   context.Context
		user  *domain.User
		group *domain.Group
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Group
		wantErr error
	}{
		{
			name: "Create group success",
			args: args{
				ctx:   context.TODO(),
				user:  user,
				group: group,
			},
			want: group,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &groupRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			got, err := repo.Create(ctx, nil, tt.args.group)

			if err != nil && err != tt.wantErr {
				t.Errorf("groupRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupRepo_Update(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_update_group",
		PassWord: "123456",
		Role:     "admin",
	}

	group := &domain.Group{
		Name: "Avenger Update",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx   context.Context
		user  *domain.User
		group *domain.Group
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Group
		wantErr error
	}{
		{
			name: "Update group success",
			args: args{
				ctx:   context.TODO(),
				user:  user,
				group: group,
			},
			want: group,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &groupRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			groupCreate, err := repo.Create(ctx, nil, tt.args.group)

			if err != nil && err != tt.wantErr {
				t.Errorf("groupRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			groupUpdate := &requestModel.UpdateGroup{
				ParamGroupID: groupCreate.ID.String(),
				Name:         "Avenger 1",
			}

			got, err := repo.Update(ctx, nil, groupUpdate)

			if err != nil && err != tt.wantErr {
				t.Errorf("groupRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Group.Name, groupUpdate.Name) {
				t.Errorf("groupRepo.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupRepo_Delete(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_update_group",
		PassWord: "123456",
		Role:     "admin",
	}

	group := &domain.Group{
		Name: "Avenger FS",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx   context.Context
		user  *domain.User
		group *domain.Group
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.Group
		wantErr error
	}{
		{
			name: "Delete group success",
			args: args{
				ctx:   context.TODO(),
				user:  user,
				group: group,
			},
			want: group,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &groupRepo{
				db: dbTest,
			}
			userRepo := userRepo.NewRepo(dbTest)

			userCreate, err := userRepo.Create(tt.args.ctx, nil, tt.args.user)

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			groupCreate, err := repo.Create(ctx, nil, tt.args.group)

			if err != nil && err != tt.wantErr {
				t.Errorf("groupRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			groupDelete := &requestModel.DeleteGroup{
				ParamGroupID: groupCreate.ID.String(),
			}

			_, err = repo.Delete(ctx, nil, groupDelete)

			if err != nil && err != tt.wantErr {
				t.Errorf("groupRepo.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
