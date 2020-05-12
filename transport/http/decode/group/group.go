package group

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/group"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.CreateGroup
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest func
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.UpdateGroup
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	groupID := chi.URLParam(r, "groupID")
	req.ParamGroupID = groupID

	return req, err
}

// DeleteRequest func
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.DeleteGroup
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	groupID := chi.URLParam(r, "groupID")
	req.ParamGroupID = groupID

	return req, err
}

// AssignHeroesToGroupRequest func
func AssignHeroesToGroupRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.AssignHeroesGroup
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, err
}
