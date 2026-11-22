package main
import "fmt"

func faktorial(n int) int {
	hasil := 1

	for i := 1; i <= n; i++ {
		hasil *= i
	}

	return hasil
}

func kombinasi(x, y int) (int, int, int) {
	fx := faktorial(x)
	fy := faktorial(y)

	var hasil int

	if x >= y {
		hasil = fx / (fy * faktorial(x-y))
	} else {
		hasil = fy / (fx * faktorial(y-x))
	}

	return fx, fy, hasil
}

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	a, b, c := kombinasi(x, y)

	fmt.Println(a, b, c)
}