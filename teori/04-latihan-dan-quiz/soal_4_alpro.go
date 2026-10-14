package main
import "fmt"

func main(){
	var n, m, x, y int

	fmt.Scan(&n,&m, &x, &y)

	cangkir := 0
	for n >= x && m >= m {
		n = n - x
		m = m - y
		cangkir++
		
	fmt.Println(cangkir)	
	}
}