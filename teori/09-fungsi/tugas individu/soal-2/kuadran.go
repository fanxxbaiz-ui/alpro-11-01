package main

import "fmt"

func cekKuadran(x int, y int) string {
	if x == 0 && y == 0 {
		return "Titik Asal"
	} else if x == 0 {
		return "Sumbu Y"
	} else if y == 0 {
		return "Sumbu X"
	} else if x > 0 && y > 0 {
		return "Kuadran I"
	} else if x < 0 && y > 0 {
		return "Kuadran II"
	} else if x < 0 && y < 0 {
		return "Kuadran III"
	} else {
		return "Kuadran IV"
	}
}

func main() {
	var x, y int

	fmt.Print("Masukkan nilai koordinat x dan y: ")
	fmt.Scan(&x, &y)

	hasil := cekKuadran(x, y)
	fmt.Printf("Keluaran = %s\n", hasil)
}
