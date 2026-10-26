package main

import (
	"fmt"
)

func main() {
	var t [5]float64
	
	for i := 0; i < 5; i++ {
		fmt.Scan(&t[i])
	}

	naik := true
	turun := true

	
	for i := 0; i < 4; i++ {
		if t[i+1] <= t[i] {
			naik = false
		}
		if t[i+1] >= t[i] {
			turun = false
		}
	}

	
	if naik {
		fmt.Println("Stabil naik")
	} else if turun {
		fmt.Println("Stabil turun")
	} else {
		fmt.Println("Tidak stabil")
	}
}