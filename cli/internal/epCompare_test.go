package internal_test

import (
	"context"
	"testing"

	"github.com/mcaci/briscola-serv/cli/internal"
)

func TestCompareCreateOK(t *testing.T) {
	_, err := internal.Compare([]string{"1", "1", "1", "2", "1"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompareCreateKO(t *testing.T) {
	testcases := []struct {
		name string
		args []string
	}{
		{name: "insufficient args", args: []string{"1", "1", "1", "2"}},
		{name: "not a number", args: []string{"1", "1", "one", "2", "1"}},
		{name: "not in range 0", args: []string{"11", "1", "1", "2", "1"}},
		{name: "not in range 1", args: []string{"1", "6", "1", "2", "1"}},
		{name: "not in range 2", args: []string{"1", "1", "11", "2", "1"}},
		{name: "not in range 3", args: []string{"1", "1", "1", "22", "1"}},
		{name: "not in range 4", args: []string{"1", "1", "1", "2", "12"}},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := internal.Compare(tc.args)
			if err == nil {
				t.FailNow()
			}
		})
	}
}

func TestCompareRun(t *testing.T) {
	ccEp, err := internal.Compare([]string{"1", "1", "1", "2", "1"})
	if err != nil {
		t.Fatal(err)
	}
	v, err := ccEp.Run(context.TODO(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if v == nil {
		t.Fatalf("expecting a return value for v but none was found")
	}
}
