# <h1 align="center">Laporan Praktikum Modul 7 - Percabangan  </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. PERCABANGAN
Percabangan (branching) adalah struktur kontrol alur dalam pemrograman yang digunakan untuk membuat keputusan. Dengan percabangan, program dapat memilih jalur eksekusi kode yang berbeda-beda berdasarkan terpenuhi atau tidaknya suatu kondisi (bernilai true atau false).
### B. struktur percabangan umumnya dibagi menjadi dua jenis utama:
#### 1. Pernyataan if, else if, dan else
Digunakan untuk mengevaluasi suatu kondisi boolean. Jika kondisi pada if bernilai true, maka blok kode di dalamnya akan dieksekusi. Jika bernilai false, program akan mengecek kondisi else if berikutnya atau langsung menuju ke blok else.

#### 2. Pernyataan switch
switch adalah alternatif percabangan yang lebih bersih dan mudah dibaca dibandingkan menggunakan if-else yang terlalu banyak atau bertingkat (nested). Biasanya digunakan untuk memeriksa nilai variabel tunggal terhadap berbagai macam kemungkinan nilai.


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. KABISAT.GO

```go
package main

import "fmt"

func main() {
	var tahun int
	fmt.Scan(&tahun)

	if (tahun%4 == 0 && tahun%100 != 0 ) || (tahun % 400 == 0) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/guided/guided-1/output%20(47).png)

#### Deskripsi
Program ini dibuat untuk memeriksa apakah suatu tahun yang diinputkan merupakan tahun kabisat atau bukan, di mana suatu tahun disebut kabisat jika habis dibagi 4 dan tidak habis dibagi 100, atau jika habis dibagi 400.
### 2. SEGITIGA.GO

```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	if a <= 0 || b <= 0 || c <= 0 || a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Bukan segitiga")
	} else if  a == b && b == c {
		fmt.Println("Segitiga sama sisi")
	} else if a == b || a == c || b == c {
		fmt.Println("Segitiga sama kaki")
	} else {
		fmt.Println("Segitiga sembarang")
	}
}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/guided/guided-2/output%20(48).png)


#### Deskripsi
Program ini menerima tiga nilai panjang sisi segitiga (a, b, dan c) untuk memeriksa validitasnya serta mengklasifikasikannya ke dalam kategori segitiga sama sisi, segitiga sama kaki, segitiga sembarang, atau menyatakan "Bukan segitiga" jika tidak memenuhi aturan panjang sisi.
<!-- Tambahkan blok file/kode lain sesuai jumlah file pada soal guided -->

### 3. KALKULATOR.GO

```go
package main

import "fmt"

func main() {
	var a, b int
	var op string

	fmt.Scan(&a, &op, &b)

	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dnegan nol")
		} else {
			fmt.Println("%.2f\n", float64(a)/float64(b))
		}
	case "%":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol")
		} else {
			fmt.Println(a % b)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/guided/guided-3/output%20(49).png)


#### Deskripsi
Program ini berfungsi sebagai kalkulator otomatis menggunakan struktur switch untuk memproses dua bilangan bulat berdasarkan operator aritmatika (+, -, *, /, %) serta menangani kesalahan seperti pembagian dengan nol atau operator yang tidak valid.

### 4. ONGKIR.GO

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var berat float64
	var zona int
	var layanan string

	fmt.Scan(&berat, &zona, &layanan)

	if berat > 30 {
		fmt.Println("Berat melebihi batas maksimum (30 kg)")
		return
	} 

	var tarifPerkg int

	switch zona {
	case 1:
		tarifPerkg = 10000
	case 2:
		tarifPerkg = 15000
	case 3:
		tarifPerkg = 25000
	default:
		fmt.Println("zona tidak valid")
		return
	}

	beratBulat := int(math.Ceil(berat))
	total := beratBulat * tarifPerkg

	if layanan == "Express" {
		total += total / 2
	} else if layanan != "Reguler" {
		fmt.Println("Layanan tidak valid")
		return
	}

	fmt.Println(beratBulat)
	fmt.Println(total)
}
```
##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/guided/guided-4/output%20(50).png)

#### Deskripsi
Program ini menghitung total biaya pengiriman paket secara otomatis berdasarkan berat barang yang dibulatkan ke atas, zona wilayah tujuan (Zona 1, 2, atau 3), serta pilihan jenis layanan Reguler atau Express yang mengenakan biaya tambahan sebesar 50%


## Unguided

### 1. BESOK.GO

```go
 package main

import "fmt"

func isKabisat(tahun int) bool {
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return false
	} else if tahun%4 == 0 {
		return true
	}
	return false
}

func getJumlahHari(bulan, tahun int) int {
	switch bulan {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isKabisat(tahun) {
			return 29
		}
		return 28
	default:
		return 0
	}
}

func main() {
	var d, m, y int


	fmt.Scan(&d, &m, &y)

	if m < 1 || m > 12 {
		fmt.Println("Tanggal tidak valid")
		return
	}

	maxHari := getJumlahHari(m, y)

	if d < 1 || d > maxHari {
		fmt.Println("Tanggal tidak valid")
		return
	}

	nextD := d + 1
	nextM := m
	nextY := y

	if nextD > maxHari {
		nextD = 1
		nextM++
		if nextM > 12 {
			nextM = 1
			nextY++
		}
	}

	fmt.Println("Valid")
	fmt.Printf("%d/%d/%d\n", nextD, nextM, nextY)
}
```

##### Output
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/unguided/soal-1/output%20(51).png)

#### Deskripsi
program besok.go yang bertujuan untuk memvalidasi input tanggal, bulan, dan tahun (termasuk memperhitungkan batasan hari dalam bulan serta status tahun kabisat) untuk menampilkan tanggal keesokan harinya

### 2. PARKIR.GO (SOAL A)

```go
package main

import "fmt"

func main() {
	var jenis string
	var jam int

	fmt.Scan(&jenis, &jam)

	totalTarif := 0

	if jenis == "Mobil" {
		totalTarif = 5000 + (jam-1)*3000

		if jam > 24 {
			totalTarif += 50000
		}
	} else if jenis == "Motor" {
		totalTarif = 3000 + (jam-1)*2000

		if jam > 24 {
			totalTarif += 30000
		}
	}

	fmt.Println(totalTarif)
}
```

##### Output

![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/unguided/soal-2/soal_A/output%20(52).png)

#### Deskripsi
Soal A [parkir.go] yang bertujuan untuk menghitung total biaya parkir kendaraan (Mobil atau Motor) berdasarkan durasi parkir dalam jam dengan memperhitungkan tarif jam pertama, tarif jam berikutnya, serta biaya denda tambahan jika durasi parkir melebihi 24 jam.   

<!-- Duplikasi blok "### [nama_soal]" sesuai jumlah folder soal di dalam unguided -->


## Kesimpulan
Percabangan dalam bahasa Go adalah bahwa struktur kontrol alur seperti `if`, `else if`, `else`, dan `switch` sangat esensial digunakan untuk membuat keputusan logis berdasarkan kondisi tertentu. Melalui pemanfaatan percabangan, program dapat mengeksekusi jalur kode yang berbeda-beda untuk menyelesaikan berbagai studi kasus komputasi, mulai dari validasi tanggal, penentuan tahun kabisat, klasifikasi jenis segitiga, operasi aritmatika kalkulator, hingga perhitungan tarif parkir dan ongkos kirim secara otomatis.

## Referensi
1. Materi Modul 7 RUNTUNAN/SEKUENSI