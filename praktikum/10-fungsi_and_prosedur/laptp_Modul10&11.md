# <h1 align="center">Tugas Pendahuluan Modul 10  & 11  - FUNGSI & PROSEDUR </h1>
<p align="center">Fathan Al Akbar - 109092630006</p>

### 1. RAPORT.GO

```go
package main

import (
	"fmt"
)

const NMAX int = 999

type rapot struct {
	matkul string
	nilai  int
}

type tabRapot [NMAX]rapot

func nilaiHuruf(x int) string {
	if x > 80 {
		return "A"
	} else if x > 70 && x <= 80 {
		return "AB"
	} else if x > 65 && x <= 70 {
		return "B"
	} else if x > 60 && x <= 65 {
		return "BC"
	} else if x > 50 && x <= 60 {
		return "C"
	} else if x > 40 && x <= 50 {
		return "D"
	} else {
		return "E"
	}
}

func rataRata(T tabRapot, n int) float64 {
	total := 0
	for i := 0; i < n; i++ {
		total += T[i].nilai
	}
	return float64(total) / float64(n)
}

func main() {
	var n int
	var T tabRapot

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].matkul, &T[i].nilai)
	}

	for i := 0; i < n; i++ {
		huruf := nilaiHuruf(T[i].nilai)
		fmt.Printf("%s %d %s\n", T[i].matkul, T[i].nilai, huruf)
	}

	avg := rataRata(T, n)
	fmt.Printf("Rata-rata nilai: %.2f\n", avg)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/tp/fungsi-10/raport/output%2025.png)


#### Deskripsi
raport.go meminta pembuatan program untuk mengonversi nilai angka mata kuliah menjadi nilai huruf dan menghitung rata-rata nilai raport menggunakan dua fungsi terpisah.   

### 2. INPUTNILAI.GO

 
```go
package main

import (
	"fmt"
)

const NMAX int = 100

type student struct {
	nim   string
	name  string
	grade float64
}

type students [NMAX]student

func cariNIM(T students, n int, nim string) int {
	for i := 0; i < n; i++ {
		if T[i].nim == nim {
			return i
		}
	}
	return -1
}

func rataRata(T students, n int) float64 {
	if n == 0 {
		return 0.0
	}
	total := 0.0
	for i := 0; i < n; i++ {
		total += T[i].grade
	}
	return total / float64(n)
}

func main() {
	var T students
	var n int = 0
	var nim, name string
	var grade float64

	for {
		fmt.Scan(&nim)
		if nim == "STOP" {
			break
		}
		fmt.Scan(&name, &grade)

		if cariNIM(T, n, nim) == -1 {
			T[n].nim = nim
			T[n].name = name
			T[n].grade = grade
			n++
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %.1f\n", T[i].nim, T[i].name, T[i].grade)
	}

	avg := rataRata(T, n)
	fmt.Printf("Rata-rata nilai: %.2f\n", avg)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/tp/fungsi-10/inputnilai/output%2026.png)


#### Deskripsi
inputnilai.go meminta pembuatan program untuk menghapus data nilai mahasiswa yang terinput ganda (duplikat) menggunakan fungsi pencarian NIM, lalu menampilkan data bersih beserta rata-rata nilainya.

### 3. SEKOLAH.GO

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1001

type tabMurid [NMAX]string

func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
	*ketemu = false
	*posisi = 0
	for i := 0; i < n; i++ {
		if T[i] == nama {
			*ketemu = true
			*posisi = i + 1 // posisi dimulai dari 1
			break
		}
	}
}

func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
	*jumlah = 0
	for i := 0; i < n; i++ {
		if len(T[i]) > 0 && T[i][0] == h {
			(*jumlah)++
		}
	}
}

