package daemon

import (
	"context"

	briscola "github.com/mcaci/briscola-serv/daemon/lib"
)

type briscolaService struct{}

func NewService() briscolaService {
	return briscolaService{}
}

func (b briscolaService) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	return briscola.Points(number), nil
}

func (b briscolaService) PointCount(ctx context.Context, numbers []uint32) (uint32, error) {
	return briscola.Count(numbers), nil
}

func (b briscolaService) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	return briscola.IsOtherWinning(firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed), nil
}
