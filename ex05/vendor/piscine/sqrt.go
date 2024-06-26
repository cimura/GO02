package piscine

func Sqrt(nb int) int {
	sqrt := 0
	for i := 2; i <= nb; i++ {
		if nb % (i * i) == 0 {
			sqrt = i
		}
	}
	return sqrt
}