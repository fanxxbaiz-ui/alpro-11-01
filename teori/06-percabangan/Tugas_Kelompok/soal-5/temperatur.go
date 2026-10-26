package main

import (
	"fmt"
)

func main() {
	var t [5]float64
	// Membaca 5 masukan temperatur[cite: 6]
	for i := 0; i < 5; i++ {
		fmt.Scan(&t[i])
	}

	naik := true
	turun := true

	// Memeriksa perubahan temperatur antar pencatatan
	for i := 0; i < 4; i++ {
		if t[i+1] <= t[i] {
			naik = false
		}
		if t[i+1] >= t[i] {
			turun = false
		}
	}

	// Menampilkan hasil berdasarkan kondisi perubahan[cite: 6]
	if naik {
		fmt.Println("Stabil naik")
	} else if turun {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}