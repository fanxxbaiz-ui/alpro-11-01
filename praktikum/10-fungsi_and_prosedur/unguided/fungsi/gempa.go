package main

import (
	"fmt"
)

const NMAX int = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(m float64) string {
	if m < 4.0 {
		return "Minor"
	} else if m >= 4.0 && m < 6.0 {
		return "Sedang"
	} else if m >= 6.0 && m < 7.0 {
		return "Kuat"
	} else {
		return "Besar"
	}
}

func rataMagnitudo(T tabGempa, n int) float64 {
	total := 0.0
	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}
	return total / float64(n)
}

func idxTerkuat(T tabGempa, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if T[i].magnitudo > T[idxMax].magnitudo {
			idxMax = i
		}
	}
	return idxMax
}

func hitungKategori(T tabGempa, n int, k string) int {
	count := 0
	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var T tabGempa

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		kat := kategori(T[i].magnitudo)
		fmt.Printf("%s %d %.1f %s\n", T[i].nama, T[i].tahun, T[i].magnitudo, kat)
	}

	avg := rataMagnitudo(T, n)
	fmt.Printf("Rata-rata magnitudo: %.2f\n", avg)

	bestIdx := idxTerkuat(T, n)
	fmt.Printf("Gempa terkuat: %s %d %.1f\n", T[bestIdx].nama, T[bestIdx].tahun, T[bestIdx].magnitudo)

	jumlahBesar := hitungKategori(T, n, "Besar")
	fmt.Printf("Jumlah gempa kategori Besar: %d\n", jumlahBesar)
}
