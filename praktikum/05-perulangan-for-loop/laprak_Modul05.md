# <h1 align="center">Laporan Praktikum Modul 5 - Perulangan For  </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. PERULANGAN FOR
Dalam bahasa Go, perulangan for adalah satu-satunya struktur kendali perulangan yang tersedia. Berbeda dengan bahasa pemrograman lain yang memiliki banyak jenis perulangan seperti while atau do-while, Go menyatukan seluruh kebutuhan iterasi tersebut ke dalam satu kata kunci, yaitu for.
### B. bentuk dan variasi penggunaan perulangan for dalam bahasa GO

#### 1. Perulangan Standar (Three-Component Loop)
Bentuk ini adalah yang paling umum digunakan, mirip dengan perulangan for pada C, C++, atau Java. Terdiri dari tiga bagian yang dipisahkan oleh titik koma:

1) Inisialisasi: Mendeklarasikan dan menginisialisasi variabel penghitung (counter).

2) Kondisi: Ekspresi boolean yang dievaluasi sebelum setiap iterasi. Jika bernilai true, perulangan dilanjutkan; jika false, perulangan berhenti.

3) Post-statement: Pernyataan yang dieksekusi di akhir setiap iterasi (biasanya berupa increment atau decrement).

#### 2. Perulangan dengan Kondisi Saja (Berperilaku seperti while)
Jika komponen inisialisasi dan post-statement dihilangkan, perulangan ini akan bertindak seperti perulangan while pada bahasa pemrograman lain. Perulangan akan terus berjalan selama kondisi bernilai true.

#### 3. Perulangan Tanpa Henti (Infinite Loop)
Jika seluruh komponen di dalam kurung siku/tanpa kondisi dikosongkan sama sekali, perulangan akan berjalan secara terus-menerus tanpa henti. Perulangan jenis ini biasanya dihentikan menggunakan pernyataan break atau return di dalam blok kodenya.
#### 4.Perulangan Menggunakan range
Bentuk for range digunakan untuk melakukan iterasi atau menelusuri elemen-elemen di dalam struktur data koleksi, seperti array, slice, string, map, atau channel.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. TABUNGAN.GO

```go
package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scan(&a, &b, &n)

	total := 0
	tabunganHariIni := a

	for i := 1; i <= n; i++ {
		total = total + tabunganHariIni
		tabunganHariIni = tabunganHariIni + b
	}

	fmt.Println(total)

}
```
##### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/guided/guided-1/output%20(19).png)

#### Deskripsi
Program dalam bahasa Go untuk menghitung total tabungan Budi setelah 
n hari. Budi menabung pada hari pertama sebesar 
a rupiah, dan setiap hari berikutnya jumlah uang yang ditabung bertambah 
b rupiah dari hari sebelumnya.
### 2.TERBESAR.GO

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var angka, max int
	fmt.Scan(&angka)
	max = angka

	for i := 2; i <= n; i++ {
		fmt.Scan(&angka)

		if angka > max {
			max = angka
		}

	}

	fmt.Println(max)
}

```
##### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/guided/guided-2/output(20).png)


#### Deskripsi
Program dalam bahasa Go untuk mencari nilai bilangan terbesar dari sekumpulan n buah bilangan bulat yang dimasukkan oleh pengguna.
<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. BIOSKOP.GO

```go
package main

import "fmt"

