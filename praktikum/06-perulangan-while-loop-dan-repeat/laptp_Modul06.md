# <h1 align="center">Tugas Pendahuluan Modul 6  - PERULANGAN WHILE </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. LULUS.GO

```go
package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	for nilai != -1 {
		if nilai >= 60 {
			fmt.Println("Anda lulus ujian.")
		} else {
			fmt.Println("Anda tidak lulus ujian.")
		}
		fmt.Print("Masukkan nilai Anda: ")
		fmt.Scan(&nilai)
	}

	fmt.Println("Program selesai.")
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/soal-1/output%20(25).png)


#### Deskripsi
Program untuk mengecek status kelulusan siswa secara berulang berdasarkan nilai ujian (lulus jika > 60, tidak lulus jika < 60) dan berhenti ketika pengguna memasukkan angka -1.   

### 2. DERET.GO

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var jumlah int = 0
	var i int = 1

	// Perulangan while menggunakan for di Go
	for i <= n {
		jumlah += i
		i++
	}

	fmt.Println(jumlah)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/soal-2/output%20(26).png)


#### Deskripsi
Program untuk menghitung hasil penjumlahan deret bilangan dari 1 hingga n (S = 1 + 2 + ... + n) berdasarkan masukan bilangan bulat positif n.  


### 3. DUA TERENDAH.GO

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var nama, min1Name, min2Name string
	var nilai, min1Val, min2Val int

	// Inisialisasi nilai minimum awal dengan angka yang sangat besar
	min1Val = 999999
	min2Val = 999999

	var i int = 1
	// Perulangan while menggunakan for di Go
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if nilai < min1Val {
			min2Val = min1Val
			min2Name = min1Name
			min1Val = nilai
			min1Name = nama
		} else if nilai < min2Val {
			min2Val = nilai
			min2Name = nama
		}

		i++
	}

	fmt.Println(min1Name)
	fmt.Println(min2Name)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/soal-3/output%20(27).png)

#### Deskripsi
Program untuk mencari dan menampilkan dua nama siswa dengan nilai terendah serta nilai terendah kedua dari total n data siswa yang diinputkan.


### 4. BIAYA.GO

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var biaya float64 = 10000000
	var i int = 1

	// Perulangan while menggunakan for di Go
	for i < n {
		biaya += biaya * 0.06
		i++
	}

	fmt.Printf("%.0f\n", biaya)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/tp/soal-4/output%20(28).png)

#### Deskripsi
Program untuk menghitung perkiraan biaya kuliah pada tahun ke-$n$ dengan biaya awal Rp 10.000.000 yang mengalami kenaikan sebesar 6% setiap tahunnya.
## Kesimpulan
Perulangan *while* adalah struktur kontrol dalam pemrograman yang digunakan untuk mengeksekusi blok kode secara berulang selama kondisi bernilai benar (`true`). Khusus dalam bahasa Go, konsep *while* diimplementasikan menggunakan kata kunci `for` yang hanya berisi kondisi (`for kondisi`) karena Go tidak menyediakan kata kunci `while` tersendiri. Struktur ini sangat fleksibel untuk berbagai penyelesaian logika—seperti sentinel berhenti, perhitungan deret, pencarian data minimum, maupun kalkulasi finansial—dengan catatan bahwa variabel pengontrol harus selalu diperbarui di dalam blok agar program terhindar dari perulangan tak terbatas (infinite loop).
