package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var r float64
	r = 1
	value := 1
	for i := 1; i <= x; i++ {
		value *= i
		r += 1 / float64(value)
	}

	fmt.Print(r)
}