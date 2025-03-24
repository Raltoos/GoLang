// go ecosystems is made up of packages
// small case and short length typically same as the folder name
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

// returns multiple values
func split(sum int) (int, int) {
	//can only init like this in funcs
	x := sum * 4 / 9
	y := sum - x

	return x, y
}

// variable initialization outside functions
// vars are initialized to the dtype zero values no nulls
var a, b, c bool
var d string = "Anuj"

func main() {
	fmt.Println("Hello World ", rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(10, 12))
	fmt.Println(split(17))
	fmt.Println(a, b, c, d)

	dtype()
	loop()
	ref()
	map_()
}
