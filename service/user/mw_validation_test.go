package user

import (
	"context"
	"reflect"
	"testing"

	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
	"github.com/lehoangthienan/marvel-heroes-backend/util/errors"
)

func Test_validationMiddleware_Create(t *testing.T) {
	t.Parallel()
	userSvcMock := &ServiceMock{
		CreateFunc: func(ctx context.Context, req requestModel.CreateUser) (*responseModel.CreateUser, error) {
			return nil, nil
		},
	}

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
			name: "Create user input data valid",
			args: args{
				ctx: context.TODO(),
				req: requestModel.CreateUser{
					Name:     "anletest",
					UserName: "anletest",
					PassWord: "anletest",
					Role:     "admin",
				},
			},
		},
		{
			name: "Create user failed by missing name",
			args: args{
				ctx: context.TODO(),
				req: requestModel.CreateUser{},
			},
			wantErr: errors.MissingNameUserError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validatingMiddleware{
				Service: userSvcMock,
			}
			got, err := mw.Create(tt.args.ctx, tt.args.req)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("validationMiddleware.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validationMiddleware_Update(t *testing.T) {
	t.Parallel()
	userSvcMock := &ServiceMock{
		UpdateFunc: func(ctx context.Context, req requestModel.UpdateUser) (*responseModel.UpdateUser, error) {
			return nil, nil
		},
	}

	var name string = "anletest"
	var userName string = "anletest"
	var passWord string = "anletest"
	var role string = "admin"

	var nameF string = "anle"

	type args struct {
		ctx context.Context
		req requestModel.UpdateUser
	}
	tests := []struct {
		name    string
		args    args
		want    *responseModel.UpdateUser
		wantErr error
	}{
		{
			name: "Update user input data valid",
			args: args{
				ctx: context.TODO(),
				req: requestModel.UpdateUser{
					Name:     &name,
					UserName: &userName,
					PassWord: &passWord,
					Role:     &role,
				},
			},
		},
		{
			name: "Update user failed by len name",
			args: args{
				ctx: context.TODO(),
				req: requestModel.UpdateUser{
					Name: &nameF,
				},
			},
			wantErr: errors.LengthNameError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validatingMiddleware{
				Service: userSvcMock,
			}
			got, err := mw.Update(tt.args.ctx, tt.args.req)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("validationMiddleware.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}
