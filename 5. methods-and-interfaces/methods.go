package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

// The method bellow should use (v *Vertex) as the receiver
// It avoids the value/struct copy when calling the method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(scale float64) {
	v.X = v.X * scale
	v.Y = v.Y * scale
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(3)
	fmt.Println(v)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
