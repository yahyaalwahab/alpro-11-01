package main

import "fmt"

func deretKuadrat(n int) int {
	var total int

	for i := 1; i <= n; i++ {
		total = total + (i * i)
	}

	return total
}

func main() {
	var n int

	fmt.Scan(&n)

	hasil := deretKuadrat(n)

	fmt.Println(hasil)
}