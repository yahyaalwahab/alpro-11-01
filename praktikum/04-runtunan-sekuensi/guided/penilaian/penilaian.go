package main
import (
	"bufio" 
	//digunakan untuk membaca masukan string yang memeiliki spasi (seperti nama lengkap)
	"fmt"
	"os"
	//menyediakan akses ke sistem operasi, dalam hal ini os.Stdin yang mempresentasikan keyboard
) 

func main(){
	//pendefinisian variabel secara eksplisit beserta tipe datanya
	var pilihan int
	var nama string
	var nilai float64
	var grade string

	//menampilkan cetakan Menu ke layar
	fmt.Println("========== Menu ==========")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("pilih opsi:")

	//membaca masukan pilihan (angka).
	//kita menggunakan Scanln agar saat user menekan 'Enter, karakter enter tersebut
	//ikut di olah/dibersihkan dan tidak melompat (mengganggu) masukan nama di bawahnya.
	fmt.Scanln(&pilihan)

	//percabangan/sekuesi bersyarat berdasarkan input pengguna
	if pilihan == 1 {
		fmt.Print("masukkan nama siswa:")

	//membuat alat pembaca (scanner) baru yang mengambil masukkan dari keyboard
	scanner := bufio.NewScanner(os.Stdin)

	//proses membaca masukan dari pengguna sampai tombol enter ditekan
	scanner.Scan()

	//mengambil text (nama) yang baru saja di baca dan menyimpannya ke variabel 
	nama = scanner.Text()

	fmt.Print("masukkan nilai siswa:")
	fmt.Scanln(&nilai)

	//menentukan grade nilai
	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	//%s adalah format (placeholder) untuk mencetak data bertipe string
	//% pertama akan digantikan oleh isi variabel 'nama', %s kedua oleh 'grade'
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)

	} else if pilihan == 0 {
		//di jalankan jika pengguna mengetik 0
		fmt.Println("keluar dari program.")

	} else {
		//dijalankan jika pengguna mengitik angka selain 1 dan 0 
		fmt.Println("pilihan tidak valid")
	}
}