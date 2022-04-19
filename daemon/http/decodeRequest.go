package srvhttp

import (
	"context"
	"encoding/json"
	"net/http"
)

func RequestDecode[T any](ctx context.Context, r *http.Request) (interface{}, error) {
	var req T
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
