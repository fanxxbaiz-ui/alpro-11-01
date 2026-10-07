package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	// Menerima input tahun dan nama bulan
	fmt.Scan(&tahun, &bulan)

	// Logika pengecekan tahun kabisat
	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	// Logika penentuan jumlah hari berdasarkan bulan
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if isKabisat {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		// Menampilkan "-" jika nama bulan tidak sesuai format/tidak valid
		fmt.Println("-")
	}
}