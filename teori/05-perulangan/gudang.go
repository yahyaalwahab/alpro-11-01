package main

import "fmt"

func main() {
	var k, v, total int
	fmt.Scan(&k)

	total = 0

	for i := 1; i <= 5; i++ {
		fmt.Scan(&v)
		total += v

		fmt.Println(total >= k)
	}
}