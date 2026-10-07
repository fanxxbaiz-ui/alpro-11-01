package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	// Blok 1
	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	} // result = 15

	// Blok 2
	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	} // result = 30

	// Blok 3
	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	} // result = 35

	// Blok 4
	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	} // result = 25

	fmt.Println("Nilai akhir result:", result)
}