package briscola

// Points gives the briscola points for a specific input card
func Points(pointGiver interface{ Number() uint8 }) uint8 {
	return map[uint8]uint8{1: 11, 3: 10, 8: 2, 9: 3, 10: 4}[pointGiver.Number()]
}
