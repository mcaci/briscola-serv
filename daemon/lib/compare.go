package briscola

import (
	"log"

	"github.com/mcaci/ita-cards/card"
)

// IsOtherWinning checks if 'other' card wins being played after the base one and
// with the specified input briscola
func IsOtherWinning(firstCardNumber, firstCardSeed, secondCardNumber, secondCardSeed, briscolaSeed uint32) bool {
	base, err := card.FromID(firstCardNumber + firstCardSeed*10)
	if err != nil {
		log.Println(err)
		return false
	}
	other, err := card.FromID(secondCardNumber + secondCardSeed*10)
	if err != nil {
		log.Println(err)
		return false
	}
	briscola, err := card.FromID(1 + briscolaSeed*10)
	if err != nil {
		log.Println(err)
		return false
	}
	switch sameSeed(base, other) {
	case true:
		return isOtherGreater(base, other)
	case false:
		return sameSeed(other, briscola)
	}
	return false
}

func sameSeed(base, other interface{ Seed() card.Seed }) bool { return base.Seed() == other.Seed() }

func isOtherGreater(base, other interface{ Number() uint8 }) bool {
	isOtherGreaterOnPoints := Points(uint32(base.Number())) < Points(uint32(other.Number()))
	isSamePoints := Points(uint32(base.Number())) == Points(uint32(other.Number()))
	isOtherGreaterOnNumber := base.Number() < other.Number()
	return (isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints
}
