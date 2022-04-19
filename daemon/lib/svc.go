package briscola

import (
	"context"
)

type briscolaService struct{}

func NewService() briscolaService {
	return briscolaService{}
}

func (b briscolaService) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	return Points(number), nil
}

func (b briscolaService) PointCount(ctx context.Context, numbers []uint32) (uint32, error) {
	return Count(numbers), nil
}

func (b briscolaService) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	return IsOtherWinning(firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed), nil
}
