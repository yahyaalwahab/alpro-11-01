package main

import "fmt"

const NMAX = 999

type rapot struct {
    matkul string
    nilai  int
}

type tabRapot [NMAX]rapot

// Fungsi menentukan nilai huruf
func nilaiHuruf(x int) string {
    if x > 80 {
        return "A"
    } else if x > 70 {
        return "AB"
    } else if x > 65 {
        return "B"
    } else if x > 60 {
        return "BC"
    } else if x > 50 {
        return "C"
    } else if x > 40 {
        return "D"
    } else {
        return "E"
    }
}

// Fungsi menghitung rata-rata
func rataRata(T tabRapot, n int) float64 {
    total := 0

    for i := 0; i < n; i++ {
        total += T[i].nilai
    }

    return float64(total) / float64(n)
}

func main() {
    var T tabRapot
    var n int

    fmt.Scan(&n)

    // Input data mata kuliah
    for i := 0; i < n; i++ {
        fmt.Scan(&T[i].matkul, &T[i].nilai)
    }

    // Menampilkan hasil
    for i := 0; i < n; i++ {
        fmt.Println(T[i].matkul, T[i].nilai, nilaiHuruf(T[i].nilai))
    }

    fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}