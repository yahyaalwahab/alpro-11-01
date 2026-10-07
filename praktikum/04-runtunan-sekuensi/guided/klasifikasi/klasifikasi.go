package main

import (
	"fmt" //hanya memerlukan fmt untuk keperluan input dan output
)

func main(){
	//pendefinisian variabel secara ekplisit beserta tipe datanya
	var usia int
	var gaji int
	var keterangan string

	//membaca dua masukan berupa angka dari pengguna (usia dan gaji)
	//fmt.Scan otomatis memisahkan input berdasarkan spasi atau baris baru (enter)
	fmt.Scan(&usia)
	fmt.Scan(&gaji)

	//menggunakan switch tanpa ekspresi.
	//cara kerjanya sama persis seperti if-else if.
	//program akan mengecek dari atas ke bawah, dan menjalankan case pertama yang bernilai benar (true)
	switch {
	case usia < 18:
		keterangan = "masih sekolah"
	case usia >= 18 && usia <= 25 && gaji >= 50:
		keterangan = "muda sukses"
	case usia >= 18 && usia <= 25 && gaji < 50:
		keterangan = "masih belajar hidup"
	case usia >= 26 && usia <= 40 && gaji >= 100:
		keterangan = "pekerja mapan"
	case usia >= 26 && usia <= 40 && gaji < 100:
		keterangan = "perlu perbaiki karier"
	case usia > 40 && gaji >= 150:
		keterangan = "profesional berpengalaman"
	case usia > 40 && gaji < 150:
		keterangan = "perllu evaluasi finansal"
	}

	//menampilkan hsail klasifikasi ke layar
	fmt.Println(keterangan)
}