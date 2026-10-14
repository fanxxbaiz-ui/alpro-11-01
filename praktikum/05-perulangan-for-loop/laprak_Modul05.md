# <h1 align="center">Laporan Praktikum Modul 4 - RUNTUNAN/SEKUENSI </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. RUNTUNAN
Runtunan dalam algoritma pemrograman adalah struktur kontrol dasar yang mengeksekusi serangkaian instruksi secara berurutan dari atas ke bawah, baris demi baris. Setiap langkah hanya akan dijalankan setelah langkah sebelumnya selesai diproses tanpa ada instruksi yang melompat atau diabaikan.

### B. MACAM-MACAM SEKUENSI

#### 1. Sekuensi Linier
Rangkaian instruksi paling dasar yang dieksekusi secara lurus dari baris pertama hingga akhir tanpa ada lompatan, pengondisian, atau pengulangan.

#### 2. Sekuensi Terblok (Dalam Struktur Kontrol)
Urutan perintah yang dikelompokkan di dalam suatu blok percabangan (seperti if-else) atau perulangan (seperti for atau while). Instruksi dalam blok tersebut tetap berjalan berurutan saat kondisi terpenuhi.

#### 3. Sekuensi Asinkron
Alur penulisan kode yang tampak berurutan (menggunakan mekanisme seperti async/await), tetapi eksekusinya dapat dihentikan sementara untuk menunggu proses lain (seperti mengambil data dari server) sebelum melanjutkan ke baris berikutnya.

#### 4.Sekuensi Data (Struktur Data Sekuensial)
Penerapan urutan pada penyimpanan elemen data dalam memori yang diakses berurutan, seperti Array, List, Queue, Stack, dan String.

#### 5.Sekuensi Matematika dalam Algoritma

Urutan angka berpola yang diterapkan dalam penyelesaian masalah komputasi, seperti Sekuensi Fibonacci, Sekuensi Aritmatika, dan Sekuensi Geometri.

<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. GRADE.GO

```go
package main

import (
	"bufio"
	"os"

	"fmt"
)

func main() {
	var nama string
	var nilai float64
	var grade string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
```
##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/soal%201/output%2012.png)

#### Deskripsi
Program menerima input nama dan nilai angka dari seorang siswa.   Nilai angka diklasifikasikan ke dalam nilai huruf A, B, C, D, atau F.   Program mencetak hasil grade huruf beserta kalimat nama siswa.   Menguji penerapan struktur percabangan dalam pengolahan nilai.

### 2.PENILAIAN.GO

```go
package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi (seperti nama lengkap)
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang merepresentasikan keyboard
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	// Menampilkan cetakan Menu ke layar
	fmt.Println("============= Menu =============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")

	// Membaca masukan pilihan (angka).
	// Kita menggunakan Scanln agar saat user menekan 'Enter', karakter enter tersebut
	// ikut diolah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	// Percabangan/Sekuensi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("Masukkan nama siswa: ")

		// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
		scanner := bufio.NewScanner(os.Stdin)

		// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
		scanner.Scan()

		// Mengambil teks (nama) yang baru saja dibaca dan menyimpannya ke variabel
		nama = scanner.Text()

		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scanln(&nilai)

		// Menentukan grade nilai
		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 && nilai < 90 {
			grade = "B"
		} else if nilai >= 70 && nilai < 80 {
			grade = "C"
		} else if nilai >= 60 && nilai < 70 {
			grade = "D"
		} else {
			grade = "F"
		}

		// %s adalah format (placeholder) untuk mencetak data bertipe string
		// %s pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
		fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		// Dijalankan jika pengguna mengetik 0
		fmt.Println("Keluar dari program.")

	} else {
		// Dijalankan jika pengguna mengetik angka selain 1 dan 0
		fmt.Println("Pilihan tidak valid.")
	}
}
```
##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/soal%202/output%20(13).png)


#### Deskripsi
Program menampilkan menu interaktif pilihan sistem penilaian atau keluar.   Opsi 1 memproses masukan nama dan nilai siswa menggunakan logika Soal 1.   Opsi 0 menghentikan eksekusi program dengan mencetak pesan keluar.   Opsi selain 1 dan 0 mencetak peringatan bahwa pilihan tidak valid.

