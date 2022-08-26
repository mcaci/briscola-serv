package briscola

// Points gives the briscola points for a specific input card
func Points(number uint32) uint32 {
	return map[uint8]uint32{1: 11, 3: 10, 8: 2, 9: 3, 0: 4}[uint8(number)%10]
}
