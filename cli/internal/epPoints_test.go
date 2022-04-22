package internal_test

import (
	"context"
	"testing"

	"github.com/mcaci/briscola-serv/cli/internal"
)

func TestPointsCreateOK(t *testing.T) {
	_, err := internal.Points([]string{"1"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestPointsCreateKO(t *testing.T) {
	testcases := []struct {
		name string
		args []string
	}{
		{name: "too many args", args: []string{"1", "1", "1", "2"}},
		{name: "not a number", args: []string{"one"}},
		{name: "number not in range", args: []string{"11"}},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := internal.Points(tc.args)
			if err == nil {
				t.FailNow()
			}
		})
	}
}

func TestPointsRun(t *testing.T) {
	cpEp, err := internal.Points([]string{"1"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := cpEp.Run(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if v == nil {
		t.Fatalf("expecting a return value for v but none was found")
	}
}
