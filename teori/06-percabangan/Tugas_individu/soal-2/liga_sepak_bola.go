package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	
	fmt.Scan(&a, &b, &c, &d)

	max := a
	min := a

	if b > max {
		max = b
	}
	if b < min {
		min = b
	}

	if c > max {
		max = c
	}
	if c < min {
		min = c
	}

	if d > max {
		max = d
	}
	if d < min {
		min = d
	}

	fmt.Printf("%d %d\n", max, min)
}