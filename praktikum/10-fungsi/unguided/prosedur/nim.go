package main

import "fmt"

const NMAX = 2002

type tabNIM [NMAX]int

func cariNIM(T tabNIM, n, x int) int {
	for i := 0; i < n; i++ {
		if T[i] == x {
			return i
		}
	}
	return -1
}

func ubahNIM(T *tabNIM, idx, baru int) {
	T[idx] = baru
}

func cetak(T tabNIM, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(T[i])
	}
}

func main() {
	var n, x, baru int
	var T tabNIM

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i])
	}

	fmt.Scan(&x)

	idx := cariNIM(T, n, x)

	if idx != -1 {
		fmt.Scan(&baru)

		ubahNIM(&T, idx, baru)

		fmt.Printf("Mahasiswa dengan NIM %d terdaftar pada posisi ke-%d\n", x, idx+1)
		fmt.Println("Daftar NIM setelah terjadi perubahan data:")
		cetak(T, n)
	} else {
		fmt.Println("NIM tidak ditemukan dalam sistem")
	}
}