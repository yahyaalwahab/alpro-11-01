package main

import "fmt"

func main(){
	var x int

	fmt.Print("masukkan nominal:")
	fmt.Scan(&x)

	var SepuluhRibuan int = x / 10000
	var sisa int = x % 10000

	var LimaRibuan int = sisa /5000
	sisa = sisa % 5000

	var seribuan int = sisa / 1000

	fmt. Println(SepuluhRibuan, LimaRibuan, seribuan)
}