func main() {
	var t int
	fmt.Print("Input t (jumlah transaksi): ")
	fmt.Scan(&t)
	totalPendapatan := 0

	for i := 1; i <= t; i++ {
		var jenisFilm, hari string
		var jumlahTiket int

		fmt.Printf("Input transaksi ke-%d (jenis, hari, tiket): ", i)
		fmt.Scan(&jenisFilm, &hari, &jumlahTiket)

		hargaSatuan := 0

		if hari == "Senin" || hari == "Selasa" || hari == "Rabu" || hari == "Kamis" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 40000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 80000
			} else if jenisFilm == "3D" {
				hargaSatuan = 60000
			}
		} else if hari == "Jumat" || hari == "Sabtu" || hari == "Minggu" {
			if jenisFilm == "Reguler" {
				hargaSatuan = 50000
			} else if jenisFilm == "IMAX" {
				hargaSatuan = 90000
			} else if jenisFilm == "3D" {
				hargaSatuan = 70000
			}
		}
		totalHarga := hargaSatuan * jumlahTiket

		if jumlahTiket >= 5 {
			diskon := (totalHarga * 10) / 100
			totalHarga = totalHarga - diskon
		}
		fmt.Printf("Output total harga transaksi ke-%d: %d\n", i, totalHarga)

		totalPendapatan = totalPendapatan + totalHarga
	}
	fmt.Println("Output total pendapatan: ", totalPendapatan)
}

```
##### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/guided/guided-3/output%20(21).png)


#### Deskripsi
Program dalam bahasa Go untuk menghitung total harga per transaksi dan total keseluruhan pendapatan bioskop berdasarkan tarif jenis film (Reguler, 3D, IMAX) dan hari pemutaran (Senin–Kamis atau Jumat–Minggu).
## Unguided

### 1. PRIMA.GO

```go
 package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var count int = 0
	// Melakukan iterasi dari 1 sampai n untuk mencari faktor
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			count++
		}
	}
	fmt.Println()

	// Menentukan apakah bilangan prima (tepat memiliki 2 faktor)
	if count == 2 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

```

##### Output
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/soal-1/output%20(22).png)

#### Deskripsi
Program prima.go merupakan latihan praktikum perulangan (for loop) dalam bahasa Go yang bertujuan untuk mencari seluruh faktor dari sebuah bilangan bulat positif n serta menentukan status bilangan tersebut sebagai bilangan prima atau bukan. Melalui iterasi terstruktur, program ini memindai setiap angka dari 1 hingga n untuk mengidentifikasi pembagi yang valid, menghitung total faktornya, dan mengevaluasi apakah jumlah faktor tersebut tepat dua untuk dikategorikan sebagai bilangan prima.

### 2. ECCT.GO

```go
func main() {
	var m int
	fmt.Scan(&m)

	var count int = 0
	// Perulangan untuk setiap mahasiswa sebanyak m kali
	for i := 0; i < m; i++ {
		var p, v, ga, fs, ci, tc float64
		fmt.Scan(&p, &v, &ga, &fs, &ci, &tc)

		// Menghitung rata-rata dari 6 nilai tes ECCT
		avg := (p + v + ga + fs + ci + tc) / 6.0

		// Menentukan apakah rata-rata memenuhi syarat minimal 3.33
		if avg >= 3.33 {
			fmt.Println("Ya")
			count++
		} else {
			fmt.Println("Tidak")
		}
	}

	// Menampilkan total mahasiswa yang mendapat sertifikat
	fmt.Println(count)
}
```

##### Output

![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/unguided/soal-2/output%20(23).png)

#### Deskripsi
program ecct.go dirancang untuk memproses data akademik secara kolektif dengan menghitung rata-rata dari enam nilai tes ECCT bagi sejumlah mahasiswa. Program ini memanfaatkan perulangan untuk mengevaluasi kelayakan setiap mahasiswa berdasarkan standar nilai minimum 3.33, sekaligus merekapitulasi dan menampilkan total keseluruhan mahasiswa dalam kelompok tersebut yang berhasil mendapatkan sertifikat.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
 program tersebut menunjukkan bagaimana struktur perulangan ("for loop") dalam bahasa Go digunakan secara efektif untuk menyelesaikan persoalan komputasi dan logika, baik dalam analisis matematis untuk menentukan faktor dan bilangan prima, maupun dalam pengolahan data institusional untuk mengevaluasi kelayakan sertifikat mahasiswa.

## Referensi
1. Materi Modul 5 RUNTUNAN/SEKUENSI
2. WEB PEMROGRAMAN
