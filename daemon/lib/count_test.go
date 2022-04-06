package briscola

import (
	"testing"
)

func TestCount(t *testing.T) {
	if Count([]uint32{1, 2, 3}) != 21 {
		t.Fatal("Points string should contain the total of 21")
	}
}
