package main
import "fmt"

func main(){
	var nilai, sebelumnya int
	var tertinggi, terendah int
	var total, jumlah int
	var rataRata float64

	fmt.Scan(&nilai)

	for {
		fmt.Scan(&nilai)

		if nilai == sebelumnya {
			break
		}

		if nilai > tertinggi {
			tertinggi = nilai
		}

		if nilai < terendah { 
			terendah = nilai
		}

		total += nilai
		jumlah++
		sebelumnya = nilai
	}

	rataRata = float64(total) / float64(jumlah)

	fmt.Printf("%d %d %.2f\n", tertinggi, terendah, rataRata)
}