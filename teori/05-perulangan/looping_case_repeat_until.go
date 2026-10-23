package main

import "fmt"

func main() {
	var i, n int
	fmt.Scan(&n)
	i = 1
	for {
		fmt.Print(i, " ")
		1++
		if i > n {
			break
		}
	}
}
