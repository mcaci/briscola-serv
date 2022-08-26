package briscola

import (
	"testing"
)

func TestPoints(t *testing.T) {
	in := uint32(1)
	p1 := Points(in)
	p2 := Points(in)
	if p1 != p2 {
		t.Errorf("Expecting %d to equal %d", p1, p2)
	}
}