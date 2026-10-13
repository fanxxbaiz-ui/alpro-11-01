package main

import "fmt"

func main() {
	// Deklarasi variabel untuk masukan n
	var n int
	fmt.Scan(&n)

	// Variabel untuk menyimpan jumlah total bilangan ganjil
	var jumlah int = 0

	// Perulangan dari 1 hingga n dengan penambahan 2 (i += 2)
	// karena bilangan ganjil selalu berselisih 2 (1, 3, 5, 7, dst.)
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		jumlah += i // Menambahkan nilai i ke total jumlah
	}

	// Pindah baris untuk mencetak total jumlah bilangan ganjil di bawahnya
	fmt.Println()
	fmt.Println(jumlah)
}
