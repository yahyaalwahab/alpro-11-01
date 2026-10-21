package main

import "fmt"

func main(){
	var n, i, total int
	fmt.Scan(&n)

	i = 1
	for i <= n {
		total += i
		i++
	}
	fmt.Println(total)
}