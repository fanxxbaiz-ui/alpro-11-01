package main

import "fmt"

func main() {
	var lebar, terkecil int
	fmt.Scan(&lebar)
	fmt.Scan(&terkecil)

	for i := 1; i < lebar; i++ {
		var lebar int
		fmt.Scan(&lebar)

		if lebar < terkecil {
			terkecil = lebar
		}
	}

	fmt.Println(terkecil)

}
