package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k) // Membaca kapasitas gudang K[cite: 3]

	var total int
	// Perulangan sebanyak 5 kali untuk menerima 5 masukan pengiriman barang V[cite: 3]
	for i := 0; i < 5; i++ {
		var v int
		fmt.Scan(&v)

		total += v // Mengakumulasi total barang[cite: 3]

		// Memeriksa apakah total barang sudah mencapai atau melebihi kapasitas K[cite: 3]
		if total >= k {
			fmt.Println(true)
		} else {
			fmt.Println(false)
		}
	}
}
