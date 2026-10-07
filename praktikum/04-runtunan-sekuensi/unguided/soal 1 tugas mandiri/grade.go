package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 1. Membaca baris pertama untuk nama (bisa mengandung spasi)
	var nama string
	if scanner.Scan() {
		nama = scanner.Text()
	}

	// 2. Membaca baris kedua untuk nilai
	var nilai float64
	if scanner.Scan() {
		inputNilai := strings.TrimSpace(scanner.Text())
		nilai, _ = strconv.ParseFloat(inputNilai, 64)
	}

	var grade string

	// 3. Penentuan grade menggunakan switch
	switch {
	case nilai > 80:
		grade = "A"
	case nilai >= 70: // Nilai 80 menghasilkan B
		grade = "B"
	case nilai >= 60:
		grade = "C"
	case nilai >= 50:
		grade = "D"
	default:
		grade = "E"
	}

	// 4. Menampilkan keluaran sesuai format soal
	fmt.Println(grade)
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}