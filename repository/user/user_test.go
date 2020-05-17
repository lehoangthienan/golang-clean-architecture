package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/lehoangthienan/marvel-heroes-backend/model/domain"
	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	pgConfig "github.com/lehoangthienan/marvel-heroes-backend/util/config/db/pg"
	"github.com/lehoangthienan/marvel-heroes-backend/util/contextkey"
	tx "github.com/lehoangthienan/marvel-heroes-backend/util/transaction"
)

func Test_userRepo_Create(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		Name:     "AnLe",
		UserName: "anle_create",
		PassWord: "123456",
		Role:     "admin",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr error
	}{
		{
			name: "Create user success",
			args: args{
				ctx:  context.TODO(),
				user: user,
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				db: dbTest,
			}
			got, err := repo.Create(tt.args.ctx, nil, tt.args.user)

			if err != nil && err != tt.wantErr {
				t.Errorf("userRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepo_Update(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	var Name string = "anle"
	var UserName string = "anle_update"
	var PassWord string = "234234"
	var Role string = "admin"

	userUpdate := &requestModel.UpdateUser{
		Name:     &Name,
		UserName: &UserName,
		PassWord: &PassWord,
		Role:     &Role,
	}

	userUpdateWant := &responseModel.UpdateUser{
		&domain.User{
			Name:     Name,
			UserName: UserName,
			PassWord: PassWord,
			Role:     Role,
		},
	}

	userCreate := &domain.User{
		Name:     "AnLe",
		UserName: "anle",
		PassWord: "123456",
		Role:     "admin",
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx        context.Context
		userUpdate *requestModel.UpdateUser
		userCreate *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.UpdateUser
		wantErr error
	}{
		{
			name: "Update user success",
			args: args{
				ctx:        context.TODO(),
				userUpdate: userUpdate,
				userCreate: userCreate,
			},
			want: userUpdateWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				db: dbTest,
			}

			userCreate, err := repo.Create(tt.args.ctx, nil, tt.args.userCreate)

			userUpdateTemp := &requestModel.UpdateUser{
				ParamUserID: userCreate.ID.String(),
				Name:        tt.args.userUpdate.Name,
				UserName:    tt.args.userUpdate.UserName,
				PassWord:    tt.args.userUpdate.PassWord,
				Role:        tt.args.userUpdate.Role,
			}

			ctx := context.WithValue(tt.args.ctx, contextkey.UserIDContextKey, userCreate.ID.String())

			got, err := repo.Update(ctx, nil, userUpdateTemp)

			if err != nil && err != tt.wantErr {
				t.Errorf("userRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.User.Name, tt.want.User.Name) {
				t.Errorf("userRepo.Update() = %v, want %v", got.User.Name, tt.want.User.Name)
			}
			if !reflect.DeepEqual(got.User.UserName, tt.want.User.UserName) {
				t.Errorf("userRepo.Update() = %v, want %v", got.User.UserName, tt.want.User.UserName)
			}
			if !reflect.DeepEqual(got.User.Role, tt.want.User.Role) {
				t.Errorf("userRepo.Update() = %v, want %v", got.User.Role, tt.want.User.Role)
			}
		})
	}
}

func Test_userRepo_SignIn(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := pgConfig.CreateTestDatabase(t)
	defer cleanup()

	var Name string = "anle"
	var UserName string = "anle_signin"
	var PassWord string = "123456"
	var Role string = "admin"

	userSignIn := &requestModel.SignInUser{
		UserName: UserName,
		PassWord: PassWord,
	}

	userCreate := &domain.User{
		Name:     Name,
		UserName: UserName,
		PassWord: PassWord,
		Role:     Role,
	}

	txSvc := tx.NewTransactionService(tx.NewConfig(dbTest))
	txSvc.TXBegin()

	type args struct {
		ctx        context.Context
		userSignIn *requestModel.SignInUser
		userCreate *domain.User
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.UpdateUser
		wantErr error
	}{
		{
			name: "SignIn user success",
			args: args{
				ctx:        context.TODO(),
				userSignIn: userSignIn,
				userCreate: userCreate,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				db: dbTest,
			}

			_, err := repo.Create(tt.args.ctx, nil, tt.args.userCreate)

			if err != nil && err != tt.wantErr {
				t.Errorf("userRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			userSignInTemp := &requestModel.SignInUser{
				UserName: tt.args.userSignIn.UserName,
				PassWord: tt.args.userSignIn.PassWord,
			}

			got, err := repo.LogIn(tt.args.ctx, nil, userSignInTemp)

			if err != nil && err != tt.wantErr {
				t.Errorf("userRepo.LogIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got.Token, "") {
				t.Errorf("userRepo.LogIn() = %v, want %v", got.Token, tt.args.userSignIn.UserName)
			}
		})
	}
}
