# <h1 align="center">Laporan Praktikum Modul 10 & 11 - FUNGSI DAN PROSEDUR  </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. PENGERTIAN FUNGSI DAN PROSEDUR
Fungsi adalah blok kode modular yang mengembalikan nilai (`return`), sedangkan prosedur adalah blok kode yang menjalankan suatu aksi atau tugas tanpa mengembalikan nilai.
### B. PERBEDAAN FUNGSI DAN PROSEDEUR
#### 1. FUNGSI
Menjalankan tugas tertentu dan mengembalikan suatu nilai hasil pemrosesan (menggunakan kata kunci seperti return) ke bagian program yang memanggilnya.

#### 2. PROSEDUR
Menjalankan sekumpulan instruksi atau aksi tertentu (seperti mencetak teks atau mengubah data) tanpa mengembalikan nilai.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. HARIKE

```go
package main

import "fmt"

// Fungsi untuk mengecek tahun kabisat
func kabisat(tahun int) bool {
	return (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
}

// Fungsi untuk mengembalikan jumlah hari pada bulan tertentu dalam suatu tahun
func jumlahHari(bulan, tahun int) int {
	var hari int
	if bulan == 2 {
		if kabisat(tahun) {
			hari = 29
		} else {
			hari = 28
		}
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		hari = 30
	} else {
		hari = 31
	}
	return hari
}

// Fungsi untuk menghitung hari ke-berapa dalam tahun tersebut menggunakan perulangan
func hariKe(tanggal, bulan, tahun int) int {
	var total int = 0
	var i int
	for i = 1; i < bulan; i++ {
		total = total + jumlahHari(i, tahun)
	}
	total = total + tanggal
	return total
}

func main() {
	var tanggal, bulan, tahun int
	fmt.Scan(&tanggal, &bulan, &tahun)

	// Menghitung urutan hari dalam tahun
	urutanHari := hariKe(tanggal, bulan, tahun)

	// Menentukan jumlah total hari dalam satu tahun (kabisat = 366, biasa = 365)
	var totalHariTahun int
	if kabisat(tahun) {
		totalHariTahun = 366
	} else {
		totalHariTahun = 365
	}

	// Menghitung sisa hari hingga akhir tahun
	sisaHari := totalHariTahun - urutanHari

	// Menampilkan keluaran sesuai format
	fmt.Printf("Hari ke-%d\n", urutanHari)
	fmt.Printf("Sisa hari dalam tahun %d: %d\n", tahun, sisaHari)
}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/guided/fungsi/harike/output%2020.png)

#### Deskripsi
Program ini dibuat untuk memeriksa hari dalam tahun tersebut dan sisa hari dalam tahun tersebut
### 2. PLAYER.GO

```go
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
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/guided/fungsi/player/output%2021.png)


#### Deskripsi
Soal player.go meminta pembuatan program untuk mengolah data skor player menggunakan fungsi terpisah guna menghitung rata-rata skor, mencari skor tertinggi, dan menghitung jumlah player di atas rata-rata.
<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. SAPI.GO

```go
package main

import "fmt"

const NMAX = 10000

type Arr [NMAX]int

func mengisiSapi(T *Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func geserKanan(T *Arr, n int) {
	var temp int = T[n-1]

	for i := n - 1; i > 0; i-- {
		T[i] = T[i-1]
	}

	T[0] = temp
}

func cetak(T Arr, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var T Arr
	mengisiSapi(&T, n)

	k = k % n

	for i := 0; i < k; i++ {
		geserKanan(&T, n)
	}

	cetak(T, n)

}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/guided/prosedur/output%2022.png)


#### Deskripsi
Soal sapi.go meminta pembuatan program menggunakan prosedur untuk mensimulasikan perpindahan posisi sapi di beberapa kandang selama $K$ hari, di mana setiap hari sapi di kandang terakhir akan berpindah ke kandang pertama sementara sapi lainnya bergeser ke kanan, lalu menampilkan susunan akhirnya.


## Unguided

### 1. GEMPA.GO

```go
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
```

##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/unguided/fungsi/output%2024.png)

#### Deskripsi
Soal gempa.go meminta pembuatan program untuk mengolah data gempa bumi (nama wilayah, tahun, dan magnitudo) menggunakan empat fungsi terpisah guna menentukan kategori setiap gempa berdasarkan magnitudonya, menghitung rata-rata magnitudo, mencari gempa terkuat, serta menghitung jumlah gempa berkategori "Besar".

### 2. NIM.GO

```go
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
```

##### Output

![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/unguided/prosedur/output%2023.png)

#### Deskripsi
Soal `[nim.go]` meminta pembuatan program menggunakan prosedur untuk mencari NIM mahasiswa dalam daftar, lalu menggantinya dengan NIM baru jika ditemukan dan menampilkan daftar NIM yang telah diperbarui.
<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Kelima soal tersebut melatih penggunaan **fungsi** (untuk perhitungan/pengembalian nilai seperti raport, pembersihan duplikat, statistik player, dan data gempa) serta **prosedur** (untuk aksi/manipulasi data seperti pencarian murid, data penduduk, simulasi sapi, dan pembaruan NIM) dalam mengolah *array* dan *struct*.
## Referensi
1. Materi Modul 10 FUNGSI
2. Materi Modul 11 PROSEDUR