package main

import "fmt"

func konversi(km, m, cm int) int {
	hasil := (km * 100000) + (m * 100) + cm
	return hasil
}

func main() {
	var km, m, cm int

	fmt.Scan(&km, &m, &cm)

	hasil := konversi(km, m, cm)

	fmt.Println("Hasil konversi =", hasil, "cm")
}