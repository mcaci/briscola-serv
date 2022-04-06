package briscola

import (
	"testing"
)

func TestCompareCases(t *testing.T) {
	tcs := []struct {
		name               string
		fN, fS, sN, sS, bS uint32
		expected           bool
		msg                string
	}{
		{"Compare With Briscola", 1, 0, 2, 2, 2, true, "Expecting 1 of Coin to lose against 2 of Sword when briscola is Sword"},
		{"Compare With no Briscola", 1, 0, 2, 1, 0, false, "Expecting 1 of Coin to lose against 2 of Cup (Coin briscola is indifferent here)"},
		{"Compare with same seed: win", 5, 2, 6, 2, 1, true, "EExpecting 5 of Sword to lose against 6 of Sword (Cup briscola is indifferent here)"},
		{"Compare with same seed: loss", 10, 3, 8, 3, 1, false, "Expecting 10 of Cudgel to win against 8 of Cudgel (Cup briscola is indifferent here)"},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsOtherWinning(tc.fN, tc.fS, tc.sN, tc.sS, tc.bS)
			if actual != tc.expected {
				t.Fatal(tc.msg)
			}
		})
	}
}
