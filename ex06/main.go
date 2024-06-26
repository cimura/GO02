package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(5)) // true
	fmt.Println(piscine.IsPrime(4)) // false
	fmt.Println(piscine.IsPrime(2)) // true
	fmt.Println(piscine.IsPrime(3)) // true
	fmt.Println(piscine.IsPrime(1)) // false
	fmt.Println(piscine.IsPrime(21)) // false
	fmt.Println(piscine.IsPrime(2147483647)) // true
}
