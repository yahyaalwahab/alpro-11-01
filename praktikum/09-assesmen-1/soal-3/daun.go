package main
import "fmt"

func main(){
	var n, lebar, terkecil int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&lebar)

		if i == 1 || lebar < terkecil {
			terkecil = lebar
		}
	}

	fmt.Println("daun paling sempit:", terkecil)
}