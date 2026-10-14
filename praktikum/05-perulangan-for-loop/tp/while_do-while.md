1. For dengan tiga komponen.
   Perulangan for dengan tiga komponen memiliki inisialisasi, kondisi, dan post statement. Biasanya digunakan ketika jumlah perulangan sudah diketahui.

   contoh kode:
   for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
   penjelasan:
   i := 1 sebagai nilai awal.
   i <= 5 sebagai kondisi perulangan.
   i++ untuk menambah nilai i setiap perulangan.

2. For yang hanya menggunakan kondisi.
   Perulangan for yang hanya menggunakan kondisi bekerja seperti while pada bahasa pemrograman lain. Perulangan akan terus berjalan selama kondisi bernilai true.

   contoh kode:
   i := 1
for i <= 5 {
    fmt.Println(i)
    i++
}
    penjelasan:
    i := 1 menentukan nilai awal di luar perulangan.
    i <= 5 menjadi kondisi perulangan.
    i++ menambah nilai i di dalam blok kode.