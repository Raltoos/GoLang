package main

import "fmt"

func ref() {
	i, j := 42, 2701

	// var p int*
	p := &i
	fmt.Println(*p)
	*p = 10
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
	temp()
}

// aggregate data type of composite data values
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} // Y:0 is implicit
	p  = &Vertex{}    // has dtype *Vertex

)

func temp() {
	t1 := Vertex{1, 2}
	p := &t1
	p.Y = 5
	fmt.Println(t1)
	fmt.Println(v1, v2, *p)

	// arrays
	a := [3]string{"rahul", "ramesh", "rakesh"}
	//slices - reference type always backs on an og array
	//slice - length, capacity and pointer
	//dynamic array
	var b []string = a[0:2]
	c := []int{1, 2, 3, 4, 5}

	fmt.Println(b, c)
}
