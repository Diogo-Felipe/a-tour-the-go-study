package main

import "fmt"

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for index, value := range pow {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	for _, value := range pow {
		fmt.Println(value)
	}

	for index, _ := range pow {
		fmt.Println(index)
	}
}
