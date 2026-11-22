package main

import "fmt"

const NMAX = 1000

type penduduk struct {
    nama  string
    tahun int
    kota  string
}

type tabPenduduk [NMAX]penduduk

// Prosedur mencari tahun lahir
func cariTahun(T tabPenduduk, n, th int) int {
    for i := 0; i < n; i++ {
        if T[i].tahun == th {
            return i
        }
    }
    return -1
}

// Prosedur mencari penduduk tertua dan termuda
func tertuaTermuda(T tabPenduduk, n int, tua, muda *penduduk) {
    *tua = T[0]
    *muda = T[0]

    for i := 1; i < n; i++ {
        if T[i].tahun < tua.tahun {
            *tua = T[i]
        }

        if T[i].tahun > muda.tahun {
            *muda = T[i]
        }
    }
}

func main() {
    var T tabPenduduk
    var n, th, idx int
    var tua, muda penduduk

    // Input jumlah penduduk
    fmt.Scan(&n)

    // Input data penduduk
    for i := 0; i < n; i++ {
        fmt.Scan(&T[i].nama, &T[i].tahun, &T[i].kota)
    }

    // Input tahun yang dicari
    fmt.Scan(&th)

    // Mencari tahun lahir
    idx = cariTahun(T, n, th)

    if idx != -1 {
        fmt.Printf("%s %d %s ditemukan di index ke-%d\n",
            T[idx].nama, T[idx].tahun, T[idx].kota, idx)
    } else {
        fmt.Println("Data tidak ditemukan")
    }

    // Mencari penduduk tertua dan termuda
    tertuaTermuda(T, n, &tua, &muda)

    fmt.Printf("Penduduk tertua: %s (%d)\n", tua.nama, tua.tahun)
    fmt.Printf("Penduduk termuda: %s (%d)\n", muda.nama, muda.tahun)
}