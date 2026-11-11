package main

import "fmt"

func main() {
	var huruf string
	var M, K, T int

	for {
		fmt.Scan(&huruf)

		if huruf == "M" {
			M++
		} else if huruf == "K" {
			K++
		} else if huruf == "T" {
			T++
		} else {
			break
		}
	}
	fmt.Println("M:", M)
	fmt.Println("K:", K)
	fmt.Println("T:", T)
}
