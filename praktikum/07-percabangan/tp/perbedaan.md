1. if
Digunakan untuk menjalankan kode jika suatu kondisi bernilai benar (true). Jika kondisi salah, kode di dalam if tidak dijalankan.

2. if-else
Digunakan untuk memilih salah satu dari dua kemungkinan. Jika kondisi if benar, blok if dijalankan. Jika salah, blok else dijalankan.

3. else-if
Digunakan untuk memeriksa beberapa kondisi secara berurutan. Kondisi diperiksa dari atas ke bawah, dan blok pertama yang kondisinya benar akan dijalankan.

4. switch
Digunakan untuk memilih tindakan berdasarkan nilai atau kondisi tertentu. switch dapat menjadi alternatif dari banyak percabangan if-else, terutama ketika membandingkan satu nilai dengan beberapa pilihan.

5. 2 perbedaan `switch` pada Go dan Java
    Go
    1. Tidak memerlukan kata kunci `break`pada setiap `case` karena secara default eksekusi berhenti setelah satu kasus yang cocok selesai.
    2. `switch` dapat di gunakan tanpa ekspresi, sehingga setiap `case` bisa berisi kondisi boolean.

    Java
    1. Pada bentuk `switch` klasik, `break` biasanya diperlukan agar eksekusi tidak berlanjut ke `case` berikutnya (fall-through).
    2. Pada `switch` klasik, setiap `case` umumnya berisi nilai atau label yang dibandingkan dengan ekspresi`switch`, bukan kondisi boolean bebas.