package main
import "fmt"

func main(){
	var a, b int

	fmt.Scan(&a, &b)
	tambah := a + b
	kurang := a - b
	kali := a * b
	bagi := a / b
	sisa := a % b

	fmt.Println("Hasil tambah:", tambah)
	fmt.Println("Hasil kurang:", kurang)
	fmt.Println("Hasil kali:", kali)
	fmt.Println("Hasil bagi:", bagi)
	fmt.Println("Sisa:", sisa)
}
