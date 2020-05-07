package clinic

import (
	"context"

	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
	userProto "github.com/lehoangthienan/marvel-heroes-backend/transport/grpc/proto/user"
)

// SignIn func
func SignIn(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*userProto.SignInUserRequest)
	return request.SignInUser{
		UserName: req.Username,
		PassWord: req.Password,
	}, nil
}
