package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var count int = 0
	// Melakukan iterasi dari 1 sampai n untuk mencari faktor
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan prima (tepat memiliki 2 faktor)
	if count == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
