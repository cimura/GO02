package piscine

func IterativeFactorial(nb int) int {
	iter := 1
	for i := 0; i < nb; {
		iter *= nb
		nb--
	}
	return iter
}