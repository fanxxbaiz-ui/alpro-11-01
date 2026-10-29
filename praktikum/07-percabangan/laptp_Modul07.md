# <h1 align="center">Tugas Pendahuluan Modul 7  - PERCABANGAN </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. Jelaskan perbedaan antara if, if-else, else if, dan switch pada bahasa pemrograman Go! Sebutkan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

JAWABAN =
 - if: Digunakan untuk mengevaluasi satu kondisi boolean. Jika kondisi bernilai benar (true), blok kode di dalamnya dieksekusi. Catatan khusus di Go: kondisi tidak perlu dibungkus dalam tanda kurung (), tetapi tanda kurung kurawal {} wajib ada meskipun hanya berisi satu baris kode.

- if-else: Perluasan dari if yang menyediakan jalur alternatif. Jika kondisi pada if bernilai salah (false), maka blok kode pada else yang akan dieksekusi.

- else if: Digunakan untuk memeriksa beberapa kondisi berbeda secara berurutan setelah kondisi if pertama gagal.

- switch: Struktur kendali alternatif yang lebih bersih dan mudah dibaca dibandingkan menggunakan rantai else if yang terlalu panjang. switch mencocokkan nilai sebuah variabel atau ekspresi terhadap beberapa pilihan (case).

2. Dua Perbedaan switch di Go Dibandingkan Java
Terdapat beberapa perbedaan mendasar cara kerja switch di Go jika dibandingkan dengan Java:

A) Perilaku Fall-Through (Automatic Break)

Java: Secara default, jika suatu case terpenuhi, eksekusi akan terus berlanjut ke case di bawahnya (fall-through) kecuali Anda secara eksplisit menambahkan pernyataan break di akhir setiap case.

Go: Bersifat sebaliknya. Di Go, setiap case memiliki break otomatis di akhir bloknya. Eksekusi tidak akan otomatis turun ke case berikutnya. Jika Anda sengaja ingin meneruskan eksekusi ke case selanjutnya, Anda harus menuliskan keyword fallthrough secara eksplisit.

B) Dukungan Switch Tanpa Ekspresi (Expressionless Switch)

Java: Pernyataan switch selalu membutuhkan variabel atau ekspresi target yang nilainya akan dicocokkan dengan setiap case.

Go: switch dapat ditulis tanpa ekspresi sama sekali (sering disebut tagless switch). Dalam hal ini, switch bertindak sebagai pengganti rantai if-else if yang lebih bersih, di mana setiap case berisi kondisi boolean independen (contoh: switch { case x > 0: ... case y < 0: ... }).



#### Deskripsi
soal ini menuntut penjelasan yang mendalam mengenai perbedaan antara if, if-else, else if, dan switch pada bahasa pemrograman Go dan juga dua perbedaan switch pada Go dibandingkan dengan switch pada Java. 

### 2. BILANGAN.GO 

 
```go
package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan == 0 {
		fmt.Println("Nol")
	} else {
		var jenis, tanda string
		
		if bilangan%2 == 0 {
			jenis = "Genap"
		} else {
			jenis = "Ganjil"
		}

		if bilangan > 0 {
			tanda = "Positif"
		} else {
			tanda = "Negatif"
		}

		fmt.Printf("%s, %s\n", jenis, tanda)
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/tp/soal-2/output%20(44).png)


#### Deskripsi
(bilangan.go) adalah program dalam bahasa Go yang dirancang untuk menerima masukan sebuah bilangan bulat dan menentukan kondisinya secara komprehensif, yaitu apakah bilangan tersebut bernilai nol, atau memiliki kombinasi sifat genap/ganjil serta positif/negatif menggunakan struktur percabangan if-else

### 3. HARI.GO

```go
package main

import "fmt"

func main() {
	var hari int
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu (akhir pekan)")
	case 7:
		fmt.Println("Minggu (akhir pekan)")
	default:
		fmt.Println("Nomor hari tidak valid")
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/tp/soal-3/output%20(45).png)

#### Deskripsi
(hari.go) adalah program dalam bahasa Go yang menggunakan struktur kendali switch untuk menerima masukan nomor hari dari 1 hingga 7, lalu menampilkan nama hari yang sesuai,termasuk tambahan keterangan akhir pekan untuk hari Sabtu dan Minggu,serta mencetak pesan kesalahan apabila nomor masukan berada di luar rentang 1 hingga 7.


### 4. TERBESAR.GO


```go
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	var terbesar int

	if a >= b && a >= c {
		terbesar = a
	} else if b >= a && b >= c {
		terbesar = b
	} else {
		terbesar = c
	}

	fmt.Println(terbesar)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/07-percabangan/tp/soal-4/output%20(46).png)

#### Deskripsi
(terbesar.go) adalah program dalam bahasa Go yang dirancang untuk menentukan nilai terbesar dari tiga bilangan bulat yang diinputkan menggunakan struktur percabangan if-else secara berjenjang tanpa menggunakan perulangan ataupun fungsi bawaan max.
## Kesimpulan
bahasa Go ini bertujuan untuk menguji dan melatih penguasaan struktur kawalan percabangan, yang merangkumi penentuan sifat nombor melalui if-else, pemetaan nama hari menggunakan switch beserta syarat tambahan, serta pencarian nilai terbesar daripada tiga nombor bulat secara logik berjenjang tanpa menggunakan fungsi bawaan max.