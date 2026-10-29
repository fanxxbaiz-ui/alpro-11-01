package main

import "fmt"

func main() {
	var jenis string
	var jam int

	fmt.Scan(&jenis, &jam)

	totalTarif := 0

	if jenis == "Mobil" {
		totalTarif = 5000 + (jam-1)*3000

		if jam > 24 {
			totalTarif += 50000
		}
	} else if jenis == "Motor" {
		totalTarif = 3000 + (jam-1)*2000

		if jam > 24 {
			totalTarif += 30000
		}
	}

	fmt.Println(totalTarif)
}