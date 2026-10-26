package main

import (
	"fmt"
)

func main() {
	var p int
	fmt.Scan(&p)

	// Kapasitas maksimum total
	// Appa dewasa: 3 ekor x @5 orang = 15 orang
	// Appa kecil: 5 ekor x @2 orang = 10 orang
	
	// Hitung Appa dewasa (maksimal 3)
	dewasa := p / 5
	if dewasa > 3 {
		dewasa = 3
	}
	
	// Sisa orang setelah diangkut Appa dewasa
	sisa := p - (dewasa * 5)
	if sisa < 0 {
		sisa = 0
	}

	// Hitung Appa kecil (maksimal 5)
	kecil := sisa / 2
	if kecil > 5 {
		kecil = 5
	}

	// Sisa orang yang tidak bisa berangkat
	sisa = sisa - (kecil * 2)
	if sisa < 0 {
		sisa = 0
	}

	// Cetak hasil sesuai kondisi
	if dewasa > 0 && kecil > 0 && sisa > 0 {
		fmt.Printf("Dewasa %d, kecil %d, dan %d tak berangkat\n", dewasa, kecil, sisa)
	} else if dewasa > 0 && kecil > 0 {
		fmt.Printf("Dewasa %d, kecil %d\n", dewasa, kecil)
	} else if dewasa > 0 && sisa > 0 {
		fmt.Printf("Dewasa %d, dan %d tak berangkat\n", dewasa, sisa)
	} else if dewasa > 0 {
		fmt.Printf("Dewasa %d\n", dewasa)
	} else if kecil > 0 {
		fmt.Printf("kecil %d\n", kecil)
	} else {
		fmt.Println("Tidak ada yang berangkat")
	}
}