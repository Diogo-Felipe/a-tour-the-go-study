package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	primes[2] = 17
	fmt.Println(s)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	q = q[:2]
	fmt.Println(q)

	q = q[1:]
	fmt.Println(q)

}
