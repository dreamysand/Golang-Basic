package main

import "fmt"

func main() {
	// const => Konstan => nilainya tidak bisa dirubah
	// Ada beberapa cara untuk mendefinisikan const
	// 1. Classic way
	// const nama_const data_type = value
	const pi float32 = 3.14
	// 2. Tanpa tipe data => Go akan mengidetifikasi tipe data nya sendiri
	// const nama_const = value
	const golden_ratio = 1.61803398875
	// Multiple decl
	const (
		warna_laut string = "Biru"
		warna_darat string = "Hijau"
	)
	// Const tidak akan error jika tidak dipakai, karena Golang berpikir mungkin akan disimpan untuk suatu operasi
	fmt.Printf("Isi constnya %f, tipe datanya %T \n", pi, pi)
	fmt.Printf("Isi constnya %f, tipe datanya %T \n", golden_ratio, golden_ratio)
	fmt.Printf("Isi constnya %v, tipe datanya %T \n", warna_laut, warna_laut)
	fmt.Printf("Isi constnya %v, tipe datanya %T \n", warna_darat, warna_darat)
}