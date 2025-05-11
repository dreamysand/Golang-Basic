package main

import "fmt"

func main() {
	// Variabel => tempat sementara untuk meletakkan value
	// Ada beberapa cara untuk mendefinisikan variabel
	// 1. Classic way
	// var nama_var tipe_data = value
	var x int = 10
	// 2. Tanpa tipe data => Go akan mengidetifikasi tipe data nya sendiri
	// var nama_var = value
	var y = "Halo"
	// 3. Tanpa value
	// var nama_var tipe_data
	var z int8
	z = 11
	// 4. Versi singkat
	w := true

	// Multiple declaration
	
	// cara 1:
	var (
		a int = 10
		b string = "Dunia"
		k = 3.14
	)
	// cara 2 :
	// Mendefinisikan variabel yang data_type nya sama
	var c,d,e,f int = 1,2,3,4
	// Mendefinisikan variabel yang data_type nya beda beda
	var g,h,i,j = "Halo", 1, 5, 8.17

	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", x, x)
	fmt.Printf("Isi variabelnya %v, tipe datanya %T \n", y, y)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", z, z)
	fmt.Printf("Isi variabelnya %v, tipe datanya %T \n", w, w)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", a, a)
	fmt.Printf("Isi variabelnya %v, tipe datanya %T \n", b, b)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", c, c)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", d, d)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", e, e)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", f, f)
	fmt.Printf("Isi variabelnya %v, tipe datanya %T \n", g, g)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", h, h)
	fmt.Printf("Isi variabelnya %d, tipe datanya %T \n", i, i)
	fmt.Printf("Isi variabelnya %f, tipe datanya %T \n", j, j)
	fmt.Printf("Isi variabelnya %f, tipe datanya %T \n", k, k)

	// Aturan penamaan variabel Go: 
	// Nama variabel harus diawali dengan huruf atau karakter garis bawah (_) 
	// Nama variabel tidak boleh diawali dengan angka 
	// Nama variabel hanya boleh berisi karakter alfanumerik dan garis bawah (a-z, A-Z, 0-9, dan _) 
	// Nama variabel peka huruf besar-kecil (age, Age, dan AGE adalah tiga variabel yang berbeda) 
	// Tidak ada batasan panjang nama variabel 
	// Nama variabel tidak boleh berisi spasi 
	// Nama variabel tidak boleh berupa kata kunci Go apa pun
}