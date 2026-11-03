package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Print("genap, ")
	} else {
		fmt.Print("ganjil, ")
	}

	if n > 0 {
		fmt.Println("positif")
	} else if n < 0 {
		fmt.Println("negatif")
	} else {
		fmt.Println("nol")
	}
}