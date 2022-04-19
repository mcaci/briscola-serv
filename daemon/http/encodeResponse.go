package srvhttp

import (
	"context"
	"encoding/json"
	"net/http"
)

func responseEncode(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
