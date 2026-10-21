package main

import "fmt"

func main(){
	var poin, total int
	fmt.Scan(&poin)

	for poin != 0 {
		total += poin
		fmt.Scan(&poin)
	}
	fmt.Println(total)
}