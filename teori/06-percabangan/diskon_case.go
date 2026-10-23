package main

import "fmt"

func main() {
	var totalBelanja float64
	var isCLO bool

	fmt.Scan(&totalBelanja, &isCLO)

	if isCLO {
		totalBelanja = totalBelanja - (totalBelanja * 0.35)
	}

	fmt.Printf("%.0f\n", totalBelanja)
}
