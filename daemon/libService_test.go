package daemon_test

import (
	"context"
	"testing"

	"github.com/mcaci/briscola-serv/daemon"
)

func TestPointsService(t *testing.T) {
	srv := daemon.NewService()
	ctx := context.Background()
	points, err := srv.CardPoints(ctx, 1)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}
