package briscola

import (
	"testing"
)

func FuzzCompare(f *testing.F) {
	testcases := [][]uint32{{1, 0, 2, 2, 2}, {1, 0, 2, 1, 0}, {5, 2, 6, 2, 1}, {1, 0, 2, 1, 0}}
	for _, tc := range testcases {
		f.Add(tc[0], tc[1], tc[2], tc[3], tc[4])
	}
	f.Fuzz(func(t *testing.T, a, b, c, d, e uint32) {
		compare1 := IsOtherWinning(a, b, c, d, e)
		compare2 := IsOtherWinning(a, b, c, d, e)
		if compare1 != compare2 {
			t.Fatalf("the count should be the same (%t vs %t) with the same input", compare1, compare2)
		}
	})

}
