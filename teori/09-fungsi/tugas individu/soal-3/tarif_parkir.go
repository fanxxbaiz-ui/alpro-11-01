package main

import "fmt"

func hitungTarif(jam int) int {
	if jam <= 0 {
		return 0
	} else if jam <= 2 {
		return 5000
	} else {
		return 5000 + (jam-2)*3000
	}
}

func main() {
	var jam int

	fmt.Print("Masukkan lama parkir (jam): ")
	fmt.Scan(&jam)

	totalTarif := hitungTarif(jam)

	fmt.Println(totalTarif)
}
