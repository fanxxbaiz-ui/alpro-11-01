package main

import (
	"fmt"
)

func main() {
	var h1, m1, h2, m2 int
	// Membaca masukan jam dan menit mulai serta selesai parkir
	fmt.Scan(&h1, &m1, &h2, &m2)

	// Konversi jam siang/sore (1 s.d. 5) ke format 24 jam
	if h1 >= 1 && h1 <= 5 {
		h1 += 12
	}
	if h2 >= 1 && h2 <= 5 {
		h2 += 12
	}

	// Menghitung total waktu dalam menit
	totalMenitMulai := h1*60 + m1
	totalMenitSelesai := h2*60 + m2
	durasiMenit := totalMenitSelesai - totalMenitMulai

	jam := durasiMenit / 60
	menit := durasiMenit % 60

	// Menampilkan durasi parkir dalam jam dan menit[cite: 3]
	fmt.Printf("%d jam %d menit\n", jam, menit)
}