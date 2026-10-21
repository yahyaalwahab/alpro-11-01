package main

import "fmt"

func main(){
	var n, m, x, y, jumlah int
	fmt.Scan(&n, &m, &x, &y)

	jumlah = 0

	for n >= x && m >= y {
		n = n - x
		m = m - y
		jumlah++
	}
	fmt.Println(jumlah)
}