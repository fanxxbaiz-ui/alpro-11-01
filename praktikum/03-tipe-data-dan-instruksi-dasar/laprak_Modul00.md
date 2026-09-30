# <h1 align="center">Laporan Praktikum Modul  - TIPE DATA DAN INTRUKSI DASAR </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. TIPE DATA
Tipe data adalah klasifikasi dalam pemrograman yang menentukan jenis nilai dari suatu variabel serta jenis operasi yang dapat dilakukan terhadapnya, sehingga komputer dapat mengalokasikan memori secara efisien dan memproses data tanpa keliru. Tipe dasar yang sering digunakan meliputi bilangan bulat (integer), desimal (float), teks (string), karakter (char), dan logika (boolean). Penentuan tipe data sangat penting untuk mencegah error operasi karena sistem memperlakukan tiap tipe secara spesifik—seperti membedakan operasi penjumlahan matematika pada angka 10 + 10 yang menghasilkan 20 dengan penggabungan teks pada "10" + "10" yang menghasilkan "1010"

### B. INTRUKSI DASAR

#### 1. Instruksi Sekuensial (Sequential Statements)
Instruksi yang dieksekusi secara berurutan baris demi baris, dari atas ke bawah.

- Pengisian Nilai (Assignment): Menyimpan data ke dalam variabel.

Contoh: umur = 20

- Input / Output: Menerima masukan dari pengguna atau menampilkan hasil ke layar.

Contoh: print("Pendaftaran Berhasil") atau input()

#### 2. Instruksi Perulangan (Iteration / Repetition / Looping)
Instruksi yang digunakan untuk menjalankan sekumpulan perintah secara berulang-ulang selama kondisi tertentu masih terpenuhi.

- for loop: Digunakan ketika jumlah perulangannya sudah diketahui pasti (misal: diulang 10 kali).

- while loop: Digunakan ketika perulangan bergantung pada kondisi tertentu yang belum pasti jumlahnya.

- do-while loop: Memastikan instruksi dijalankan setidaknya satu kali sebelum memeriksa kondisi perulangan.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. KASIR.GO

```go
package main

import "fmt"

func main() {
	var x int

	fmt.Println("Masukkan nominal: ")
	fmt.Scan(&x)

	var sepuluhRibuan int + x / 10000
	var sisa int = x % 10000

	var limaRibuan int = sisa / 5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt.Println(sepuluhRibuan, limaRibuan, seribuan)
}
```
#### Deskripsi
Logika program ini memecah jumlah total uang kembalian ($x$) dari nilai pecahan terbesar ke terkecil dengan membagi total kembalian dengan 10.000 (x / 10000) untuk mendapatkan jumlah lembar sepuluh ribuan, lalu mengambil sisa uang menggunakan operasi modulo (x % 10000). Sisa uang tersebut kemudian dibagi dengan 5.000 (sisa / 5000) untuk menghitung lembar lima ribuan dan dimodulo kembali (sisa % 5000) untuk memperbarui sisa uang. Akhirnya, sisa uang terakhir dibagi dengan 1.000 (sisa / 1000) untuk menentukan jumlah lembar seribuan, lalu ketiga hasil jumlah lembaran tersebut dicetak secara berurutan dipisahkan oleh spasi.

### 2. KONVERSI.GO

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Print("Masukan suhu: ")
	fmt.Scan(&celcius)

	fmt.Println(celcius + 273)
}
```
#### Deskripsi
Logika program ini diawali dengan menerima masukan nilai suhu dalam derajat Celsius bertipe bilangan real (float64), kemudian mengonversinya ke satuan Kelvin dengan menambahkan konstanta 273 ke dalam nilai Celsius tersebut (K = C + 273), dan diakhiri dengan mencetak hasil kalkulasi suhu dalam satuan Kelvin tersebut ke layar.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. TUKAR.GO

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	temp := x
	x = z
	z = y
	y = temp

	fmt.Println(x, y, z)
}
```
#### Deskripsi
Logika program ini diawali dengan menerima masukan tiga bilangan bulat untuk variabel x, y, dan z. Pertukaran nilai kemudian dilakukan sesuai ketentuan di mana nilai x yang baru diisi oleh nilai z awal, nilai y yang baru diisi oleh nilai x awal, dan nilai z yang baru diisi oleh nilai y awal. Proses pertukaran tersebut dapat memanfaatkan variabel bantuan sementara atau fitur penugasan ganda secara simultan di Go (x, y, z = z, x, y), lalu diakhiri dengan mencetak nilai x, y, dan z yang telah berhasil dipertukarkan ke layar.

## Unguided

### 1. KONVERSI.GO

```go
 package main

import "fmt"

func main() {
	var totalHari, tahun, bulan, minggu, sisaHari int

    // Input jumlah hari
    fmt.Scan(&totalHari)

    // Hitung tahun (1 tahun = 360 hari) dan update sisa hari
    tahun = totalHari / 360
    sisaHari = totalHari % 360

    // Hitung bulan (1 bulan = 30 hari) dan update sisa hari
    bulan = sisaHari / 30
    sisaHari = sisaHari % 30

    // Hitung minggu (1 minggu = 7 hari) dan sisa hari akhir
    minggu = sisaHari / 7
    sisaHari = sisaHari % 7

    // Tampilkan hasil
    fmt.Println(tahun)
    fmt.Println(bulan)
    fmt.Println(minggu)
    fmt.Println(sisaHari)
}
```

##### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi/output%20(7).png)

#### Deskripsi
Logika program ini mengonversi total hari ke dalam satuan tahun, bulan, minggu, dan sisa hari dengan melakukan kalkulasi bertahap menggunakan pembagian bulat (/) dan sisa bagi (%). Pertama, total hari dibagi 360 (hasil dari 12 bulan  30 hari) untuk memperoleh jumlah tahun, lalu sisanya disimpan menggunakan operator modulo. Sisa hari tersebut kemudian dibagi 30 untuk mendapatkan jumlah bulan dan dihitung sisa harinya kembali. Terakhir, sisa hari tersisa dibagi 7 untuk menghitung jumlah minggu, sementara sisa akhir dari modulo 7 menjadi sisa hari, lalu keempat hasil tersebut dicetak secara berurutan.

### 2. SUHU.GO

```go
package main

import "fmt"

func main() {
	var celcius float64

	fmt.Println("Masukan suhu: ")
	fmt.Scan(&celcius)

	r := 4.0 / 5.0 * celcius
	fmt.Println("Reamur: ", r)
}
```

##### Output

![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/suhu/output%20(6).png)

#### Deskripsi
Logika program ini diawali dengan menerima masukan nilai suhu dalam derajat Celsius bertipe bilangan real (float64), kemudian mengonversinya ke satuan Reamur dengan mengalikan nilai Celsius tersebut dengan faktor skala 4.0 / 5.0 atau 0.8 (R = (4/5) C), dan diakhiri dengan mencetak hasil kalkulasi suhu dalam satuan Reamur yang juga berupa bilangan real ke layar.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Secara keseluruhan, materi ini membahas penerapan dasar variabel sebagai penampung data dan operator sebagai pemrosesnya dalam pemrograman Go untuk menyelesaikan berbagai persoalan komputasi logis. Melalui implementasi konversi suhu, pertukaran nilai variabel, hingga pemecahan satuan bertahap seperti konversi hari dan pecahan uang kembalian memanfaatkan kombinasi pembagian bulat serta sisa bagi (modulo), ditunjukkan bagaimana logika aritmatika dibangun secara sistematis.

## Referensi
1. Materi Modul 3- Variabel Dan Operator
