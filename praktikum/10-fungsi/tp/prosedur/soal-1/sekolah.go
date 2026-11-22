package main

import "fmt"

const NMAX = 1001

type tabMurid [NMAX]string

// Prosedur mencari nama murid
func cariMurid(T tabMurid, n int, nama string, posisi *int, ketemu *bool) {
    *ketemu = false
    *posisi = 0

    for i := 0; i < n; i++ {
        if T[i] == nama {
            *ketemu = true
            *posisi = i + 1
            return
        }
    }
}

// Prosedur menghitung huruf awal
func hitungHuruf(T tabMurid, n int, h byte, jumlah *int) {
    *jumlah = 0

    for i := 0; i < n; i++ {
        if T[i][0] == h {
            *jumlah++
        }
    }
}

func main() {
    var T tabMurid
    var n int
    var nama string
    var posisi, jumlah int
    var ketemu bool

    // Input jumlah murid
    fmt.Scan(&n)

    // Input nama murid
    for i := 0; i < n; i++ {
        fmt.Scan(&T[i])
    }

    // Input nama yang dicari
    fmt.Scan(&nama)

    // Memanggil prosedur
    cariMurid(T, n, nama, &posisi, &ketemu)

    huruf := nama[0]
    hitungHuruf(T, n, huruf, &jumlah)

    // Menampilkan hasil
    if ketemu {
        fmt.Printf("Murid terdaftar dan berada di urutan ke-%d\n", posisi)
    } else {
        fmt.Println("Murid tidak terdaftar")
    }

    fmt.Printf("Jumlah murid dengan huruf awal %c: %d\n", huruf, jumlah)
}