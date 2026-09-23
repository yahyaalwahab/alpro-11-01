package main
import "fmt"

func main(){
	var nama string
	var skormatematika, skorbahasainggris int

	// Membaca input
	fmt.Scan(&nama, &skorbahasainggris, &skormatematika)

	// Menghitung total & rata-rata 
	total := skormatematika + skorbahasainggris
	rataRata := total / 2

	// Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
}