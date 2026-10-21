package main

import "fmt"

func main() {
	var x, kuadrat, digit, total int
	fmt.Scan(&x)

	total = 0

	for x > 0 {
		digit = x % 10
		kuadrat = digit * digit

		fmt.Print(kuadrat, " ")
		total += kuadrat
		x = x / 10
	}

	fmt.Println()
	fmt.Println(total)
}