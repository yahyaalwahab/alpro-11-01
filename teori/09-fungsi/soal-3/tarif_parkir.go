package main

import "fmt"

func tarifParkir(jam int) int {
	var total int

	if jam <= 2 {
		total = 5000
	} else {
		total = 5000 + (jam-2)*3000
	}

	return total
}

func main() {
	var jam int

	fmt.Scan(&jam)

	hasil := tarifParkir(jam)

	fmt.Println(hasil)
}