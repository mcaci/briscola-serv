package briscola_test

import (
	"context"
	"testing"

	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

func TestPointsEP(t *testing.T) {
	ctx := context.Background()
	points, err := briscola.PointsEP(ctx, struct {
		CardNumber uint32 `json:"number"`
	}{CardNumber: 1})
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}

func TestCountEP(t *testing.T) {
	ctx := context.Background()
	count, err := briscola.CountEP(ctx, struct {
		CardNumbers []uint32 `json:"numbers"`
	}{CardNumbers: []uint32{1, 2, 3, 10}})
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(count)
}

func TestCompareEP(t *testing.T) {
	ctx := context.Background()
	points, err := briscola.CompareEP(ctx, struct {
		FirstCardNumber  uint32 `json:"firstCardNumber"`
		FirstCardSeed    uint32 `json:"firstCardSeed"`
		SecondCardNumber uint32 `json:"secondCardNumber"`
		SecondCardSeed   uint32 `json:"secondCardSeed"`
		BriscolaSeed     uint32 `json:"briscolaSeed"`
	}{
		FirstCardNumber:  1,
		FirstCardSeed:    2,
		SecondCardNumber: 3,
		SecondCardSeed:   1,
		BriscolaSeed:     1,
	})
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	t.Log(points)
}
