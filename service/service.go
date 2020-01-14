package serv

import (
	"context"

	"github.com/mcaci/ita-cards/card"
	"github.com/mcaci/briscola-serv/briscola"
)

type Service interface {
	CardPoints(ctx context.Context, number uint32) (uint32, error)
	PointCount(ctx context.Context, number []uint32) (uint32, error)
	CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error)
}

type briscolaService struct{}

func (b briscolaService) CardPoints(ctx context.Context, number uint32) (uint32, error) {
	return uint32(briscola.Points(cardnumber(number))), nil
}

func (b briscolaService) PointCount(ctx context.Context, numbers []uint32) (uint32, error) {
	var a []interface{ Number() uint8 }
	for _, n := range numbers {
		a = append(a, cardnumber(n))
	}
	return uint32(briscola.CountWithIntf(a)), nil
}

func (b briscolaService) CardCompare(ctx context.Context, firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) (bool, error) {
	first, second := *card.MustID(uint8(firstCardNumber + firstCardSeed*10)), *card.MustID(uint8(secondCardNumber + secondCardSeed*10))
	br := cardseed(briscolaSeed)
	return briscola.IsOtherWinning(first, second, br), nil
}

func NewService() Service {
	return briscolaService{}
}

type cardnumber uint8

func (n cardnumber) Number() uint8 { return uint8(n) }

type cardseed card.Seed

func (s cardseed) Seed() card.Seed { return card.Seed(s) }
