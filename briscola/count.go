package briscola

// Count counts the number of points in a set of cards
func Count[T interface{ Number() uint8 }](cards []T) (sum uint8) {
	for _, c := range cards {
		sum += Points(c)
	}
	return
}
