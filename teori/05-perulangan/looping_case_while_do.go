package main

import "fmt"

func main() {
	var i, n int
	fmt.Scan(&n)
	i = 1
	for i <= n {
		fmt.Print(i, " ")
		i++
	}
}
