package briscola

import (
	"github.com/mcaci/ita-cards/card"
)

// IsOtherWinning checks if 'other' card wins being played after the base one and
// with the specified input briscola
func IsOtherWinning(firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) bool {
	base := *card.MustID(firstCardNumber + firstCardSeed*10)
	other := *card.MustID(secondCardNumber + secondCardSeed*10)
	briscola := *card.MustID(1 + briscolaSeed*10)
	return (!sameSeed(base, other) && sameSeed(other, briscola)) || (sameSeed(base, other) && isOtherGreater(base, other))
}

func sameSeed(base, other interface{ Seed() card.Seed }) bool { return base.Seed() == other.Seed() }

func isOtherGreater(base, other interface{ Number() uint8 }) bool {
	isOtherGreaterOnPoints := Points(uint32(base.Number())) < Points(uint32(other.Number()))
	isSamePoints := Points(uint32(base.Number())) == Points(uint32(other.Number()))
	isOtherGreaterOnNumber := base.Number() < other.Number()
	return (isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints
}
