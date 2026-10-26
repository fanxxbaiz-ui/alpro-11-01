package main

import (
	"fmt"
	"math"
)

func main() {
	var bulanSebelumnya, bulanIni float64
	fmt.Scan(&bulanSebelumnya, &bulanIni)

	selisih := bulanIni - bulanSebelumnya

	if selisih > 0 {
		
		fmt.Printf("Peningkatan sebesar %.0f\n", selisih)
	} else if selisih < 0 {
		fmt.Printf("Penurunan sebesar %.0f\n", math.Abs(selisih))
	} else {
		fmt.Println("Tetap")
	}
}