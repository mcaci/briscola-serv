package httpmdw_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpmdw "github.com/mcaci/briscola-serv/daemon/http/mdw"
)

func TestLoggingMdw(t *testing.T) {
	logHndl := httpmdw.Logged(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	}))
	ts := httptest.NewServer(logHndl)
	defer ts.Close()

	res, err := http.Post(ts.URL+"/points", "", strings.NewReader(`{"number":1}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
