package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	slice := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		internalSlice := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			internalSlice[j] = uint8((i + j) / 2)
		}
		slice[i] = internalSlice
	}

	return slice
}

func main() {
	pic.Show(Pic)
}
