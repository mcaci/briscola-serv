package briscola

import "github.com/mcaci/ita-cards/card"

func sameSeed(base, other interface{ Seed() card.Seed }) bool { return base.Seed() == other.Seed() }

// IsOtherWinning checks if 'other' card wins being played after the base one and
// with the specified input briscola
func IsOtherWinning(base, other card.Item, briscola interface{ Seed() card.Seed }) bool {
	return (!sameSeed(base, other) && sameSeed(other, briscola)) || (sameSeed(base, other) && isOtherGreater(base, other))
}

func isOtherGreater(base, other card.Item) bool {
	isOtherGreaterOnPoints := Points(base) < Points(other)
	isSamePoints := Points(base) == Points(other)
	isOtherGreaterOnNumber := base.Number() < other.Number()
	return (isSamePoints && isOtherGreaterOnNumber) || isOtherGreaterOnPoints
}
