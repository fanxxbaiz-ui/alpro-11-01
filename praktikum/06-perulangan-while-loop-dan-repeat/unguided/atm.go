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
