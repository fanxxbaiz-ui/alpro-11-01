package main

import (
	"fmt"
)

func main() {
	var h1, m1, h2, m2 int
	
	fmt.Scan(&h1, &m1, &h2, &m2)

	
	if h1 >= 1 && h1 <= 5 {
		h1 += 12
	}
	if h2 >= 1 && h2 <= 5 {
		h2 += 12
	}

	
	totalMenitMulai := h1*60 + m1
	totalMenitSelesai := h2*60 + m2
	durasiMenit := totalMenitSelesai - totalMenitMulai

	jam := durasiMenit / 60
	menit := durasiMenit % 60

	
	fmt.Printf("%d jam %d menit\n", jam, menit)
}