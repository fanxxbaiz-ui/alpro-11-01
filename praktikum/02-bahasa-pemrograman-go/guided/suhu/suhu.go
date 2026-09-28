package main

import "fmt"

func main() {
	// Deklarasi variabel (pastikan ejaan 'celcius' konsisten)
	var celcius, reamur, fahrenheit, kelvin float64

	// Membaca masukan
	fmt.Scan(&celcius)

	// Perhitungan konversi suhu
	reamur = celcius * 4.0 / 5.0
	fahrenheit = celcius * 9.0 / 5.0 + 32.0
	kelvin = celcius + 273.15

	// Menampilkan hasil
	fmt.Println(reamur, fahrenheit, kelvin)
}