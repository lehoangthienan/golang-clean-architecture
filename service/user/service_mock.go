package user

import (
	"context"

	requestModel "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	responseModel "github.com/lehoangthienan/marvel-heroes-backend/model/response/user"
)

// Make sure ServiceMock implement Service interface
var _ Service = &ServiceMock{}

// ServiceMock struct used to mock test
type ServiceMock struct {
	CreateFunc func(ctx context.Context, req requestModel.CreateUser) (*responseModel.CreateUser, error)
	LogInFunc  func(ctx context.Context, req requestModel.SignInUser) (*responseModel.SignInUser, error)
	UpdateFunc func(ctx context.Context, req requestModel.UpdateUser) (*responseModel.UpdateUser, error)
}

// Create mock function
func (sm *ServiceMock) Create(ctx context.Context, req requestModel.CreateUser) (*responseModel.CreateUser, error) {
	if sm.CreateFunc == nil {
		panic("ServiceMock not declare Create function")
	}
	return sm.CreateFunc(ctx, req)
}

// LogIn mock function
func (sm *ServiceMock) LogIn(ctx context.Context, req requestModel.SignInUser) (*responseModel.SignInUser, error) {
	if sm.LogInFunc == nil {
		panic("ServiceMock not declare Create function")
	}
	return sm.LogInFunc(ctx, req)
}

// Update mock function
func (sm *ServiceMock) Update(ctx context.Context, req requestModel.UpdateUser) (*responseModel.UpdateUser, error) {
	if sm.UpdateFunc == nil {
		panic("ServiceMock not declare Create function")
	}
	return sm.UpdateFunc(ctx, req)
}
