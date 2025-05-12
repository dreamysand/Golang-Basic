package main

import "fmt"

func main() {
	// Slice pada dasarnya sama seperti array hanya saja lebih fleksibel dan dinamis
	// berbeda dengan array, slice memiliki length yang dapat bertambah dan berkurang sesuai keinginan
	// cara membuat slice ada 4 yaitu:
	// 1. Classic way
	// var nama_slice = []data_type{value1, value2}
	var tumbuhan = []string{"Pohon", "Rumput", "Rotan"}
	// 2. Short way
	// nama_slice := []data_type{value1, value2}
	hewan_langka := []string{"Anoa", "Badak Jawa", "Komodo", "Burung Cendrawasih"}
	// 3. Membuat dari array
	angka := [...]int{-2,-1,0,1,2,3,4}
	// nama_slice := nama_array[index_start:index_end]
	angka_prima := angka[3:6] 
	//Kalau index_end nya kosong maka sampai akhir, begitu pula kalau index awalnya kosong maka dari awal
	angka_pos := angka[3:]
	angka_neg := angka[:2]

	// angka_neg = append(angka_neg, -9, -9, -9, -9, -9, -9)
	// 4. Menggunakan make()
	// nama_slice := make([]data_type, length, capacity)
	maketes := make([]int, 5, 10)
	fmt.Printf("Isi dari slicenya adalah %v, lalu lengthnya %d, dan kapasitasnya %d \n", tumbuhan, len(tumbuhan), cap(tumbuhan))
	fmt.Printf("Isi dari slicenya adalah %v, lalu lengthnya %d, dan kapasitasnya %d \n", hewan_langka, len(hewan_langka), cap(hewan_langka))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima, len(angka_prima), cap(angka_prima))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_pos, len(angka_pos), cap(angka_pos))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_neg, len(angka_neg), cap(angka_neg))
	fmt.Printf("Isi dari arraynya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka, len(angka), cap(angka))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", maketes, len(maketes), cap(maketes))
	println()
	/*
	Modifikasi Slice
	 */
	// Mengganti isi slice => mengganti isi existing array nya
	angka_pos[3] = 6
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_pos, len(angka_pos), cap(angka_pos))
	fmt.Printf("Isi dari arraynya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka, len(angka), cap(angka))
	println()
	// Append ke slice => menambahkan data ke dalam slice
	angka_prima2 := append(angka_prima, 5) // Jika tidak melebihi capacity maka akan mengganti data di dalam existing array dan slice yang terkait
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima2, len(angka_prima2), cap(angka_prima2))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima, len(angka_prima), cap(angka_prima))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_pos, len(angka_pos), cap(angka_pos))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_neg, len(angka_neg), cap(angka_neg))
	fmt.Printf("Isi dari arraynya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka, len(angka), cap(angka))
	println()
	angka_prima2 = append(angka_prima, 5, 7, 11) // Jika melebihi maka akan membuat array baru
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima2, len(angka_prima2), cap(angka_prima2))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima, len(angka_prima), cap(angka_prima))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_pos, len(angka_pos), cap(angka_pos))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_neg, len(angka_neg), cap(angka_neg))
	fmt.Printf("Isi dari arraynya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka, len(angka), cap(angka))
	println()
	// Copy slice
	angka_prima3 := make([]int, len(angka_prima2), cap(angka_prima2)) //Pastikan length dan cap nya sama agar tidak terpotong
	// Cara copy : copy(to, ref)
	copy(angka_prima3, angka_prima2)
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima3, len(angka_prima3), cap(angka_prima3))
	fmt.Printf("Isi dari slicenya adalah %d, lalu lengthnya %d, dan kapasitasnya %d \n", angka_prima2, len(angka_prima2), cap(angka_prima2))
}