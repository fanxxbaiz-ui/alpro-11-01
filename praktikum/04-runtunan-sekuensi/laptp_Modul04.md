# <h1 align="center">Tugas Pendahuluan Modul 4  - RUNTUNAN/SEKUENSI </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. Evaluasi ekspresi kontrol dalam GO

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	// Menampilkan langsung hasil evaluasi boolean (true/false)
	fmt.Println("1 :", intNum > 5)
	fmt.Println("2 :", intNum >= 5 && intOther < 11)
	fmt.Println("3 :", sngNum != -1 || intOther < 0)
	fmt.Println("4 :", !(intNum > 3) || intNum <= 5)
	fmt.Println("5 :", !(intOther >= intNum))
	fmt.Println("6 :", 0-sngNum > 0)
	fmt.Println("7 :", 4/2 == intOther/intNum)
	fmt.Println("8 :", intOther%2 == 0)
	fmt.Println("9 :", intOther+2*intNum != 30 || !(sngNum > 0))
	fmt.Println("10:", intOther > 0 && intNum > 0 || sngNum > 0)
	fmt.Println("11:", sngNum > 0 || (intNum >= 0 && -1*intOther == -10))
	fmt.Println("12:", intNum == 5)
	fmt.Println("13:", intNum > 0 || (sngNum <= 0 && intOther == 13))
	fmt.Println("14:", !(!(!(!(intNum > 0)))))
}
```

##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/evaluasi_ekspresi_Kontroldalamgo/output%20(8).png)



#### Deskripsi
Mengevaluasi 14 bentuk ekspresi logika dan perbandingan berdasarkan variabel intNum, intOther, dan sngNum.   Menguji pemahaman operator relasional (>, <, ==, !=), aritmatika (%, /), dan logika (&&, ||, !).   Menentukan nilai kebenaran akhir dari setiap ekspresi kontrol dalam bentuk true atau false

### 2. TRACING EVALUASI PERNYATAAN KONDISI

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	// Blok 1
	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	} // result = 15

	// Blok 2
	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	} // result = 30

	// Blok 3
	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	} // result = 35

	// Blok 4
	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	} // result = 25

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/evaluasi_pernyataan_kondisi/output%20(9).png)

#### Deskripsi
Melakukan penelusuran manual (tracing) alur eksekusi program yang memiliki beberapa blok if-else bertingkat.   Menentukan nilai akhir variabel result serta teks luaran (output) yang dicetak program.   Menganalisis dan menjelaskan dampak evaluasi dari 4 kondisi logika spesifik terhadap perubahan variabel result


### 3. MENENTUKAN JUMLAH HARI DALAM SEBULAN BERDASARKAN TAHUNDAN BULAN

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	// Menerima input tahun dan nama bulan
	fmt.Scan(&tahun, &bulan)

	// Logika pengecekan tahun kabisat
	isKabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	// Logika penentuan jumlah hari berdasarkan bulan
	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		fmt.Println(31)
	case "Apr", "Jun", "Sep", "Nov":
		fmt.Println(30)
	case "Feb":
		if isKabisat {
			fmt.Println(29)
		} else {
			fmt.Println(28)
		}
	default:
		// Menampilkan "-" jika nama bulan tidak sesuai format/tidak valid
		fmt.Println("-")
	}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/Menentukan_jumlah_hari_dalam_sebulan/output%20(10).png)

#### Deskripsi
Membuat program Go yang menerima masukan tahun dan 3 huruf nama bulan.   Menerapkan logika perhitungan tahun kabisat khusus untuk penentuan jumlah hari pada bulan Februari.   Memvalidasi format masukan nama bulan secara case-sensitive (harus diawali huruf kapital, contoh: "Jan").   Menampilkan total hari jika masukan valid atau mencetak tanda - jika format bulan tidak sesuai.

### 4. SWITCH CASE

```go
package main

import "fmt"

func main() {
	var nilai string

	fmt.Print("Masukkan nilai huruf (A/B/C/D/E): ")
	fmt.Scan(&nilai)

	// Penerapan switch case untuk menentukan predikat
	switch nilai {
	case "A":
		fmt.Println("Predikat: Sangat Memuaskan")
	case "B":
		fmt.Println("Predikat: Memuaskan")
	case "C":
		fmt.Println("Predikat: Cukup")
	case "D":
		fmt.Println("Predikat: Kurang")
	case "E":
		fmt.Println("Predikat: Tidak Lulus")
	default:
		fmt.Println("Nilai yang dimasukkan tidak valid!")
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/tp/switch_case/output%2011%20.png)

#### Deskripsi
Membuat satu contoh program sederhana dalam bahasa Go yang menerapkan percabangan switch case.   Melatih penggunaan switch case sebagai struktur kontrol alternatif dari percabangan if-else
## Kesimpulan
Rangkaian soal tersebut bertujuan untuk menguji pemahaman mengenai struktur kontrol percabangan (if-else dan switch case), evaluasi ekspresi logika, serta kemampuan penelusuran alur eksekusi program dalam bahasa Go.