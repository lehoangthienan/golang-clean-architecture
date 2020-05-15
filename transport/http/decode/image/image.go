package image

import (
	"context"
	"encoding/json"
	"net/http"

	request "github.com/lehoangthienan/marvel-heroes-backend/model/request/image"
)

// CreateRequest func
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req request.Images
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
