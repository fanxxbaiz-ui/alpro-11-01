package main

import "fmt"

func main() {
	var nilai string

	fmt.Print("Masukkan nilai huruf (A/B/C/D/E): ")
	fmt.Scan(&nilai)

	// Penerapan switch case untuk menentukan predikat
	switch nilai {
	case "A":
		fmt.Println("Predikat: Sangat Memuaskan")
	case "B":
		fmt.Println("Predikat: Memuaskan")
	case "C":
		fmt.Println("Predikat: Cukup")
	case "D":
		fmt.Println("Predikat: Kurang")
	case "E":
		fmt.Println("Predikat: Tidak Lulus")
	default:
		fmt.Println("Nilai yang dimasukkan tidak valid!")
	}
}