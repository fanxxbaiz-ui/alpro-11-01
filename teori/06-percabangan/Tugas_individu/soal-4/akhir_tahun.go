package main

import (
	"fmt"
)

func main() {
	var totalBelanja int
	var bersediaKartu bool

	// Membaca masukan total belanja dan status kesediaan kartu[cite: 3, 4]
	fmt.Scan(&totalBelanja, &bersediaKartu)

	// Menentukan diskon 10% jika belanja minimal Rp 100.000[cite: 3]
	dapatDiskon := totalBelanja >= 100000

	// Memperoleh kartu apabila bersedia dan mendapatkan diskon[cite: 3]
	dapatKartu := bersediaKartu && dapatDiskon

	// Memperoleh cashback jika belanja minimal Rp 200.000 dan mendapatkan kartu[cite: 3]
	dapatCashback := totalBelanja >= 200000 && dapatKartu

	// Menghitung total belanja yang harus dibayar setelah diskon dan cashback[cite: 3]
	bayar := float64(totalBelanja)
	if dapatDiskon {
		bayar -= bayar * 0.10
	}
	if dapatCashback {
		bayar -= 75000
	}

	// Menampilkan keluaran sesuai format soal[cite: 4]
	fmt.Printf("Kartu? %t\n", dapatKartu)
	fmt.Printf("Diskon? %t\n", dapatDiskon)
	fmt.Printf("Cashback? %t\n", dapatCashback)
	fmt.Printf("Rp. %.0f\n", bayar)
}