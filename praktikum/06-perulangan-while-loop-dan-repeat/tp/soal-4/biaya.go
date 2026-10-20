package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var biaya float64 = 10000000
	var i int = 1

	// Perulangan while menggunakan for di Go
	for i < n {
		biaya += biaya * 0.06
		i++
	}

	fmt.Printf("%.0f\n", biaya)
}