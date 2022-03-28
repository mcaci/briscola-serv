package briscola

// Count counts the number of points in a set of cards
func Count(numbers []uint32) uint32 {
	var sum uint32
	for _, n := range numbers {
		sum += Points(n)
	}
	return sum
}
