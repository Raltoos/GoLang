package main

import (
	"fmt"
	"math/cmplx"
)

// int, string, bool
// aggrergate dtypes : arrays, struct
// reference dtypes : pointers, slices, functions, channel, maps
// interfaces
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func dtype() {
	fmt.Printf("Type of %T Value : %v\n", ToBe, ToBe)
	fmt.Printf("Type of %T Value : %v\n", MaxInt, MaxInt)
	fmt.Printf("Type of %T Value : %v\n", z, z)
}
