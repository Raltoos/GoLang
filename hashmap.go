package main

import "fmt"

// var m map[string]int

//declaring map syntax -> use make() or literal syntax with values beforehand
func map_() {
	// m = make(map[string]int)
	// m := map[string]int{
	// 	"Anuj":  18,
	// 	"Sahas": 1,
	// }
	m := map[string]int{}

	m["Anuj"] = 18
	// delete(m, "Anuj")

	v, ok := m["Anuj"] //ok is a bool that confirms the existence
	fmt.Println(v, ok)

	for key, value := range m {
		fmt.Println(key, value)
	}
}
