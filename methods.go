package main

import (
	"fmt"
	"math"
)

type Vert struct {
	X, Y float64
}

// methods have a special part known as receiver - value, reference
func (v Vert) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func met() {
	v := Vert{
		X: 3,
		Y: 4,
	}

	fmt.Println(v.Abs())
}
