package briscola

import (
	"github.com/mcaci/ita-cards/set"
)

// Count counts the number of points in a set of cards
func Count(cards set.Cards) (sum uint8) {
	for _, c := range cards {
		sum += Points(c)
	}
	return
}

// CountWithIntf counts the number of points in a set of cards
func CountWithIntf(pointGivers []interface{ Number() uint8 }) (sum uint8) {
	for _, c := range pointGivers {
		sum += Points(c)
	}
	return
}
