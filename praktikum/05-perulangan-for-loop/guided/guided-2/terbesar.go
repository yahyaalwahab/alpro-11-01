package main

import "fmt"

func main() {
	var n, bilangan, terbesar int
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&bilangan)

		if i == 1 || bilangan > terbesar {
			terbesar = bilangan
		}
	}

	fmt.Println(terbesar)
}