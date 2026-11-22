package main

import (
	"fmt"
)

const NMAX int = 999

type player struct {
	nickname string
	score    int
}

type tabPlayer [NMAX]player

func rataRata(t tabPlayer, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += t[i].score
	}
	return float64(total) / float64(n)
}

func idxTertinggi(t tabPlayer, n int) int {
	idxMax := 0
	for i := 1; i < n; i++ {
		if t[i].score > t[idxMax].score {
			idxMax = i
		}
	}
	return idxMax
}

func hitungDiatas(t tabPlayer, n int, batas float64) int {
	count := 0
	for i := 0; i < n; i++ {
		if float64(t[i].score) > batas {
			count++
		}
	}
	return count
}

func main() {
	var n int
	var t tabPlayer

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t[i].nickname, &t[i].score)
	}

	avg := rataRata(t, n)
	bestIdx := idxTertinggi(t, n)
	aboveAvgCount := hitungDiatas(t, n, avg)

	fmt.Printf("Rata-rata skor: %.2f\n", avg)
	fmt.Printf("Player terbaik: %s (%d)\n", t[bestIdx].nickname, t[bestIdx].score)
	fmt.Printf("Jumlah player di atas rata-rata: %d\n", aboveAvgCount)
}
