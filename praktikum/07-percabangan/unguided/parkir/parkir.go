package main

import "fmt"

func main() {
	var jenis string
	var jam int
	var tarif int
	var denda int = 0

	fmt.Print("Masukkan jenis kendaraan: ")
	fmt.Scan(&jenis)

	fmt.Print("Masukkan lama parkir (jam): ")
	fmt.Scan(&jam)

	if jenis == "Mobil" || jenis == "mobil" {
		if jam <= 1 {
			tarif = 5000
		} else {
			tarif = 5000 + (jam-1)*3000
		}

		if jam > 24 {
			denda = 50000
		}
	} else if jenis == "Motor" || jenis == "motor" {
		if jam <= 1 {
			tarif = 3000
		} else {
			tarif = 3000 + (jam-1)*2000
		}

		if jam > 24 {
			denda = 30000
		}
	} else {
		fmt.Println("Jenis kendaraan tidak valid")
		return
	}

	fmt.Println(tarif + denda)
}