package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	aux := (z*z - x) / (2 * z)

	for z != z-aux {
		aux = (z*z - x) / (2 * z)
		z -= aux
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
