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

func main() {
	fmt.Println("Hello World ", rand.Intn(10))
	fmt.Println(math.Pi)
}
