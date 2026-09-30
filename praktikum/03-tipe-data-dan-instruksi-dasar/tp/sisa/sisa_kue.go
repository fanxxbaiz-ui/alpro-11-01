package main

import "fmt"

func main() {
	var y, x int

    // Membaca input y (jumlah kue) dan x (jumlah anggota keluarga)
    fmt.Scan(&y, &x)

    // Menghitung sisa kue menggunakan operator modulo (%)
    sisa := y % x

    // Mencetak hasil sisa kue
    fmt.Println(sisa)
}