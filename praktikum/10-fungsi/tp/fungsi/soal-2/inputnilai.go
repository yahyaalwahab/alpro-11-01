package main

import "fmt"

const NMAX = 100

type student struct {
    nim   string
    name  string
    grade float64
}

type students [NMAX]student

// Fungsi mencari NIM
func cariNIM(T students, n int, nim string) int {
    for i := 0; i < n; i++ {
        if T[i].nim == nim {
            return i
        }
    }
    return -1
}

// Fungsi menghitung rata-rata
func rataRata(T students, n int) float64 {
    if n == 0 {
        return 0
    }

    total := 0.0

    for i := 0; i < n; i++ {
        total += T[i].grade
    }

    return total / float64(n)
}

func main() {
    var T students
    var n int

    var nim, nama string
    var nilai float64

    n = 0

    // Input data mahasiswa
    for {
        fmt.Scan(&nim)

        if nim == "STOP" {
            break
        }

        fmt.Scan(&nama, &nilai)

        // Memeriksa apakah NIM sudah ada
        if cariNIM(T, n, nim) == -1 {
            T[n].nim = nim
            T[n].name = nama
            T[n].grade = nilai
            n++
        }
    }

    // Menampilkan data mahasiswa yang sudah bersih
    for i := 0; i < n; i++ {
        fmt.Printf("%s %s %.2f\n",
            T[i].nim, T[i].name, T[i].grade)
    }

    fmt.Printf("Rata-rata nilai: %.2f\n", rataRata(T, n))
}