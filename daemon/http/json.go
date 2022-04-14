package briscolahttp

import (
	"context"
	"encoding/json"
	"net/http"
)

func RequestDecodeJSON[T any](ctx context.Context, r *http.Request) (interface{}, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func ResponseEncodeJSON(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
