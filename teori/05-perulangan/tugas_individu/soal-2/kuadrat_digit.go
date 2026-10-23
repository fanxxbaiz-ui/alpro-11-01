package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	sum := 0
	first := true

	for x > 0 {
		digit := x % 10
		sq := digit * digit
		sum += sq

		if !first {
			fmt.Print(" ")
		}
		fmt.Print(sq)
		first = false

		x /= 10
	}
	fmt.Println()
	fmt.Println(sum)
}
