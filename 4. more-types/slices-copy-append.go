package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 9, 11, 13}
	fmt.Println(cap(primes))

	doubleSliceCap(&primes)
	fmt.Println(cap(primes))

	primes = append(primes, 17)
	fmt.Println(primes)

	morePrimes := []int{23, 27}
	primes = append(primes, morePrimes...)
	fmt.Println(primes)
}

func doubleSliceCap(slice *[]int) {
	tempSlice := make([]int, len(*slice), (cap(*slice) * 2))
	copy(tempSlice, *slice)
	*slice = tempSlice
}
