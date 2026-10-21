package main

import "fmt"

func main() {
	var password string
	var percobaan int
	var saldo int = 10000000
	var pilihan int

	// Proses login
	percobaan = 0

	for password != "1234" && percobaan < 3 {
		fmt.Print("Masukkan password ATM: ")
		fmt.Scan(&password)

		percobaan++

		if password != "1234" {
			fmt.Println("Password salah. Silakan coba lagi.")
		}
	}

	if password != "1234" {
		fmt.Println("ATM terblokir. Program selesai.")
		return
	}

	fmt.Println("Login berhasil!")

	// Menu ATM
	pilihan = 0

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

			if tarik <= 0 {
				fmt.Println("Jumlah penarikan tidak valid.")
			} else if tarik <= saldo {
				saldo -= tarik
				fmt.Printf("Anda berhasil menarik uang sebesar Rp %d\n", tarik)
				fmt.Printf("Saldo tersisa: Rp %d\n", saldo)
			} else {
				fmt.Println("Saldo tidak mencukupi.")
			}
		} else if pilihan == 3 {
			fmt.Println("Keluar dari sistem. Terima kasih!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}