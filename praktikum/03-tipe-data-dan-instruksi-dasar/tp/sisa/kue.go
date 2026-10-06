package main

import "fmt"

func main(){

	var y, x int

	fmt.Print("masukkan jumlah kue:")
	fmt.Scan(&y)

	fmt.Print("masukkan jumlah anggota keluarga:")
	fmt.Scan(&x)

	fmt.Print("jumlah kue yang tersisa:")
	fmt.Println(y % x)
}