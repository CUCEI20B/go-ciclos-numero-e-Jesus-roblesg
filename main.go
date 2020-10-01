package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var r float64
	r = 1
	for i := 1; i <= x; i++ {
		r += 1 / (r * float64(i))
	}

	fmt.Print(r)
}