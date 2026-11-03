package main

import (
	"fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

var reader = bufio.NewReader(os.Stdin)

func masukan(msg string) string {
	fmt.Print(msg)
	text,_:=reader. ReadString('\n')
	return strings.TrimSpace(text)
}

func daftarmenu(){
	fmt.Println("1. Tampilkan  daftar tabungan")
	fmt.Println("2. Buat tabungan baru")
	fmt.Println("3. Isi tabungan")
	fmt.Println("4. Ambil isi tabungan")
	fmt.Println("5. Edit tabungan")
	fmt.Println("6. Hapus tabungan")
	fmt.Println("0. keluar")

}
type tabungan struct {
	nama string
	target float64
	terkumpul float64
	status string
}

func buatTabunganBaru(daftar []tabungan, namaTabungan string, targetTabungan float64) []tabungan {
	tabunganBaru := tabungan{nama: namaTabungan, target: targetTabungan, terkumpul:0.0, status: "proses",
}
fmt.Printf("\nTabungan berhasil di buat\n")

	return append(daftar, tabunganBaru)
}

//Tampilkan tabungan 
func tampilkanTabungan(daftar []tabungan) {
	if len(daftar) == 0 {
		fmt.Println("tabungan belum pernah di buat")
	} else {
		fmt.Printf("%-4s %-25s %-18s %-18s %-18s %-10s", "No.", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")
		fmt.Println(strings.Repeat("-", 85))
	}

	for i, v := range daftar {
		fmt.Printf("%-4d %-25s %-18.0f %-18.0f %-18.0f %-10s", i+1, v.nama, v.target, v.terkumpul, (v.target - v.terkumpul), v.status)
	}
}

// Isi tabungan yang sudah dibuat
func isiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			nominal, err := strconv.ParseFloat(masukan("Masukkan nominal:"), 64)

			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}

			daftar[i].terkumpul += nominal
			fmt.Printf("\nTabungan %q berhasil diisi\n", v.nama)
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Ambil isi Tabungan
func ambilIsiTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			nominal, err := strconv.ParseFloat(masukan("Nominal penarikan: "), 64)

			if err != nil {
				fmt.Println("Nominal harus berupa angka!")
				return
			}

			if nominal > v.terkumpul {
				fmt.Printf("\nSaldo tidak mencukupi\n")
				return
			}

			daftar[i].terkumpul -= nominal
			fmt.Printf("\nTabungan berhasil diambil\n")
			break
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Edit tabungan
func editTabungan(daftar []tabungan, namaTabungan string) {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			namaBaru := masukan("Masukkan nama tabungan baru: ")
			targetBaru, err := strconv.ParseFloat(masukan("Masukkan target tabungan baru: Rp."), 64)

			if err != nil {
				fmt.Printf("\nNominal harus berupa angka\n")
				return
			}

			daftar[i].nama = namaBaru
			daftar[i].target = targetBaru
			fmt.Printf("\nTabungan berhasil diedit\n")

			return
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
}

// Hapus tabungan
func hapusTabungan(daftar []tabungan, namaTabungan string) []tabungan {
	for i, v := range daftar {
		if strings.EqualFold(v.nama, namaTabungan) {
			fmt.Printf("\nTabungan %q berhasil dihapus\n", v.nama)
			return append(daftar[:i], daftar[i+1:]...)
		} else {
			fmt.Printf("\nNama tabungan tidak ditemukan\n")
		}
	}
	return daftar
}

func main() {
	daftarTabungan := []tabungan{}
	fmt.Println("Tabungan Gueh")
	fmt.Println()

	for {
		daftarmenu()

		pilihan, _ := strconv.Atoi(masukan("Masukkan nomor perintah: "))

		if pilihan == 0 {
			break
		}

		switch pilihan {
		case 1:
			tampilkanTabungan(daftarTabungan)
		case 2:
			nama := masukan("Masukkan nama: ")
			target, err := strconv.ParseFloat(masukan("Masukan target: Rp."), 64)

			if err != nil {
				fmt.Println("Target tabungan harus berupa angka")
			} else {
				daftarTabungan = buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			nama := masukan("Masukkan nama tabungan: ")
			isiTabungan(daftarTabungan, nama)
		case 4:
			nama := masukan("Masukkan nama tabungan yang ingin diambil: ")
			ambilIsiTabungan(daftarTabungan, nama)
		case 5:
			nama := masukan("Masukkan nama tabungan yang ingin diedit: ")
			editTabungan(daftarTabungan, nama)
		case 6:
			nama := masukan("Masukkan nama tabungan yang ingin dihapus: ")
			daftarTabungan = hapusTabungan(daftarTabungan, nama)
		}
	}
}
