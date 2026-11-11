package main
import "fmt"

func main(){
	var kendaraan string
	var M, K, T int

	for {
		fmt.Scan(&kendaraan)
		
		if kendaraan == "M" {
			M++
		} else if kendaraan == "K" {
			K++
		} else if kendaraan == "T" {
			T++
		} else {
			break
		}

	}

	fmt.Println("Motor =", M)
	fmt.Println("Mobil =", K)
	fmt.Println("Truk =", T)
}