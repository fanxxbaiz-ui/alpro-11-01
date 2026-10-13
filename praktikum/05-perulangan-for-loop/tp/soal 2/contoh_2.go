package main

import "fmt"

func main() {
	i := 1 // Inisialisasi di luar

	// Hanya menggunakan kondisi
	for i <= 5 {
		fmt.Println("Perulangan ke-", i)
		i++ // Post statement di dalam blok kode
	}
}
