package main

import "fmt"

func main() {
	// Deklarasi variabel untuk masukan a (bilangan yang dikalikan) dan b (batas perkalian)
	var a, b int
	fmt.Scan(&a, &b)

	// Perulangan dari i = 1 hingga i <= b
	for i := 1; i <= b; i++ {
		// Menampilkan hasil perkalian dengan format: a x i = hasil
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}
