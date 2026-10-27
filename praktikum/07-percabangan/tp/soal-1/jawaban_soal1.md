1. Penjelasan Struktur Percabangan di Go
if: Digunakan untuk mengevaluasi satu kondisi boolean. Jika kondisi bernilai benar (true), blok kode di dalamnya dieksekusi. Catatan khusus di Go: kondisi tidak perlu dibungkus dalam tanda kurung (), tetapi tanda kurung kurawal {} wajib ada meskipun hanya berisi satu baris kode.

if-else: Perluasan dari if yang menyediakan jalur alternatif. Jika kondisi pada if bernilai salah (false), maka blok kode pada else yang akan dieksekusi.

else if: Digunakan untuk memeriksa beberapa kondisi berbeda secara berurutan setelah kondisi if pertama gagal.

switch: Struktur kendali alternatif yang lebih bersih dan mudah dibaca dibandingkan menggunakan rantai else if yang terlalu panjang. switch mencocokkan nilai sebuah variabel atau ekspresi terhadap beberapa pilihan (case).

2. Dua Perbedaan switch di Go Dibandingkan Java
Terdapat beberapa perbedaan mendasar cara kerja switch di Go jika dibandingkan dengan Java:

A) Perilaku Fall-Through (Automatic Break)

Java: Secara default, jika suatu case terpenuhi, eksekusi akan terus berlanjut ke case di bawahnya (fall-through) kecuali Anda secara eksplisit menambahkan pernyataan break di akhir setiap case.

Go: Bersifat sebaliknya. Di Go, setiap case memiliki break otomatis di akhir bloknya. Eksekusi tidak akan otomatis turun ke case berikutnya. Jika Anda sengaja ingin meneruskan eksekusi ke case selanjutnya, Anda harus menuliskan keyword fallthrough secara eksplisit.

B) Dukungan Switch Tanpa Ekspresi (Expressionless Switch)

Java: Pernyataan switch selalu membutuhkan variabel atau ekspresi target yang nilainya akan dicocokkan dengan setiap case.

Go: switch dapat ditulis tanpa ekspresi sama sekali (sering disebut tagless switch). Dalam hal ini, switch bertindak sebagai pengganti rantai if-else if yang lebih bersih, di mana setiap case berisi kondisi boolean independen (contoh: switch { case x > 0: ... case y < 0: ... }).