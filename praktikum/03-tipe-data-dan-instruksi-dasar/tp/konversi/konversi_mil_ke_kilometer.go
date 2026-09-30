package main

import "fmt"

func main() {
	var mil float64

    // Membaca masukan nilai mil bertipe float64
    fmt.Scan(&mil)

    // Menghitung konversi ke kilometer
    km := mil * 1.6

    // Mencetak hasil konversi dengan format 1 angka di belakang koma
    fmt.Printf("%.1f\n", km)
}