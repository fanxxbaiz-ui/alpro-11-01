package main

import (
	"fmt"
)

const NMAX int = 2002

type tabNIM [NMAX]int

func cariNIM(T tabNIM, n int, x int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i] == x {
			*idx = i
			break
		}
	}
}

func ubahNIM(T *tabNIM, idx int, baru int) {
	T[idx] = baru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n int
	var T tabNIM
	var cari, pengganti int
	var idx int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&cari)

	cariNIM(T, n, cari, &idx)

	if idx != -1 {
		fmt.Scan(&pengganti)

		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", cari, idx+1)

		ubahNIM(&T, idx, pengganti)

		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {

		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}
