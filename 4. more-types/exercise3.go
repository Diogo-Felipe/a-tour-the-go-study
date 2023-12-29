package main

import "fmt"

func fibonacci() func() int {
	previous := 1
	antiPrevious := 0

	return func() int {
		actual := antiPrevious
		antiPrevious = previous
		previous = actual + antiPrevious
		return actual
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
