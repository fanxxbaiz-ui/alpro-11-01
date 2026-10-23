# <h1 align="center">Laporan Praktikum Modul 6 - Perulangan While   </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

## Dasar Teori

### A. PERULANGAN FOR
Perulangan while (atau while loop) adalah struktur kontrol dalam pemrograman yang digunakan untuk menjalankan sekumpulan blok kode secara berulang selama kondisi tertentu bernilai benar (true).

Berbeda dengan perulangan for yang biasanya digunakan ketika jumlah perulangan sudah diketahui sejak awal, perulangan while ideal digunakan ketika jumlah perulangan belum pasti dan bergantung pada suatu kondisi yang berubah secara dinamis saat program berjalan.
### B. while-loop dan repeat-until
#### 1. while-loop
- Evaluasi kondisi yang dilakukan diawal(sebelum masuk)
- kondisi di awal FALSE,loop tidak akan pernah dieksekusi



#### 2. Repeat-until

- kerjakan dahulu, evaluasi kondisi diakhir

- Blok intruksi pasti dieksekusi minimal satu kali


<!-- Tambahkan poin A, B, C, ... atau sub-topik 1, 2, 3, ... sesuai kebutuhan modul -->

## Guided

### 1. TERKECIL.GO

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n*n <= batas {
        n++
    }

    fmt.Println(n)
}

```
##### Output
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-1/output%20(29).png)

#### Deskripsi
Program terkecil.go dirancang untuk memproses sekumpulan data numerik guna mengidentifikasi dan menampilkan nilai terkecil dari deretan angka yang diinputkan atau tersimpan dalam array/slice.
### 2.TERBESAR.GO

```go
package main

import "fmt"

func main() {
    var batas int
    fmt.Scan(&batas)

    n := 1
    for n*n < batas {
        n++
    }

    fmt.Println(n - 1)
}

```
##### Output
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-2/output%20(30).png)


#### Deskripsi
Program terbesar.go bertugas untuk membandingkan deretan elemen angka dan menemukan nilai maksimum atau terbesar dari kumpulan data yang diproses.

### 3. VOKAL.GO

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	kalimat, _ := reader.ReadString('\n')

	i := 0
	jumlahVokal := 0

	for i < len(kalimat) {
		ch := strings.ToLower(string(kalimat[i]))
		if ch == "a" || ch == "i" || ch == "u" || ch == "e" || ch == "o" {
			jumlahVokal++
		}
		i++
	}

	fmt.Println(jumlahVokal)
}
```
##### Output
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/guided/guided-3/output%20(31).png)


#### Deskripsi
Program vokal.go berfungsi untuk menganalisis karakter dalam suatu teks guna memfilter, memeriksa, atau menghitung jumlah huruf vokal (a, e, i, o, u) dari kalimat atau kata yang dimasukkan
## Unguided

### 1. ATM.GO

```go
package main

import "fmt"

func main() {
	var password string
	var percobaan int = 0
	var loginBerhasil bool = false
	var saldo int = 10000000
	var pilihan int = 0

	// Perulangan while untuk proses login (maksimal 3 kesempatan dengan password "1234")[cite: 1, 2]
	for percobaan < 3 && !loginBerhasil {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		if password == "1234" {
			loginBerhasil = true
			fmt.Println("Login berhasil!")
		} else {
			percobaan++
			if percobaan < 3 {
				fmt.Println("Password salah. Silakan coba lagi.")
			}
		}
	}

	// Jika salah memasukkan password sebanyak 3 kali berturut-turut[cite: 1]
	if !loginBerhasil {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	// Perulangan while untuk Menu ATM (berulang selama pengguna belum memilih keluar)[cite: 2]
	for pilihan != 3 {
		fmt.Println("Menu ATM:")
		fmt.Println("1. Lihat Saldo")
		fmt.Println("2. Tarik Uang")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			fmt.Printf("Saldo Anda: Rp %d\n", saldo)
		} else if pilihan == 2 {
			var tarik int
			fmt.Print("Masukkan jumlah uang yang ingin ditarik: Rp ")
			fmt.Scan(&tarik)

			// Mengurangi saldo jika mencukupi, jika tidak tampilkan pesan[cite: 2]
			if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan menu tidak valid.")
		}
	}
}

```

##### Output
![Screenshot Output Unguided](/praktikum/06-perulangan-while-loop-dan-repeat/unguided/output%20(37).png)

#### Deskripsi
Program simulasi sistem ATM berbasis bahasa Go (atm.go) ini meminta pengguna memasukkan password "1234" dengan batas maksimal 3 kesempatan dan inisialisasi saldo awal sebesar Rp 10.000.000. Jika salah memasukkan password sebanyak 3 kali berturut-turut, ATM akan terblokir dan program berakhir, tetapi jika benar, program akan menampilkan menu ATM yang terus berulang hingga pengguna memilih keluar. Spesifikasi menu tersebut mencakup pilihan untuk melihat saldo, menarik uang dengan pengecekan ketersediaan saldo, serta keluar dari sistem.


## Kesimpulan
Perulangan `while-loop` dan `repeat-until` merupakan struktur kontrol perulangan dalam pemrograman yang dibedakan berdasarkan waktu evaluasi kondisinya. Pada `while-loop` (perulangan di awal), kondisi diperiksa sebelum blok kode dijalankan sehingga perulangan bisa tidak dieksekusi sama sekali jika kondisi bernilai salah sejak awal. Sebaliknya, pada `repeat-until` atau do-while (perulangan di akhir), blok kode akan dieksekusi setidaknya satu kali terlebih dahulu sebelum kondisinya diperiksa di akhir perulangan.
## Referensi
1. Materi Modul 6 PERULANGAN WHILE
