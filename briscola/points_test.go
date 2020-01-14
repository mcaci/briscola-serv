package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestPoints(t *testing.T) {
	if Points(*card.MustID(1)) != Points(*card.MustID(21)) {
		t.Fatal("unexpected result")
	}
}
