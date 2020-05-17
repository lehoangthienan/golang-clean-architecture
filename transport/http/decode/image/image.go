package image

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.Images
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// GetImageFileRequest func
func GetImageFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	filePath := chi.URLParam(r, "filePath")
	return filePath, nil
}
