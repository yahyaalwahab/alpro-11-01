package main

import "fmt"

func main(){
	var n, nilai int
	var nama string
	var terendah, kedua int
	var namaTerendah, namaKedua string

	fmt.Scan(&n)

	i := 1
	for i <= n {
		fmt.Scan(&nama, &nilai)

		if i == 1 {
			terendah = nilai
			kedua = nilai
			namaTerendah = nama
			namaKedua = nama
		} else if nilai < terendah {
			kedua = terendah
			namaKedua = namaTerendah
			terendah = nilai
			namaTerendah = nama
		} else if i == 2 || nilai < kedua {
			kedua = nilai
			namaKedua = nama
		}

		i++
	}

	fmt.Println(namaTerendah)
	fmt.Println(namaKedua)
	
}