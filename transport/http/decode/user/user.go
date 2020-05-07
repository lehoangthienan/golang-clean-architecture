package user

import (
	"context"
	"encoding/json"
	"net/http"

	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/user"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.CreateUser
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// SignInRequest func
func SignInRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.SignInUser
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest func
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.UpdateUser
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, err
}
