# <h1 align="center">Laporan Praktikum Modul [2] - [Pemrograman Bahasa Go]</h1>
<p align="center">[Fathan Al Akbar] - [109092630006]</p>

## Dasar Teori

### A. BAHASA PEMROGRAMAN GO
Bahasa pemrograman Go (Golang) adalah bahasa pemrograman bersumber terbuka (open-source), terkompilasi (compiled), dan berketikan statis (statically typed). Bahasa ini dirancang pada tahun 2007 oleh Robert Griesemer, Rob Pike, dan Ken Thompson di Google, lalu dipublikasikan secara resmi ke publik pada tahun 2009.

Landasan teori perancangan Go berakar pada kebutuhan untuk memecahkan masalah rekayasa perangkat lunak skala besar (software engineering at scale), seperti kompilasi lambat, ketergantungan kode yang rumit, serta tingginya kompleksitas pemrograman konkuren pada arsitektur prosesor multiteras (multicore).
### B. PACKAGE DAN STRUKTUR PROGRAM DI GO

#### 1. PENGERTIAN PACKAGE MAIN DAN FUNC MAIN
Package main =
Merupakan deklarasi paket khusus yang memberi tahu kompilator Go bahwa berkas tersebut ditujukan untuk dikompilasi menjadi program yang dapat dieksekusi (berkas biner / .exe), bukan sebagai pustaka (library) atau modul pendukung yang diimpor oleh berkas lain. Setiap proyek Go yang ingin dijalankan secara langsung wajib memiliki minimal satu berkas dengan package main.

Func main =Merupakan titik masuk utama (entry point) tempat eksekusi program dimulai. Ketika berkas biner dijalankan, runtime Go akan langsung memanggil fungsi main() ini. Ciri khas fungsi main() di Go adalah tidak menerima parameter masukan dan tidak mengembalikan nilai (return value). Ketika baris terakhir dalam fungsi main() selesai dieksekusi, program akan berhenti secara otomatis.

#### 2. TIPE DATA
Tipe data adalah klasifikasi atau atribut yang memberi tahu kompilator atau interpreter mengenai jenis nilai yang disimpan oleh suatu variabel, berapa banyak kapasitas memori yang dibutuhkan, serta operasi apa saja yang valid dilakukan terhadap nilai tersebut.

Tanpa tipe data, komputer hanya melihat informasi sebagai urutan bit (0 dan 1) tanpa tahu cara menginterpretasikannya_apakah bit tersebut merepresentasikan angka, teks, alamat memori, atau instruksi logika.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. TUGAR.GO

```go
package main

import "fmt"

func main() {
	var (
		a int
		b int
	)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Scan(&a)
	fmt.Scan(&b)

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
```
#### Deskripsi
Program tukar.go merupakan program Go yang bertujuan untuk mensimulasikan pertukaran nilai antara dua variabel integer (a dan b). Alur eksekusi diawali dengan mencetak nilai bawaan (zero value) dari kedua variabel yang terdeklarasi, dilanjutkan dengan menerima dua input bilangan bulat dari pengguna menggunakan fmt.Scan. Setelah input diterima, pertukaran nilai dilakukan secara langsung menggunakan fitur multiple assignment khas Go (a, b = b, a) tanpa membutuhkan variabel penampung sementara (temporary variable), hingga akhirnya program mencetak nilai a dan b yang telah berhasil ditukar ke layar.

### 2. SUHU.GO

```go
package main

import "fmt"

func main() {
	// Deklarasi variabel (pastikan ejaan 'celcius' konsisten)
	var celcius, reamur, fahrenheit, kelvin float64

	// Membaca masukan
	fmt.Scan(&celcius)

	// Perhitungan konversi suhu
	reamur = celcius * 4.0 / 5.0
	fahrenheit = celcius * 9.0 / 5.0 + 32.0
	kelvin = celcius + 273.15

	// Menampilkan hasil
	fmt.Println(reamur, fahrenheit, kelvin)
}
```
#### Deskripsi
Program suhu.go dalam bahasa Go ini berfungsi untuk mengonversi nilai suhu dari skala Celsius ke skala Reamur, Fahrenheit, dan Kelvin. Kode diawali dengan deklarasi empat variabel bertipe data float64, dilanjutkan dengan membaca input masukan suhu Celsius dari pengguna melalui perintah fmt.Scan. Setelah masukan diterima, program secara otomatis melakukan kalkulasi konversi suhu sesuai dengan rumus matematika masing-masing skala dan mencetak ketiga nilai hasilnya ke layar secara berurutan menggunakan fungsi fmt.Println.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

## Unguided

### 1. CACAHUANG.GO

```go
package main

import "fmt"

func main() {
	var nominal int
	// Membaca input nominal uang
	fmt.Scan(&nominal)

	// Menghitung jumlah lembar pecahan 10.000
	sepuluhRibu := nominal / 10000
	sisa := nominal % 10000

	// Menghitung jumlah lembar pecahan 5.000 dari sisa
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	// Menghitung jumlah lembar pecahan 1.000 dari sisa
	seribu := sisa / 1000

	// Menampilakan banyaknya lembar pecahan 10rb, 5rb, dan 1rb
	fmt.Println(sepuluhRibu, limaRibu,seribu)
}
```

##### Output
12 1 4


#### Deskripsi
Program menghitung pecahan uang yang di butuhkan dari nominal rupiah. Menentukan beberapa jumlah lembar pecahan uang sesedikit mungkin dari Rp.10000, Rp.5000, Rp.1000. Menggunakan operasi / untuk menentukan jumlah lembar setiap pecahan dan % untuk mendapatkan sisa dari pembagian.

### 2. KALKULATOR.GO

```go
package main

import "fmt"

func main() {
	var a, b int
	// Membaca dua input  bilangan bulat
	fmt.Scan(&a, &b)

	// Menghitung hasil operasi aritmatika
	tambah := a + b
	kurang := a - b 
	kali := a * b
	bagi := a / b
	modulo := a % b 

	// Menampilkan hasil keluaran sesuai urutan
	fmt.Println(tambah, kurang, kali, bagi, modulo)
}
```

##### Output
27 13 140 2 6

#### Deskripsi
Kalkulator sederhana input bilangan bulat angka A dan B, dan program menghitung tambah, kurang, kali, bagi, dan sisa dari kedua bilangan tersebut. Bilangan B tidak boleh 0.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
* Tidak berbeda dengan penulisan program sumber dalam bahasa lain, program Go harus 
dibuat menggunakan penyunting teks dan disimpan dalam format teks, bukan dalam 
format dokumen (doc, docx, atau lainnya).
* Setiap program go disimpan dalam file teks dengan ekstensi .go, dengan nama bebas. 
Sebaiknya nama file adalah nama untuk program tersebut.
* Setiap satu program lengkap Go disimpan dalam satu folder tersendiri. Nama folder 
merupakan nama program tersebut. Karena itu secara prinsip, satu program Go dapat 
dipecah dalam beberapa file dengan ekstensi .go selama disimpan dalam folder yang 
sama.

## Referensi
1. Materi Modul 2-pemrograman Bahasa Go
