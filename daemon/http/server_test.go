package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpsrv "github.com/mcaci/briscola-serv/daemon/http"
)

func TestPointsViaHTTP(t *testing.T) {
	ts := httptest.NewServer(httpsrv.NewHandler(context.TODO()))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/points", "", strings.NewReader(`{"number":1}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestCountViaHTTP(t *testing.T) {
	ts := httptest.NewServer(httpsrv.NewHandler(context.TODO()))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/count", "", strings.NewReader(`{"numbers":[1, 2, 3]}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestCompareViaHTTP(t *testing.T) {
	ts := httptest.NewServer(httpsrv.NewHandler(context.TODO()))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/compare", "", strings.NewReader(`{"firstCardNumber":1, "firstCardSeed":2, "secondCardNumber":3, "secondCardSeed":1, "briscolaSeed":1}`))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
