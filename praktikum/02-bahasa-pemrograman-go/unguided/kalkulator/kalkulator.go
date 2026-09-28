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