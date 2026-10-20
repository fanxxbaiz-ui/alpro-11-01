package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var jumlah int = 0
	var i int = 1

	// Perulangan while menggunakan for di Go
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}
