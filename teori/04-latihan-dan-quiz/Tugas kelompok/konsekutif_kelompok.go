package main

import ("fmt")

func main() {
	var x string
	
	// Membaca masukan X sebagai string untuk memudahkan pemrosesan per digit
	if _, err := fmt.Scan(&x); err != nil {
		return
	}

	isKonsekutif := true

	// Memeriksa setiap digit yang bersebelahan
	for i := 0; i < len(x)-1; i++ {
		// Mengubah karakter digit menjadi integer
		d1 := int(x[i] - '0')
		d2 := int(x[i+1] - '0')

		// Menghitung selisih mutlak antara dua digit yang bersebelahan
		diff := d1 - d2
		if diff < 0 {
			diff = -diff
		}

		// Jika selisihnya bukan 1, maka bukan bilangan konsekutif
		if diff != 1 {
			isKonsekutif = false
			break
		}
	}

	// Menampilkan keluaran berupa nilai boolean (true/false)
	fmt.Println(isKonsekutif)
}