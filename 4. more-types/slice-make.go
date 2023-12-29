package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Println(a)

	b := make([]int, 4, 6)
	fmt.Println(b)
	fmt.Println(cap(b))
}