<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. KLASIFIKASI.GO

```go
package main

import (
	"fmt" // Hanya memerlukan fmt untuk keperluan input dan output
)

func main() {
	// Pendefinisian variabel secara eksplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	// Membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	// fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	// Menggunakan switch tanpa ekspresi.
	// Cara kerjanya sama persis seperti deretan if - else if.
	// Program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true).
	switch {
	case usia < 18:
		keterangan = "Masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "Muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "Masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "Pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "Perlu perbaikan karier"
	case usia > 40 && gaji >= 150:
		keterangan = "Profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "Perlu evaluasi finansial"
	}

	// Menampilkan hasil klasifikasi ke layar
	fmt.Println(keterangan)
}
```
##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/soal%203/output%20(14).png)


#### Deskripsi
Program menerima input dua variabel berupa usia dan gaji tahunan.   Menerapkan pemrosesan logika switch tanpa ekspresi pada setiap kondisi.   Memetakan kombinasi usia dan gaji ke dalam salah satu dari 7 kategori.   Mencetak luaran berupa teks keterangan klasifikasi yang sesuai tabel.
## Unguided

### 1. GRADE.GO (TUGAS MANDIRI)

```go
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
```

##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/soal%201%20tugas%20mandiri/output%20(15).png)

#### Deskripsi
Memodifikasi program klasifikasi nilai siswa dengan mengubah struktur if-else menjadi switch.   Menerima masukan berupa nama lengkap siswa yang dapat memuat spasi serta nilai angka.   Mengkategorikan nilai ke dalam penentuan grade sesuai rentang kriteria yang ditentukan.   Mencetak keluaran berupa huruf grade beserta kalimat pernyataan nilai siswa.

### 2. PAJAK.GO

```go
package main

import "fmt"

func main() {
	var penghasilan float64

	// 1. Menerima input penghasilan (dalam juta) dari pengguna
	fmt.Scanln(&penghasilan)

	var pajak float64

	// 2. Percabangan if-else untuk menentukan perhitungan pajak progresif
	if penghasilan <= 50 {
		// Kategori 1: Penghasilan <= 50 juta
		pajak = penghasilan * 0.05
	} else if penghasilan <= 100 {
		// Kategori 2: Penghasilan > 50 - 100 juta
		pajak = (50 * 0.05) + ((penghasilan - 50) * 0.10)
	} else if penghasilan <= 200 {
		// Kategori 3: Penghasilan > 100 - 200 juta
		pajak = (50 * 0.05) + (50 * 0.10) + ((penghasilan - 100) * 0.15)
	} else {
		// Kategori 4: Penghasilan > 200 juta
		pajak = (50 * 0.05) + (50 * 0.10) + (100 * 0.15) + ((penghasilan - 200) * 0.20)
	}

	// 3. Mencetak total pajak yang harus dibayar
	fmt.Println(pajak)
}
```

##### Output

![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/soal%202%20tugas%20mandiri/output%20(16).png)

#### Deskripsi
Menghitung total pajak penghasilan seseorang berdasarkan kategori rentang (bracket) penghasilan.   Menerima input nilai penghasilan dalam satuan juta rupiah berjenis tipe data float64.   Menggunakan percabangan if-else bertingkat untuk menerapkan perhitungan tarif berjenjang.   Mencetak total nominal pajak yang wajib dibayar pengguna sesuai aturan perhitungan progresif.

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Program nomor 1 dan 2 melatih pemahaman struktur kontrol percabangan serta pengolahan tipe data pada bahasa Go melalui dua kasus yang berbeda. Program nomor 1 mengimplementasikan switch untuk mengklasifikasikan nilai angka menjadi grade huruf dan menangani masukan nama berspasi, sedangkan program nomor 2 memanfaatkan percabangan if-else bertingkat dengan tipe data float64 untuk menghitung pajak progresif berjenjang berdasarkan bracket penghasilan. Kombinasi keduanya memperkuat kemampuan penyusunan alur logika kondisional, konversi data, dan kalkulasi angka presisi secara efektif.

## Referensi
1. Materi Modul 4 RUNTUNAN/SEKUENSI
2. BUKU PEMROGRAMAN
