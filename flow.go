package main

import (
	"fmt"
	"math"
	"time"
)

func pow(x, n, lim float64) float64 {
	// if statements can have local vars
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func loop() {
	defer fmt.Println("I was ready to be printed")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	// while version of go
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)
	fmt.Println(pow(2, 5, 100))

	//switch statements - a litte different
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow ")
	case today + 2:
		fmt.Println("In Two Days ")
	}
}
