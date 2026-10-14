package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	var count int = 0
	// Perulangan untuk setiap mahasiswa sebanyak m kali
	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc float64
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		// Menghitung rata-rata dari 6 nilai tes ECCT
		avg := (p + v + ga + fs + ci + tc) / 6.0

		// Menentukan apakah rata-rata memenuhi syarat minimal 3.33
		if avg >= 3.33 {
			fmt.Println("Ya")
			count++
		} else {
			fmt.Println("Tidak")
		}
	}

	// Menampilkan total mahasiswa yang mendapat sertifikat
	fmt.Println(count)
}
