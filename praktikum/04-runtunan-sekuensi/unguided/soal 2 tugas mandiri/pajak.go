package main

import "fmt"

func main() {
	var penghasilan float64

	// 1. Menerima input penghasilan (dalam juta) dari pengguna
	fmt.Scanln(&penghasilan)

	var pajak float64

	// 2. Percabangan if-else untuk menentukan perhitungan pajak progresif
	if penghasilan <= 50 {
		// Kategori 1: Penghasilan <= 50 juta
		pajak = penghasilan * 0.05
	} else if penghasilan <= 100 {
		// Kategori 2: Penghasilan > 50 - 100 juta
		pajak = (50 * 0.05) + ((penghasilan - 50) * 0.10)
	} else if penghasilan <= 200 {
		// Kategori 3: Penghasilan > 100 - 200 juta
		pajak = (50 * 0.05) + (50 * 0.10) + ((penghasilan - 100) * 0.15)
	} else {
		// Kategori 4: Penghasilan > 200 juta
		pajak = (50 * 0.05) + (50 * 0.10) + (100 * 0.15) + ((penghasilan - 200) * 0.20)
	}

	// 3. Mencetak total pajak yang harus dibayar
	fmt.Println(pajak)
}