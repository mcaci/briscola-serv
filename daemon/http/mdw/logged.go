package httpmdw

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Logged(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
		}
		log.Printf("handling request at URL %q, method %q and body %v", r.URL, r.Method, string(b))
		fmt.Fprintf(w, "request accepted with URL %q, method %q and body %v\n", r.URL, r.Method, string(b))

		rClone := r.Clone(r.Context())
		rClone.Body = io.NopCloser(bytes.NewReader(b))
		handler.ServeHTTP(w, rClone)
	})
}
