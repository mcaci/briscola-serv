package serv

import (
	"context"

	"github.com/mcaci/briscola-serv/briscola"
)

type Service interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
	PointCount(ctx context.Context, number []uint32) (uint32, error)
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}

type briscolaService struct{}

func NewService() Service {
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
