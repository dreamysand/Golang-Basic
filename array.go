package main

import "fmt"

func main() {
	// Array => sekumpulan data
	// Ada beberapa cara mendeklarasikannya antara lain:
	// 1. Classic way
	// var nama_array = [length]data_type{value1, value2}
	var buah_buahan = [2]string{"Manggis", "Alpukat"}
	// 2. Classic way tannpa length
	// var nama_array = [...]data_type{value1, value2} => Jadi panjangnya bisa bebas
	var hewan = [...]string{"Kucing", "Tupai", "Kuda nil", "Burung Beo"}

	// 1. Short way
	// nama_array := [length]data_type{value1, value2}
	merek_laptop := [2]string{"Asus", "Acer"}
	// 2. Short way tannpa length
	// nama_array := [...]data_type{value1, value2} => Jadi panjangnya bisa bebas
	makanan := [...]string{"Nasi goreng", "Ikan bakar", "Rendang"}

	fmt.Printf("Isi dari array nya adalah %v dan lengthnya %d \n", buah_buahan, len(buah_buahan))
	fmt.Printf("Isi dari array nya adalah %v dan lengthnya %d \n", hewan, len(hewan))
	fmt.Printf("Isi dari array nya adalah %v dan lengthnya %d \n", merek_laptop, len(merek_laptop))
	fmt.Printf("Isi dari array nya adalah %v dan lengthnya %d \n", makanan, len(makanan))

	// Ganti isi array
	// Isi dari array bisa diganti dengan cara
	// nama_array[index] = value_baru
	hewan[0] = "Kijang"
	fmt.Printf("Isi dari array nya adalah %v dan lengthnya %d \n", hewan, len(hewan))
}