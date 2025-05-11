package main

import "fmt"

 func main() {
 	// tipe data number ada banyak, dibagi menjadi beberapa golongan yaitu int, uint dan float
 	var (
 		// Int
 		intval1 int8 = 127
 		intval2 int16 = 32767
 		intval3 int32 = 2147483647
 		intval4 int64 = 9223372036854775807
 		intval5 int = 9223372036854775807 //Tergantung pengaturan OS
 		intval6 int8 = int8(intval4) //Contoh konversi tipe data number
 		// Uint => Unsigned integer => Tidak ada nilai negatif, nilai terkecilnya 0
 		uintval1 uint8 = 255
 		uintval2 uint16 = 65535
 		uintval3 uint32 = 4294967295
 		uintval4 uint64 = 18446744073709551615
 		uintval5 uint = 18446744073709551615 //Tergantung pengaturan OS
 		uintval6 uint8 = uint8(uintval4) //Contoh konversi tipe data number
 		// Float => nilai desimal => angka koma
 		floatval1 float32 = 21.474836
 		floatval2 float64 = 922337.20368547758
 		floatval3 float32 = float32(floatval2)
 		// Tipe data yang aslinya number tapi bukan number
 		byteval byte = 255 // Ini alias untuk uint8
 		runeval rune = 2147483647 // Ini alias untuk int32
 		
 	)

 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval1, intval1)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval2, intval2)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval3, intval3)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval4, intval4)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval5, intval5)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", intval6, intval6)
 	println()
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval1, uintval1)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval2, uintval2)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval3, uintval3)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval4, uintval4)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval5, uintval5)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", uintval6, uintval6)
 	println()
 	fmt.Printf("Angka : %f, tipe datanya: %T \n", floatval1, floatval1)
 	fmt.Printf("Angka : %f, tipe datanya: %T \n", floatval2, floatval2)
 	fmt.Printf("Angka : %f, tipe datanya: %T \n", floatval3, floatval3)
 	println()
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", byteval, byteval)
 	fmt.Printf("Angka : %d, tipe datanya: %T \n", runeval, runeval)
 	println()
 }