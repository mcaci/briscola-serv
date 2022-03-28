package briscola

import "github.com/mcaci/ita-cards/card"

// Points gives the briscola points for a specific input card
func Points(number uint32) uint32 {
	c := *card.MustID(number)
	return map[uint8]uint32{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}[c.Number()]
}
