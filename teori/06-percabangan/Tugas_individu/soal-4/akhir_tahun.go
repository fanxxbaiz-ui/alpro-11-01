package main

import (
	"fmt"
)

func main() {
	var totalBelanja int
	var bersediaKartu bool

	fmt.Scan(&totalBelanja, &bersediaKartu)

	dapatDiskon := totalBelanja >= 100000

	dapatKartu := bersediaKartu && dapatDiskon

	dapatCashback := totalBelanja >= 200000 && dapatKartu

	bayar := float64(totalBelanja)
	if dapatDiskon {
		bayar -= bayar * 0.10
	}
	if dapatCashback {
		bayar -= 75000
	}

	fmt.Printf("Kartu? %t\n", dapatKartu)
	fmt.Printf("Diskon? %t\n", dapatDiskon)
	fmt.Printf("Cashback? %t\n", dapatCashback)
	fmt.Printf("Rp. %.0f\n", bayar)
}