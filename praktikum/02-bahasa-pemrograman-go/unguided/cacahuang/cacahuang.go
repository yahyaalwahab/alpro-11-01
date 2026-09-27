package main
import "fmt"

func main(){
	var nominal int

	//Membaca nominal masukan uang
	fmt.Scanln(&nominal)

	//Menghitung lembar sepuluh ribu
	sepuluhRibu := nominal / 10000
	sisa := nominal % 10000

	//Menghitung lembar lima ribu dari sisa sebelumnya
	limaRibu := sisa / 5000
	sisa = sisa % 5000

	//Menghitung lembar seribu dari sisa sebelumnya
	seribu := sisa / 1000

	//Tampilkan hasil berupa tiga bilangan bulat dipisahkan spasi
	fmt.Printf("%d %d %d\n", sepuluhRibu, limaRibu, seribu)
}