package briscola_test

import (
	"context"
	"testing"

	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

func TestPointsService(t *testing.T) {
	ctx := context.Background()
	points, err := briscola.PointsEP(ctx, struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: 1})
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}
