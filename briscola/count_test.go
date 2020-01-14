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

func TestCountWithIntf(t *testing.T) {
	fakeHand := *set.NewMust(1, 2, 3)
	var fakeHandInterface []interface{ Number() uint8 }
	for _, c := range fakeHand {
		fakeHandInterface = append(fakeHandInterface, c)
	}
	if score1 := CountWithIntf(fakeHandInterface); score1 != 21 {
		t.Fatal("Points string should contain the total of 21")
	}
}
