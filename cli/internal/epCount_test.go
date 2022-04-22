package internal_test

import (
	"context"
	"testing"

	"github.com/mcaci/briscola-serv/cli/internal"
)

func TestCountCreateOK(t *testing.T) {
	_, err := internal.Count([]string{"1", "2", "10", "3", "5", "8", "6", "4", "4", "2", "5"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCountCreateKO(t *testing.T) {
	testcases := []struct {
		name string
		args []string
	}{
		{name: "not a number", args: []string{"1", "1", "one", "2", "1"}},
		{name: "number not in range", args: []string{"1", "1", "1", "2", "12"}},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := internal.Count(tc.args)
			if err == nil {
				t.FailNow()
			}
		})
	}
}

func TestCountRun(t *testing.T) {
	pcEp, err := internal.Count([]string{"1", "2", "10", "3", "5", "8", "6", "4", "4", "2", "5"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := pcEp.Run(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if v == nil {
		t.Fatalf("expecting a return value for v but none was found")
	}
}
