package main

import ("fmt")

func main() {
	var n int
	
	// Membaca jumlah persegi (n)
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	// Perulangan sebanyak n kali untuk setiap sisi persegi
	for i := 0; i < n; i++ {
		var sisi int
		fmt.Scan(&sisi)

		// Menghitung luas dan keliling persegi
		luas := sisi * sisi
		keliling := 4 * sisi

		// Menampilkan hasil luas dan keliling dipisahkan spasi
		fmt.Printf("%d %d\n", luas, keliling)
	}
}