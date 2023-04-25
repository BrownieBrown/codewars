package kata

func NbrOfLaps(x, y uint16) [2]uint16 {
	n, gcd := x, y
	for n != 0 {
		n, gcd = gcd%n, n
	}
	return [2]uint16{y / gcd, x / gcd}
}
