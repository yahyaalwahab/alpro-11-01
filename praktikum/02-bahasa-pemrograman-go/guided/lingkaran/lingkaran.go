package main
import "fmt"

func main(){
	var pi = 3.14
	var r float64
	
	fmt.Scanln(&r)

	var luas float64 = pi * r * r
	
	fmt.Println(pi, r, luas)
}