package main

import "fmt"

func main(){
	var hari int
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("senin")
	case 2:
		fmt.Println("selasa")
	case 3:
		fmt.Println("rabu")
	case 4:
		fmt.Println("kamis")
	case 5:
		fmt.Println("jum'at")
	case 6:
		fmt.Println("sabtu")
	case 7:
		fmt.Println("minggu")
	default:
		fmt.Println("tidak valid")
	}
}