func main() {
	var n int
	var T tabMurid
	var cariNama string

	reader := bufio.NewReader(os.Stdin)
	if _, err := fmt.Scan(&n); err != nil {
		return
	}

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		T[i] = strings.TrimSpace(line)

		if T[i] == "" {
			i--
		}
	}

	line, _ := reader.ReadString('\n')
	cariNama = strings.TrimSpace(line)
	if cariNama == "" {
		fmt.Scan(&cariNama)
	}

	var posisi int
	var ketemu bool
	var jumlahHuruf int

	cariMurid(T, n, cariNama, &posisi, &ketemu)

	if ketemu {
		hurufPertama := cariNama[0]
		hitungHuruf(T, n, hurufPertama, &jumlahHuruf)

		fmt.Printf("Murid terdaftar dan berada di urutan absen ke-%d\n", posisi)
		fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", hurufPertama, jumlahHuruf)
	} else {
		var hurufPertama byte
		if len(cariNama) > 0 {
			hurufPertama = cariNama[0]
			hitungHuruf(T, n, hurufPertama, &jumlahHuruf)
		}

		fmt.Println("Murid tidak terdaftar")
		fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", hurufPertama, jumlahHuruf)
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/tp/prosedur-11/sekolah/output%2027.png)

#### Deskripsi
sekolah.go meminta pembuatan program menggunakan prosedur untuk mencari status pendaftaran serta nomor urut absen seorang murid, sekaligus menghitung jumlah murid yang memiliki huruf awal nama yang sama dengan nama yang dicari.


### 4. PENDUDUK.GO


```go
package main

import (
	"fmt"
)

const NMAX int = 1000

type penduduk struct {
	nama  string
	tahun int
	kota  string
}

type tabPenduduk [NMAX]penduduk

func cariTahun(T tabPenduduk, n int, th int, idx *int) {
	*idx = -1
	for i := 0; i < n; i++ {
		if T[i].tahun == th {
			*idx = i
			break
		}
	}
}

func tertuaTermuda(T tabPenduduk, n int, tua *penduduk, muda *penduduk) {
	*tua = T[0]
	*muda = T[0]
	for i := 1; i < n; i++ {
		if T[i].tahun < tua.tahun {
			*tua = T[i]
		}
		if T[i].tahun > muda.tahun {
			*muda = T[i]
		}
	}
}

func main() {
	var n int
	var T tabPenduduk
	var cariTh int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
	}

	fmt.Scan(&cariTh)

	var idx int
	cariTahun(T, n, cariTh, &idx)

	if idx != -1 {
		fmt.Printf("%s %d %s ditemukan di index ke-%d\n", T[idx].nama, T[idx].tahun, T[idx].kota, idx)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

	var tua, muda penduduk
	tertuaTermuda(T, n, &tua, &muda)

	fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
	fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/alpro-11-01/praktikum/10-fungsi_and_prosedur/tp/prosedur-11/penduduk/output%2028.png)

#### Deskripsi
penduduk.go meminta pembuatan program untuk mengelola data penduduk (nama, tahun lahir, dan kota asal) menggunakan prosedur guna mencari data penduduk pertama berdasarkan tahun lahir tertentu serta menampilkan identitas penduduk tertua dan termuda dari kumpulan data tersebut.
## Kesimpulan
Dalam pemrograman, **fungsi** dan **prosedur** merupakan subprogram modular yang dirancang untuk memecah kode menjadi bagian-bagian kecil yang terstruktur dan dapat digunakan kembali (*reusable*) guna mempermudah pengelolaan program. Perbedaan mendasar di antara keduanya terletak pada nilai kembalian (*return value*), di mana **fungsi** bertugas mengolah data dan wajib mengembalikan suatu nilai hasil perhitungan menggunakan kata kunci `return`, sedangkan **prosedur** hanya menjalankan serangkaian instruksi atau aksi tertentu (seperti mencari, mengurutkan, atau mencetak data) tanpa mengembalikan nilai secara langsung. Dengan menerapkan keduanya secara tepat, kode program menjadi jauh lebih efisien, rapi, mudah dibaca, serta memudahkan proses pelacakan kesalahan (debugging).