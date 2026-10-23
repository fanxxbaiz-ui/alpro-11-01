package main

import "fmt"

func main() {
	var n, i, bilangan, d1, jumlah int
	fmt.Scan(&n)
	for i = 1; i <= N; I++ {
		fmt.Scan(&bilangan)
		d1 = bilangan / 1000
		// d4 = bilangan % 10
		jumlah += d1
	}
	fmt.Println(jumlah)
}
