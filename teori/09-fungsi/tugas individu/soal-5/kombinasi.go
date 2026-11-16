package main

import "fmt"

func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x dan y: ")
	fmt.Scan(&x, &y)

	fakX := faktorial(x)
	fakY := faktorial(y)

	var hasilKomb int
	if x >= y {
		hasilKomb = kombinasi(x, y)
	} else {
		hasilKomb = kombinasi(y, x)
	}

	fmt.Printf("%d %d %d\n", fakX, fakY, hasilKomb)
}
