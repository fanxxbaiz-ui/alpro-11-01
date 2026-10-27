package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("Nol")
	} else {
		var jenis, tanda string
		
		if bilangan%2 == 0 {
			jenis = "Genap"
		} else {
			jenis = "Ganjil"
		}

		if bilangan > 0 {
			tanda = "Positif"
		} else {
			tanda = "Negatif"
		}

		fmt.Printf("%s, %s\n", jenis, tanda)
	}
}