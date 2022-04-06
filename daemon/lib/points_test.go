package briscola

import (
	"testing"
)

func TestPoints(t *testing.T) {
	if Points(1) != Points(1) {
		t.Fatal("unexpected result")
	}
}
