# <h1 align="center">Tugas Pendahuluan Modul 5  - PERULANGAN FOR </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. Apa fungsi utama dari struktur perulangan for dalam program Go? Sebutkan dan jelaskan tiga komponen pada perulangan for (inisialisasi, kondisi, dan post statement)! 

JAWABAN =
Fungsi Utama:
Fungsi utama dari struktur perulangan for dalam program Go adalah untuk mengeksekusi blok kode secara 
berulang-ulang (repetitive execution) selama kondisi tertentu terpenuhi (true), atau berdasarkan
 rentang/jumlah perulangan yang telah ditentukan.

 Tiga Komponen Utama Perulangan for:

1) Inisialisasi (Init Statement):
Penjelasan: Komponen yang dieksekusi sekali saja sebelum perulangan dimulai.
 Biasanya digunakan untuk mendeklarasikan dan memberi nilai awal pada 
 variabel counter/penghitung.
Contoh: i := 1

2) Kondisi (Condition):
Penjelasan: Ekspresi boolean yang dievaluasi sebelum setiap iterasi/putaran perulangan.
 Jika bernilai true, kode di dalam perulangan akan dijalankan;
  jika false, perulangan akan berhenti.
Contoh: i <= 10

3) Post Statement (Perubahan Nilai):
Penjelasan: Komponen yang dieksekusi di akhir setiap iterasi/putaran perulangan. 
Biasanya digunakan untuk mengubah (menambah atau mengurangi) nilai variabel counter.
Contoh: i++ (artinya i = i + 1)




#### Deskripsi
soal ini menuntut penjelasan yang mendalam mengenai tiga komponen utama yang menyusun sintaks perulangan for, yaitu inisialisasi untuk mendefinisikan nilai awal variabel pencacah, kondisi boolean sebagai syarat berjalannya perulangan, serta post statement untuk mengubah nilai variabel pada setiap akhir iterasi.

### 2. Bahasa Go tidak memiliki perulangan while dan do-while. Jelaskan perbedaan antara for dengan tiga komponen dan for yang hanya menggunakan kondisi! Berikan contoh kode masing-masing. 

JAWABAN =
1) for dengan Tiga Komponen (Standard For Loop):
Digunakan ketika jumlah perulangan sudah diketahui dengan pasti sejak awal. 
Struktur ini mencakup inisialisasi, kondisi, dan post statement dalam satu baris.

2) for hanya Menggunakan Kondisi (While-like For Loop):
Digunakan ketika jumlah perulangan tidak diketahui secara pasti di awal, melainkan bergantung pada suatu kondisi yang harus
dipenuhi (berfungsi sama seperti perulangan while pada bahasa pemrogramam lain). Inisialisasi dilakukan di luar
sintaks for, dan penambahan/perubahan nilai variabel dilakukan di dalam blok kode perulangan.

CONTOH 1 = 

```go
package main

import "fmt"

func main() {
    // Memperlihatkan perulangan 1 sampai 5
    for i := 1; i <= 5; i++ {
        fmt.Println("Perulangan ke-", i)
    }
}
```
CONTOH 2 =
```go
package main

import "fmt"

func main() {
	i := 1 // Inisialisasi di luar

	// Hanya menggunakan kondisi
	for i <= 5 {
		fmt.Println("Perulangan ke-", i)
		i++ // Post statement di dalam blok kode
	}
}

```


#### Deskripsi
soal nomer 2 diminta menguraikan perbedaan mendasar antara perulangan for dengan tiga komponen lengkap versus perulangan for yang hanya menyertakan kondisi tunggal (yang dapat berfungsi menyerupai perulangan while).


### 3. GANJIL.GO

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk masukan n
	var n int
	fmt.Scan(&n)

	// Variabel untuk menyimpan jumlah total bilangan ganjil
	var jumlah int = 0

	// Perulangan dari 1 hingga n dengan penambahan 2 (i += 2)
	// karena bilangan ganjil selalu berselisih 2 (1, 3, 5, 7, dst.)
	for i := 1; i <= n; i += 2 {
		fmt.Print(i, " ")
		jumlah += i // Menambahkan nilai i ke total jumlah
	}

	// Pindah baris untuk mencetak total jumlah bilangan ganjil di bawahnya
	fmt.Println()
	fmt.Println(jumlah)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/soal%203/output%20(17).png)

#### Deskripsi
 Soal ketiga mengarah pada studi kasus pembuatan program praktis bernama ganjil.go yang menerapkan perulangan for. Program ini berfokus pada pemrosesan masukan berupa sebuah bilangan bulat positif n yang diinputkan oleh pengguna. Keluaran yang harus dihasilkan program mencakup deret bilangan ganjil berurutan dari angka 1 hingga mencapai nilai n, serta menampilkan total penjumlahan keseluruhan dari deret bilangan ganjil tersebut di akhir baris.   
### 4. PERKALIAN.GO

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel untuk masukan a (bilangan yang dikalikan) dan b (batas perkalian)
	var a, b int
	fmt.Scan(&a, &b)

	// Perulangan dari i = 1 hingga i <= b
	for i := 1; i <= b; i++ {
		// Menampilkan hasil perkalian dengan format: a x i = hasil
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}
}

```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/05-perulangan-for-loop/tp/soal%204/output%20(18).png)

#### Deskripsi
Program ini dirancang guna menampilkan tabel perkalian dari suatu bilangan dasar a dengan memanfaatkan perulangan for. Masukan yang diterima berupa dua buah bilangan bulat positif, yakni a sebagai angka yang dikalikan dan b sebagai pembatas akhir perkalian. Program kemudian mencetak sebanyak b baris keluaran hasil perkalian yang tersusun secara rapi mulai dari a x 1 hingga mencapai batas akhir a x b.
## Kesimpulan
Materi perulangan `for` dalam bahasa Go mencakup pemahaman teoretis mengenai fungsi serta tiga komponen utamanya—yaitu inisialisasi, kondisi, dan post statement—sekaligus fleksibilitas penggunaannya sebagai pengganti struktur `while`. Lebih lanjut, materi ini mengimplementasikan konsep tersebut secara praktis melalui penyusunan program untuk menampilkan deret bilangan ganjil beserta total penjumlahannya serta pembuatan tabel perkalian dinamis berdasarkan masukan data dari pengguna.
