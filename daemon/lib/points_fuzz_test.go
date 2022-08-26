package briscola

import (
	"testing"
)

func FuzzPoints(f *testing.F) {
	testcases := []uint32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, in uint32) {
		p1 := Points(in)
		p2 := Points(in)
		if p1 != p2 {
			t.Errorf("Expecting %d to equal %d", p1, p2)
		}
	})
}
