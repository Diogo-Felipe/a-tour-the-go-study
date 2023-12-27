package main

import "fmt"

func main() {
	defer fmt.Println("world")

	countingStack()

	fmt.Println("Hello, ")
}

func countingStack() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
