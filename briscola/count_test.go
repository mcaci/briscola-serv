package briscola

import (
	"testing"
)

func TestCount(t *testing.T) {
	if score1 := Count([]uint32{1, 2, 3}); score1 != 21 {
		t.Fatal("Points string should contain the total of 21")
	}
}
