package main

import "fmt"

func isKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	}
	return false
}

func getJumlahHari(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isKabisat(tahun) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

func main() {
	var d, m, y int


	fmt.Scan(&d, &m, &y)

	if m < 1 || m > 12 {
		fmt.Println("Tanggal tidak valid")
		return
	}

	maxHari := getJumlahHari(m, y)

	if d < 1 || d > maxHari {
		fmt.Println("Tanggal tidak valid")
		return
	}

	nextD := d + 1
	nextM := m
	nextY := y

	if nextD > maxHari {
		nextD = 1
		nextM++
		if nextM > 12 {
			nextM = 1
			nextY++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", nextD, nextM, nextY)
}