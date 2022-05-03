package daemon_test

import (
	"context"
	"testing"
	"time"

	"github.com/mcaci/briscola-serv/daemon"
)

func TestStartWithTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	err := daemon.NewServer(&daemon.Opts{HTTPAddr: "localhost:8080", Ctx: ctx})
	if err != nil {
		t.Fatal(err)
	}
}
