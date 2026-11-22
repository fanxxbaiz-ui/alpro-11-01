package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1001

type tabMurid [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	*posisi = 0
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1 // posisi dimulai dari 1
			break
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			(*jumlah)++
		}
	}
}

func main() {
	var n int
	var T tabMurid
	var cariNama string

	reader := bufio.NewReader(os.Stdin)
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		T[i] = strings.TrimSpace(line)

		if T[i] == "" {
			i--
		}
	}

	line, _ := reader.ReadString('\n')
	cariNama = strings.TrimSpace(line)
	if cariNama == "" {
		fmt.Scan(&cariNama)
	}

	var posisi int
	var ketemu bool
	var jumlahHuruf int

	cariMurid(T, n, cariNama, &posisi, &ketemu)

	if ketemu {
		hurufPertama := cariNama[0]
		hitungHuruf(T, n, hurufPertama, &jumlahHuruf)

		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
		fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", hurufPertama, jumlahHuruf)
	} else {
		var hurufPertama byte
		if len(cariNama) > 0 {
			hurufPertama = cariNama[0]
			hitungHuruf(T, n, hurufPertama, &jumlahHuruf)
		}

		fmt.Println("Murid tidak terdaftar")
		fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", hurufPertama, jumlahHuruf)
	}
}
