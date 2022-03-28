package briscola

import (
	"testing"
)

func TestCompareWithBriscola(t *testing.T) {
	if !IsOtherWinning(1, 0, 2, 2, 2) {
		t.Fatal("Expecting 1 of Coin to lose against 2 of Sword when briscola is Sword")
	}
}

func TestCompareWithNoBriscola(t *testing.T) {
	if IsOtherWinning(1, 0, 2, 1, 0) {
		t.Fatal("Expecting 1 of Coin to lose against 2 of Cup (Coin briscola is indifferent here)")
	}
}

func TestCompareSameSeedWin(t *testing.T) {
	if !IsOtherWinning(5, 2, 6, 2, 1) {
		t.Fatal("Expecting 5 of Sword to lose against 6 of Sword (Cup briscola is indifferent here)")
	}
}

func TestCompareSameSeedLoss(t *testing.T) {
	if IsOtherWinning(10, 3, 8, 3, 1) {
		t.Fatal("Expecting 10 of Cudgel to win against 8 of Cudgel (Cup briscola is indifferent here)")
	}
}
