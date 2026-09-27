package main
import "fmt"

func main(){
	var celcius float64

	//Menerima suhu dalam celcius
	fmt.Scanln(&celcius)

	//Menghitung suhu dengan rumus
	reamur := celcius * 4.0 / 5.0
	fahrenheit := celcius * 9.0 / 5.0 + 32.0
	kelvin := celcius + 273.15

	//Mencetak tiga nilai keluaran di pisahkan oleh spasi
	fmt.Println(reamur, fahrenheit, kelvin)
}