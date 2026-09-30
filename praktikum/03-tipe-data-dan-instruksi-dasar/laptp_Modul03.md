# <h1 align="center">Tugas Pendahuluan Modul  - VARIABEL DAN OPERATOR </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {
	var y, x int

    // Membaca input y (jumlah kue) dan x (jumlah anggota keluarga)
    fmt.Scan(&y, &x)

    // Menghitung sisa kue menggunakan operator modulo (%)
    sisa := y % x

    // Mencetak hasil sisa kue
    fmt.Println(sisa)
}
```

##### Output
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa/output%20(4).png)



#### Deskripsi
Keluaran berupa jumlah kue yang tersisa setelah kue dibagikan sama rata kepada setiap anggota keluarga. 

### 2. KONVERSI MIL KE KILOMETER

```go
package main

import "fmt"

func main() {
	var mil float64

    // Membaca masukan nilai mil bertipe float64
    fmt.Scan(&mil)

    // Menghitung konversi ke kilometer
    km := mil * 1.6

    // Mencetak hasil konversi dengan format 1 angka di belakang koma
    fmt.Printf("%.1f\n", km)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi/output.png)

#### Deskripsi
Program ini mampu membaca nilai dalam mil berupa bilangan desimal (float64) dari input pengguna, mengonversinya ke kilometer, lalu menampilkan hasilnya.


### 3. BOOL GO

```go
package main

import "fmt"

func main() {
	var b bool

    // Membaca masukan bertipe boolean (true / false)
    fmt.Scan(&b)

    // Mencetak kembali nilai boolean
    fmt.Println(b)
}
```
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/bool/output%20(5).png)

#### Deskripsi
Program untuk mengkonversi nilai jarak mil ke kilometer dengan operasi mil * kilometer dan menggunakan %.1f agar hasil desimal di belakang kome hanya 1 angka.



## Kesimpulan
Variabel dan operator merupakan dua elemen dasar dalam pemrograman yang bekerja sama untuk mengolah, menyimpan, dan memanipulasi data di dalam program. Variabel berperan sebagai wadah di dalam memori komputer yang menampung nilai tertentu agar dapat diakses atau diperbarui kapan saja. Setiap variabel memiliki nama sebagai identitas uniknya serta tipe data spesifik yang menentukan jenis nilai yang disimpannya, seperti angka, teks, maupun nilai kebenaran (boolean).