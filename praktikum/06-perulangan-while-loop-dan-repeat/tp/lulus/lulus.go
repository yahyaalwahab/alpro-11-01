package main 

import "fmt"

func main(){
	var nilai int

	for {
		fmt.Print("masukkan nilai: ")
		fmt.Scan(&nilai)

		if nilai == -1 {
			fmt.Println("program selesai")
			break
		}

		if nilai >= 60 {
			fmt.Println("anda lulu ujian")
		} else {
			fmt.Println("anda tidak lulus ujian")
		}
	}
}