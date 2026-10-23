package main

import "fmt"

func main() {
	var xi, total int
	total = 0
	fmt.Scan(&xi)
	for xi%2 == 0 {
		total = total + xi
		fmt.Scan(&xi)
	}
	fmt.Println(total)
}

// Task: 10 10 20 5 -> 40

// Case 1
// Line 8:Input 5 -> Total = 0

// Case 2
// Line 8: Input 10 -> enter to looping condition
// Line 9 - 12
// Line 10 Total = total + xi -> 0 + 10 + 10
// Line 11 Input xi berikutnya
