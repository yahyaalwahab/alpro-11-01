package main 

import "fmt"

func main(){
	var n int
	var biaya float64 = 10000000

	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		biaya = biaya * 1.06
	}

	fmt.Printf("%.0f\n", biaya)
}