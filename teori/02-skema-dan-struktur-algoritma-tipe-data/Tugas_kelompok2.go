package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x, &y)

	fxy := (1.0 / (3.0*x*x + 10.0)) + (10.0*y) + 7.0

	fmt.Println(fxy)
}