package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var r float64
	var residuo float64

	r = 1
	for i := 1; i <= x; i++ {
		value := float64(i)
		residuo = r * value
		res := 1 / residuo
		r = r +  res
	}

	fmt.Print(r)
}