package main
import "fmt"

func main() {
	var x, digit, sebelumnya int
	var cantik bool
	fmt.Scan(&x)

	cantik = true
	sebelumnya = x % 10
	x = x / 10

	for x > 0 {
		digit = x % 10

		if digit > sebelumnya {
			cantik = false
		}

		sebelumnya = digit
		x = x / 10
	}

	fmt.Println(cantik)

}