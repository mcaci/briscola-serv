package briscola_test

import (
	"context"
	"testing"

	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

func TestPointsService(t *testing.T) {
	srv := briscola.NewService()
	ctx := context.Background()
	points, err := srv.CardPoints(ctx, 1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}
