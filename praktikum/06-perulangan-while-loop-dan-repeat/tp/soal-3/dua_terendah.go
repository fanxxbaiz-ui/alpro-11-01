package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama, min1Name, min2Name string
	var nilai, min1Val, min2Val int

	// Inisialisasi nilai minimum awal dengan angka yang sangat besar
	min1Val = 999999
	min2Val = 999999

	var i int = 1
	// Perulangan while menggunakan for di Go
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if nilai < min1Val {
			min2Val = min1Val
			min2Name = min1Name
			min1Val = nilai
			min1Name = nama
		} else if nilai < min2Val {
			min2Val = nilai
			min2Name = nama
		}

		i++
	}

	fmt.Println(min1Name)
	fmt.Println(min2Name)
}
