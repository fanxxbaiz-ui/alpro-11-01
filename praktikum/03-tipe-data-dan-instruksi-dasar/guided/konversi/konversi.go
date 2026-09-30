package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan suhu: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}