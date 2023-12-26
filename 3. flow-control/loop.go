package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
