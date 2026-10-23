package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x) // Membaca bilangan bulat positif X[cite: 4]

	isBeautiful := true

	if x > 0 {
		// Mengambil digit paling kanan sebagai pembanding awal
		prevDigit := x % 10
		x /= 10

		// Perulangan untuk membandingkan setiap digit dengan digit berikutnya[cite: 4]
		for x > 0 {
			currentDigit := x % 10
			// Jika digit dari kiri ke kanan menurun (artinya digit saat ini > digit sebelumnya dari kanan), maka bukan bilangan cantik
			if currentDigit > prevDigit {
				isBeautiful = false
				break
			}
			prevDigit = currentDigit
			x /= 10
		}
	}

	// Mencetak nilai boolean yang menyatakan X adalah bilangan cantik atau tidak[cite: 4]
	fmt.Println(isBeautiful)
}
