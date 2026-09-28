package main

import "fmt"

func main() {
	var nominal int
	// Membaca input nominal uang
	fmt.Scan(&nominal)

	// Menghitung jumlah lembar pecahan 10.000
	sepuluhRibu := nominal / 10000
	sisa := nominal % 10000

	// Menghitung jumlah lembar pecahan 5.000 dari sisa
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	// Menghitung jumlah lembar pecahan 1.000 dari sisa
	seribu := sisa / 1000

	// Menampilakan banyaknya lembar pecahan 10rb, 5rb, dan 1rb
	fmt.Println(sepuluhRibu, limaRibu,seribu)
}