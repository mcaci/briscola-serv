package briscola

import (
	"testing"
)

func FuzzCount(f *testing.F) {
	testcases := [][]uint32{{1, 2, 3}, {5, 7, 9}}
	for _, tc := range testcases {
		f.Add(tc[0], tc[1], tc[2])
	}
	f.Fuzz(func(t *testing.T, a uint32, b uint32, c uint32) {
		count1 := Count([]uint32{a, b, c})
		count2 := Count([]uint32{a, b, c})
		if count1 != count2 {
			t.Fatalf("the count should be the same (%d vs %d) with the same input", count1, count2)
		}
	})

}
