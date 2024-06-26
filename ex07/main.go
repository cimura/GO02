package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.FindNextPrime(-1)) // 負の場合は2を返す
	fmt.Println(piscine.FindNextPrime(1))
	fmt.Println(piscine.FindNextPrime(2))
	fmt.Println(piscine.FindNextPrime(3))
	fmt.Println(piscine.FindNextPrime(4))
	fmt.Println(piscine.FindNextPrime(5))
	fmt.Println(piscine.FindNextPrime(40))
}
