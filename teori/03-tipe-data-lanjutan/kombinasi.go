package main
import "fmt"

func main(){
	var p, q int
	fmt.Scan(&p, &q)

	genapSalahSatu := (p%2 == 0) || (q%2 == 0)
	ganjilKeduanya := (p%2 != 0) && (q%2 != 0)
	tidakSama := !(p == q)

	fmt.Println(genapSalahSatu, ganjilKeduanya, tidakSama)
}