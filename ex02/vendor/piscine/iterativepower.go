package	piscine

func IterativePower(nb int, power int) int {
	iter := 1
	for i := 0; i < power; {
		iter *= nb
	}
	return iter
}