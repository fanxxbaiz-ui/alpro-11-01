package main

import "fmt"

func main() {
	var r, luas float64
	pi := 3.14 // Sesuai petunjuk: assigment variabel pi :=3.14

	// Membaca masukan jari-jari lingkaran  (bilangan real)
	fmt.Scan(&r)

	// Menghitung luas lingkaran
	luas = pi * r * r

	// Menampilkan hasil luas
	fmt.Println(luas)
}