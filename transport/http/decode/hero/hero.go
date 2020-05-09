package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/hero"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.CreateHero
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest func
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.UpdateHero
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	heroID := chi.URLParam(r, "heroID")
	req.ParamHeroID = heroID

	return req, err
}

// DeleteRequest func
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.DeleteHero
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	heroID := chi.URLParam(r, "heroID")
	req.ParamHeroID = heroID

	return req, err
}

// GetHeroesRequest func
func GetHeroesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.GetHeroes

	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	req.Skip = skip
	req.Limit = limit

	return req, nil
}
