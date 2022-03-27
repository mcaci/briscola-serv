package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/set"
)

func TestCount(t *testing.T) {
	fakeHand := *set.NewMust(1, 2, 3)
	if score1 := Count(fakeHand); score1 != 21 {
		t.Fatal("Points string should contain the total of 21")
	}
}
