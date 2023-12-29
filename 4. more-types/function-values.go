package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	multiply := func(x, y float64) float64 {
		return x * y
	}

	fmt.Println(multiply(2, 3))

	fmt.Println(compute(multiply))
	fmt.Println(compute(math.Pow))
}
