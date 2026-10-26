package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	// Membaca empat masukan jumlah gol[cite: 2]
	fmt.Scan(&a, &b, &c, &d)

	// Inisialisasi awal nilai maksimum dan minimum
	max := a
	min := a

	// Membandingkan dengan bilangan kedua (b)
	if b > max {
		max = b
	}
	if b < min {
		min = b
	}

	// Membandingkan dengan bilangan ketiga (c)
	if c > max {
		max = c
	}
	if c < min {
		min = c
	}

	// Membandingkan dengan bilangan keempat (d)
	if d > max {
		max = d
	}
	if d < min {
		min = d
	}

	// Menampilkan hasil gol terbanyak dan tersedikit[cite: 2]
	fmt.Printf("%d %d\n", max, min)
}