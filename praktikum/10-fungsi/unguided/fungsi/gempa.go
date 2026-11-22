package main

import "fmt"

const NMAX = 999

type gempa struct {
	nama      string
	tahun     int
	magnitudo float64
}

type tabGempa [NMAX]gempa

func kategori(m float64) string {
	if m < 4.0 {
		return "minor"
	} else if m < 6.0 {
		return "sedang"
	} else if m < 7.0 {
		return "kuat"
	} else {
		return "besar"
	}
}

func rataMagnitudo(T tabGempa, n int) float64 {
	total := 0.0

	for i := 0; i < n; i++ {
		total += T[i].magnitudo
	}

	return total / float64(n)
}

func idxTerkuat(T tabGempa, n int) int {
	idx := 0

	for i := 0; i < n; i++ {
		if T[i].magnitudo > T[idx].magnitudo {
			idx = i
		}
	}

	return idx
}

func hitungKategori(T tabGempa, n int, k string) int {
	jumlah := 0

	for i := 0; i < n; i++ {
		if kategori(T[i].magnitudo) == k {
			jumlah++
		}
	}

	return jumlah
}

func main() {
	var T tabGempa
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].magnitudo)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s %d %.1f %s\n",
			T[i].nama,
			T[i].tahun,
			T[i].magnitudo,
			kategori(T[i].magnitudo))
	}

	fmt.Printf("Rata-rata magnitudo: %.2f\n", rataMagnitudo(T, n))

	terkuat := idxTerkuat(T, n)

	fmt.Printf("Gempa terkuat: %s %d %.1f\n",
		T[terkuat].nama,
		T[terkuat].tahun,
		T[terkuat].magnitudo)

	fmt.Printf("Jumlah gempa kategori besar: %d\n",
		hitungKategori(T, n, "besar"))
}