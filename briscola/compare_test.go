package briscola

import (
	"testing"

	"github.com/mcaci/ita-cards/card"
)

func TestCompareWithBriscola(t *testing.T) {
	if !IsOtherWinning(*card.MustID(1), *card.MustID(22), *card.MustID(21)) {
		t.Fatal("Expecting 1 of Coin to lose against 2 of Sword when briscola is Sword")
	}
}

func TestCompareWithNoBriscola(t *testing.T) {
	if IsOtherWinning(*card.MustID(1), *card.MustID(12), *card.MustID(3)) {
		t.Fatal("Expecting 1 of Coin to lose against 2 of Cup (Coin briscola is indifferent here)")
	}
}

func TestCompareSameSeedWin(t *testing.T) {
	if !IsOtherWinning(*card.MustID(25), *card.MustID(26), *card.MustID(13)) {
		t.Fatal("Expecting 5 of Sword to lose against 6 of Sword (Cup briscola is indifferent here)")
	}
}

func TestCompareSameSeedLoss(t *testing.T) {
	if IsOtherWinning(*card.MustID(40), *card.MustID(38), *card.MustID(13)) {
		t.Fatal("Expecting 10 of Cudgel to win against 8 of Cudgel (Cup briscola is indifferent here)")
	}
}
