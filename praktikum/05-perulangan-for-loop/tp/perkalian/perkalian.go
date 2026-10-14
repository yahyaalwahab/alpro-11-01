package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a, &b)

	for i := 1; i <= b; i++ {
		fmt.Println(a, "x", i, "=", a*i)
	}
}