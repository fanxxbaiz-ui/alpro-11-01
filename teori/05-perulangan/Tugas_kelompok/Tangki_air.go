package main

import "fmt"

func main() {
	var T, V, total int
	fmt.Scan(&T)
	for total < T {
		fmt.Scan(&V)
		total += V
		if total >= T {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
