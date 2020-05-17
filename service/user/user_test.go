package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	repo "github.com/lehoangthienan/marvel-heroes-backend/repository"
	userRepo "github.com/lehoangthienan/marvel-heroes-backend/repository/user"
	pgConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
	"github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
	tx "github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

func Test_userService_Create(t *testing.T) {
	t.Parallel()

	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))

	userRepoMock := &userRepo.RepositoryMock{
		CreateFunc: func(ctx context.Context, pool *transaction.Pool, user *domain.User) (*domain.User, error) {
			return user, nil
		},
	}

	repoMock := repo.NewRepositoryMock(repo.RepositoryMock{
		User: userRepoMock,
	})

	type args struct {
		ctx context.Context
		req requestModel.CreateUser
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.CreateUser
		wantErr error
	}{
		{
			name: "Create user success",
			args: args{
				ctx: context.TODO(),
				req: requestModel.CreateUser{
					Name:     "AnLe",
					UserName: "anle_create",
					PassWord: "123456",
					Role:     "admin",
				},
			},
			want: &responseModel.CreateUser{
				User: &domain.User{
					Name:     "AnLe",
					UserName: "anle_create",
					PassWord: "123456",
					Role:     "admin",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSvc := &userService{
				repo: repoMock,
				tx:   txSvc,
			}

			got, err := userSvc.Create(tt.args.ctx, tt.args.req)

			if err != nil && err != tt.wantErr {
				t.Errorf("userService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
