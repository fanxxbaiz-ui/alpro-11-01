package main

import "fmt"

func main() {
	var totalHari, tahun, bulan, minggu, sisaHari int

    // Input jumlah hari
    fmt.Scan(&totalHari)

    // Hitung tahun (1 tahun = 360 hari) dan update sisa hari
    tahun = totalHari / 360
    sisaHari = totalHari % 360

    // Hitung bulan (1 bulan = 30 hari) dan update sisa hari
    bulan = sisaHari / 30
    sisaHari = sisaHari % 30

    // Hitung minggu (1 minggu = 7 hari) dan sisa hari akhir
    minggu = sisaHari / 7
    sisaHari = sisaHari % 7

    // Tampilkan hasil
    fmt.Println(tahun)
    fmt.Println(bulan)
    fmt.Println(minggu)
    fmt.Println(sisaHari)
}